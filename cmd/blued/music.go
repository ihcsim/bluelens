package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

// MusicController implements the music resource.
type MusicController struct {
	*goa.Controller
}

// NewMusicController creates a music controller.
func NewMusicController(service *goa.Service) *MusicController {
	return &MusicController{Controller: service.NewController("MusicController")}
}

// Create runs the create action.
func (c *MusicController) Create(ctx *app.CreateMusicContext) error {
	music := &core.Music{
		ID:   ctx.Payload.ID,
		Tags: ctx.Payload.Tags,
	}
	updated, err := store().UpdateMusic(music)
	if err != nil {
		return err
	}

	res := &app.BluelensMusicLink{
		Href: "/music/" + updated.ID,
	}
	return ctx.CreatedLink(res)
}

// List runs the list action.
func (c *MusicController) List(ctx *app.ListMusicContext) error {
	ml, err := store().ListMusic(ctx.Limit, ctx.Offset)
	if err != nil {
		return err
	}

	var res app.BluelensMusicCollection
	for _, music := range ml {
		res = append(res, mediaTypeMusic(music))
	}

	return ctx.OK(res)
}

// Show runs the show action.
func (c *MusicController) Show(ctx *app.ShowMusicContext) error {
	m, err := store().FindMusic(ctx.ID)
	if err != nil {
		return ctx.NotFound(err)
	}

	res := &app.BluelensMusicFull{
		ID:   m.ID,
		Href: fmt.Sprintf("/music/%s", m.ID),
		Tags: m.Tags,
	}
	return ctx.OKFull(res)
}