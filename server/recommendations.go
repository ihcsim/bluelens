package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens"
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

// Recommend runs the recommend action.
func (c *RecommendationsController) Recommend(ctx *app.RecommendRecommendationsContext) error {
	recommendations, err := core.RecommendSort(ctx.UserID, ctx.MaxCount, store())
	if err != nil {
		return err
	}

	res := recommendationsMediaType(recommendations)
	return ctx.OK(res)
}

func recommendationsMediaType(r *core.Recommendations) *app.BluelensRecommendations {
	musicIDs := []string{}
	musicLinks := app.BluelensMusicLinkCollection{}
	for _, m := range r.List {
		musicIDs = append(musicIDs, m.ID)
		link := &app.BluelensMusicLink{Href: fmt.Sprintf("/music/%s", m.ID), ID: m.ID}
		musicLinks = append(musicLinks, link)
	}

	links := &app.BluelensRecommendationsLinks{
		List: musicLinks,
		User: &app.BluelensUserLink{Href: fmt.Sprintf("/users/%s", r.UserID), ID: r.UserID},
	}
	return &app.BluelensRecommendations{
		MusicID: musicIDs,
		Links:   links,
	}
}
