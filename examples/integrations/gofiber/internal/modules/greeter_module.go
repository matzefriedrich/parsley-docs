package modules

import (
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal/services"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigureGreeter Configures the services.Greeter service dependencies.
func ConfigureGreeter(registry types.ServiceRegistry) error {
	registry.Register(services.NewGreeter, types.LifetimeTransient)
	return nil
}
