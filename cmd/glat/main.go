package main

import (
	"log"

	"github.com/evercyan/brick/cmd/glat/internal"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "glat",
		Short:   "glat: git log analysis toolkit",
		Version: "v0.0.1",
	}
	app.AddCommand(internal.HtmlCommand)
	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
