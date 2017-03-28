package main

import (
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
func (ctrl *MusicController) Create(ctx *app.CreateMusicContext) error {
	c, e := make(chan *core.Music), make(chan error)
	create := func() {
		defer func() {
			close(c)
			close(e)
		}()

		music := &core.Music{
			ID:   ctx.Payload.ID,
			Tags: ctx.Payload.Tags,
		}
		updated, err := store().UpdateMusic(music)
		c <- updated
		e <- err
	}

	go invoke(ctx, create)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return err
			}
		case updated := <-c:
			if updated != nil {
				return ctx.CreatedLink(mediaTypeMusicLink(updated))
			}
		}
	}
}

// List runs the list action.
func (ctrl *MusicController) List(ctx *app.ListMusicContext) error {
	c, e := make(chan core.MusicList), make(chan error)
	list := func() {
		defer func() {
			close(c)
			close(e)
		}()

		ml, err := store().ListMusic(ctx.Limit, ctx.Offset)
		c <- ml
		e <- err
	}

	go invoke(ctx, list)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return err
			}
		case ml := <-c:
			if ml != nil {
				var res app.BluelensMusicCollection
				for _, music := range ml {
					res = append(res, mediaTypeMusic(music))
				}
				return ctx.OK(res)
			}
		}
	}
}

// Show runs the show action.
func (ctrl *MusicController) Show(ctx *app.ShowMusicContext) error {
	c, e := make(chan *core.Music), make(chan error)
	show := func() {
		defer func() {
			close(c)
			close(e)
		}()

		m, err := store().FindMusic(ctx.ID)
		c <- m
		e <- err
	}

	go invoke(ctx, show)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return ctx.NotFound(err)
			}
		case m := <-c:
			if m != nil {
				return ctx.OKFull(mediaTypeMusicFull(m))
			}
		}
	}
}
