package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureFiber

// ConfigureFiber Configures all services required by the Fiber app.
func ConfigureFiber(registry types.ServiceRegistry) error {
	registration.RegisterInstance(registry, fiber.Config{
		AppName:   "parsley-fiber-demo",
		Immutable: true,
	})
	registry.Register(newFiber, types.LifetimeSingleton)
	registry.RegisterModule(RegisterRouteHandlers)
	return nil
}

func newFiber(config fiber.Config) *fiber.App {
	return fiber.New(config)
}
