package main

import (
	"github.com/evercyan/brick/cmd/fimg/internal"
	"github.com/spf13/cobra"
)

// grm join -d ~/Downloads/ss
func main() {
	root := &cobra.Command{
		Use:     "fimg",
		Short:   "fimg: image toolkit",
		Version: "v0.0.1",
	}

	root.AddCommand(internal.JoinCommand)

	root.Execute()
}
