package main

import (
	"context"
	"github.com/matzefriedrich/parsley-docs/examples/registration-concepts/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()
	registry.RegisterModule(greeterModule)

	resolver := resolving.NewResolver(registry)
	greeter, _ := resolving.ResolveRequiredService[internal.Greeter](resolver, resolving.NewScopedContext(context.Background()))
	greeter.SayHello("John", false)
}

func greeterModule(registry types.ServiceRegistry) error {
	registry.Register(internal.NewGreeterFactory("Hi"), types.LifetimeTransient)
	return nil
}
