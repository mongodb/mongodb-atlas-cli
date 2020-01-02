package cli

import (
	"github.com/spf13/cobra"
)

func AtlasClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clusters",
		Aliases: []string{"cluster"},
		Short:   "Command for working with atlas clusters",
	}
	cmd.AddCommand(AtlasClustersCreateBuilder())
	cmd.AddCommand(AtlasClustersListBuilder())

	return cmd
}
