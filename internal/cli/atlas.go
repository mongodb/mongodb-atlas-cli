package cli

import (
	"sync"

	"github.com/10gen/mcli/internal/config"
	"github.com/spf13/cobra"
)

type atlasOpts struct {
	config.Config
	profile   string
	projectID string
	once      sync.Once
}

// newAtlasOpts returns an atlasOpts
func newAtlasOpts() *atlasOpts {
	return new(atlasOpts)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *atlasOpts) ProjectID() string {
	opts.loadConfig()
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = opts.Config.ProjectID()
	return opts.projectID
}

func (opts *atlasOpts) loadConfig() {
	opts.once.Do(func() {
		opts.Config = config.New(opts.profile)
	})
}

func AtlasBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atlas",
		Short: "Command for working with atlas",
	}
	cmd.AddCommand(AtlasClustersBuilder())
	cmd.AddCommand(AtlasDBUsersBuilder())
	cmd.AddCommand(AtlasWhitelistBuilder())

	return cmd
}
