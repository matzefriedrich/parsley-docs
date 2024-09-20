package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gofiber/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {

	ctx := context.Background()

	err := bootstrap.RunParsleyApplication(ctx, internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureGreeter)

	if err != nil {
		panic(err)
	}
}
