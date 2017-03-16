package main

import (
	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/server/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Follow runs the follow action.
func (c *UserController) Follow(ctx *app.FollowUserContext) error {
	user, err := store().FindUser(ctx.UserID)
	if err != nil {
		return ctx.NotFound(err)
	}

	// don't follow self and don't add an existing followee
	if ctx.UserID == ctx.FolloweeID || user.HasFollowee(ctx.FolloweeID) {
		return ctx.OK(mediaTypeUser(user))
	}

	updated, err := store().Follow(ctx.UserID, ctx.FolloweeID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(updated))
}

// Get runs the get action.
func (c *UserController) Get(ctx *app.GetUserContext) error {
	user, err := store().FindUser(ctx.UserID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(user))
}

// Listen runs the listen action.
func (c *UserController) Listen(ctx *app.ListenUserContext) error {
	user, err := store().FindUser(ctx.UserID)
	if err != nil {
		return ctx.NotFound(err)
	}

	// skip if already part of the user's history
	if user.HasHistory(ctx.MusicID) {
		return ctx.OK(mediaTypeUser(user))
	}

	updated, err := store().Listen(ctx.UserID, ctx.MusicID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(updated))
}
