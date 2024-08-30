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

	registry.Register(internal.NewLocalDataService, types.LifetimeTransient)
	registry.Register(internal.NewRemoteDataService, types.LifetimeTransient)
	features.RegisterList[internal.DataService](registry)

	registry.Register(newAggregator, types.LifetimeTransient)

	resolver := resolving.NewResolver(registry)
	ctx := resolving.NewScopedContext(context.Background())
	aggregator, _ := resolving.ResolveRequiredService[*dataAggregationService](resolver, ctx)

	results := aggregator.fetchAll()
	for _, result := range results {
		fmt.Println(result)
	}
}

type dataAggregationService struct {
	dataServices []internal.DataService
}

func newAggregator(dataServices []internal.DataService) *dataAggregationService {
	return &dataAggregationService{
		dataServices: dataServices,
	}
}

func (a *dataAggregationService) fetchAll() []string {
	results := make([]string, 0)
	for _, dataService := range a.dataServices {
		data := dataService.FetchData()
		results = append(results, data)
	}
	return results
}
