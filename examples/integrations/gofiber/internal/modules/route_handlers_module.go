package modules

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// RegisterRouteHandlers Registers all route handlers of the GoFiber app.
func RegisterRouteHandlers(ctx context.Context, registry types.ServiceRegistry) error {
	features.RegisterList[route_handlers.RouteHandler](ctx, registry)
	registration.RegisterTransient(registry, route_handlers.NewGreeterRouteHandler)
	return nil
}
