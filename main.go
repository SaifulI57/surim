package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SaifulI57/surim/app"
	"github.com/fatih/color"
)

func main() {
	if err := app.New().Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		os.Exit(1)
	}
}
