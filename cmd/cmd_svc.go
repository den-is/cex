package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// DC lookup specific datacenter
	DC string
	// SERVICE return list of hosts for requested service
	SERVICE string
	// HOSTNAME return hostname instead of IP
	HOSTNAME bool

	svcCmd = &cobra.Command{
		Use:   "svc",
		Short: "Query services",
		Long: `Returns list of services if used without "-s" flag.
Otherwise returns list of hosts for requested service.`,
		Run: func(cmd *cobra.Command, args []string) {

			if SERVICE == "" {
				GetServices()
			} else {
				GetService(SERVICE)
			}

		},
	}
)

func init() {
	svcCmd.Flags().StringVarP(&DC, "dc", "d", "", "lookup specific datacenter")
	svcCmd.Flags().StringVarP(&SERVICE, "service", "s", "", "return list of hosts for requested service")
	svcCmd.Flags().BoolVarP(&HOSTNAME, "hostname", "f", false, "return Node name instead of IP (in most cases that would be FQDN)")
	CexCmd.AddCommand(svcCmd)
}
