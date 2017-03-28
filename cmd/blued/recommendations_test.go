package main

import (
	"context"
	"reflect"
	"testing"
	"time"

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
	t.Run("ok", func(t *testing.T) {
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
	})

	t.Run("not found", func(t *testing.T) {
		userID, limit := "user-00", 10
		expectedErr := core.NewEntityNotFound(userID, "user")
		if _, actual := test.RecommendRecommendationsNotFound(t, nil, nil, ctrl, userID, limit); !reflect.DeepEqual(actual, expectedErr) {
			t.Errorf("Error mismatch. Expected %v, but got %v", expectedErr, actual)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
		defer cancel()

		userID, limit := "user-00", 10
		if _, actual := test.RecommendRecommendationsInternalServerError(t, ctx, nil, ctrl, userID, limit); actual != context.DeadlineExceeded {
			t.Errorf("Error mismatch. Expected %v, but got %v", context.DeadlineExceeded, actual)
		}
	})
}
