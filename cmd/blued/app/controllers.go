// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/ihcsim/bluelens/design
// --out=$(GOPATH)/src/github.com/ihcsim/bluelens/cmd/blued
// --version=v1.1.0
//
// API "bluelens": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// MusicController is the controller interface for the Music actions.
type MusicController interface {
	goa.Muxer
	Create(*CreateMusicContext) error
	List(*ListMusicContext) error
	Show(*ShowMusicContext) error
}

// MountMusicController "mounts" a Music resource controller on the given service.
func MountMusicController(service *goa.Service, ctrl MusicController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateMusicContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Music)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("POST", "/bluelens/music", ctrl.MuxHandler("Create", h, unmarshalCreateMusicPayload))
	service.LogInfo("mount", "ctrl", "Music", "action", "Create", "route", "POST /bluelens/music", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListMusicContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("GET", "/bluelens/music", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Music", "action", "List", "route", "GET /bluelens/music", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowMusicContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("GET", "/bluelens/music/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Music", "action", "Show", "route", "GET /bluelens/music/:id", "security", "APIKey")
}

// unmarshalCreateMusicPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateMusicPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &music{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// RecommendationsController is the controller interface for the Recommendations actions.
type RecommendationsController interface {
	goa.Muxer
	Recommend(*RecommendRecommendationsContext) error
}

// MountRecommendationsController "mounts" a Recommendations resource controller on the given service.
func MountRecommendationsController(service *goa.Service, ctrl RecommendationsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRecommendRecommendationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Recommend(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("GET", "/bluelens/recommendations/:userID/:limit", ctrl.MuxHandler("Recommend", h, nil))
	service.LogInfo("mount", "ctrl", "Recommendations", "action", "Recommend", "route", "GET /bluelens/recommendations/:userID/:limit", "security", "APIKey")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/bluelens/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/bluelens/swagger.yaml", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/bluelens/swagger.json", "cmd/blued/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/bluelens/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "cmd/blued/swagger/swagger.json", "route", "GET /bluelens/swagger.json")

	h = ctrl.FileHandler("/bluelens/swagger.yaml", "cmd/blued/swagger/swagger.yaml")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/bluelens/swagger.yaml", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "cmd/blued/swagger/swagger.yaml", "route", "GET /bluelens/swagger.yaml")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Follow(*FollowUserContext) error
	List(*ListUserContext) error
	Listen(*ListenUserContext) error
	Show(*ShowUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*User)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("POST", "/bluelens/user", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /bluelens/user", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewFollowUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*FollowUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Follow(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("POST", "/bluelens/user/:id/follows/:followeeID", ctrl.MuxHandler("Follow", h, unmarshalFollowUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Follow", "route", "POST /bluelens/user/:id/follows/:followeeID", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("GET", "/bluelens/user", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /bluelens/user", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListenUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListenUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Listen(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("POST", "/bluelens/user/:id/listen/:musicID", ctrl.MuxHandler("Listen", h, unmarshalListenUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Listen", "route", "POST /bluelens/user/:id/listen/:musicID", "security", "APIKey")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("APIKey", h)
	service.Mux.Handle("GET", "/bluelens/user/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Show", "route", "GET /bluelens/user/:id", "security", "APIKey")
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &user{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalFollowUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalFollowUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &followUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalListenUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalListenUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listenUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}