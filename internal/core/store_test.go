package core

import (
	"reflect"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := FixtureStore{
		&InMemoryStore{
			musicList:      MusicList{},
			musicMap:       make(map[string]*Music),
			musicMapByTags: make(map[string]MusicList),

			userList: UserList{},
			userMap:  make(map[string]*User),
		},
	}

	t.Run("users", func(t *testing.T) {
		t.Run("load", func(t *testing.T) {
			users := UserList{
				&User{ID: "user-00"},
				&User{ID: "user-01"},
				&User{ID: "user-02"},
				&User{ID: "user-03"},
				&User{ID: "user-04"},
				&User{ID: "user-05"},
				&User{ID: "user-06"},
				&User{ID: "user-07"},
				&User{ID: "user-08"},
				&User{ID: "user-09"},
			}

			if err := store.LoadUsers(users); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			t.Run("find", func(t *testing.T) {
				for _, user := range users {
					actual, err := store.FindUser(user.ID)
					if err != nil {
						t.Fatal("Unexpected error: ", err)
					}

					if !reflect.DeepEqual(user, actual) {
						t.Errorf("User mismatch. Expected %+v, but got %+v", user, actual)
					}
				}

				_, err := store.FindUser("non-existent")
				if err == nil {
					t.Error("Expected error didn't occur. Should have received an EntityNotFound error.")
				}
			})

			t.Run("list", func(t *testing.T) {
				tests := []struct {
					limit  int
					offset int
				}{
					{limit: -1, offset: 0},
					{limit: 0, offset: 0},
					{limit: 5, offset: 0},
					{limit: 10, offset: 0},
					{limit: 20, offset: 0},
				}

				for id, tc := range tests {
					actual, err := store.ListUsers(tc.limit)
					if err != nil {
						t.Fatal("Unexpected error with test case %d: ", id, err)
					}

					var expected UserList
					if tc.limit > len(users) || tc.limit <= 0 {
						expected = users
					} else {
						expected = users[:tc.limit]
					}

					if !reflect.DeepEqual(expected, actual) {
						t.Errorf("User list mismatch. Test case: %d\nExpected: %v\nBut got: %v", id, users, actual)
					}
				}
			})
		})
	})

	t.Run("music", func(t *testing.T) {
		t.Run("load", func(t *testing.T) {
			musicList := MusicList{
				&Music{ID: "song-01", Tags: []string{"rock", "top-10"}},
				&Music{ID: "song-02", Tags: []string{"instrument", "rock"}},
				&Music{ID: "song-03", Tags: []string{"pop"}},
				&Music{ID: "song-04", Tags: []string{"rock", "top-10"}},
				&Music{ID: "song-05", Tags: []string{"instrument", "rock"}},
				&Music{ID: "song-06", Tags: []string{"pop"}},
				&Music{ID: "song-07", Tags: []string{"rock", "top-10"}},
				&Music{ID: "song-08", Tags: []string{"instrument", "rock"}},
				&Music{ID: "song-09", Tags: []string{"pop"}},
			}
			if err := store.LoadMusic(musicList); err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			t.Run("find", func(t *testing.T) {
				for _, music := range musicList {
					actual, err := store.FindMusic(music.ID)
					if err != nil {
						t.Fatal("Unexpected error: ", err)
					}

					if !reflect.DeepEqual(music, actual) {
						t.Errorf("Music resource mismatch. Expected: %s, But got: %s", musicList[0], actual)
					}
				}
			})

			t.Run("find by tags", func(t *testing.T) {
				musicListByTags := map[string]MusicList{
					"rock":       MusicList{musicList[0], musicList[1], musicList[3], musicList[4], musicList[6], musicList[7]},
					"top-10":     MusicList{musicList[0], musicList[3], musicList[6]},
					"instrument": MusicList{musicList[1], musicList[4], musicList[7]},
					"pop":        MusicList{musicList[2], musicList[5], musicList[8]},
				}

				for tag, expected := range musicListByTags {
					actual, err := store.FindMusicByTags(tag)
					if err != nil {
						t.Error("Unexpected error when looking up music list for tag ", tag)
					}
					if !reflect.DeepEqual(expected, actual) {
						t.Errorf("Music list for tag %q mismatch. Expected:\n%v\nBut got:\n%v", tag, expected, actual)
					}
				}
			})

			t.Run("list", func(t *testing.T) {
				tests := []struct {
					limit  int
					offset int
				}{
					{limit: -1, offset: 0},
					{limit: 0, offset: 0},
					{limit: 5, offset: 0},
					{limit: 10, offset: 0},
					{limit: 20, offset: 0},
				}

				for id, tc := range tests {
					actual, err := store.ListMusic(tc.limit)
					if err != nil {
						t.Fatal("Unexpected error at test case %d: ", id, err)
					}

					var expected MusicList
					if tc.limit > len(musicList) || tc.limit <= 0 {
						expected = musicList
					} else {
						expected = musicList[:tc.limit]
					}

					if !reflect.DeepEqual(expected, actual) {
						t.Errorf("Music list mismatch. Test case %d.\nExpected: %v\nBut got: %v", id, musicList, actual)
					}
				}
			})
		})
	})
}
