package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

func TestMediaTypeRecommendations(t *testing.T) {
	recommendations := &core.Recommendations{
		List:   core.MusicList{&core.Music{ID: "song-00"}, &core.Music{ID: "song-01"}},
		UserID: "user-00",
	}

	expected := &app.BluelensRecommendations{
		MusicID: []string{"song-00", "song-01"},
		Links: &app.BluelensRecommendationsLinks{
			List: app.BluelensMusicLinkCollection{
				&app.BluelensMusicLink{Href: "/music/song-00"},
				&app.BluelensMusicLink{Href: "/music/song-01"},
			},
			User: &app.BluelensUserLink{Href: "/users/user-00"},
		},
	}
	actual := mediaTypeRecommendations(recommendations)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Failed to convert to the correct media type. Expected %+v, but got %+v", expected, actual)
	}
}

func TestMediaTypeMusic(t *testing.T) {
	music := &core.Music{
		ID:   "song-00",
		Tags: []string{"rock", "80's"},
	}

	t.Run("view=default", func(t *testing.T) {
		expected := &app.BluelensMusic{
			ID:   music.ID,
			Href: fmt.Sprintf("/music/%s", music.ID),
			Tags: music.Tags,
		}
		if actual := mediaTypeMusic(music); !reflect.DeepEqual(expected, actual) {
			t.Errorf("media type mismatch. Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("view=full", func(t *testing.T) {
		expected := &app.BluelensMusicFull{
			ID:   music.ID,
			Href: fmt.Sprintf("/music/%s", music.ID),
			Tags: music.Tags,
		}
		if actual := mediaTypeMusicFull(music); !reflect.DeepEqual(expected, actual) {
			t.Errorf("Media type mismatch. Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("view=link", func(t *testing.T) {
		expected := &app.BluelensMusicLink{
			Href: "/music/song-00",
		}

		if actual := mediaTypeMusicLink(music); !reflect.DeepEqual(expected, actual) {
			t.Errorf("media type mismatch. Expected %v, but got %v\n", expected, actual)
		}
	})
}

func TestMediaTypeUser(t *testing.T) {
	user := &core.User{
		ID: "user-00",
		Followees: []*core.User{
			&core.User{ID: "user-01"},
			&core.User{ID: "user-02"},
		},
		History: core.MusicList{
			&core.Music{ID: "song-01"},
			&core.Music{ID: "song-02"},
		},
	}

	t.Run("view=default", func(t *testing.T) {
		followeesLinks := app.BluelensUserLinkCollection{
			&app.BluelensUserLink{Href: "/users/user-01"},
			&app.BluelensUserLink{Href: "/users/user-02"},
		}
		historyLinks := app.BluelensMusicLinkCollection{
			&app.BluelensMusicLink{Href: "/music/song-01"},
			&app.BluelensMusicLink{Href: "/music/song-02"},
		}
		links := &app.BluelensUserLinks{
			Followees: followeesLinks,
			History:   historyLinks,
		}
		expected := &app.BluelensUser{
			ID:    user.ID,
			Href:  fmt.Sprintf("/users/%s", user.ID),
			Links: links,
		}
		if actual := mediaTypeUser(user); !reflect.DeepEqual(expected, actual) {
			t.Errorf("Media type mismatch. Expected %+v, but got %+v", expected, actual)
		}
	})

	t.Run("view=full", func(t *testing.T) {
		followees := app.BluelensUserCollection{
			&app.BluelensUser{ID: "user-01", Href: "/users/user-01",
				Links: &app.BluelensUserLinks{
					Followees: app.BluelensUserLinkCollection{},
					History:   app.BluelensMusicLinkCollection{},
				}},
			&app.BluelensUser{ID: "user-02", Href: "/users/user-02",
				Links: &app.BluelensUserLinks{
					Followees: app.BluelensUserLinkCollection{},
					History:   app.BluelensMusicLinkCollection{},
				}},
		}
		history := app.BluelensMusicCollection{
			&app.BluelensMusic{ID: "song-01", Href: "/music/song-01"},
			&app.BluelensMusic{ID: "song-02", Href: "/music/song-02"},
		}

		expected := &app.BluelensUserFull{
			ID:        user.ID,
			Href:      fmt.Sprintf("/users/%s", user.ID),
			Followees: followees,
			History:   history,
		}

		if actual := mediaTypeUserFull(user); !reflect.DeepEqual(expected, actual) {
			t.Errorf("Media type mismatch.\nExpected %s\nBut got %s", expected, actual)
		}
	})

	t.Run("view=link", func(t *testing.T) {
		expected := &app.BluelensUserLink{
			Href: "/users/user-00",
		}
		if actual := mediaTypeUserLink(user); !reflect.DeepEqual(expected, actual) {
			t.Errorf("Media type mismatched. Expected %v, but got %v", expected, actual)
		}
	})
}
