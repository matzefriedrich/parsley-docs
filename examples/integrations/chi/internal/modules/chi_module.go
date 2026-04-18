package modules

import (
	"github.com/go-chi/chi/v5"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureChi

// ConfigureChi Configures all services required by the chi app.
func ConfigureChi(registry types.ServiceRegistry) error {
	_ = registry.Register(newChi, types.LifetimeSingleton)
	_ = registry.RegisterModule(RegisterRouteHandlers)
	return nil
}

func newChi() chi.Router {
	return chi.NewRouter()
}
