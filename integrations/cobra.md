---
icon: flame
label: "Walkthrough: Parsley Integration with Cobra"
tags: [  integrations, guide, walkthrough, examples ]
---
# Walkthrough: Parsley Integration with Cobra

This article provides an example of integrating Parsley with [Cobra](https://github.com/spf13/cobra) to build a CLI application. The example in the walkthrough uses the [cobra-extension package](https://github.com/matzefriedricg/cobra-extensions), which provides functionality to implement Cobra command handlers based on typed command services, which is the foundation for integrating with Parsley. The goal is to showcase how Parsley’s service registration and resolving features can streamline the creation of a modular and maintainable CLI setup.

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

The main entry point of the application is in the `cmd/main.go` file:

:::code language="golang" source="/examples/integrations/cobra/cmd/main.go" :::

This file registers all dependencies with a `ServiceRegistry` instance. To keep things clean, the registration of application services and commands is organized in `ModuleFunc` methods. After all application dependencies are registered, a `*charmer.CommandLineApplication` service is resolved through a `Resolver`; Finally, the `Execute` method is called to process the command.

### Modules

The `modules` package contains the service configurations required to set up the Cobra application.

The `cobra_module.go` module configures the Cobra application, registers it as a singleton service, and configures the `*cobra.Command` type to be resolvable as a list. The latter is required so that all registered services can be resolved at once and added to the Cobra application.

:::code language="golang" source="/examples/integrations/cobra/internal/modules/cobra_module.go" :::

### Services

Next, the `internal/services/greeter.go` module file defines the `Greeter` service, which returns a greeting message based on the provided name and politeness flag.

:::code language="golang" source="/examples/integrations/cobra/internal/services/greeter.go" :::

The `Greeter` service is registered by the `services_module.go` module.

:::code language="golang" source="/examples/integrations/cobra/internal/modules/services_module.go" :::

### Command Handlers

In this example, command handlers are services,  i.e., structs implementing the `TypedCommand` interface. The constructor functions of command handler services must return a `*cobra.Command` object so that the application factory function can resolve and add them. The example project provides the `helloCommand` service that consumes a `Greeter` service to render the greeting message.

With **cobra-extensions** and **Parsley**, command declaration (a struct with `flag` struct tags), the demand of required services via a constructor function, and command logic are all defined in the same module. This helps to organize aspects that belong together as a single unit.

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