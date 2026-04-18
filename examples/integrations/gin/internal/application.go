package internal

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

type parsleyApplication struct {
	engine *gin.Engine
}

var _ bootstrap.Application = &parsleyApplication{}

// NewApp Creates the main application service instance. This constructor function gets invoked by Parsley; add parameters for all required services.
func NewApp(engine *gin.Engine, routeHandlers []route_handlers.RouteHandler) bootstrap.Application {
	for _, routeHandler := range routeHandlers {
		routeHandler.Register(engine)
	}
	return &parsleyApplication{
		engine: engine,
	}
}

// Run The entrypoint for the Parsley application.
func (a *parsleyApplication) Run(_ context.Context) error {
	return a.engine.Run(":5503")
}
