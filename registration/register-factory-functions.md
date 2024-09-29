---
meta:
  title: Parsley - Advanced Service Registration using Factory Functions
description: Learn how to use factory functions in Parsley to gain more control over service creation. Discover the difference between factory and constructor functions, and see how factory functions enable dynamic service configuration with additional parameters.
icon: file
label: Register Factory Functions
tags: [ registration, factory function, service factory ]
category:
  - Registration
---
# Register Factory Functions

In Parsely, you can register a factory function to provide more control and customization when creating service instances. A factory function differs from a constructor function because it allows you to pass additional parameters to configure the service creation process.

Instead of directly registering a constructor function that returns a service instance, you register a factory function that returns another function. This returned function is then used to create the actual service instances.

Here’s a detailed look at how to register and use a factory function:

Suppose you want to create a `Greeter` service where the salutation can be customized. You can use a factory function to achieve this. Here’s how you would define and register a factory function:

:::code language="golang" source="/examples/registration-concepts/internal/greeter.go" :::

This is how you use it:

:::code language="golang" source="/examples/registration-concepts/cmd/factory-functions/main.go" range="10-21" :::
