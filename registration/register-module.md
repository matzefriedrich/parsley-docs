---
meta:
  title: Parsley - Organize Service Registrations using Modules
description: Group related services into modules in Parsley to improve code organization, reusability, and separation of concerns.
icon: file
label: Register Module
tags: [ registration ]
order: 75
category:
  - Registration
---
# Register Service Modules

Parsley provides a convenient way to group related services into modules using the `RegisterModule` method. This method allows you to register a function that bundles service registrations, making your code more organized and maintainable. Here's how to use it:

:::code language="golang" source="/examples/registration-concepts/cmd/modules/main.go" :::

### Conditional Module Registration

Sometimes, you might want to register a module only if a certain condition is met (e.g., based on environment variables or configuration). For this, Parsley provides the `RegisterModuleIf` method:

```go
registry := registration.NewServiceRegistry()
isDev := os.Getenv("ENV") == "development"

// Register the DebugModule only in development environment
_ = registry.RegisterModuleIf(isDev, DebugModule)
```

## Benefits

Using `RegisterModule` offers several benefits. It helps maintain a clean and organized code structure by grouping related services, making the codebase more straightforward to manage.

Modules allow you to define and reuse service groupings across different parts of your application, promoting modularity and reusability. Additionally, this approach keeps the service registration logic separate from the application logic, enhancing the separation of concerns and improving overall code maintainability and clarity as your application scales.

Another use case could be a package that keeps all service implementation types private to the package but exports nothing else, like interfaces and module registration functions. This way, services can be integrated into apps without exposure.