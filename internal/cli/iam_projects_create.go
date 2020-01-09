package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/utils"
	"github.com/spf13/cobra"
)

type iamProjectsCreateOpts struct {
	*globalOpts
	orgID string
	name  string
	store store.ProjectCreator
}

func (opts *iamProjectsCreateOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *iamProjectsCreateOpts) Run() error {
	projects, err := opts.store.CreateProject(opts.name, opts.orgID)

	if err != nil {
		return err
	}

	return utils.PrettyJSON(projects)
}

// mcli iam project(s) create name [--orgId orgId]
func IAMProjectsCreateBuilder() *cobra.Command {
	opts := &iamProjectsCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a project",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.orgID, flags.OrgID, "", "Organization ID for the project")
	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
