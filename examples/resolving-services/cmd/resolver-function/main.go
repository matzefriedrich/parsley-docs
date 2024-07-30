package main

import (
	"context"
	"fmt"
	"github.com/matzefriedrich/parsley-docs/examples/resolving-services/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
	"reflect"
)

func main() {

	registry := registration.NewServiceRegistry()
	err := registry.Register(resolveGreeter("Hi"), types.LifetimeSingleton)
	if err != nil {
		panic(err)
	}

	resolver := resolving.NewResolver(registry)

	for i := 0; i < 3; i++ {
		ctx := resolving.NewScopedContext(context.Background())
		greeter, _ := resolving.ResolveRequiredService[internal.Greeter](resolver, ctx)
		greeter.SayHello("John", false)
	}
}

func resolveGreeter(salutation string) func(resolver types.Resolver) internal.Greeter {
	factory := internal.NewGreeterFactory(salutation)
	return func(resolver types.Resolver) internal.Greeter {
		greeter := factory()
		traceResolveEventFor(greeter)
		return greeter
	}
}

func traceResolveEventFor(service any) {
	value := reflect.ValueOf(service)
	instancePointer := value.Pointer()
	typeName := value.Elem().Type().String()
	fmt.Printf("New %s instance created: %d\n", typeName, instancePointer)
}
