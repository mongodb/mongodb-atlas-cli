package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type ListProjectOpts struct {
	profile string
	store   store.ProjectLister
}

func (opts *ListProjectOpts) Run() error {
	projects, err := opts.store.GetAllProjects()

	if err != nil {
		return err
	}

	return prettyJSON(projects)
}

func ProjectsListBuilder() *cobra.Command {
	opts := new(ListProjectOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := config.New(opts.profile)
			s, err := store.New(conf)

			if err != nil {
				return err
			}

			opts.store = s
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
