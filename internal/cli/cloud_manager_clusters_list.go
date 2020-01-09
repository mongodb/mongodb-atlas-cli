package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/utils"
	"github.com/spf13/cobra"
)

type cmClustersListOpts struct {
	*globalOpts
	store store.AutomationGetter
}

func (opts *cmClustersListOpts) init() error {
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

func (opts *cmClustersListOpts) Run() error {
	result, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	return utils.PrettyJSON(result)
}

// mcli cloud-manager cluster(s) list --projectId projectId
func CloudManagerClustersListBuilder() *cobra.Command {
	opts := &cmClustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Command to list a Cloud Manager cluster",
		Args:    cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
