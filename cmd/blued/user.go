package main

import (
	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
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
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	// UserController_Create: end_implement
	return nil
}

// Follow runs the follow action.
func (c *UserController) Follow(ctx *app.FollowUserContext) error {
	user, err := store().FindUser(ctx.ID)
	if err != nil {
		return ctx.NotFound(err)
	}

	// don't follow self and don't add an existing followee
	if ctx.ID == ctx.FolloweeID || user.HasFollowee(ctx.FolloweeID) {
		return ctx.OK(mediaTypeUser(user))
	}

	updated, err := store().Follow(ctx.ID, ctx.FolloweeID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(updated))

}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here

	// UserController_List: end_implement
	res := app.BluelensUserCollection{}
	return ctx.OK(res)
}

// Listen runs the listen action.
func (c *UserController) Listen(ctx *app.ListenUserContext) error {
	user, err := store().FindUser(ctx.ID)
	if err != nil {
		return ctx.NotFound(err)
	}

	// skip if already part of the user's history
	if user.HasHistory(ctx.MusicID) {
		return ctx.OK(mediaTypeUser(user))
	}

	updated, err := store().Listen(ctx.ID, ctx.MusicID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(updated))
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	user, err := store().FindUser(ctx.ID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OKFull(mediaTypeUserFull(user))
}
