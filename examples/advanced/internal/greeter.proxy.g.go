// Code generated by parsley-cli; DO NOT EDIT.
//
// This file was automatically generated and any changes to it will be overwritten.
// To extend or modify the behavior of this code, implement the MethodInterceptor interface and provide your custom logic there.

package internal

import (
	"github.com/matzefriedrich/parsley/pkg/features"
)

// GreeterProxyImpl A generated proxy service type for Greeter objects.
type GreeterProxyImpl struct {
	features.ProxyBase
	target Greeter
}

// GreeterProxy An interface type for Greeter objects. Parsley needs this to distinguish the proxy from the actual implementation.
type GreeterProxy interface {
	Greeter
}

// NewGreeterProxyImpl Creates a new GreeterProxy object. Register this constructor method with the registry.
func NewGreeterProxyImpl(target Greeter, interceptors []features.MethodInterceptor) GreeterProxy {
	return &GreeterProxyImpl{
		ProxyBase: features.NewProxyBase(target, interceptors),
		target:    target,
	}
}

func (p *GreeterProxyImpl) SayHello(name string, polite bool) {

	const methodName = "SayHello"
	parameters := map[string]interface{}{
		"name": name,

		"polite": polite,
	}

	callContext := features.NewMethodCallContext(methodName, parameters)
	p.InvokeEnterMethodInterceptors(callContext)
	defer func() {
		p.InvokeExitMethodInterceptors(callContext)
	}()

	p.target.SayHello(name, polite)
}

var _ Greeter = &GreeterProxyImpl{}
