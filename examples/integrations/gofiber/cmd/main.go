package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {
	context := context.Background()
	bootstrap.RunParsleyApplication(context, internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureGreeter)
}