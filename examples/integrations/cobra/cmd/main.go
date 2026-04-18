package main

import (
	"context"

	"github.com/matzefriedrich/cobra-extensions/pkg/charmer"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/cobra/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
)

func main() {

	registry := registration.NewServiceRegistry()

	_ = registry.RegisterModule(modules.CobraApplicationModule("cobra-app", "A Parsley-enabled Cobra app"))
	_ = registry.RegisterModule(modules.CobraApplicationCommandsModule)
	_ = registry.RegisterModule(modules.ServicesModule)

	resolver := resolving.NewResolver(registry)

	ctx := context.Background()
	scopedContext := resolving.NewScopedContext(ctx)
	app, _ := resolving.ResolveRequiredService[*charmer.CommandLineApplication](scopedContext, resolver)
	err := app.Execute(ctx)
	if err != nil {
		panic(err)
	}
}
