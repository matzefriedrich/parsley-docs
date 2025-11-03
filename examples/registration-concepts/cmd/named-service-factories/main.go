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

	ctx := context.Background()

	registry := registration.NewServiceRegistry()

	_ = features.RegisterNamed[func(string) internal.DataService](ctx, registry,
		registration.NamedServiceRegistration("remote", internal.NewRemoteDataServiceFactory, types.LifetimeTransient),
		registration.NamedServiceRegistration("local", internal.NewLocalDataServiceFactory, types.LifetimeTransient))

	resolver := resolving.NewResolver(registry)
	resolverContext := resolving.NewScopedContext(ctx)

	serviceFactory, _ := resolving.ResolveRequiredService[func(string) (func(string) internal.DataService, error)](resolverContext, resolver)
	localDataService, err := serviceFactory("remote")
	if err != nil {
		panic(err)
	}

	dataService := localDataService("John")
	data := dataService.FetchData()
	fmt.Println(data)
}
