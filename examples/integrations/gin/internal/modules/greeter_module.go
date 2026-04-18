package modules

import (
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal/services"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigureGreeter Configures the services.Greeter service dependencies.
func ConfigureGreeter(registry types.ServiceRegistry) error {
	_ = registry.Register(services.NewGreeter, types.LifetimeTransient)
	return nil
}
