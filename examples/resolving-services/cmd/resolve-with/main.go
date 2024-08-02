package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()
	registry.Register(newClient, types.LifetimeTransient)

	resolver := resolving.NewResolver(registry)
	ctx := resolving.NewScopedContext(context.Background())

	resolveServiceType := types.MakeServiceType[*client]()
	transportInstance := resolving.WithInstance[*transport](&transport{})
	instance, _ := resolver.ResolveWithOptions(ctx, resolveServiceType, transportInstance)

	fmt.Println(reflect.ValueOf(instance).Pointer())
}

type transport struct {
}

type client struct {
	transport *transport
}

func newClient(transport *transport) *client {
	return &client{
		transport: transport,
	}
}
