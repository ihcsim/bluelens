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
	// UserController_Follow: start_implement

	// Put your logic here

	// UserController_Follow: end_implement
	return nil
}

// Listen runs the listen action.
func (c *UserController) Listen(ctx *app.ListenUserContext) error {
	// UserController_Listen: start_implement

	// Put your logic here

	// UserController_Listen: end_implement
	return nil
}
