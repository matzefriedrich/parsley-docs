package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/quick-start/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
)

func main() {

	registry := registration.NewServiceRegistry()
	_ = registration.RegisterTransient(registry, internal.NewGreeter)

	resolver := resolving.NewResolver(registry)
	scope := resolving.NewScopedContext(context.Background())

	greeter, _ := resolving.ResolveRequiredService[internal.Greeter](resolver, scope)

	const polite = true
	greeter.SayHello("John", polite)
}
