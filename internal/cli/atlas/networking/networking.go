package networking

import (
	"github.com/mongodb/mongocli/internal/cli/atlas/networking/containers"
	"github.com/mongodb/mongocli/internal/cli/atlas/networking/peering"
	"github.com/spf13/cobra"
)

const short = "Networking operations."

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "networking",
		Short: short,
	}
	cmd.AddCommand(containers.Builder())
	cmd.AddCommand(peering.Builder())

	return cmd
}
