---
label: Register Lists
tags: [ registration, service lists ]
---
# Register service lists

Parsley supports registering multiple services under the same contract type, allowing you to inject a list of these services into your application. While you can always resolve multiple services of the same contract using the `ResolveRequiredServices[T]` method, injecting them as an array or slice requires a list registration via the `RegisterList[T]`  method.

This method facilitates the automatic injection of all registered services that share the same contract type, providing greater flexibility and scalability in managing multiple service implementations.

## Example

This example demonstrates registering and injecting a list of services sharing the same contract type using Parsley's `RegisterList[T]` feature. In the code, two data service constructor functions (`NewLocalDataService` and `NewRemoteDataService`) are registered, and then `RegisterList[T]` is used to group the service registrations under the `DataService` contract. Please note that both constructor functions must return the service instances as `DataService` objects to make this automation work.

:::code language="golang" source="/examples/registration-concepts/cmd/service-lists/main.go" :::

The `dataAggregationService` aggregates data from all registered `DataService` instances. When `fetchAll()` is called, it iterates over the injected list of `DataService` implementations, collects their data, and prints the results.

The printed output confirms that local and remote data services are successfully aggregated and utilized.

## Benefits and use cases

This feature is handy when working with multiple implementations of the same service contract, such as aggregating results from various sources or supporting different strategies for a given operation. 

The `RegisterList[T]` method simplifies managing these services and ensures that all relevant implementations are easily accessible in a single injection. This is especially valuable in modular systems or scenarios requiring dynamic extension of service capabilities.