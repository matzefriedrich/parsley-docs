package route_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal/services"
)

type greeterRouteHandler struct {
	greeter services.Greeter
}

// Register Registers all greeterRouteHandler route handlers.
func (h *greeterRouteHandler) Register(engine *gin.Engine) {
	engine.GET("/say-hello", h.HandleSayHelloRequest)
}

// HandleSayHelloRequest Handles "GET /say-hello" requests.
func (h *greeterRouteHandler) HandleSayHelloRequest(ctx *gin.Context) {
	name := ctx.Query("name")
	msg := h.greeter.SayHello(name, true)
	ctx.String(http.StatusOK, msg)
}

var _ RouteHandler = &greeterRouteHandler{}

func NewGreeterRouteHandler(greeter services.Greeter) RouteHandler {
	return &greeterRouteHandler{
		greeter: greeter,
	}
}
