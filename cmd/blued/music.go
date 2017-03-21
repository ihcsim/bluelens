package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
)

// MusicController implements the music resource.
type MusicController struct {
	*goa.Controller
}

// NewMusicController creates a music controller.
func NewMusicController(service *goa.Service) *MusicController {
	return &MusicController{Controller: service.NewController("MusicController")}
}

// Get runs the get action.
func (c *MusicController) Get(ctx *app.GetMusicContext) error {
	m, err := store().FindMusic(ctx.ID)
	if err != nil {
		return ctx.NotFound(err)
	}

	res := &app.BluelensMusic{
		ID:   m.ID,
		Href: fmt.Sprintf("/music/%s", m.ID),
		Tags: m.Tags,
	}
	return ctx.OK(res)
}
