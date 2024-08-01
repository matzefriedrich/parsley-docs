---
label:  Activating unregistered dependencies
order: 20
tags: [ resolving, instances, complex dependencies, dynamic object creation ]
---
# Live services: Activating unregistered dependencies

Parsley allows you to create live service instances with dependencies using factory functions and the `Activate` method. This feature enables you to dynamically instantiate objects using registered services without prior registration of the object types in question.

:::code language="golang" source="/examples/resolving-services/cmd/resolve-live-services/main.go" :::

In this example, a `Greeter` service is registered with a transient lifetime. The `Activate` method is used to create an instance of the `ouchie` struct on the fly, consuming the `Greeter` dependency within the factory function.

This is helpful for use cases like the following:

* **Dynamic object creation:** When you need to create objects that depend on registered services without having to register these objects themselves.
 
* **Complex dependencies:** For scenarios where objects require a complex set of dependencies that are not straightforward to register individually.

* **Third-party integrations:** When integrating with third-party libraries that require on-the-fly object creation based on dynamically resolved dependencies.

Activating unregistered dependencies enhances flexibility in managing dependencies and streamlines the process of integrating various services dynamically and efficiently.