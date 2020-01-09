package cli

import (
	"fmt"

	"github.com/10gen/mcli/internal/utils"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type cmClustersDescribeOpts struct {
	*globalOpts
	name  string
	store store.AutomationGetter
}

func (opts *cmClustersDescribeOpts) init() error {
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

func (opts *cmClustersDescribeOpts) Run() error {
	result, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	for _, rs := range result.ReplicaSets {
		if rs.ID == opts.name {
			return utils.PrettyJSON(rs)
		}

	}
	return fmt.Errorf("replicaset %s not found", opts.name)
}

// mcli cloud-manager cluster(s) describe [name] --projectId projectId
func CloudManagerClustersDescribeBuilder() *cobra.Command {
	opts := &cmClustersDescribeOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "describe [name]",
		Short: "Command to describe a Cloud Manager cluster",
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
