package processor

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "processors"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage Atlas Stream Processors.",
		Long:    `Create, list, update and delete your Atlas Stream Processors.`,
	}
	cmd.AddCommand(ListBuilder(), DescribeBuilder(), StartBuilder(), StopBuilder(), DeleteBuilder(), CreateBuilder())

	return cmd
}
