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

// CobraApplicationModule registers a Cobra-based command-line application with the given name and description in the service registry.
func CobraApplicationModule(appName string, appDescription string) func(registry types.ServiceRegistry) error {

	return func(registry types.ServiceRegistry) error {

		// Enable resolution of *cobra.Command objects at once (as a list)
		_ = features.RegisterList[*cobra.Command](registry)

		// Register a factory function for *CommandLineApplication
		_ = registration.RegisterSingleton(registry, func(resolver types.Resolver) *charmer.CommandLineApplication {

			application := charmer.NewCommandLineApplication(appName, appDescription)

			// Resolve all command handlers and add them to the app instance
			typeCommandServices, _ := resolving.ResolveRequiredServices[*cobra.Command](resolver, context.Background())
			for _, command := range typeCommandServices {
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
