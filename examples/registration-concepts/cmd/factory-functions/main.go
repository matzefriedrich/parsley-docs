package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/registration-concepts/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
)

func main() {
	registry := registration.NewServiceRegistry()
	_ = registration.RegisterTransient(registry, internal.NewGreeterFactory("Hi"))

	resolver := resolving.NewResolver(registry)
	scope := resolving.NewScopedContext(context.Background())

	greeter, _ := resolving.ResolveRequiredService[internal.Greeter](scope, resolver)

	const polite = false
	greeter.SayHello("John", polite)
}
