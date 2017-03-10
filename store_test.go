package core

import (
	"reflect"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()

	t.Run("manage users", func(t *testing.T) {
		user := &User{ID: "user-01"}
		store.AddUser(user)
		actual, err := store.FindUser("user-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if !reflect.DeepEqual(user, actual) {
			t.Errorf("User mismatch. Expected %+v, but got %+v", user, actual)
		}

		_, err = store.FindUser("non-existent")
		if err == nil {
			t.Error("Expected error didn't occur. Should have received an EntityNotFound error.")
		}
	})

	t.Run("manage music", func(t *testing.T) {
		musicList := MusicList{
			&Music{ID: "song-01", Tags: []string{"rock", "top-10"}},
			&Music{ID: "song-02", Tags: []string{"instrument", "rock"}},
			&Music{ID: "song-03", Tags: []string{"pop"}},
		}
		if err := store.LoadMusic(musicList); err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		actualList, err := store.ListMusic(0)
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if !reflect.DeepEqual(musicList, actualList) {
			t.Errorf("Music list mismatch. Expected %v, but got %v", musicList, actualList)
		}

		actualResource, err := store.FindMusic("song-01")
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}

		if !reflect.DeepEqual(musicList[0], actualResource) {
			t.Errorf("Music resource mismatch. Expected: %s, But got: %s", musicList[0], actualResource)
		}

		t.Run("by tags", func(t *testing.T) {
			musicListByTags := map[string]MusicList{
				"rock":       MusicList{musicList[0], musicList[1]},
				"top-10":     MusicList{musicList[0]},
				"instrument": MusicList{musicList[1]},
				"pop":        MusicList{musicList[2]},
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
	})
}
