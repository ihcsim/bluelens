package ctrl

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app/test"
	"github.com/ihcsim/bluelens/server/store"
)

func TestMusicController(t *testing.T) {
	// mock the store() function
	storeFunc := store.Instance
	storeFuncMock := func() core.Store {
		s, err := core.NewFixtureStore()
		if err != nil {
			t.Fatal("Unexpected error: ", err)
		}
		return s
	}
	store.Instance = storeFuncMock
	defer func() {
		store.Instance = storeFunc
	}()

	svc := goa.New("goatest")
	ctrl := NewMusicController(svc)

	t.Run("get", func(t *testing.T) {
		t.Run("found", func(t *testing.T) {
			musicID := "song-00"
			music, err := store.Instance().FindMusic(musicID)
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}

			expected := mediaTypeMusic(music)
			if _, actual := test.GetMusicOK(t, nil, nil, ctrl, musicID); !reflect.DeepEqual(expected, actual) {
				t.Errorf("Music mismatch. Expected %s, but got %s", expected, actual)
			}
		})

		t.Run("not found", func(t *testing.T) {
			musicID := "example"
			if _, err := test.GetMusicNotFound(t, nil, nil, ctrl, musicID); err == nil {
				t.Error("Expected error to occur", err)
			}
		})
	})
}
