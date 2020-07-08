package onlinearchive

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "onlineArchives",
		Aliases: []string{"onlineArchive", "onlinearchives", "onlinearchive", "online-archives", "online-archive"},
		Short:   description.Events,
	}

	cmd.AddCommand(ListBuilder())

	return cmd
}
