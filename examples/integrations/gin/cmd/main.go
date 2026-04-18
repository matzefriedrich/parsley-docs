package main

import (
	"context"

	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal"
	"github.com/matzefriedrich/parsley-docs/examples/integrations/gin/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {

	ctx := context.Background()

	err := bootstrap.RunParsleyApplication(ctx, internal.NewApp,
		modules.ConfigureGin,
		modules.ConfigureGreeter)

	if err != nil {
		panic(err)
	}
}
