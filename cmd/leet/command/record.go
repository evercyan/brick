package command

import (
	"github.com/spf13/cobra"
)

var (
	// RecordCommand ...
	RecordCommand = &cobra.Command{
		Use:   "record",
		Short: "生成答题记录文件 RECORD.md",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
