package main

import (
	"context"
	"fmt"
	"github.com/matzefriedrich/parsley-docs/examples/resolving-services/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()

	registry.Register(func() internal.Greeter {
		factory := internal.NewGreeterFactory("Hi")
		return factory()
	}, types.LifetimeTransient)

	resolver := resolving.NewResolver(registry)
	ctx := resolving.NewScopedContext(context.Background())

	instance, _ := resolving.Activate[*ouchie](resolver, ctx, func(greeter internal.Greeter) *ouchie {
		return &ouchie{
			msg: greeter.Ouch(),
		}
	})

	fmt.Println(instance.msg)
}

type ouchie struct {
	msg string
}
