package internal

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/chi/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

type parsleyApplication struct {
	router chi.Router
}

var _ bootstrap.Application = &parsleyApplication{}

// NewApp Creates the main application service instance. This constructor function gets invoked by Parsley; add parameters for all required services.
func NewApp(router chi.Router, routeHandlers []route_handlers.RouteHandler) bootstrap.Application {
	for _, routeHandler := range routeHandlers {
		routeHandler.Register(router)
	}
	return &parsleyApplication{
		router: router,
	}
}

// Run The entrypoint for the Parsley application.
func (a *parsleyApplication) Run(_ context.Context) error {
	return http.ListenAndServe(":5504", a.router)
}
