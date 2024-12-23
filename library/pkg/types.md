---
meta:
  title: Parsley Docs - "types" package
description: Official documentation of the Parsley´s "types" package
icon: '<svg viewBox="0 0 24 24"><path d="M10.68,10.16c-.28.08-.57.16-.9.25-.16.05-.2.06-.35-.13-.18-.22-.31-.36-.57-.48-.76-.39-1.5-.28-2.18.19-.82.56-1.24,1.38-1.23,2.41.01,1.02.68,1.86,1.63,1.99.82.11,1.51-.19,2.05-.84.11-.14.21-.29.33-.47h-2.33c-.25,0-.31-.17-.23-.38.16-.39.45-1.05.62-1.38.04-.08.12-.2.3-.2h3.88c.17-.58.46-1.13.83-1.65.88-1.22,1.94-1.86,3.38-2.12,1.23-.23,2.39-.1,3.44.65.95.69,1.54,1.61,1.7,2.83.21,1.72-.27,3.11-1.39,4.31-.8.85-1.77,1.39-2.89,1.63-.21.04-.43.06-.64.08-.11.01-.22.02-.33.03-1.1-.03-2.1-.36-2.94-1.12-.59-.54-1-1.21-1.21-1.98-.14.3-.31.59-.51.86-.87,1.21-2,1.96-3.44,2.16-1.18.17-2.28-.08-3.24-.84-.89-.71-1.4-1.65-1.53-2.82-.16-1.39.23-2.63,1.03-3.72.86-1.18,1.99-1.93,3.38-2.2,1.13-.22,2.22-.08,3.2.62.64.44,1.1,1.05,1.4,1.79.07.11.02.18-.12.22-.42.11-.77.21-1.13.31h-.01ZM18.66,11.61v.14c-.06,1.09-.58,1.91-1.53,2.43-.64.34-1.3.38-1.97.08-.87-.41-1.33-1.41-1.11-2.4.27-1.19.99-1.94,2.11-2.21,1.15-.28,2.24.43,2.46,1.69.02.09.02.18.03.28h.01Z" style="fill-rule: evenodd;"/></svg>'
label: types
tags: [ "packages", "types" ]
category:
  - Package Reference
---
# types

```go
import "github.com/matzefriedrich/parsley/pkg/types"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func NewDependencyError\(msg string\) error](<#NewDependencyError>)
- [func NewReflectionError\(msg string, initializers ...ParsleyErrorFunc\) error](<#NewReflectionError>)
- [func NewRegistryError\(msg string, initializers ...ParsleyErrorFunc\) error](<#NewRegistryError>)
- [func NewResolverError\(msg string, initializers ...ParsleyErrorFunc\) error](<#NewResolverError>)
- [type DependencyError](<#DependencyError>)
- [type DependencyInfo](<#DependencyInfo>)
- [type FunctionInfo](<#FunctionInfo>)
- [type FunctionParameterInfo](<#FunctionParameterInfo>)
- [type LifetimeScope](<#LifetimeScope>)
- [type ModuleFunc](<#ModuleFunc>)
- [type NamedService](<#NamedService>)
- [type ParsleyAggregateError](<#ParsleyAggregateError>)
  - [func \(f ParsleyAggregateError\) Error\(\) string](<#ParsleyAggregateError.Error>)
  - [func \(f ParsleyAggregateError\) Errors\(\) \[\]error](<#ParsleyAggregateError.Errors>)
  - [func \(f ParsleyAggregateError\) Is\(err error\) bool](<#ParsleyAggregateError.Is>)
- [type ParsleyError](<#ParsleyError>)
  - [func \(f ParsleyError\) Error\(\) string](<#ParsleyError.Error>)
  - [func \(f ParsleyError\) Is\(err error\) bool](<#ParsleyError.Is>)
  - [func \(f ParsleyError\) Unwrap\(\) error](<#ParsleyError.Unwrap>)
- [type ParsleyErrorFunc](<#ParsleyErrorFunc>)
  - [func ForServiceType\(serviceType string\) ParsleyErrorFunc](<#ForServiceType>)
  - [func WithAggregatedCause\(errs ...error\) ParsleyErrorFunc](<#WithAggregatedCause>)
  - [func WithCause\(err error\) ParsleyErrorFunc](<#WithCause>)
- [type ParsleyErrorWithServiceTypeName](<#ParsleyErrorWithServiceTypeName>)
- [type ReflectionError](<#ReflectionError>)
- [type RegistryError](<#RegistryError>)
  - [func \(r \*RegistryError\) MatchesServiceType\(name string\) bool](<#RegistryError.MatchesServiceType>)
  - [func \(r \*RegistryError\) ServiceTypeName\(name string\)](<#RegistryError.ServiceTypeName>)
- [type Resolver](<#Resolver>)
- [type ResolverError](<#ResolverError>)
  - [func \(r \*ResolverError\) ServiceTypeName\(name string\)](<#ResolverError.ServiceTypeName>)
- [type ResolverOptionsFunc](<#ResolverOptionsFunc>)
- [type ServiceKey](<#ServiceKey>)
  - [func NewServiceKey\(value string\) ServiceKey](<#NewServiceKey>)
  - [func \(s ServiceKey\) String\(\) string](<#ServiceKey.String>)
- [type ServiceRegistration](<#ServiceRegistration>)
- [type ServiceRegistrationList](<#ServiceRegistrationList>)
- [type ServiceRegistrationSetup](<#ServiceRegistrationSetup>)
- [type ServiceRegistry](<#ServiceRegistry>)
- [type ServiceRegistryAccessor](<#ServiceRegistryAccessor>)
- [type ServiceType](<#ServiceType>)
  - [func MakeServiceType\[T any\]\(\) ServiceType](<#MakeServiceType>)
  - [func ServiceTypeFrom\(t reflect.Type\) ServiceType](<#ServiceTypeFrom>)


## Constants

<a name="ErrorRequiresFunctionValue"></a>

```go
const (
    ErrorRequiresFunctionValue               = "the given value is not function"
    ErrorCannotRegisterModule                = "failed to register module"
    ErrorTypeAlreadyRegistered               = "type already registered"
    ErrorServiceAlreadyLinkedWithAnotherList = "service already linked with another list"
    ErrorFailedToRegisterType                = "failed to register type"
)
```

<a name="ErrorServiceTypeNotRegistered"></a>

```go
const (
    ErrorServiceTypeNotRegistered               = "service type is not registered"
    ErrorRequiredServiceNotRegistered           = "required service type is not registered"
    ErrorCannotResolveService                   = "cannot resolve service"
    ErrorAmbiguousServiceInstancesResolved      = "the resolve operation resulted in multiple service instances"
    ErrorActivatorFunctionInvalidReturnType     = "activator function has an invalid return type"
    ErrorCircularDependencyDetected             = "circular dependency detected"
    ErrorCannotBuildDependencyGraph             = "failed to build dependency graph"
    ErrorInstanceCannotBeNil                    = "instance cannot be nil"
    ErrorServiceTypeMustBeInterface             = "service type must be an interface"
    ErrorCannotRegisterTypeWithResolverOptions  = "cannot register type with resolver options"
    ErrorCannotCreateInstanceOfUnregisteredType = "failed to create instance of unregistered type"
)
```

<a name="ErrorInstanceAlreadySet"></a>

```go
const (
    ErrorInstanceAlreadySet = "instance already set"
)
```

## Variables

<a name="ErrRequiresFunctionValue"></a>

```go
var (

    // ErrRequiresFunctionValue indicates that the provided value is not a function.
    ErrRequiresFunctionValue = errors.New(ErrorRequiresFunctionValue)

    // ErrCannotRegisterModule indicates that the module registration process has failed.
    ErrCannotRegisterModule = errors.New(ErrorCannotRegisterModule)

    // ErrTypeAlreadyRegistered indicates that an attempt was made to register a type that is already registered.
    ErrTypeAlreadyRegistered = errors.New(ErrorTypeAlreadyRegistered)

    // ErrFailedToRegisterType indicates that the attempt to register a type has failed.
    ErrFailedToRegisterType = errors.New(ErrorFailedToRegisterType)
)
```

<a name="ErrServiceTypeNotRegistered"></a>

```go
var (

    // ErrServiceTypeNotRegistered is returned when attempting to resolve a service type that has not been registered.
    ErrServiceTypeNotRegistered = errors.New(ErrorServiceTypeNotRegistered)

    // ErrRequiredServiceNotRegistered is returned when a required service type is not registered.
    ErrRequiredServiceNotRegistered = errors.New(ErrorRequiredServiceNotRegistered)

    // ErrActivatorFunctionInvalidReturnType is returned when an activator function has an invalid return type.
    ErrActivatorFunctionInvalidReturnType = errors.New(ErrorCannotResolveService)

    // ErrCannotBuildDependencyGraph is returned when the resolver fails to build a dependency graph due to missing dependencies or other issues.
    ErrCannotBuildDependencyGraph = errors.New(ErrorCannotBuildDependencyGraph)

    // ErrCircularDependencyDetected is returned when a circular dependency is detected during the resolution process.
    ErrCircularDependencyDetected = errors.New(ErrorCircularDependencyDetected)

    // ErrInstanceCannotBeNil is returned when an instance provided is nil, but a non-nil value is required.
    ErrInstanceCannotBeNil = errors.New(ErrorInstanceCannotBeNil)

    // ErrServiceTypeMustBeInterface is returned when a service type is not an interface.
    ErrServiceTypeMustBeInterface = errors.New(ErrorServiceTypeMustBeInterface)

    // ErrCannotRegisterTypeWithResolverOptions is returned when the resolver failed to register a type via resolver options.
    ErrCannotRegisterTypeWithResolverOptions = errors.New(ErrorCannotRegisterTypeWithResolverOptions)

    // ErrCannotCreateInstanceOfUnregisteredType is returned when the resolver fails to instantiate a type that has not been registered.
    ErrCannotCreateInstanceOfUnregisteredType = errors.New(ErrorCannotCreateInstanceOfUnregisteredType)
)
```

<a name="ErrInstanceAlreadySet"></a>

```go
var (
    // ErrInstanceAlreadySet is returned when there is an attempt to set an instance that is already set.
    ErrInstanceAlreadySet = errors.New(ErrorInstanceAlreadySet)
)
```

<a name="NewDependencyError"></a>
## func [NewDependencyError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/dependency_error.go#L21>)

```go
func NewDependencyError(msg string) error
```

NewDependencyError creates a new DependencyError with the provided message.

<a name="NewReflectionError"></a>
## func [NewReflectionError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/reflection_error.go#L9>)

```go
func NewReflectionError(msg string, initializers ...ParsleyErrorFunc) error
```

NewReflectionError creates a new ReflectionError with a specified message and optional initializers.

<a name="NewRegistryError"></a>
## func [NewRegistryError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/registry_error.go#L51>)

```go
func NewRegistryError(msg string, initializers ...ParsleyErrorFunc) error
```

NewRegistryError creates a new RegistryError with the given message and initializers to modify the error.

<a name="NewResolverError"></a>
## func [NewResolverError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/resolver_error.go#L64>)

```go
func NewResolverError(msg string, initializers ...ParsleyErrorFunc) error
```

NewResolverError creates a new ResolverError with the provided message and applies optional ParsleyErrorFunc initializers.

<a name="DependencyError"></a>
## type [DependencyError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/dependency_error.go#L16-L18>)

DependencyError represents an error that occurs due to a missing or failed dependency. This error type encapsulates a ParsleyError.

```go
type DependencyError struct {
    ParsleyError
}
```

<a name="DependencyInfo"></a>
## type [DependencyInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L159-L189>)

DependencyInfo provides functionality to manage dependency information.

```go
type DependencyInfo interface {
    // AddRequiredServiceInfo adds a child dependency to the current dependency info.
    AddRequiredServiceInfo(child DependencyInfo)

    // CreateInstance creates an instance of the service associated with this dependency info.
    CreateInstance() (interface{}, error)

    // Consumer returns the parent dependency for the current dependency info.
    Consumer() DependencyInfo

    // HasInstance checks if an instance has already been created for the dependency represented by the current DependencyInfo object.
    HasInstance() bool

    // Instance retrieves the created instance of the service associated with this dependency info.
    Instance() interface{}

    // Registration gets the service registration of the current dependency info.
    Registration() ServiceRegistration

    // RequiredServiceTypes gets the service types required by this dependency info.
    RequiredServiceTypes() []ServiceType

    // RequiredServices retrieves the instances of services required by this dependency info.
    RequiredServices() ([]interface{}, error)

    // ServiceTypeName gets the name of the service type associated with this dependency info.
    ServiceTypeName() string

    // SetInstance sets the instance for the current dependency info.
    SetInstance(instance interface{}) error
}
```

<a name="FunctionInfo"></a>
## type [FunctionInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L10-L16>)

FunctionInfo Stores information about a service activator function. This interface supports the internal infrastructure.

```go
type FunctionInfo interface {
    fmt.Stringer
    Name() string
    Parameters() []FunctionParameterInfo
    ReturnType() ServiceType
    ParameterTypes() []ServiceType
}
```

<a name="FunctionParameterInfo"></a>
## type [FunctionParameterInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L18-L21>)



```go
type FunctionParameterInfo interface {
    fmt.Stringer
    Type() ServiceType
}
```

<a name="LifetimeScope"></a>
## type [LifetimeScope](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L192>)

LifetimeScope represents the duration for which a service or object instance is retained.

```go
type LifetimeScope uint
```

<a name="LifetimeTransient"></a>

```go
const (

    // LifetimeTransient represents a transient lifetime where a new instance is created each time it is requested.
    LifetimeTransient LifetimeScope = iota

    // LifetimeScoped represents a scoped lifetime where a single instance is created per scope.
    LifetimeScoped

    // LifetimeSingleton represents a single instance scope that persists for the lifetime of the application.
    LifetimeSingleton
)
```

<a name="ModuleFunc"></a>
## type [ModuleFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L78>)

ModuleFunc defines a function used to register services with the given service registry.

```go
type ModuleFunc func(registry ServiceRegistry) error
```

<a name="NamedService"></a>
## type [NamedService](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L140-L143>)

NamedService is a generic interface defining a service with a name and an activator function.

```go
type NamedService[T any] interface {
    Name() string
    ActivatorFunc() any
}
```

<a name="ParsleyAggregateError"></a>
## type [ParsleyAggregateError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L56-L59>)

ParsleyAggregateError represents an aggregate of multiple errors.

```go
type ParsleyAggregateError struct {
    Msg string
    // contains filtered or unexported fields
}
```

<a name="ParsleyAggregateError.Error"></a>
### func \(ParsleyAggregateError\) [Error](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L67>)

```go
func (f ParsleyAggregateError) Error() string
```

Error returns the message associated with the ParsleyAggregateError.

<a name="ParsleyAggregateError.Errors"></a>
### func \(ParsleyAggregateError\) [Errors](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L62>)

```go
func (f ParsleyAggregateError) Errors() []error
```

Errors returns the slice of errors contained within ParsleyAggregateError.

<a name="ParsleyAggregateError.Is"></a>
### func \(ParsleyAggregateError\) [Is](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L72>)

```go
func (f ParsleyAggregateError) Is(err error) bool
```

Is checks if the given error is equivalent to any error within the ParsleyAggregateError.

<a name="ParsleyError"></a>
## type [ParsleyError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L6-L9>)

ParsleyError represents an error with an associated message and optional underlying cause.

```go
type ParsleyError struct {
    Msg string
    // contains filtered or unexported fields
}
```

<a name="ParsleyError.Error"></a>
### func \(ParsleyError\) [Error](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L12>)

```go
func (f ParsleyError) Error() string
```

Error returns the message associated with the ParsleyError.

<a name="ParsleyError.Is"></a>
### func \(ParsleyError\) [Is](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L22>)

```go
func (f ParsleyError) Is(err error) bool
```

Is compares the current ParsleyError's message with another error's message to determine if they are the same.

<a name="ParsleyError.Unwrap"></a>
### func \(ParsleyError\) [Unwrap](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L17>)

```go
func (f ParsleyError) Unwrap() error
```

Unwrap returns the underlying cause of the ParsleyError, allowing for error unwrapping functionality.

<a name="ParsleyErrorFunc"></a>
## type [ParsleyErrorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L27>)

ParsleyErrorFunc is a function type that modifies a given error.

```go
type ParsleyErrorFunc func(e error)
```

<a name="ForServiceType"></a>
### func [ForServiceType](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L46>)

```go
func ForServiceType(serviceType string) ParsleyErrorFunc
```

ForServiceType creates a ParsleyErrorFunc that sets the service type name on errors that implement the ParsleyErrorWithServiceTypeName interface.

<a name="WithAggregatedCause"></a>
### func [WithAggregatedCause](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L85>)

```go
func WithAggregatedCause(errs ...error) ParsleyErrorFunc
```

WithAggregatedCause returns a ParsleyErrorFunc that sets an aggregated error cause with the provided errors.

<a name="WithCause"></a>
### func [WithCause](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L30>)

```go
func WithCause(err error) ParsleyErrorFunc
```

WithCause wraps a given error within a ParsleyError.

<a name="ParsleyErrorWithServiceTypeName"></a>
## type [ParsleyErrorWithServiceTypeName](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/parsley_error.go#L41-L43>)

ParsleyErrorWithServiceTypeName defines an interface for setting the service type name on errors.

```go
type ParsleyErrorWithServiceTypeName interface {
    ServiceTypeName(name string)
}
```

<a name="ReflectionError"></a>
## type [ReflectionError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/reflection_error.go#L4-L6>)

ReflectionError represents an error specifically related to reflection operations, extending ParsleyError.

```go
type ReflectionError struct {
    ParsleyError
}
```

<a name="RegistryError"></a>
## type [RegistryError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/registry_error.go#L29-L32>)

RegistryError represents an error that gets returned for failing registry operations.

```go
type RegistryError struct {
    ParsleyError
    // contains filtered or unexported fields
}
```

<a name="RegistryError.MatchesServiceType"></a>
### func \(\*RegistryError\) [MatchesServiceType](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/registry_error.go#L43>)

```go
func (r *RegistryError) MatchesServiceType(name string) bool
```

MatchesServiceType checks if the service type name of the RegistryError matches the specified name.

<a name="RegistryError.ServiceTypeName"></a>
### func \(\*RegistryError\) [ServiceTypeName](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/registry_error.go#L38>)

```go
func (r *RegistryError) ServiceTypeName(name string)
```

ServiceTypeName sets the service type name of the RegistryError.

<a name="Resolver"></a>
## type [Resolver](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L149-L156>)

Resolver provides methods to resolve registered services based on types.

```go
type Resolver interface {

    // Resolve attempts to resolve all registered services of the specified ServiceType.
    Resolve(ctx context.Context, serviceType ServiceType) ([]interface{}, error)

    // ResolveWithOptions resolves services of the specified type using additional options and returns a list of resolved services or an error.
    ResolveWithOptions(ctx context.Context, serviceType ServiceType, options ...ResolverOptionsFunc) ([]interface{}, error)
}
```

<a name="ResolverError"></a>
## type [ResolverError](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/resolver_error.go#L50-L53>)

ResolverError represents an error that gets returned for failing service resolver operations.

```go
type ResolverError struct {
    ParsleyError
    // contains filtered or unexported fields
}
```

<a name="ResolverError.ServiceTypeName"></a>
### func \(\*ResolverError\) [ServiceTypeName](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/resolver_error.go#L59>)

```go
func (r *ResolverError) ServiceTypeName(name string)
```

ServiceTypeName sets the service type name for the ResolverError instance.

<a name="ResolverOptionsFunc"></a>
## type [ResolverOptionsFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L146>)

ResolverOptionsFunc represents a function that configures a service registry used by the resolver.

```go
type ResolverOptionsFunc func(registry ServiceRegistry) error
```

<a name="ServiceKey"></a>
## type [ServiceKey](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L24-L26>)

ServiceKey represents a unique key for identifying services in the service registry.

```go
type ServiceKey struct {
    // contains filtered or unexported fields
}
```

<a name="NewServiceKey"></a>
### func [NewServiceKey](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L34>)

```go
func NewServiceKey(value string) ServiceKey
```

NewServiceKey creates a new ServiceKey with the given value.

<a name="ServiceKey.String"></a>
### func \(ServiceKey\) [String](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L29>)

```go
func (s ServiceKey) String() string
```

String Gets the value of the current ServiceKey instance.

<a name="ServiceRegistration"></a>
## type [ServiceRegistration](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L93-L112>)

ServiceRegistration represents a service registrations.

```go
type ServiceRegistration interface {

    // Id Returns the unique identifier of the service registration.
    Id() uint64

    // InvokeActivator calls the activator function with the provided parameters and returns the resulting instance and any error.
    InvokeActivator(params ...interface{}) (interface{}, error)

    // IsSame checks if the provided ServiceRegistration equals the current ServiceRegistration.
    IsSame(other ServiceRegistration) bool

    // LifetimeScope returns the LifetimeScope associated with the service registration.
    LifetimeScope() LifetimeScope

    // RequiredServiceTypes returns a slice of ServiceType, containing all service types required by the service registration.
    RequiredServiceTypes() []ServiceType

    // ServiceType retrieves the type of the service being registered.
    ServiceType() ServiceType
}
```

<a name="ServiceRegistrationList"></a>
## type [ServiceRegistrationList](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L115-L129>)

ServiceRegistrationList provides functionality to manage a list of service registrations. This interface supports internal infrastructure services.

```go
type ServiceRegistrationList interface {

    // AddRegistration adds a new service registration to the list.
    AddRegistration(registration ServiceRegistrationSetup) error

    // Id returns the unique identifier of the service registration list.
    Id() uint64

    // Registrations returns a slice of ServiceRegistration, containing all registrations in the list.
    Registrations() []ServiceRegistration

    // IsEmpty checks if the service registration list contains any registrations.
    // It returns true if the list is empty, otherwise false.
    IsEmpty() bool
}
```

<a name="ServiceRegistrationSetup"></a>
## type [ServiceRegistrationSetup](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L132-L137>)

ServiceRegistrationSetup extends ServiceRegistration and supports internal infrastructure services.

```go
type ServiceRegistrationSetup interface {
    ServiceRegistration

    // SetId sets the unique identifier for the service registration. This method supports internal infrastructure and is not intended to be used by your code.
    SetId(id uint64) error
}
```

<a name="ServiceRegistry"></a>
## type [ServiceRegistry](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L55-L75>)

ServiceRegistry provides methods to map service types to activator functions. The service registration organizes and stores the metadata required by the service resolver.

```go
type ServiceRegistry interface {
    ServiceRegistryAccessor

    // CreateLinkedRegistry creates and returns a new ServiceRegistry instance linked to the current registry. A linked service registry is an empty service registry.
    CreateLinkedRegistry() ServiceRegistry

    // CreateScope creates and returns a scoped ServiceRegistry instance which inherits all service registrations from the current ServiceRegistry instance.
    CreateScope() ServiceRegistry

    // GetServiceRegistrations retrieves all service registrations.
    GetServiceRegistrations() ([]ServiceRegistration, error)

    // IsRegistered checks if a service of the specified ServiceType is registered in the service registry.
    IsRegistered(serviceType ServiceType) bool

    // Register registers a service with its activator function and defines its lifetime scope with the service registry.
    Register(activatorFunc any, scope LifetimeScope) error

    // RegisterModule registers one or more modules, encapsulated as ModuleFunc, with the service registry. A module is a logical unit of service registrations.
    RegisterModule(modules ...ModuleFunc) error
}
```

<a name="ServiceRegistryAccessor"></a>
## type [ServiceRegistryAccessor](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L81-L90>)

ServiceRegistryAccessor provides methods to access and retrieve service registrations from the registry.

```go
type ServiceRegistryAccessor interface {

    // TryGetServiceRegistrations attempts to retrieve all service registrations for the given service type.
    // Returns the service registration list and true if found, otherwise returns false.
    TryGetServiceRegistrations(serviceType ServiceType) (ServiceRegistrationList, bool)

    // TryGetSingleServiceRegistration attempts to retrieve a single service registration for the given service type.
    // Returns the service registration and true if found, otherwise returns false.
    TryGetSingleServiceRegistration(serviceType ServiceType) (ServiceRegistration, bool)
}
```

<a name="ServiceType"></a>
## type [ServiceType](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/types.go#L39-L52>)

ServiceType represents a service type.

```go
type ServiceType interface {

    // Name returns the name of the service type.
    Name() string

    // PackagePath returns the package path of the service type.
    PackagePath() string

    // ReflectedType returns the underlying reflect.Type representation of the service type.
    ReflectedType() reflect.Type

    // LookupKey retrieves the ServiceKey associated with the service type.
    LookupKey() ServiceKey
}
```

<a name="MakeServiceType"></a>
### func [MakeServiceType](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/service_type.go#L44>)

```go
func MakeServiceType[T any]() ServiceType
```

MakeServiceType creates a ServiceType instance for the specified generic type T.

<a name="ServiceTypeFrom"></a>
### func [ServiceTypeFrom](<https://github.com/matzefriedrich/parsley/blob/main/pkg/types/service_type.go#L51>)

```go
func ServiceTypeFrom(t reflect.Type) ServiceType
```

ServiceTypeFrom creates a ServiceType from the given reflect.Type. Supports pointer, interface, function, slice, and struct types. The function panics, if t is of an unsupported kind is given.

