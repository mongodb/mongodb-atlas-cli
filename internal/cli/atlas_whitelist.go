package cli

import (
	"github.com/spf13/cobra"
)

func AtlasWhitelistBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Command for working with atlas whitelist",
	}
	cmd.AddCommand(AtlasWhitelistCreateBuilder())

	return cmd
}
