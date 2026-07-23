# Bootstrap Application

The Parsley bootstrap package provides a streamlined way to manage your application's lifecycle. It automates the process of initializing the service registry, configuring modules, and running the main application entry point.

## The Application Interface

To use the bootstrap package, your application must implement the `bootstrap.Application` interface. This interface defines a single method:

```golang
type Application interface {
    Run(ctx context.Context) error
}
```

The `Run` method serves as the entry point for your application logic, such as starting a web server or a background worker.

## Running the Application

Use `bootstrap.RunParsleyApplication` to initialize and start your application. This function takes a context, an application factory function, and optional module configuration functions.

```golang
func RunParsleyApplication(ctx context.Context, appFactoryFunc any, configure ...types.ModuleFunc) error
```

### Example

The following example demonstrates how to bootstrap a web application using the `chi` router.

First, define your application struct and its constructor:

:::code language="golang" source="/examples/integrations/chi/internal/application.go" :::

Then, in your `main` function, use `bootstrap.RunParsleyApplication` to start the application:

:::code language="golang" source="/examples/integrations/chi/cmd/main.go" :::

## How It Works

When you call `bootstrap.RunParsleyApplication`, the following steps are performed:

1. **Registry Initialization:** A new service registry is created.
2. **Module Configuration:** All provided `ModuleFunc` functions are executed to register services and configurations.
3. **Application Registration:** The `appFactoryFunc` is registered as a singleton service of type `bootstrap.Application`.
4. **Resolution:** The `bootstrap.Application` instance is resolved from the registry.
5. **Execution:** The `Run` method of the resolved application instance is invoked with the provided context.
