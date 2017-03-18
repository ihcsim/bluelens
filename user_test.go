package core

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ihcsim/bluelens/json"
)

func TestUser(t *testing.T) {
	t.Run("IsNew", func(t *testing.T) {
		user := &User{ID: "user-01"}
		if !user.IsNew() {
			t.Error("Expected user with no followees and history to be labeled as 'new'")
		}

		user.Followees = append(user.Followees, &User{ID: "user-02"})
		if user.IsNew() {
			t.Error("Expected user with followees and history to not be labeled as 'new'")
		}

		user.History = append(user.History, &Music{ID: "song-01"})
		if user.IsNew() {
			t.Error("Expected user with followees and history to not be labeled as 'new'")
		}
	})

	t.Run("HasFollowee", func(t *testing.T) {
		user := &User{ID: "user-01"}
		followee := &User{ID: "user-02"}
		if user.HasFollowee(followee.ID) {
			t.Error("Expected user to not have a followee")
		}

		user.Followees = append(user.Followees, followee)
		if !user.HasFollowee(followee.ID) {
			t.Error("Expected user to have a followee")
		}
	})

	t.Run("HasHistory", func(t *testing.T) {
		user := &User{ID: "user-01"}
		history := &Music{ID: "song-01"}
		if user.HasHistory(history.ID) {
			t.Error("Expected user to not have a music history")
		}

		user.History = append(user.History, history)
		if !user.HasHistory(history.ID) {
			t.Error("Expected user to have a music history")
		}
	})
}

func TestUserList(t *testing.T) {
	t.Run("BuildFrom", func(t *testing.T) {
		t.Run("user's history", func(t *testing.T) {
			tests := []struct {
				jsonObj  *json.UserHistory
				expected UserList
			}{
				{jsonObj: &json.UserHistory{}, expected: nil},
				{jsonObj: &json.UserHistory{
					Listens: map[string][]string{
						"user-01": []string{"song-01"}}},
					expected: UserList{
						&User{ID: "user-01", History: MusicList{&Music{ID: "song-01"}}}},
				},
				{jsonObj: &json.UserHistory{
					Listens: map[string][]string{
						"user-01": []string{"song-01"},
						"user-02": []string{"song-01", "song-02"}}},
					expected: UserList{
						&User{ID: "user-01", History: MusicList{&Music{ID: "song-01"}}},
						&User{ID: "user-02", History: MusicList{&Music{ID: "song-01"}, &Music{ID: "song-02"}}},
					},
				},
			}

			for id, test := range tests {
				var actual UserList
				if err := actual.Unmarshal(test.jsonObj); err != nil {
					t.Fatal("Unexpected error: ", err)
				}

				sort.Slice(actual, func(i, j int) bool {
					return strings.Compare(actual[i].ID, actual[j].ID) == -1
				})

				if !reflect.DeepEqual(test.expected, actual) {
					t.Errorf("User list mismatch. Test case %d. Expected:\n%s\nBut got:\n%s", id, test.expected, actual)
				}
			}
		})

		t.Run("user's followees", func(t *testing.T) {
			tests := []struct {
				jsonObj  *json.UserFollowees
				expected UserList
			}{
				{jsonObj: &json.UserFollowees{
					Follows: [][]string{{"user-01", "user-02"}}},
					expected: UserList{
						&User{ID: "user-01", Followees: UserList{&User{ID: "user-02"}}}},
				},
				{jsonObj: &json.UserFollowees{
					Follows: [][]string{
						{"user-01", "user-02"},
						{"user-01", "user-03"}}},
					expected: UserList{
						&User{ID: "user-01",
							Followees: UserList{
								&User{ID: "user-02"},
								&User{ID: "user-03"}}},
					},
				},
				{jsonObj: &json.UserFollowees{
					Follows: [][]string{
						{"user-01", "user-02"},
						{"user-01", "user-03"},
						{"user-02", "user-01"},
						{"user-02", "user-04"}}},
					expected: UserList{
						&User{ID: "user-01",
							Followees: UserList{
								&User{ID: "user-02"},
								&User{ID: "user-03"}}},
						&User{ID: "user-02",
							Followees: UserList{
								&User{ID: "user-01"},
								&User{ID: "user-04"}}},
					},
				},
			}

			for id, test := range tests {
				var actual UserList
				if err := actual.Unmarshal(test.jsonObj); err != nil {
					t.Fatal("Unexpected error: ", err)
				}

				sort.Slice(actual, func(i, j int) bool {
					return strings.Compare(actual[i].ID, actual[j].ID) == -1
				})

				if !reflect.DeepEqual(test.expected, actual) {
					t.Errorf("User list mismatch. Test case %d.\nExpected: %s\nBut got: %s", id, test.expected, actual)
				}
			}
		})
	})
}
