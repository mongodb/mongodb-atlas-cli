package cli

import (
	"github.com/spf13/cobra"
)

func AtlasBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atlas",
		Short: "Command for working with atlas",
	}
	cmd.AddCommand(AtlasClustersBuilder())
	cmd.AddCommand(AtlasDBUsersBuilder())

	return cmd
}
