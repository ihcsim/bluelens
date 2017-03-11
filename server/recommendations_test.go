package main

import (
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	core "github.com/ihcsim/bluelens"
	"github.com/ihcsim/bluelens/server/app"
	"github.com/ihcsim/bluelens/server/app/test"
)

const maxCount = 20

func TestRecommendationsController(t *testing.T) {
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
	fixtureStore := store().(*core.FixtureStore)

	// retrieve the expected recommendations from the fixture store
	recommendations, err := fixtureStore.Recommendations(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	// invoke the recommendations controller
	service := goa.New("goatest")
	ctrl := NewRecommendationsController(service)

	users, err := fixtureStore.ListUsers(maxCount)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	for _, user := range users {
		_, actual := test.RecommendRecommendationsOK(t, nil, nil, ctrl, user.ID, maxCount)
		expected := recommendationsMediaType(recommendations[user.ID])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Recommendations response mismatch. Expected:\n%+v\nBut got:\n%+v", expected, actual)
		}
	}
}

func TestRecommendationsMediaType(t *testing.T) {
	recommendations := &core.Recommendations{
		List:   core.MusicList{&core.Music{ID: "song-00"}, &core.Music{ID: "song-01"}},
		UserID: "user-00",
	}

	expected := &app.BluelensRecommendations{
		MusicID: []string{"song-00", "song-01"},
		Links: &app.BluelensRecommendationsLinks{
			List: app.BluelensMusicLinkCollection{
				&app.BluelensMusicLink{Href: "/music/song-00", ID: "song-00"},
				&app.BluelensMusicLink{Href: "/music/song-01", ID: "song-01"},
			},
			User: &app.BluelensUserLink{
				Href: "/users/user-00",
				ID:   "user-00",
			},
		},
	}
	actual := recommendationsMediaType(recommendations)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Failed to convert to the correct media type. Expected %+v, but got %+v", expected, actual)
	}
}
