package cli

import (
	"github.com/spf13/cobra"
)

func ProjectBuilder() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "projects",
		Short:   "Projects operations",
		Long:    "Create, list and manage your MongoDB Cloud projects.",
		Aliases: []string{"project"},
	}
	cmd.AddCommand(ProjectsListBuilder())
	cmd.AddCommand(ProjectCreateBuilder())
	return cmd
}
