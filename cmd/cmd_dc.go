package cmd

import (
	"github.com/spf13/cobra"
)

var (
	dcCmd = &cobra.Command{
		Use:   "dc",
		Short: "Returns list of all discovarable datacenters",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			GetDatacenters()
		},
	}
)

func init() {
	CexCmd.AddCommand(dcCmd)
}
