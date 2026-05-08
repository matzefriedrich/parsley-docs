# Quick Start

This quick start guide walks you through structuring your application and using Parsley to dynamically register and resolve services (or components) at runtime. In this documentation, the terms "services" and "components" are used interchangeably—they both refer to object instances of a specific type. To effectively understand dependency mapping in Parsley, let's look at a practical example involving types, interfaces, and constructor functions. Once you are familiar with the basics, you can check out the rest of this documentation for more advanced usage and integration examples.

## Structuring the Application

Inversion of Control (IoC) is a design principle that flips your application's traditional approach to handling dependencies. Instead of having components directly instantiate their dependencies, you structure your application so that dependencies are provided to components at the time of their creation. In Parsley, object instances are created through constructor functions, whereby dependencies are expressed as arguments. This approach decouples your components and promotes more flexible and maintainable code.

In the quick start example, we define a `DataService` interface and provide two implementations for it. The application maps the constructor functions to the abstraction and resolves all registered services to call their `FetchData` method in a loop.

### Add the Parsley Reference

Use the following command to add a reference to the latest version of the Parsley library to your project:

```bash
go get github.com/matzefriedrich/parsley
```

### Define Interfaces

Interfaces in Go are a powerful tool for decoupling dependencies and defining contracts between different components of an application or library. While the Go community often emphasizes composition over inheritance and favors concrete types over interfaces in many scenarios, defining interfaces remains viable, particularly when integrating with existing enterprise patterns—such as dependency injection—from other languages.

```go
type DataService interface {
    FetchData() string
}
```

In the provided example, `DataService` is an interface that mandates the implementation of the `FetchData` method, which should return a string. This contract ensures that any implementing `DataService` will provide a method to fetch data, regardless of how it is internally implemented (e.g., fetching from a remote server, database, or local storage).

### Create Implementation Types

Implementation types are the concrete structs that fulfill the contracts defined by the interfaces. If combined with constructor functions (whose return types are interfaces), those structs can be kept private.

```go
type remoteDataService struct {}

func (s *remoteDataService) FetchData() string {
    return "Data from remote server"
}

type localDataService struct {}

func (s *localDataService) FetchData() string {
    return "Data from local database"
}
```

### Define Constructor Functions

Constructor functions are responsible for creating instances of the implementation types. These functions specify the dependencies required by the implementation type through their parameters. Parsley supports various signatures for constructor functions, including those that accept a `context.Context` and those that return an `error`.

To keep things simple for now, the services in this example are not dependent on other services; thus, the constructor functions do not have any parameters.

```go
func NewRemoteDataService() DataService {
    return &remoteDataService{}
}

func NewLocalDataService() DataService {
    return &localDataService{}
}
```

### Configure Dependency Mapping in Parsley

With the types, interfaces, and constructor functions defined, you can now configure Parsley to map these dependencies. This involves setting up a service registry and registering the services with appropriate lifetimes. The example below registers the `NewRemoteDataService` and `NewLocalDataService` constructor functions as transient services.

* **Create a service registry**: The `NewServiceRegistry` function initializes a new service registry, which will hold the configuration for all your service mappings.
* **Register services**: Services are registered with the registry using functions such as `RegisterTransient`, `RegisterScoped`, and `RegisterSingleton`. In this example, we're registering services using the constructor functions `NewRemoteDataService` and `NewLocalDataService` with a transient lifetime behavior, causing the resolver to instantiate services each time they are requested.

```go
registry := registration.NewServiceRegistry()
_ = registration.RegisterTransient(registry, NewRemoteDataService)
_ = registration.RegisterTransient(registry, NewLocalDataService)
```

### Resolve Dependencies

Finally, services can be resolved via Parsley. This involves using the `resolving` package to obtain instances of your services as needed. Here's a more detailed breakdown of the process:

* **Create a resolver**: The `NewResolver` function initializes a resolver with the provided registry, which contains the configuration for all your service mappings.
* **Create a context to manage scoped services**: The `NewScopedContext` function creates a new scope context. A scope context defines the lifetime and scope of the services being resolved. Here, we're using the background context from the `context` package, which is a common way to initialize a root context in Go.
* **Resolve the required services**: The `ResolveRequiredServices[T]` function is used to retrieve instances of the specified service type (`DataService` in this case). This function requires the resolver and scope context as parameters. It returns a list of object instances compatible with the specified type `T` and an error, allowing you to handle any issues that might occur during the resolution process.

```go
resolver := resolving.NewResolver(registry)
scope := resolving.NewScopedContext(context.Background())
dataServices, _ := resolving.ResolveRequiredServices[DataService](scope, resolver)

for _, service := range dataServices {
    data := service.FetchData()
    fmt.Println(data)
}
```
