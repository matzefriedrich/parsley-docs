---
label:  Resolving services with specific instances
order: 30
tags: [ resolving, instances ]
---
# Passing specific instances to the resolver 

Parsley's `ResolveWithOptions` method allows you to pass specific instances into the resolver, providing certain dependencies dynamically at resolution time. This method is handy when overriding registered dependencies or injecting unregistered instances into your objects.

In the following example, the `newClient` constructor function registers a `client` service with a transient lifetime. During the resolution, an existing transport instance is passed to the resolver using `WithInstance`, ensuring that this specific instance is used for the client's transport dependency.

:::code language="golang" source="/examples/resolving-services/cmd/resolve-with/main.go" :::

Parsley's ability to inject specific instances dynamically improves flexibility in managing dependencies and provides greater control over the instantiation process. This feature is particularly beneficial for the following use cases:

* **Runtime configuration:** Dynamically configure services based on runtime conditions by passing specific instances.

* **Dependency injection in testing:** Use mock objects or specific instances during tests to simulate various scenarios without altering the registration setup.

* **Third-party integrations:** Integrate with third-party services or libraries that require preconfigured instances, ensuring they are used as dependencies when needed.

By using `ResolveWithOptions`, you can push unregistered instances into the resolver, ensuring the correct instances are used when resolving dependencies.