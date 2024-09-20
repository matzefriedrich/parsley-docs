---
icon: file
label: Register Constructor Functions
tags: [ registration, constructor function ]
---
# Register Constructor Functions

In Parsely, the fundamental way to register services is by providing a constructor function to the service registry. This method enables you to create instances of your services in a controlled manner and inject them into your application as needed.

### What is a constructor function?

In Go, a **constructor function** is a regular function used to initialize and return an instance of a type. It typically follows the pattern `NewTypeName` and allows for setting up a struct with necessary values or configurations before returning it.  Unlike languages with built-in constructors, Go doesn't have special constructor syntax, so developers create these functions explicitly to manage object initialization and ensure type safety. For example, a constructor for the `greeter` struct might be `NewGreeter() *greeter` or `NewGreeter() Greeter` if interfaces are preferred.

When registering a constructor function for a service type, Parsley inspects the given function upon service resolution and automatically determines the required services. Constructor functions must also be registered for required services so that Parsley can construct the whole tree of dependencies and pass service instances as needed.

## Example

Here’s a basic example of registering a service using a constructor function:

:::code language="golang" source="/examples/quick-start/cmd/main.go" range="13-14" :::

In this example, the `NewGreeter` function is used to create instances of the `Greeter` interface. This function returns an implementation of `Greeter` while hiding the concrete type `*greeter`.

The `Greeter` interface and its implementation are defined as follows:

:::code language="golang" source="/examples/quick-start/internal/greeter.go" :::

While it’s possible for `NewGreeter` to return a `*greeter` directly, it is generally advisable to return the interface type `Greeter`. This approach maintains flexibility and abstraction, making it easier to change implementations without modifying dependent code.