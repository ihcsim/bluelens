package main

import (
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
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	user := &core.User{
		ID: ctx.Payload.ID,
	}

	var followees core.UserList
	for _, f := range ctx.Payload.Followees {
		followee, err := store().FindUser(f.ID)
		if err != nil {
			return err
		}

		followees = append(followees, followee)
	}
	user.Followees = followees

	var history core.MusicList
	for _, h := range ctx.Payload.History {
		music, err := store().FindMusic(h.ID)
		if err != nil {
			return err
		}
		history = append(history, music)
	}
	user.History = history

	updated, err := store().UpdateUser(user)
	if err != nil {
		return err
	}

	res := mediaTypeUserLink(updated)
	return ctx.CreatedLink(res)
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
