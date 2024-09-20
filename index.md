---
icon: home
tags: [ introduction, guide ]
---

# Welcome to Parsley

Parsley is an easy-to-use, reflection-based dependency injection package that fits seamlessly into any Go application. It bridges the gap between dependency configuration and service activation, providing automated lifetime management and a clean, organized way to handle dependencies.

## Key Features

### Type Registration & Resolution

- **Constructor Functions:** [Register types via constructor functions](registration/register-constructor-functions.md) and automatically inject dependencies into constructors.
- **Resolve by Type:** Support for both interfaces and pointer types.
- **Safe Casting:** Use `ResolveRequiredService[T]` for type-safe resolutions.
- **Lifetime Management:** [Register types with singleton, scoped, or transient lifetimes](resolving/lifetime-scopes.md).

### Advanced Registrations

- **Modular Registrations:** [Bundle registrations into modules](registration/register-module.md) for cleaner organization.
- **Lazy Loading:** [Dependencies are injected only when needed](resolving/resolve-lazy-proxy.md) using `Lazy[T]`.
- **Custom Factories:** [Define custom factory functions](registration/register-factory-functions.md) for dynamic resolution.
- **Provide to service instance to the resolver:** [Provide service instances to the resolver](resolving/resolve-with.md) that cannnot be automatically constructed.

### Multiple Registrations

- **Named Services:** [Register and resolve multiple services for the same interface](registration/register-named-services.md).
- **Service Lists:** [Resolve services as a list](registration/register-lists.md).

### Proxy & Mocking Support

- **Proxy Services:** [Generate proxies as drop-in replacements with method interception](advanced-features/generate-proxy-services.md).
- **Mock Generation:** [Create configurable mocks](advanced-features/generate-mocks.md) to enhance testing.

Want to know more? Read on to discover how Parsley can enhance your Go application's architecture and maintainability.