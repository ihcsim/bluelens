package main

import (
	"fmt"

	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

func mediaTypeRecommendations(r *core.Recommendations) *app.BluelensRecommendations {
	musicIDs := []string{}
	musicLinks := app.BluelensMusicLinkCollection{}
	for _, m := range r.List {
		musicIDs = append(musicIDs, m.ID)
		link := &app.BluelensMusicLink{Href: href("music", m.ID)}
		musicLinks = append(musicLinks, link)
	}

	links := &app.BluelensRecommendationsLinks{
		List: musicLinks,
		User: &app.BluelensUserLink{Href: href("users", r.UserID)},
	}
	return &app.BluelensRecommendations{
		MusicID: musicIDs,
		Links:   links,
	}
}

func mediaTypeMusic(m *core.Music) *app.BluelensMusic {
	return &app.BluelensMusic{
		ID:   m.ID,
		Href: href("music", m.ID),
		Tags: m.Tags,
	}
}

func mediaTypeMusicFull(m *core.Music) *app.BluelensMusicFull {
	return &app.BluelensMusicFull{
		ID:   m.ID,
		Href: href("music", m.ID),
		Tags: m.Tags,
	}
}

func mediaTypeMusicLink(m *core.Music) *app.BluelensMusicLink {
	return &app.BluelensMusicLink{
		Href: href("music", m.ID),
	}
}

func mediaTypeUser(u *core.User) *app.BluelensUser {
	followeesLinks := app.BluelensUserLinkCollection{}
	for _, followee := range u.Followees {
		link := &app.BluelensUserLink{
			Href: href("users", followee.ID),
		}
		followeesLinks = append(followeesLinks, link)
	}

	historyLinks := app.BluelensMusicLinkCollection{}
	for _, music := range u.History {
		link := &app.BluelensMusicLink{
			Href: href("music", music.ID),
		}
		historyLinks = append(historyLinks, link)
	}

	links := &app.BluelensUserLinks{
		Followees: followeesLinks,
		History:   historyLinks,
	}

	return &app.BluelensUser{
		ID:    u.ID,
		Href:  href("users", u.ID),
		Links: links,
	}
}

func mediaTypeUserFull(u *core.User) *app.BluelensUserFull {
	followees := app.BluelensUserCollection{}
	for _, followee := range u.Followees {
		user := mediaTypeUser(followee)
		followees = append(followees, user)
	}

	history := app.BluelensMusicCollection{}
	for _, music := range u.History {
		music := mediaTypeMusic(music)
		history = append(history, music)
	}

	return &app.BluelensUserFull{
		ID:        u.ID,
		Followees: followees,
		History:   history,
		Href:      href("users", u.ID),
	}
}

func mediaTypeUserLink(u *core.User) *app.BluelensUserLink {
	return &app.BluelensUserLink{
		Href: href("users", u.ID),
	}
}

func href(kind, id string) string {
	return fmt.Sprintf("/bluelens/%s/%s", kind, id)
}
