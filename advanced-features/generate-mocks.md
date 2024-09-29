---
meta:
  title: Parsley CLI - Mocking Made Easy
description: This article introduces the new `generate mocks` CLI command in Parsley, a feature that simplifies the creation of mock implementations for service interfaces in Go. You’ll learn how to generate mocks quickly, configure their behavior, and assert expectations for method calls within your test suites.
icon: file
label: Generate Configurable Mocks
tags: [ service mocks, testing, code generation, mocking, CLI ]
category:
  - Parsley CLI
  - Mocking
  - Automated testing
---
# Mocking Made Easy with Parsley

In [version v0.9](https://github.com/matzefriedrich/parsley/releases/tag/v0.9.0) of Parsley, a powerful new feature is introduced: the `generate mocks` CLI command. This command generates configurable mock implementations from your service interfaces, making testing faster and more efficient.

In this article, you learn how to use this feature to generate mock objects for interfaces, configure their behavior, and assert expectations for method calls in tests. It also showcases a real-world example of integrating these mocks into your Go test suites.

## Why Mocks?

Mock objects are an essential part of writing unit tests. They allow you to replace the real implementations of your services with objects that simulate their behavior, giving you more control over your test environment. With Parsley’s new `generate mocks` command, you can generate mocks that trace method calls, verify parameters, and count invocations without writing additional boilerplate code.

## Defining a Service Interface

Let's begin with a simple service interface, `Greeter`, which defines two methods: `SayHello` and `SayNothing`.

```go
package features

//go:generate parsley-cli generate mocks

type Greeter interface {
    SayHello(name string) (string, error)
    SayNothing()
}
```

By adding the `//go:generate` directive at the top, we instruct Parsley’s CLI to generate a mock implementation for this interface.

### The Generated Mock Code

After running the `parsley-cli generate mocks` command, Parsley generates the following mock implementation code. A `mock.g.go` file is automatically created, and any manual changes will be overwritten if the command is rerun.

```go
package features

import (
    "github.com/matzefriedrich/parsley/pkg/features"
)

type greeterMock struct {
    features.MockBase
    SayHelloFunc   SayHelloFunc
    SayNothingFunc SayNothingFunc
}

type SayHelloFunc func(name string) (string, error)
type SayNothingFunc func()

const (
    FunctionSayHello   = "SayHello"
    FunctionSayNothing = "SayNothing"
)

func (m *greeterMock) SayHello(name string) (string, error) {
    m.TraceMethodCall(FunctionSayHello, name)
    return m.SayHelloFunc(name)
}

func (m *greeterMock) SayNothing() {
    m.TraceMethodCall(FunctionSayNothing)
    m.SayNothingFunc()
}

var _ Greeter = (*greeterMock)(nil)

// NewGreeterMock Creates a new configurable greeterMock object.
func NewGreeterMock() *greeterMock {
    mock := &greeterMock{
        MockBase: features.NewMockBase(),
        SayHelloFunc: func(name string) (string, error) {
            var result0 string
            var result1 error
            return result0, result1
        },
        SayNothingFunc: func() {},
    }
    mock.AddFunction(FunctionSayHello, "SayHello(name string) (string, error)")
    mock.AddFunction(FunctionSayNothing, "SayNothing()")
    return mock
}
```

## Configuring Mocks in Tests

Once the mock has been generated, it can be used in your tests as a drop-in replacement for the actual service. You can configure its behavior by overriding the functions for each method in the interface.

Here’s an example of how to test the `GreeterMock` type by configuring the behavior of `SayHello` and using Parsley’s `Verify` assertion helpers:

```go
package features

import (
    "fmt"
    "github.com/matzefriedrich/parsley/pkg/features"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_GreeterMock_SayHello(t *testing.T) {
    // Arrange
    mock := NewGreeterMock()
    mock.SayHelloFunc = func(name string) (string, error) {
        return fmt.Sprintf("Hi, %s", name), nil
    }

    const expectedName = "John"

    // Act
    _, _ = mock.SayHello("Max")
    actual, err := mock.SayHello(expectedName)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Hi, John", actual)

    assert.True(t, mock.Verify(FunctionSayHello, features.TimesOnce(), features.Exact(expectedName)))
    assert.True(t, mock.Verify(FunctionSayHello, features.TimesNever(), features.Exact("Jane")))
    assert.True(t, mock.Verify(FunctionSayHello, features.TimesExactly(2)))
}
```

## Verifying Method Calls (and Arguments)

The generated mocks automatically trace method calls and allow you to verify how often methods were invoked and with which arguments. Parsley provides a set of powerful assertion helpers to verify method calls:

### Counter verification functions

- **TimesOnce**: Verifies that a method was called exactly once.
- **TimesNever**: Verifies that a method was never called.
- **TimesExactly**: Verifies that a method was called exactly `n` times.

### Argument matching

- **Exact**: Matches arguments exactly.
- **IsAny**: Matches any given argument - use this as a placeholder.


For example, in the test case above:
- The mock verifies that `SayHello` was called **once** with the exact argument `"John"`.

- The mock checks that `SayHello` was **never** called with the argument `"Jane"`.

- Lastly, it verifies that `SayHello` was called **exactly twice** in total.

> **Note:** You can also provide custom `TimesFunc` and `ArgMatch` callbacks to evaluate counter values and method call parameters.


## Conclusion

With the new `generate mocks` command, Parsley makes it easy to create fully configurable and traceable mock objects for your services.

The mock generation feature is especially useful for writing tests, where you want to simulate different service behaviors, verify method calls, and ensure your components interact correctly with each other.

### Using Parsley’s generator commands without Runtime Dependency Injection

Parsley’s generator commands, like the `generate mocks` command, can be used independently of Parsley’s runtime dependency injection. This flexibility allows developers to leverage powerful code generation features, such as creating mock implementations for interfaces, without adopting Parsley’s full DI system.

For instance, you may prefer to manually wire your dependencies in a traditional Go setup while still using Parsley’s mock generation capabilities for testing. This makes Parsley a versatile tool that can be integrated into various workflows, whether or not you choose to use runtime DI for your projects.