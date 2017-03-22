package core

import (
	"reflect"
	"testing"
)

func TestRecommend(t *testing.T) {
	limit := 20
	store, err := NewFixtureStore()
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	recommendations, err := store.Recommendations(limit)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	actual, err := store.ListUsers(limit)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range actual {
		actual, err := RecommendSort(user.ID, limit, store)
		if err != nil {
			t.Fatalf("Unexpected err: %s. Fixture user: %q", err, user.ID)
		}

		expected := recommendations[user.ID]
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations mismatch for user %q. Expected:\n%+v\nBut got:\n%+v", user.ID, expected, actual)
		}
	}
}
