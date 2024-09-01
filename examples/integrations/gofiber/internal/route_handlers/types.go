package route_handlers

import "github.com/gofiber/fiber/v2"

type RouteHandler interface {
	Register(app *fiber.App)
}
