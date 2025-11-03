package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/resolving-services/internal"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()
	features.RegisterLazy[internal.Greeter](registry, internal.NewGreeterFactory("Hi"), types.LifetimeTransient)

	resolver := resolving.NewResolver(registry)
	ctx := resolving.NewScopedContext(context.Background())
	lazy, _ := resolving.ResolveRequiredService[features.Lazy[internal.Greeter]](ctx, resolver)

	greeter := lazy.Value()
	greeter.SayHello("John", true)
}
