package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app/test"
	"github.com/ihcsim/bluelens/internal/core"
)

const (
	limit  = 20
	offset = 0
)

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
	recommendations, err := fixtureStore.Recommendations(limit)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	service := goa.New("goatest")
	ctrl := NewRecommendationsController(service)

	// test the recommendations of all users in the store
	users, err := store().ListUsers(limit, offset)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range users {
		_, actual := test.RecommendRecommendationsOK(t, nil, nil, ctrl, user.ID, limit)
		expected := mediaTypeRecommendations(recommendations[user.ID])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations response mismatch. Expected:\n%+v\nBut got:\n%+v", expected, actual)
		}
	}
}
