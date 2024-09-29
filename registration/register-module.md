---
meta:
  title: Parsley - Organize Service Registrations using Modules
description: Learn how to group related services with the `RegisterModule` method in Parsley to maintain modular code structure. Discover how service modules improve code organization, reusability, and separation of concerns, making your applications easier to scale and manage.
icon: file
label: Register Module
tags: [ registration, modules, separation of concerns, service groupings ]
category:
  - Registration
---
# Register Service Modules

Parsley provides a convenient way to group related services into modules using the `RegisterModule` method. This method allows you to register a function that bundles service registrations, making your code more organized and maintainable. Here’s how to use it:

:::code language="golang" source="/examples/registration-concepts/cmd/modules/main.go" :::

## Benefits

Using `RegisterModule` offers several benefits. It helps maintain a clean and organized code structure by grouping related services, making the codebase more straightforward to manage. 

Modules allow you to define and reuse service groupings across different parts of your application, promoting modularity and reusability. Additionally, this approach keeps the service registration logic separate from the application logic, enhancing the separation of concerns and improving overall code maintainability and clarity as your application scales.

Another use case could be a package that keeps all service implementation types private to the package but exports nothing else, like interfaces and module registration functions. This way, services can be integrated into apps without exposure. 