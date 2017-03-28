package main

import (
	"fmt"
	"sync"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (ctrl *UserController) Create(ctx *app.CreateUserContext) error {
	c, e := make(chan *core.User), make(chan error)
	create := func() {
		defer func() {
			close(c)
			close(e)
		}()

		// construct the user object based on the payload data
		user := &core.User{ID: ctx.Payload.ID}

		var wg sync.WaitGroup

		// find all the followees
		wg.Add(1)
		go func() {
			defer wg.Done()

			var followees core.UserList
			for _, f := range ctx.Payload.Followees {
				followee, err := store().FindUser(f.ID)
				if err != nil {
					e <- fmt.Errorf("Can't create user %q. Followee %q doesn't exist.", user.ID, f.ID)
				}

				followees = append(followees, followee)
			}
			user.Followees = followees
		}()

		// find all the music resources
		wg.Add(1)
		go func() {
			defer wg.Done()

			var history core.MusicList
			for _, h := range ctx.Payload.History {
				music, err := store().FindMusic(h.ID)
				if err != nil {
					e <- fmt.Errorf("Can't create user %q. Music resource %q doesn't exist.", user.ID, music.ID)
				}
				history = append(history, music)
			}
			user.History = history
		}()

		wg.Wait()
		updated, err := store().UpdateUser(user)
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
				return ctx.CreatedLink(mediaTypeUserLink(updated))
			}
		}
	}
}

// Follow runs the follow action.
func (ctrl *UserController) Follow(ctx *app.FollowUserContext) error {
	c, e := make(chan *core.User), make(chan error)
	follow := func() {
		defer func() {
			close(c)
			close(e)
		}()

		updated, err := store().Follow(ctx.ID, *ctx.Payload.FolloweeID)
		c <- updated
		e <- err
	}

	go invoke(ctx, follow)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return ctx.NotFound(err)
			}
		case updated := <-c:
			if updated != nil {
				return ctx.OK(mediaTypeUser(updated))
			}
		}
	}
}

// List runs the list action.
func (ctrl *UserController) List(ctx *app.ListUserContext) error {
	c, e := make(chan core.UserList), make(chan error)
	list := func() {
		defer func() {
			close(c)
			close(e)
		}()

		ul, err := store().ListUsers(ctx.Limit, ctx.Offset)
		c <- ul
		e <- err
	}

	go invoke(ctx, list)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			return err
		case userList := <-c:
			var res app.BluelensUserCollection
			for _, user := range userList {
				res = append(res, mediaTypeUser(user))
			}
			return ctx.OK(res)
		}
	}
}

// Listen runs the listen action.
func (ctrl *UserController) Listen(ctx *app.ListenUserContext) error {
	c, e := make(chan *core.User), make(chan error)
	listen := func() {
		defer func() {
			close(c)
			close(e)
		}()

		updated, err := store().Listen(ctx.ID, *ctx.Payload.MusicID)
		c <- updated
		e <- err
	}

	go invoke(ctx, listen)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return ctx.NotFound(err)
			}
		case updated := <-c:
			if updated != nil {
				return ctx.OK(mediaTypeUser(updated))
			}
		}
	}
}

// Show runs the show action.
func (ctrl *UserController) Show(ctx *app.ShowUserContext) error {
	c, e := make(chan *core.User), make(chan error)
	show := func() {
		defer func() {
			close(c)
			close(e)
		}()

		u, err := store().FindUser(ctx.ID)
		c <- u
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
		case user := <-c:
			if user != nil {
				return ctx.OKFull(mediaTypeUserFull(user))
			}
		}
	}
}
