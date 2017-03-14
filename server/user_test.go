package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
)

func TestUserController(t *testing.T) {
	// mock the store() function
	storeFunc := store
	storeFuncMock := func() core.Store {
		s, err := core.NewFixtureStore()
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		return s
	}
	store = storeFuncMock
	defer func() {
		store = storeFunc
	}()

	svc := goa.New("goatest")
	ctrl := NewUserController(svc)

	userID := "user-01"
	user, err := store().FindUser(userID)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	expected := mediaTypeUser(user)
	if _, actual := test.GetUserOK(t, nil, nil, ctrl, userID); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Media type mismatch. Expected %+v, but got %+v", expected, actual)
	}
}
