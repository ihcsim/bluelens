package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/cmd/blued/app/test"
	"github.com/ihcsim/bluelens/internal/core"
)

func TestMusicController(t *testing.T) {
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
	ctrl := NewMusicController(svc)

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
			musicList, err := store().ListMusic(tc.limit, tc.offset)
			if err != nil {
				t.Error("Unexpected error with test case %d: ", id, err)
			}

			var expected app.BluelensMusicCollection
			for _, music := range musicList {
				expected = append(expected, mediaTypeMusic(music))
			}

			if _, actual := test.ListMusicOK(t, nil, nil, ctrl, tc.limit, tc.offset); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Response mismatch. Test case: %d\nExpected %s\nBut got %s", id, expected, actual)
			}
		}
	})

	t.Run("show", func(t *testing.T) {
		t.Run("found", func(t *testing.T) {
			musicID := "song-00"
			music, err := store().FindMusic(musicID)
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			expected := mediaTypeMusicFull(music)
			if _, actual := test.ShowMusicOKFull(t, nil, nil, ctrl, musicID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Music mismatch. Expected %s, but got %s", expected, actual)
			}
		})

		t.Run("not found", func(t *testing.T) {
			musicID := "example"
			if _, err := test.ShowMusicNotFound(t, nil, nil, ctrl, musicID); err == nil {
				t.Error("Expected error to occur", err)
			}
		})
	})
}
