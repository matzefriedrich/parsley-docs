package modules

import (
	"context"
	"github.com/matzefriedrich/cobra-extensions/pkg/charmer"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/cobra/internal/commands"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/resolving"
	"github.com/matzefriedrich/parsley/pkg/types"
	"github.com/spf13/cobra"
)

// CobraApplicationModule creates a module for a command-line application using Cobra with the given name and description.
func CobraApplicationModule(appName string, appDescription string) func(registry types.ServiceRegistry) error {
	return func(registry types.ServiceRegistry) error {
		_ = features.RegisterList[*cobra.Command](registry)
		_ = registration.RegisterSingleton(registry, func(resolver types.Resolver) *charmer.CommandLineApplication {
			application := charmer.NewCommandLineApplication(appName, appDescription)
			commands, _ := resolving.ResolveRequiredServices[*cobra.Command](resolver, context.Background())
			for _, command := range commands {
				application.AddCommand(command)
			}
			return application
		})
		return nil
	}
}

// CobraApplicationCommandsModule registers typed command handler services in the given service registry.
func CobraApplicationCommandsModule(registry types.ServiceRegistry) error {
	_ = registration.RegisterTransient(registry, commands.NewHelloCommand)
	return nil
}
