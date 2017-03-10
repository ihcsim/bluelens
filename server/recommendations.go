package main

import (
	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/server/app"
)

// RecommendationsController implements the recommendations resource.
type RecommendationsController struct {
	*goa.Controller
}

// NewRecommendationsController creates a recommendations controller.
func NewRecommendationsController(service *goa.Service) *RecommendationsController {
	return &RecommendationsController{Controller: service.NewController("RecommendationsController")}
}

// List runs the list action.
func (c *RecommendationsController) List(ctx *app.ListRecommendationsContext) error {
	// RecommendationsController_List: start_implement

	// Put your logic here

	// RecommendationsController_List: end_implement
	res := &app.BluelensRecommendations{}
	return ctx.OK(res)
}
