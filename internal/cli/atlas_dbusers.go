package cli

import (
	"github.com/spf13/cobra"
)

func AtlasDBUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dbusers",
		Aliases: []string{"dbuser", "databaseUsers", "databaseUser"},
		Short:   "Command for working with atlas database users",
	}
	cmd.AddCommand(AtlasDBUsersCreateBuilder())

	return cmd
}
