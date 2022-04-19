package main

import (
	"github.com/evercyan/brick/cmd/leet/command"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:     "leet",
		Short:   "leet: LeetCode Toolkit",
		Version: "v0.0.1",
	}
	root.AddCommand(command.ConfigCommand)
	root.AddCommand(command.ListCommand)
	root.AddCommand(command.QuestionCommand)
	root.AddCommand(command.RecordCommand)
	root.AddCommand(command.TagCommand)
	root.Execute()
}
