package internal

import (
	"github.com/spf13/cobra"
)

var (
	// HtmlCommand ...
	HtmlCommand = &cobra.Command{
		Use:   "html",
		Short: "Generate git log analysis html",
		Run: func(cmd *cobra.Command, args []string) {
			g.Html()
		},
	}
)
