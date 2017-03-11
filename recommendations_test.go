package core

import (
	"reflect"
	"testing"
)

func TestRecommend(t *testing.T) {
	maxCount := 20
	store, err := NewFixtureStore()
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	recommendations, err := store.Recommendations(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	actual, err := store.ListUsers(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range actual {
		actual, err := RecommendSort(user.ID, maxCount, store)
		if err != nil {
			t.Fatalf("Unexpected err: %s. Fixture user: %q", err, user.ID)
		}

		expected := recommendations[user.ID]
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations mismatch for user %q. Expected:\n%+v\nBut got:\n%+v", user.ID, expected, actual)
		}
	}
}
