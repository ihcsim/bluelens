package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
)

var (
	// store singleton
	userCtrlStoreFixture core.Store

	// ensures test store is init only once
	testStoreInit sync.Once
)

func TestUserController(t *testing.T) {
	// mock the store() function
	storeFunc := store
	storeFuncMock := func() core.Store {
		var err error
		testStoreInit.Do(func() {
			userCtrlStoreFixture, err = core.NewFixtureStore()
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})

		return userCtrlStoreFixture
	}
	store = storeFuncMock
	defer func() {
		store = storeFunc
		userCtrlStoreFixture = nil
	}()

	svc := goa.New("goatest")
	ctrl := NewUserController(svc)

	t.Run("get", func(t *testing.T) {
		t.Run("found", func(t *testing.T) {
			user, err := store().FindUser("user-01")
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			expected := mediaTypeUser(user)
			if _, actual := test.GetUserOK(t, nil, nil, ctrl, user.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Media type mismatch. Expected %+v, but got %+v", expected, actual)
			}
		})

		t.Run("not found", func(t *testing.T) {
			if _, err := test.GetUserNotFound(t, nil, nil, ctrl, "example"); err == nil {
				t.Errorf("Expected EntityNotFound error to occur")
			}
		})
	})

	t.Run("follow", func(t *testing.T) {
		user, err := store().FindUser("user-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		t.Run("self", func(t *testing.T) {
			// no changes to expected media type result if followee is self
			expected := mediaTypeUser(user)
			if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, user.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}
		})

		t.Run("new followee", func(t *testing.T) {
			followee, err := store().FindUser("user-03")
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}
			user.Followees = append(user.Followees, followee)
			expected := mediaTypeUser(user)

			if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, followee.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}

			t.Run("repeat", func(t *testing.T) {
				// following the same followee again should have no effects
				if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, followee.ID); !reflect.DeepEqual(expected, actual) {
					t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
				}
			})
		})

		t.Run("not found", func(t *testing.T) {
			t.Run("user", func(t *testing.T) {
				if _, err := test.FollowUserNotFound(t, nil, nil, ctrl, "example", "user-01"); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})

			t.Run("followee", func(t *testing.T) {
				if _, err := test.FollowUserNotFound(t, nil, nil, ctrl, user.ID, "example"); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})
		})
	})

	t.Run("listen", func(t *testing.T) {
		user, err := store().FindUser("user-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		t.Run("same music", func(t *testing.T) {
			// re-listening to music already in the history list should have no effects
			expected := mediaTypeUser(user)
			if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, user.History[0].ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%s\nBut got:\n%s\n", expected.Links, actual.Links)
			}
		})

		t.Run("new music", func(t *testing.T) {
			music, err := store().FindMusic("song-05")
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}
			user.History = append(user.History, music)

			expected := mediaTypeUser(user)
			if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, music.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%s\nBut got:\n%s\n", expected.Links, actual.Links)
			}

			t.Run("repeat", func(t *testing.T) {
				// listening to the same music should have no effects on the user's history list
				if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, music.ID); !reflect.DeepEqual(expected, actual) {
					t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
				}
			})
		})

		t.Run("not found", func(t *testing.T) {
			t.Run("user", func(t *testing.T) {
				if _, err := test.ListenUserNotFound(t, nil, nil, ctrl, "example", "song-00"); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})

			t.Run("music", func(t *testing.T) {
				if _, err := test.ListenUserNotFound(t, nil, nil, ctrl, user.ID, "example"); err == nil {
					t.Error("Expected EntityNotFound error to occur")
				}
			})

		})
	})
}
