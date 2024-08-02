---
label:  Activating unregistered dependencies
order: 20
tags: [ resolving, instances, complex dependencies, dynamic object creation ]
---
# Live services: Activating unregistered dependencies

The `Activate` method allows you to dynamically create instances using registered services, even if the requested service type is not registered. This approach is helpful for scenarios where you need to instantiate objects on the fly with specific dependencies provided by registered services.

:::code language="golang" source="/examples/resolving-services/cmd/resolve-live-services/main.go" :::

In this example, a `Greeter` service is registered with a transient lifetime. The `Activate` method is used to create an instance of the `ouchie` struct on the fly, consuming the `Greeter` dependency within the factory function.

This is helpful for use cases like the following:

* **Dynamic object creation:** When you need to create objects that depend on registered services without having to register these objects themselves.
 
* **Complex dependencies:** For scenarios where objects require a complex set of dependencies that are not straightforward to register individually.

* **Third-party integrations:** When integrating with third-party libraries that require on-the-fly object creation based on dynamically resolved dependencies.

The `Activate` method allows for flexible and efficient dependency injection by dynamically pulling registered instances from the resolver and injecting them into your factory function, ensuring that your dynamically created objects have the necessary dependencies.