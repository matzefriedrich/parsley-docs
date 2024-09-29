package modules

import (
	"github.com/matzefriedrich/parsley-docs/examples/integrations/cobra/internal/services"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ServicesModule registers a transient service for creating new instances of Greeter in the provided service registry.
func ServicesModule(registry types.ServiceRegistry) error {
	_ = registration.RegisterTransient(registry, services.NewGreeter)
	return nil
}
