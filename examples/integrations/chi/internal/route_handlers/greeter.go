package route_handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/chi/internal/services"
)

type greeterRouteHandler struct {
	greeter services.Greeter
}

// Register Registers all greeterRouteHandler route handlers.
func (h *greeterRouteHandler) Register(router chi.Router) {
	router.Get("/say-hello", h.HandleSayHelloRequest)
}

// HandleSayHelloRequest Handles "GET /say-hello" requests.
func (h *greeterRouteHandler) HandleSayHelloRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := h.greeter.SayHello(name, true)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(msg))
}

var _ RouteHandler = &greeterRouteHandler{}

func NewGreeterRouteHandler(greeter services.Greeter) RouteHandler {
	return &greeterRouteHandler{
		greeter: greeter,
	}
}
