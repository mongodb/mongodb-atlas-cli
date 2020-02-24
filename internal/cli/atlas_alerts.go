package cli

import (
	"github.com/spf13/cobra"
)

func AtlasAlertBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerts",
		Aliases: []string{"alert"},
		Short:   "Manage Atlas alert for your project.",
		Long:    "The alertConfigs command provides access to your alerts configurations. You can create, edit, and delete alert configurations.",
	}

	cmd.AddCommand(AtlasAlertConfigBuilder())

	return cmd
}
