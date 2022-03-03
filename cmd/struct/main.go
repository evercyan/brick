package main

import (
	"github.com/evercyan/brick/cmd/struct/internal"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:     "struct",
		Short:   "Golang Struct Toolkit",
		Version: "v0.0.2",
	}

	root.AddCommand(internal.GormCommand)
	root.AddCommand(internal.CommonCommand)
	root.AddCommand(internal.EnumCommand)
	root.AddCommand(internal.SqlCommand)

	root.PersistentFlags().BoolVarP(
		&internal.FlagJSONUseSnake, "snake", "s", true, "JSON field with snake",
	)
	root.PersistentFlags().BoolVarP(
		&internal.FlagComment, "comment", "c", false, "export comment tag",
	)

	root.Execute()
}
