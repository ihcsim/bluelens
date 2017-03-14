package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	core "github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
)

const maxCount = 20

func TestRecommendationsController(t *testing.T) {
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
	fixtureStore := store().(*core.FixtureStore)

	// retrieve the expected recommendations from the fixture store
	recommendations, err := fixtureStore.Recommendations(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	service := goa.New("goatest")
	ctrl := NewRecommendationsController(service)

	users, err := store().ListUsers(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range users {
		_, actual := test.RecommendRecommendationsOK(t, nil, nil, ctrl, user.ID, maxCount)
		expected := mediaTypeRecommendations(recommendations[user.ID])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations response mismatch. Expected:\n%+v\nBut got:\n%+v", expected, actual)
		}
	}
}
