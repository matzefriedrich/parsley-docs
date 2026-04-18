package modules

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// RegisterRouteHandlers Registers all route handlers of the Gin app.
func RegisterRouteHandlers(registry types.ServiceRegistry) error {
	_ = features.RegisterList[route_handlers.RouteHandler](context.Background(), registry)
	_ = registration.RegisterTransient(registry, route_handlers.NewGreeterRouteHandler)
	return nil
}
