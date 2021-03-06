// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/ihcsim/bluelens/design
// --out=$(GOPATH)/src/github.com/ihcsim/bluelens/cmd/blued
// --version=v1.1.0
//
// API "bluelens": Application Security
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseAPIKeyMiddleware mounts the APIKey auth middleware onto the service.
func UseAPIKeyMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("APIKey"), middleware)
}

// NewAPIKeySecurity creates a APIKey security definition.
func NewAPIKeySecurity() *goa.APIKeySecurity {
	def := goa.APIKeySecurity{
		In:   goa.LocHeader,
		Name: "Authorization",
	}
	def.Description = "API key"
	return &def
}

// UseBasicAuthMiddleware mounts the BasicAuth auth middleware onto the service.
func UseBasicAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("BasicAuth"), middleware)
}

// NewBasicAuthSecurity creates a BasicAuth security definition.
func NewBasicAuthSecurity() *goa.BasicAuthSecurity {
	def := goa.BasicAuthSecurity{}
	def.Description = "Basic Auth"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}
