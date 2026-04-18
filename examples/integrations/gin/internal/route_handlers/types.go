package route_handlers

import "github.com/gin-gonic/gin"

type RouteHandler interface {
	Register(engine *gin.Engine)
}
