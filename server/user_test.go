package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
)

// store singleton
var userCtrlStoreFixture core.Store

func TestUserController(t *testing.T) {
	// mock the store() function
	storeFunc := store
	storeFuncMock := func() core.Store {
		var err error
		once.Do(func() {
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
		t.Run("exist", func(t *testing.T) {
			user, err := store().FindUser("user-01")
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			expected := mediaTypeUser(user)
			if _, actual := test.GetUserOK(t, nil, nil, ctrl, user.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Media type mismatch. Expected %+v, but got %+v", expected, actual)
			}
		})
	})

	t.Run("follow", func(t *testing.T) {
		user, err := store().FindUser("user-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		t.Run("self", func(t *testing.T) {
			// no changes if followee is self
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

			// adding the same followee again should have no effects
			if _, actual := test.FollowUserOK(t, nil, nil, ctrl, user.ID, followee.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}
		})
	})

	t.Run("listen", func(t *testing.T) {
		user, err := store().FindUser("user-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		t.Run("self", func(t *testing.T) {
			_, err := store().FindMusic("example")
			if err == nil {
				t.Fatal("Expected error to occur")
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

			// adding the same music to the history list should have no effects
			if _, actual := test.ListenUserOK(t, nil, nil, ctrl, user.ID, music.ID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("User mismatch. Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
			}
		})
	})
}
