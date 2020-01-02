package cli

import (
	"github.com/spf13/cobra"
)

func IAMBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iam",
		Short: "Command for working with authentication",
	}
	cmd.AddCommand(IAMProjectsBuilder())

	return cmd
}
