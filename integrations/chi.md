# Walkthrough: Parsley Integration with Chi

This guide demonstrates how to integrate the Parsley dependency injection framework with the Chi web framework. By following this example, you'll learn how to set up a Chi application with dependency injection managed by Parsley, making your codebase more modular, testable, and maintainable.

## Project Structure

The code for this example can be found at [examples/integrations/chi](https://github.com/matzefriedrich/parsley-docs/tree/main/examples/integrations/chi). It has the following structure:

```text
chi
├── cmd
│   └── main.go
└── internal
    ├── application.go
    ├── modules
    │   ├── chi_module.go
    │   ├── greeter_module.go
    │   └── route_handlers_module.go
    ├── route_handlers
    │   ├── greeter.go
    │   └── types.go
    └── services
        └── greeter.go
```

This article describes the project's structure and the purpose of each module in detail.

### Main Application

The main entry point of the application is in the `cmd/main.go` file:

:::code language="golang" source="/examples/integrations/chi/cmd/main.go" :::

In this file, the `RunParsleyApplication` function is called to bootstrap the application. It initializes the Parsley application context and configures the Chi server with the necessary services and route handlers.

### Modules

The `modules` package contains the service configurations required to set up the Chi application. The `chi_module.go` module configures the Chi router and registers it as a singleton service within the Parsley framework:

:::code language="golang" source="/examples/integrations/chi/internal/modules/chi_module.go" :::

This configuration ensures that the Chi router is initialized and available for dependency injection.

### Services

Next, the `internal/services/greeter.go` file defines the `Greeter` service, which returns a greeting message based on the provided name and politeness flag.

:::code language="golang" source="/examples/integrations/chi/internal/services/greeter.go" :::

The `Greeter` service is registered by the `greeter_module.go` module.

:::code language="golang" source="/examples/integrations/chi/internal/modules/greeter_module.go" :::

### Route Handlers

In this example, route handlers are also services, structs implementing the `RouteHandler` interface, which register one or more route handlers with the Chi router. The interface is defined as follows:

:::code language="golang" source="/examples/integrations/chi/internal/route_handlers/types.go" :::

The `internal/route_handlers/greeter.go` file registers the route handler for the `/say-hello` endpoint, which returns a greeting message based on the query parameters provided in the request. The logic for the message generation is handled by the `Greeter` service, which is injected into the `NewGreeterRouteHandler` method.

:::code language="golang" source="/examples/integrations/chi/internal/route_handlers/greeter.go" :::

The `route_handler_module.go` file handles the registration of the `RouteHandler` services themselves.

:::code language="golang" source="/examples/integrations/chi/internal/modules/route_handlers_module.go" :::

This configuration ensures that all route handlers the application requires are correctly registered and injected into the Chi router instance. Since the application service expects a set of route handler services, the `RegisterList` method must be used to register a list activator for the `RouteHandler` type.

### Application Logic

The `internal/application.go` file contains the main application service:

:::code language="golang" source="/examples/integrations/chi/internal/application.go" :::

This file defines the `parsleyApplication` struct as the main application service. It registers the route handlers and starts the Chi server on port `5504`. However, aspects like having the listener port configurable or a graceful server shutdown are omitted here, but they could be addressed here as well.

### Running the Application

To run the application, navigate to the root directory and execute the following command:

```bash
go run main.go
```

As configured in the code, the application will start a Chi server on `http://localhost:5504`. You can then access the `/say-hello` endpoint:

```bash
curl "http://localhost:5504/say-hello?name=John"
```

This should return:

```text
Good day, John!
```

For more details on Parsley, check out the Parsley GitHub repository.
