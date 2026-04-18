package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureGin

// ConfigureGin Configures all services required by the Gin app.
func ConfigureGin(registry types.ServiceRegistry) error {
	_ = registry.Register(newGin, types.LifetimeSingleton)
	_ = registry.RegisterModule(RegisterRouteHandlers)
	return nil
}

func newGin() *gin.Engine {
	return gin.Default()
}
