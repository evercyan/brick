package main

import (
	"github.com/evercyan/brick/cmd/struct/internal"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:     "struct",
		Short:   "struct: Golang struct toolkit",
		Version: "v0.0.1",
	}

	root.AddCommand(internal.GormCommand)
	root.AddCommand(internal.CommonCommand)
	root.AddCommand(internal.EnumCommand)

	root.PersistentFlags().BoolVarP(
		&internal.FlagJSONUseSnake, "snake", "s", false, "JSON field with snake",
	)

	root.Execute()
}
