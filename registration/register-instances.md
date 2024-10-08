---
meta:
  title: Parsley - Register service instances
description: Learn how to use the `RegisterInstance` method in Parsley to manually register and manage complex service instances. Ensure consistent state and availability across your application by treating registered instances as singletons, making them accessible to other services as dependencies.
icon: file
label: Register Instances
tags: [ registration, instances, singleton ]
category:
  - Registration
---
# Register Instances

In Parsely, the `RegisterInstance` method registers a preexisting service instance, particularly useful for objects that cannot be created automatically due to complex configuration, data dependencies, or third-party APIs. Once registered, these instances behave as singletons, meaning the same instance is reused throughout the application. This ensures the registered object is consistently available as a dependency for other services, maintaining consistency and state across your application.

Let's go through the process of registering a service instance manually and using it within your application. 

Assume you have an interface `DataService` and an implementation `LocalDataService`.

```golang
package internal

type DataService interface {
	FetchData() string
}

type localDataService struct{}

func (s *localDataService) FetchData() string {
	return "Data from local database"
}

func NewLocalDataService() DataService {
	return &localDataService{}
}
````

Manually create an instance of the service and register it using `RegisterInstance`. This method is ideal for instances that require specific configuration or are provided by an external factory. To resolve and use the registered service instance, simply request it from the resolver:

:::code language="golang" source="/examples/registration-concepts/cmd/instances/main.go" range="13-21" :::

By registering service instances manually, you can integrate complex, pre-configured, or third-party-provided objects into your application, ensuring they are available as dependencies for other services managed by Parsely. This method is handy when dealing with objects that cannot be easily instantiated within the standard service registration flow.