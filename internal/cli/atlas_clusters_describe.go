package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type atlasClustersDescribeOpts struct {
	*globalOpts
	name  string
	store store.ClusterDescriber
}

func (opts *atlasClustersDescribeOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasClustersDescribeOpts) Run() error {
	result, err := opts.store.Cluster(opts.ProjectID(), opts.name)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

// mcli atlas cluster(s) describe --projectId projectId
func AtlasClustersDescribeBuilder() *cobra.Command {
	opts := &atlasClustersDescribeOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "describe [name]",
		Short: "Command to describe an Atlas cluster",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
