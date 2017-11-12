package cmd

import (
	"github.com/spf13/cobra"
)

// CexCmd - root cex command
var (
	URL  string
	JSON bool

	CexCmd = &cobra.Command{
		Use:   "cex",
		Short: "cex - Consul EXplorer",
	}
)

func init() {
	CexCmd.PersistentFlags().StringVarP(&URL, "url", "u", "http://localhost:8500", "consul url")
	CexCmd.PersistentFlags().BoolVarP(&JSON, "json", "j", false, "JSON output")
}
