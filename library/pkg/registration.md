---
meta:
  title: Parsley Docs - "registration" package
description: Official documentation of the Parsley´s "registration" package
icon: package
label: registration
tags: [ "packages", "registration" ]
category:
  - Package Reference
---
# registration

```go
import "github.com/matzefriedrich/parsley/pkg/registration"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func CreateServiceActivatorFrom\[T any\]\(instance T\) \(func\(\) T, error\)](<#CreateServiceActivatorFrom>)
- [func CreateServiceRegistration\(activatorFunc any, lifetimeScope types.LifetimeScope\) \(types.ServiceRegistrationSetup, error\)](<#CreateServiceRegistration>)
- [func NewDependencyInfo\(registration types.ServiceRegistration, instance interface\{\}, consumer types.DependencyInfo\) types.DependencyInfo](<#NewDependencyInfo>)
- [func NewMultiRegistryAccessor\(registries ...types.ServiceRegistryAccessor\) types.ServiceRegistryAccessor](<#NewMultiRegistryAccessor>)
- [func NewServiceRegistrationList\(sequence core.ServiceIdSequence\) types.ServiceRegistrationList](<#NewServiceRegistrationList>)
- [func NewServiceRegistry\(\) types.ServiceRegistry](<#NewServiceRegistry>)
- [func RegisterInstance\[T any\]\(registry types.ServiceRegistry, instance T\) error](<#RegisterInstance>)
- [func RegisterScoped\(registry SupportsRegisterActivatorFunc, activatorFunc ...any\) error](<#RegisterScoped>)
- [func RegisterSingleton\(registry SupportsRegisterActivatorFunc, activatorFunc ...any\) error](<#RegisterSingleton>)
- [func RegisterTransient\(registry SupportsRegisterActivatorFunc, activatorFunc ...any\) error](<#RegisterTransient>)
- [type NamedServiceRegistrationFunc](<#NamedServiceRegistrationFunc>)
  - [func NamedServiceRegistration\(name string, activatorFunc any, scope types.LifetimeScope\) NamedServiceRegistrationFunc](<#NamedServiceRegistration>)
- [type SupportsRegisterActivatorFunc](<#SupportsRegisterActivatorFunc>)
- [type Validator](<#Validator>)
  - [func NewServiceRegistrationsValidator\(\) Validator](<#NewServiceRegistrationsValidator>)


## Constants

<a name="ErrorFailedToRetrieveServiceRegistrations"></a>

```go
const (
    ErrorFailedToRetrieveServiceRegistrations       = "failed to retrieve service registrations"
    ErrorRegistryMissesRequiredServiceRegistrations = "the registry misses required service registrations"
    ErrorCircularServiceRegistrationDetected        = "circular service registration detected"
)
```

## Variables

<a name="ErrFailedToRetrieveServiceRegistrations"></a>

```go
var (
    // ErrFailedToRetrieveServiceRegistrations signifies an error encountered while attempting to retrieve service registrations.
    ErrFailedToRetrieveServiceRegistrations = types.NewRegistryError(ErrorFailedToRetrieveServiceRegistrations)

    // ErrRegistryMissesRequiredServiceRegistrations indicates that required service registrations are missing.
    ErrRegistryMissesRequiredServiceRegistrations = types.NewRegistryError(ErrorRegistryMissesRequiredServiceRegistrations)

    // ErrCircularServiceRegistrationDetected signifies that a circular service registration was encountered.
    ErrCircularServiceRegistrationDetected = types.NewResolverError(ErrorCircularServiceRegistrationDetected)
)
```

<a name="CreateServiceActivatorFrom"></a>
## func [CreateServiceActivatorFrom](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/activator.go#L10>)

```go
func CreateServiceActivatorFrom[T any](instance T) (func() T, error)
```

CreateServiceActivatorFrom creates a service activator function for a given instance of type T.

<a name="CreateServiceRegistration"></a>
## func [CreateServiceRegistration](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/service_registration.go#L111>)

```go
func CreateServiceRegistration(activatorFunc any, lifetimeScope types.LifetimeScope) (types.ServiceRegistrationSetup, error)
```

CreateServiceRegistration creates a service registration instance from the given activator function and lifetime scope.

<a name="NewDependencyInfo"></a>
## func [NewDependencyInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/dependency.go#L19>)

```go
func NewDependencyInfo(registration types.ServiceRegistration, instance interface{}, consumer types.DependencyInfo) types.DependencyInfo
```

NewDependencyInfo creates a new instance of types.DependencyInfo with the provided service registration, instance, and parent dependency.

<a name="NewMultiRegistryAccessor"></a>
## func [NewMultiRegistryAccessor](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/registry_accessor.go#L36>)

```go
func NewMultiRegistryAccessor(registries ...types.ServiceRegistryAccessor) types.ServiceRegistryAccessor
```

NewMultiRegistryAccessor creates a new ServiceRegistryAccessor that aggregates multiple registries.

<a name="NewServiceRegistrationList"></a>
## func [NewServiceRegistrationList](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/service_registration_list.go#L59>)

```go
func NewServiceRegistrationList(sequence core.ServiceIdSequence) types.ServiceRegistrationList
```

NewServiceRegistrationList creates a new service registration list instance.

<a name="NewServiceRegistry"></a>
## func [NewServiceRegistry](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/registry.go#L93>)

```go
func NewServiceRegistry() types.ServiceRegistry
```

NewServiceRegistry creates a new types.ServiceRegistry instance.

<a name="RegisterInstance"></a>
## func [RegisterInstance](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/activator.go#L30>)

```go
func RegisterInstance[T any](registry types.ServiceRegistry, instance T) error
```

RegisterInstance registers an instance of type T. A registered instance behaves like a service registration with a singleton lifetime scope.

<a name="RegisterScoped"></a>
## func [RegisterScoped](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/register_functions.go#L22>)

```go
func RegisterScoped(registry SupportsRegisterActivatorFunc, activatorFunc ...any) error
```

RegisterScoped registers services with a scoped lifetime in the provided service registry.

<a name="RegisterSingleton"></a>
## func [RegisterSingleton](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/register_functions.go#L33>)

```go
func RegisterSingleton(registry SupportsRegisterActivatorFunc, activatorFunc ...any) error
```

RegisterSingleton registers services with a singleton lifetime in the provided service registry.

<a name="RegisterTransient"></a>
## func [RegisterTransient](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/register_functions.go#L11>)

```go
func RegisterTransient(registry SupportsRegisterActivatorFunc, activatorFunc ...any) error
```

RegisterTransient registers services with a transient lifetime in the provided service registry.

<a name="NamedServiceRegistrationFunc"></a>
## type [NamedServiceRegistrationFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/named_service_registration.go#L8>)

NamedServiceRegistrationFunc defines a function that returns a service name, its activator function, and its lifetime scope. This type supports the internal infrastructure.

```go
type NamedServiceRegistrationFunc func() (name string, activatorFunc any, scope types.LifetimeScope)
```

<a name="NamedServiceRegistration"></a>
### func [NamedServiceRegistration](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/named_service_registration.go#L11>)

```go
func NamedServiceRegistration(name string, activatorFunc any, scope types.LifetimeScope) NamedServiceRegistrationFunc
```

NamedServiceRegistration registers a service with a specified name, activator function, and lifetime scope.

<a name="SupportsRegisterActivatorFunc"></a>
## type [SupportsRegisterActivatorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/register_functions.go#L6-L8>)

SupportsRegisterActivatorFunc allows the registration of activator functions with different lifetime scopes.

```go
type SupportsRegisterActivatorFunc interface {
    Register(activatorFunc any, scope types.LifetimeScope) error
}
```

<a name="Validator"></a>
## type [Validator](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/validator.go#L28-L32>)

Validator defines an interface to validate service registries..

```go
type Validator interface {

    // Validate checks the provided ServiceRegistry for missing, invalid, or circular service dependencies. Returns an error if any issues are found.
    Validate(registry types.ServiceRegistry) error
}
```

<a name="NewServiceRegistrationsValidator"></a>
### func [NewServiceRegistrationsValidator](<https://github.com/matzefriedrich/parsley/blob/main/pkg/registration/validator.go#L130>)

```go
func NewServiceRegistrationsValidator() Validator
```

NewServiceRegistrationsValidator creates a new Validator instance.

