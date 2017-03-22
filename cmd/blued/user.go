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
	updated, err := store().Follow(ctx.ID, *ctx.Payload.FolloweeID)
	if err != nil {
		return ctx.NotFound(err)
	}

	return ctx.OK(mediaTypeUser(updated))

}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	userList, err := store().ListUsers(ctx.Limit, ctx.Offset)
	if err != nil {
		return err
	}

	res := app.BluelensUserCollection{}
	for _, user := range userList {
		res = append(res, mediaTypeUser(user))
	}

	return ctx.OK(res)
}

// Listen runs the listen action.
func (c *UserController) Listen(ctx *app.ListenUserContext) error {
	updated, err := store().Listen(ctx.ID, *ctx.Payload.MusicID)
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
