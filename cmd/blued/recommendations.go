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
func (ctrl *RecommendationsController) Recommend(ctx *app.RecommendRecommendationsContext) error {
	c, e := make(chan *core.Recommendations), make(chan error)
	recommend := func() {
		defer func() {
			close(c)
			close(e)
		}()

		recommendations, err := core.RecommendSort(ctx.UserID, ctx.Limit, store())
		c <- recommendations
		e <- err
	}

	go invoke(ctx, recommend)

	for {
		select {
		case <-ctx.Done():
			return ctx.InternalServerError(ctx.Err())
		case err := <-e:
			if err != nil {
				return ctx.NotFound(err)
			}
		case recommendations := <-c:
			if recommendations != nil {
				return ctx.OK(mediaTypeRecommendations(recommendations))
			}
		}
	}
}
