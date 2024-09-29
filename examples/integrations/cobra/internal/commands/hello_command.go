package commands

import (
	"fmt"
	"github.com/matzefriedrich/cobra-extensions/pkg/commands"
	"github.com/matzefriedrich/cobra-extensions/pkg/types"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/cobra/internal/services"
	"github.com/spf13/cobra"
)

type helloCommand struct {
	use     types.CommandName `flag:"hello" short:"Prints a greeting to the specified name."`
	Name    string            `flag:"name" usage:"The name to greet." required:"true"`
	Polite  bool              `flag:"polite"`
	greeter services.Greeter
}

func (h *helloCommand) Execute() {
	message := h.greeter.SayHello(h.Name, h.Polite)
	fmt.Println(message)
}

var _ types.TypedCommand = (*helloCommand)(nil)

func NewHelloCommand(greeter services.Greeter) *cobra.Command {
	command := &helloCommand{
		greeter: greeter,
	}
	return commands.CreateTypedCommand(command)
}
