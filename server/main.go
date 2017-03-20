//go:generate goagen bootstrap -d github.com/ihcsim/bluelens/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ihcsim/bluelens/server/app"
	"github.com/ihcsim/bluelens/server/ctrl"
	"github.com/ihcsim/bluelens/server/store"
)

func main() {
	config, err := parseFlags(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	if err := store.Initialize(config); err != nil {
		os.Exit(1)
	}

	// Create service
	service := goa.New("bluelens")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "music" controller
	c := ctrl.NewMusicController(service)
	app.MountMusicController(service, c)
	// Mount "recommendations" controller
	c2 := ctrl.NewRecommendationsController(service)
	app.MountRecommendationsController(service, c2)
	// Mount "swagger" controller
	c3 := ctrl.NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "user" controller
	c4 := ctrl.NewUserController(service)
	app.MountUserController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
