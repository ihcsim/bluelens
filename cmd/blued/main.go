//go:generate goagen bootstrap -d github.com/ihcsim/bluelens/design

package main

import (
	"flag"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/basicauth"
	"github.com/ihcsim/bluelens/cmd/blued/app"
	"github.com/ihcsim/bluelens/cmd/blued/middleware/security/apikey"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := parseFlags(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}

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

	// mount security middleware
	app.UseBasicAuthMiddleware(service, basicauth.New(config.user, config.password))
	app.UseAPIKeyMiddleware(service, apikey.New(config.apiKey))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(middleware.Timeout(config.timeout))

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
	if err := service.ListenAndServeTLS(":443", config.certFile, config.keyFile); err != nil {
		service.LogError("startup", "err", err)
	}
}
