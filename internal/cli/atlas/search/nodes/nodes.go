package nodes

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "nodes"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage Atlas Search nodes for your cluster.",
	}
	cmd.AddCommand(
		ListBuilder(),
	)
	return cmd
}
