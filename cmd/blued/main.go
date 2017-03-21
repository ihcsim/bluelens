//go:generate goagen bootstrap -d github.com/ihcsim/bluelens/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ihcsim/bluelens/cmd/blued/app"

	log "github.com/sirupsen/logrus"
)

func main() {

	config, err := parseFlags(os.Args[1:])
	if err != nil {
		log.WithFields(log.Fields{
			"Cause": err},
		).Fatal("Error while parsing runtime configuration flags")
	}

	if err := initStore(config); err != nil {
		log.WithFields(log.Fields{
			"Cause": err},
		).Fatal("Error while initializing data store")
	}

	// Create service
	service := goa.New("bluelens")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "music" controller
	c := NewMusicController(service)
	app.MountMusicController(service, c)
	// Mount "recommendations" controller
	c2 := NewRecommendationsController(service)
	app.MountRecommendationsController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "user" controller
	c4 := NewUserController(service)
	app.MountUserController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
