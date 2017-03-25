package apikey

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/ihcsim/bluelens/cmd/blued/app"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

// New returns a middleware that performs authentication based on the provided handler scheme.
func New(apiKey string) goa.Middleware {
	mv, _ := goa.NewMiddleware(func(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
		scheme := app.NewAPIKeySecurity()
		key := req.Header.Get(scheme.Name)

		key = strings.TrimPrefix(key, "Bearer")
		key = strings.TrimSpace(key)
		if len(key) == 0 || key != apiKey {
			return ErrUnauthorized("Invalid API key")
		}

		return nil
	})
	return mv
}
