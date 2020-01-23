package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:       "completion [name]",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"bash", "zsh"},
	Hidden:    true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == "bash" {
			return rootCmd.GenBashCompletion(os.Stdout)
		}
		return rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
