---
meta:
  title: Parsley - Advanced Dependency Injection with Generated Proxies
description: This article explores the advanced dependency injection capabilities offered by the Parsley CLI through its `generate proxy` command. You’ll learn how to automatically generate proxy services and interfaces for your service contracts using `//go:generate` annotations.
icon: file
label: Generated Proxies
tags: [ registration, service proxy, service decorator, code generation ]
category:
  - Parsley CLI
  - Registration
  - Extensibility
---
# Advanced Dependency Injection with generated Proxies

The Parsley CLI introduces a powerful feature for advanced dependency injection through its `generate proxy` command. This command, combined with `//go:generate` annotations, automatically generates proxy services and interfaces for your service contracts.

Use the following command to install the `parsley-cli` utility:
```sh
go install github.com/matzefriedrich/parsley/cmd/parsley-cli
```

## Example

Consider a `Greeter` interface and its implementation. By adding the `//go:generate parsley-cli generate proxy` annotation in your code, you enable the Parsley CLI to generate a corresponding proxy class. This proxy wraps the original service and allows for method interception, enabling additional behavior to be injected around service calls.

:::code language="golang" source="/examples/advanced/internal/greeter.go" :::

The code generator creates a `greeter.proxy.g.go` that contains a type that also implements the `Greeter` interface and thus can be used as drop-in replacements for the actual `Greeter` service. The proxy type supports intercepting method calls, allowing custom logic to be added before or after a method is invoked.

:::code language="golang" source="/examples/advanced/internal/greeter.proxy.g.go" :::

In this example, the `GreeterProxy` wraps the `Greeter` service, and any registered `MethodInterceptor` services can act upon method invocations, adding custom behavior such as logging, validation, or modification of method parameters.

The boilerplate code to register the generated proxies and a custom `MethodInterceptor` for logging purposes looks as follows:

:::code language="golang" source="/examples/advanced/cmd/service-proxy/main.go" :::

## Benefits and use cases

* **Separation of Concerns:** Proxies can separate core business logic from cross-cutting concerns like logging or security.

* **Dynamic Interception:** Proxies allow dynamic interception of method calls, making adding or modifying behavior easier without altering the service's core logic.

* **Extensibility:** This feature provides a flexible mechanism to extend the functionality of services, making it ideal for scenarios where services require dynamic behavior adjustments. This advanced feature of Parsley CLI simplifies complex dependency injection scenarios, providing developers with robust tools to manage and extend service behavior effortlessly.