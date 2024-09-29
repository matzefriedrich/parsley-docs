---
icon: flame
label: "Walkthrough: Parsley Integration with Cobra"
tags: [  integrations, guide, walkthrough, examples ]
---
# Walkthrough: Parsley Integration with Cobra

This guide demonstrates how to integrate Parsley with [Cobra](https://github.com/spf13/cobra) to build a modular and maintainable CLI application. It leverages the [cobra-extension package](https://github.com/matzefriedrich/cobra-extensions) package to implement command handlers based on typed command services, making it a natural fit for Parsley’s dependency injection capabilities.

## Project Structure

The code for this example can be found at [examples/integrations/cobra](https://github.com/matzefriedrich/parsley-docs/tree/main/examples/integrations/cobra) it has the following structure:

```text
cobra
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    ├── commands
    │   └── hello_command.go
    ├── modules
    │   ├── cobra_module.go
    │   └── services_module.go
    └── services
        └── greeter.go
```

### Main Application

The `cmd/main.go` file serves as the application's entry point. Here, the service dependencies are registered using a `ServiceRegistry` instance. Dependency registration is organized in `ModuleFunc` methods, keeping the main setup clean. Finally, a `*charmer.CommandLineApplication` service is resolved, and its Execute method is called to run the application.

:::code language="golang" source="/examples/integrations/cobra/cmd/main.go" :::

### Modules

The `internal/modules` package defines the service configurations required for the CLI setup.

The `cobra_module.go` module configures the Cobra application and registers it as a singleton service. It also enables the resolution of all `*cobra.Command` types as a list, which is necessary to populate the Cobra application with registered commands.

:::code language="golang" source="/examples/integrations/cobra/internal/modules/cobra_module.go" :::

The `services_module.go` registers application-specific services, such as the `Greeter` service.

:::code language="golang" source="/examples/integrations/cobra/internal/modules/services_module.go" :::

### Services

The Greeter service in `internal/services/greeter.go` generates a greeting message based on the provided name and politeness level. This service is used by command handlers to produce responses.

:::code language="golang" source="/examples/integrations/cobra/internal/services/greeter.go" :::

### Command Handlers

Command handlers in this example are services that implement the `TypedCommand` interface, providing structured logic and dependency injection. For instance, the `helloCommand` service uses `Greeter` to output personalized greetings. With cobra-extensions and Parsley, command declaration, service resolution, and command logic are consolidated in the same module, making the codebase easier to navigate and maintain.

:::code language="golang" source="/examples/integrations/cobra/internal/commands/hello_command.go" :::


### Running the Application

As configured in the code, the application will start a Cobra application and call it´s `Execute` method. To run the application, navigate to the root directory and execute the following command:

```sh
go run main.go hello --name John --polite
```

This should return:

```text
Good day, John!
```

For more details on Parsley, check out the Parsley GitHub repository.