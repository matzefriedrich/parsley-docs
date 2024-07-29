package main

import (
	"context"
	"fmt"
	"github.com/matzefriedrich/parsley-docs/examples/registration-concepts/internal"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
)

func main() {

	registry := registration.NewServiceRegistry()
	registration.RegisterInstance[internal.DataService](registry, internal.NewLocalDataService())

	resolver := resolving.NewResolver(registry)
	scope := resolving.NewScopedContext(context.Background())

	actual, _ := resolving.ResolveRequiredService[internal.DataService](resolver, scope)
	data := actual.FetchData()
	fmt.Println(data)
}
