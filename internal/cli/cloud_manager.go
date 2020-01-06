package cli

import (
	"github.com/spf13/cobra"
)

func CloudManagerBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-manager",
		Aliases: []string{"cm", "ops-manager", "om", "deployments"},
		Short:   "Command for working with atlas",
	}

	cmd.AddCommand(CloudManagerClustersBuilder())

	return cmd
}
