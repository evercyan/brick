package main

import (
	"github.com/evercyan/brick/cmd/money/command"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:     "money",
		Short:   "money: Make More Money",
		Version: "v0.0.1",
	}
	root.AddCommand(command.InterestCommand)
	root.AddCommand(command.TaxCommand)
	root.Execute()
}
