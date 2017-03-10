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

// Follows runs the follows action.
func (c *UserController) Follows(ctx *app.FollowsUserContext) error {
	// UserController_Follows: start_implement

	// Put your logic here

	// UserController_Follows: end_implement
	return nil
}

// Listens runs the listens action.
func (c *UserController) Listens(ctx *app.ListensUserContext) error {
	// UserController_Listens: start_implement

	// Put your logic here

	// UserController_Listens: end_implement
	return nil
}
