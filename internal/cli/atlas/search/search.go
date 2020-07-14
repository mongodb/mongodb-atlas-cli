package search

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search",
		Aliases: []string{"fts"},
		Short:   description.Search,
	}
	cmd.AddCommand(ListBuilder())
	return cmd
}
