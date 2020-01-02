package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type AtlasClustersDescribeOpts struct {
	profile   string
	projectID string
	name      string
	config    config.Config
	store     store.ClusterDescriber
}

func (opts *AtlasClustersDescribeOpts) Run() error {
	result, err := opts.store.Cluster(opts.projectID, opts.name)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

// mcli atlas cluster(s) describe --projectId projectId
func AtlasClustersDescribeBuilder() *cobra.Command {
	opts := new(AtlasClustersDescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe [name]",
		Short: "Command to describe an Atlas cluster",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.config = config.New(opts.profile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := store.New(opts.config)

			if err != nil {
				return err
			}

			opts.store = s
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.ProjectID)

	return cmd
}
