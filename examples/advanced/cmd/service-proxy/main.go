package main

import (
	"context"
	"log"

	"github.com/matzefriedrich/parsley-docs/examples/advanced/internal"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
)

func main() {

	registry := registration.NewServiceRegistry()
	registry.Register(internal.NewGreeter, types.LifetimeTransient)
	registry.Register(internal.NewGreeterProxyImpl, types.LifetimeTransient)

	features.RegisterList[features.MethodInterceptor](context.Background(), registry)
	registry.Register(newLoggingInterceptor, types.LifetimeSingleton)

	resolver := resolving.NewResolver(registry)
	ctx := resolving.NewScopedContext(context.Background())
	proxy, _ := resolving.ResolveRequiredService[internal.GreeterProxy](ctx, resolver)

	proxy.SayHello("John", true)
}

type loggingInterceptor struct {
	features.InterceptorBase
}

func (l loggingInterceptor) Enter(_ any, methodName string, parameters []features.ParameterInfo) {
	log.Default().Printf("Enter: %s, Parameters: %v\n", methodName, parameters)
}

func (l loggingInterceptor) Exit(_ any, methodName string, returnValues []features.ReturnValueInfo) {
	log.Default().Printf("Exit: %s\n", methodName)
}

func (l loggingInterceptor) OnError(_ any, methodName string, err error) {
	log.Default().Printf("OnError: %s, MethodName: %s, Error: %v", methodName, methodName, err)
}

var _ features.MethodInterceptor = &loggingInterceptor{}

func newLoggingInterceptor() features.MethodInterceptor {
	return &loggingInterceptor{
		InterceptorBase: features.NewInterceptorBase("logging", 0),
	}
}
