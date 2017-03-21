package core

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ihcsim/bluelens/internal/core/json"
)

func TestUser(t *testing.T) {
	t.Run("IsNew", func(t *testing.T) {
		user := &User{ID: "user-01"}
		if !user.IsNew() {
			t.Error("Expected user with no followees and history to be labeled as 'new'")
		}

		user.AddFollowees(&User{ID: "user-02"})
		if user.IsNew() {
			t.Error("Expected user with followees and history to not be labeled as 'new'")
		}

		user.AddHistory(&Music{ID: "song-01"})
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

		user.AddFollowees(followee)
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

		user.AddHistory(history)
		if !user.HasHistory(history.ID) {
			t.Error("Expected user to have a music history")
		}
	})

	t.Run("AddFollowees", func(t *testing.T) {
		user := &User{ID: "user-01"}
		tests := []struct {
			followees UserList
		}{
			{followees: UserList{&User{ID: "user-02"}}},
			{followees: UserList{&User{ID: "user-02"}, &User{ID: "user-03"}}},
		}

		for id, test := range tests {
			for _, f := range test.followees {
				if err := user.AddFollowees(f); err != nil {
					t.Error("Unexpected error occurred. Test case ", id)
				}

				if !user.HasFollowee(f.ID) {
					t.Error("Expected followees to be added. Test case ", id)
				}
			}
		}

		t.Run("malformed", func(t *testing.T) {
			user := &User{ID: "user-01"}
			followees := []*User{nil, &User{}}
			for _, f := range followees {
				if err := user.AddFollowees(f); err != ErrMalformedEntity {
					t.Error("Expected 'malformed entity' error to occur")
				}
			}
		})

		t.Run("duplicates", func(t *testing.T) {
			user := &User{ID: "user-01"}
			followees := UserList{&User{ID: "user-02"}, &User{ID: "user-02"}}
			for id, f := range followees {
				if err := user.AddFollowees(f); err != nil {
					t.Error("Unexpected error occurred. Test case ", id)
				}

				if !user.HasFollowee(f.ID) {
					t.Error("Expected followees to be added. Test case ", id)
				}
			}

			var found bool
			for _, f := range user.Followees {
				if f.ID == "user-02" {
					if !found {
						found = true
					} else {
						t.Error("Expect no duplicates in the list")
					}
				}
			}
		})
	})

	t.Run("AddHistory", func(t *testing.T) {
		user := &User{ID: "user-01"}
		tests := []struct {
			history MusicList
		}{
			{history: MusicList{&Music{ID: "song-02"}}},
			{history: MusicList{&Music{ID: "song-02"}, &Music{ID: "song-03"}}},
		}

		for id, test := range tests {
			for _, h := range test.history {
				if err := user.AddHistory(h); err != nil {
					t.Error("Unexpected error occurred. Test case ", id)
				}

				if !user.HasHistory(h.ID) {
					t.Error("Expected followees to be added. Test case ", id)
				}
			}
		}

		t.Run("malformed", func(t *testing.T) {
			user := &User{ID: "user-01"}
			history := []*Music{nil, &Music{}}
			for _, h := range history {
				if err := user.AddHistory(h); err != ErrMalformedEntity {
					t.Error("Expected 'malformed entity' error to occur")
				}
			}
		})

		t.Run("duplicates", func(t *testing.T) {
			user := &User{ID: "user-01"}
			history := MusicList{&Music{ID: "song-02"}, &Music{ID: "music-02"}}
			for id, h := range history {
				if err := user.AddHistory(h); err != nil {
					t.Error("Unexpected error occurred. Test case ", id)
				}

				if !user.HasHistory(h.ID) {
					t.Error("Expected followees to be added. Test case ", id)
				}
			}

			var found bool
			for _, h := range user.History {
				if h.ID == "song-02" {
					if !found {
						found = true
					} else {
						t.Error("Expect no duplicates in the list")
					}
				}
			}
		})
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
