package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type iamProjectsListOpts struct {
	*globalOpts
	store store.ProjectLister
}

func (opts *iamProjectsListOpts) init() error {
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

func (opts *iamProjectsListOpts) Run() error {
	projects, err := opts.store.GetAllProjects()

	if err != nil {
		return err
	}

	return prettyJSON(projects)
}

// mcli iam project(s) list [--orgId orgId]
func IAMProjectsListBuilder() *cobra.Command {
	opts := &iamProjectsListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
