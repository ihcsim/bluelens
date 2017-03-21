package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
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
