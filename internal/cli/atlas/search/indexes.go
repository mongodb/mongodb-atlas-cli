package search

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

func IndexesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "indexes",
		Aliases: []string{"index"},
		Short:   description.Indexes,
	}
	cmd.AddCommand(ListBuilder())
	cmd.AddCommand(CreateBuilder())
	cmd.AddCommand(DeleteBuilder())

	return cmd
}
