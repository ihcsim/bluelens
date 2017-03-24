package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/cmd/blued/app/test"
	"github.com/ihcsim/bluelens/internal/core"
)

var (
	// store fixture used for the music controller tests
	storeFixtureMusic core.Store

	// ensure the store fixture used for the music controller tests are initilized once
	initStoreFixtureMusic sync.Once

	// helper function to retrieve music fixtures from store
	musicFixture = func(t *testing.T, id string) *core.Music {
		m, err := store().FindMusic(id)
		if err != nil {
			t.Fatalf("Unexpected error with user ID %s: %", id, err)
		}

		return m
	}
)

func TestMusicController(t *testing.T) {
	// mock the store() function
	storeFunc := store
	storeFuncMock := func() core.Store {
		initStoreFixtureMusic.Do(func() {
			var err error
			storeFixtureMusic, err = core.NewFixtureStore()
			if err != nil {
				t.Fatal("Unexpected error: ", err)
			}
		})
		return storeFixtureMusic
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
			{offset: 0, limit: -1},
			{offset: 0, limit: 0},
			{offset: 0, limit: 5},
			{offset: 0, limit: 10},
			{offset: 0, limit: 20},
			{offset: 10, limit: 20},
			{offset: 5, limit: 10},
			{offset: 10, limit: 10},
			{offset: 15, limit: 10},
			{offset: -1, limit: 0},
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
			expected := mediaTypeMusicFull(musicFixture(t, musicID))
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

	t.Run("create", func(t *testing.T) {
		fixture := &core.Music{ID: "song-00.v2", Tags: []string{"rock", "90's"}}
		payload := &app.Music{ID: fixture.ID, Tags: fixture.Tags}
		expected := mediaTypeMusicLink(fixture)
		if _, actual := test.CreateMusicCreatedLink(t, nil, nil, ctrl, payload); !reflect.DeepEqual(expected, actual) {
			t.Errorf("Created music mismatch. Expected %+v, but got %+v", expected, actual)
		}

		m := musicFixture(t, "song-00.v2")
		if !reflect.DeepEqual(fixture, m) {
			t.Errorf("Resource mismatch. Expected %s, but got %s", fixture, m)
		}

		t.Run("bad request", func(t *testing.T) {
			tests := []struct {
				payload *app.Music
			}{
				{payload: &app.Music{}},
				{payload: &app.Music{ID: ""}},
			}

			for _, tc := range tests {
				if _, err := test.CreateMusicBadRequest(t, nil, nil, ctrl, tc.payload); err == nil {
					t.Error("Expected error to occur. Should have failed with a 400 response status.")
				}
			}
		})

	})
}
