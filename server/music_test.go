package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
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

	musicID := "song-00"
	music, err := store().FindMusic(musicID)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	expected := mediaTypeMusic(music)
	if _, actual := test.GetMusicOK(t, nil, nil, ctrl, musicID); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Music mismatch. Expected %s, but got %s", expected, actual)
	}
}
