package cli

import (
	"github.com/spf13/cobra"
)

func IAMProjectsBuilder() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "projects",
		Short:   "Projects operations",
		Long:    "Create, list and manage your MongoDB Cloud projects.",
		Aliases: []string{"project"},
	}
	cmd.AddCommand(IAMProjectsListBuilder())
	cmd.AddCommand(IAMProjectsCreateBuilder())
	return cmd
}
