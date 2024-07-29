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

	_ = features.RegisterNamed[func(string) internal.DataService](registry,
		registration.NamedServiceRegistration("remote", internal.NewRemoteDataServiceFactory, types.LifetimeTransient),
		registration.NamedServiceRegistration("local", internal.NewLocalDataServiceFactory, types.LifetimeTransient))

	resolver := resolving.NewResolver(registry)
	scope := resolving.NewScopedContext(context.Background())

	serviceFactory, _ := resolving.ResolveRequiredService[func(string) (func(string) internal.DataService, error)](resolver, scope)
	localDataService, err := serviceFactory("remote")
	if err != nil {
		panic(err)
	}

	dataService := localDataService("John")
	data := dataService.FetchData()
	fmt.Println(data)
}
