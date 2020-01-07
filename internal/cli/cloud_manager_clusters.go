package cli

import (
	"github.com/spf13/cobra"
)

func CloudManagerClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clusters",
		Aliases: []string{"cluster"},
		Short:   "Command for working with cloud manager clusters",
	}

	cmd.AddCommand(CloudManagerClustersListBuilder())
	cmd.AddCommand(CloudManagerClustersDescribeBuilder())

	return cmd
}
