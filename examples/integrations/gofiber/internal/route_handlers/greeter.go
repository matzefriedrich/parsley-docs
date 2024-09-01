package route_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal/services"
)

type greeterRouteHandler struct {
	greeter services.Greeter
}

// Register Registers all greeterRouteHandler route handlers.
func (h *greeterRouteHandler) Register(app *fiber.App) {
	app.Get("/say-hello", h.HandleSayHelloRequest)
}

// HandleSayHelloRequest Handles "GET /say-hello" requests.
func (h *greeterRouteHandler) HandleSayHelloRequest(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	msg := h.greeter.SayHello(name, true)
	ctx.Send([]byte(msg))
	return ctx.SendStatus(fiber.StatusOK)
}

var _ RouteHandler = &greeterRouteHandler{}

func NewGreeterRouteHandler(greeter services.Greeter) RouteHandler {
	return &greeterRouteHandler{
		greeter: greeter,
	}
}
