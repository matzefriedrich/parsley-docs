package main

import (
	"context"
	"fmt"

	"github.com/matzefriedrich/parsley-docs/examples/registration-concepts/internal"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()

	_ = features.RegisterNamed[internal.DataService](registry,
		registration.NamedServiceRegistration("remote", internal.NewRemoteDataService, types.LifetimeTransient),
		registration.NamedServiceRegistration("local", internal.NewLocalDataService, types.LifetimeTransient))

	resolver := resolving.NewResolver(registry)
	scope := resolving.NewScopedContext(context.Background())

	serviceFactory, _ := resolving.ResolveRequiredService[func(string) (internal.DataService, error)](resolver, scope)
	localDataService, err := serviceFactory("local")
	if err != nil {
		panic(err)
	}

	data := localDataService.FetchData()
	fmt.Println(data)
}
