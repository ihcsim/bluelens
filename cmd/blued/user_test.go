package main

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/cmd/blued/app/test"
	"github.com/ihcsim/bluelens/internal/core"
)

var (
	// store fixture used for the user controller tests
	storeFixtureUser core.Store

	// ensures the store fixture used for the user controller tests are initialized once
	initStoreFixtureUser sync.Once

	// helper function to retrieve user fixtures from store
	userFixture = func(t *testing.T, id string) *core.User {
		u, err := store().FindUser(id)
		if err != nil {
			t.Fatalf("Unexpected error with user ID %s: %", id, err)
		}

		return u
	}
)

func TestUserController(t *testing.T) {
	// mock the store() function
	storeFunc := store
	storeFuncMock := func() core.Store {
		var err error
		initStoreFixtureUser.Do(func() {
			storeFixtureUser, err = core.NewFixtureStore()
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})

		return storeFixtureUser
	}
	store = storeFuncMock
	defer func() {
		store = storeFunc
		storeFixtureUser = nil
	}()

	svc := goa.New("goatest")
	ctrl := NewUserController(svc)

	t.Run("list", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			tests := []struct {
				limit  int
				offset int
			}{
				{offset: 0, limit: -1},
				{offset: 0, limit: 0},
				{offset: 0, limit: 5},
				{offset: 0, limit: 10},
				{offset: 0, limit: 20},
				{offset: 10, limit: 20},
				{offset: 5, limit: 10},
				{offset: 10, limit: 10},
				{offset: 15, limit: 10},
				{offset: -1, limit: 0},
			}

			for id, tc := range tests {
				userList, err := store().ListUsers(tc.limit, tc.offset)
				if err != nil {
					t.Error("Unexpected error with test case %d: ", id, err)
				}

				var expected app.BluelensUserCollection
				for _, user := range userList {
					expected = append(expected, mediaTypeUser(user))
				}

				if _, actual := test.ListUserOK(t, nil, nil, ctrl, tc.limit, tc.offset); !reflect.DeepEqual(expected, actual) {
					t.Errorf("Response mismatch. Test case: %d\nExpected %s\nBut got %s", id, expected, actual)
				}
			}
		})

		t.Run("timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
			defer cancel()

			if _, actual := test.ListUserInternalServerError(t, ctx, nil, ctrl, limit, offset); actual != context.DeadlineExceeded {
				t.Errorf("Error mismatch. Expected %v, byut got %v", context.DeadlineExceeded, actual)
			}
		})
	})

	t.Run("show", func(t *testing.T) {
		t.Run("found", func(t *testing.T) {
			user := userFixture(t, "user-01")
			expected := mediaTypeUserFull(user)
			if _, actual := test.ShowUserOKFull(t, nil, nil, ctrl, user.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Media type mismatch. Expected %+v, but got %+v", expected, actual)
			}
		})

		t.Run("not found", func(t *testing.T) {
			if _, err := test.ShowUserNotFound(t, nil, nil, ctrl, "example"); err == nil {
				t.Errorf("Expected EntityNotFound error to occur")
			}
		})

		t.Run("timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
			defer cancel()

			if _, actual := test.ShowUserInternalServerError(t, ctx, nil, ctrl, "example"); actual != context.DeadlineExceeded {
				t.Errorf("Error mismatch. Expected %v, byut got %v", context.DeadlineExceeded, actual)
			}
		})
	})

	t.Run("create", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			followeesID := []string{"user-02", "user-03"}
			musicID := []string{"song-07", "song-03"}
			fixture := &core.User{
				ID: "user-00.v2",
				Followees: core.UserList{
					userFixture(t, followeesID[0]),
					userFixture(t, followeesID[1]),
				},
				History: core.MusicList{
					musicFixture(t, musicID[0]),
					musicFixture(t, musicID[1]),
				},
			}

			payload := &app.User{
				ID: fixture.ID,
				Followees: []*app.User{
					&app.User{ID: followeesID[0]},
					&app.User{ID: followeesID[1]},
				},
				History: []*app.Music{
					&app.Music{ID: musicID[0]},
					&app.Music{ID: musicID[1]},
				},
			}

			expected := &app.BluelensUserLink{Href: "/bluelens/users/user-00.v2"}
			if _, actual := test.CreateUserCreatedLink(t, nil, nil, ctrl, payload); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Resource mismatch. Expected %+v\nBut got:%+v", expected, actual)
			}

			u := userFixture(t, fixture.ID)
			if !reflect.DeepEqual(fixture, u) {
				t.Errorf("Resource mismatch. Expected %s\nBut got %s", fixture, u)
			}
		})

		t.Run("bad request", func(t *testing.T) {
			tests := []struct {
				payload *app.User
			}{
				{payload: &app.User{}},
				{payload: &app.User{ID: ""}},
			}

			for _, tc := range tests {
				if _, err := test.CreateUserBadRequest(t, nil, nil, ctrl, tc.payload); err == nil {
					t.Error("Expected error to occur. Should have failed with a 400 response status.")
				}
			}
		})

		t.Run("timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
			defer cancel()

			payload := &app.User{ID: "user-02"}
			if _, actual := test.CreateUserInternalServerError(t, ctx, nil, ctrl, payload); actual != context.DeadlineExceeded {
				t.Errorf("Error mismatch. Expected %v, but got %v", context.DeadlineExceeded, actual)
			}
		})
	})

	t.Run("follow", func(t *testing.T) {
		user := userFixture(t, "user-01")

		t.Run("self", func(t *testing.T) {
			// no changes to expected media type result if followee is self
			expected := mediaTypeUser(user)
			payload := &app.FollowUserPayload{FolloweeID: &user.ID}
			if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, user.ID, payload); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}
		})

		t.Run("new followee", func(t *testing.T) {
			followee := userFixture(t, "user-03")
			if err := user.AddFollowees(followee); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			payload := &app.FollowUserPayload{FolloweeID: &followee.ID}
			expected := mediaTypeUser(user)
			if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, followee.ID, payload); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}

			t.Run("repeat", func(t *testing.T) {
				// following the same followee again should have no effects
				if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, followee.ID, payload); !reflect.DeepEqual(expected, actual) {
					t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
				}
			})
		})

		t.Run("not found", func(t *testing.T) {
			t.Run("user", func(t *testing.T) {
				userID := "user-01"
				payload := &app.FollowUserPayload{FolloweeID: &userID}
				if _, err := test.FollowUserNotFound(t, nil, nil, ctrl, "example", userID, payload); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})

			t.Run("followee", func(t *testing.T) {
				userID := "example"
				payload := &app.FollowUserPayload{FolloweeID: &userID}
				if _, err := test.FollowUserNotFound(t, nil, nil, ctrl, user.ID, userID, payload); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})
		})

		t.Run("timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
			defer cancel()

			userID := "example"
			payload := &app.FollowUserPayload{FolloweeID: &userID}
			if _, actual := test.FollowUserInternalServerError(t, ctx, nil, ctrl, user.ID, userID, payload); actual != context.DeadlineExceeded {
				t.Errorf("Error mismatch. Expected %v, but got %v", context.DeadlineExceeded, actual)
			}
		})
	})

	t.Run("listen", func(t *testing.T) {
		user := userFixture(t, "user-01")
		t.Run("same music", func(t *testing.T) {
			// re-listening to music already in the history list should have no effects
			payload := &app.ListenUserPayload{MusicID: &user.History[0].ID}
			expected := mediaTypeUser(user)
			if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, user.History[0].ID, payload); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%s\nBut got:\n%s\n", expected.Links, actual.Links)
			}
		})

		t.Run("new music", func(t *testing.T) {
			music := musicFixture(t, "song-05")
			if err := user.AddHistory(music); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			payload := &app.ListenUserPayload{MusicID: &music.ID}
			expected := mediaTypeUser(user)
			if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, music.ID, payload); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%s\nBut got:\n%s\n", expected.Links, actual.Links)
			}

			t.Run("repeat", func(t *testing.T) {
				// listening to the same music should have no effects on the user's history list
				if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, music.ID, payload); !reflect.DeepEqual(expected, actual) {
					t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
				}
			})
		})

		t.Run("not found", func(t *testing.T) {
			t.Run("user", func(t *testing.T) {
				musicID := "song-00"
				payload := &app.ListenUserPayload{MusicID: &musicID}
				if _, err := test.ListenUserNotFound(t, nil, nil, ctrl, "example", musicID, payload); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})

			t.Run("music", func(t *testing.T) {
				musicID := "example"
				payload := &app.ListenUserPayload{MusicID: &musicID}
				if _, err := test.ListenUserNotFound(t, nil, nil, ctrl, user.ID, musicID, payload); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})
		})

		t.Run("timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
			defer cancel()

			musicID := "example"
			payload := &app.ListenUserPayload{MusicID: &musicID}
			if _, actual := test.ListenUserInternalServerError(t, ctx, nil, ctrl, user.ID, musicID, payload); actual != context.DeadlineExceeded {
				t.Errorf("Error mismatch. Expected %v, but got %v", context.DeadlineExceeded, actual)
			}
		})
	})
}
