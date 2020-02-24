package cli

import (
	"github.com/spf13/cobra"
)

func AtlasAlertsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerts",
		Aliases: []string{"alert"},
		Short:   "Manage Atlas alerts for your project.",
	}

	cmd.AddCommand(AtlasAlertConfigsBuilder())

	return cmd
}
