package main

import (
	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/internal/core"
)

// RecommendationsController implements the recommendations resource.
type RecommendationsController struct {
	*goa.Controller
}

// NewRecommendationsController creates a recommendations controller.
func NewRecommendationsController(service *goa.Service) *RecommendationsController {
	return &RecommendationsController{Controller: service.NewController("RecommendationsController")}
}

// Recommend runs the recommend action.
func (c *RecommendationsController) Recommend(ctx *app.RecommendRecommendationsContext) error {
	recommendations, err := core.RecommendSort(ctx.UserID, ctx.Limit, store())
	if err != nil {
		return err
	}

	res := mediaTypeRecommendations(recommendations)
	return ctx.OK(res)
}
