package cli

import (
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type ListProjectOpts struct {
	store store.ProjectLister
}

func (opts *ListProjectOpts) Run() error {
	projects, err := opts.store.GetAllProjects()

	if err != nil {
		return err
	}

	err = prettyJSON(projects)

	if err != nil {
		return err
	}

	return nil
}

func ProjectsListBuilder() *cobra.Command {
	opts := &ListProjectOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := NewConfig()
			s, err := store.New(config.GetService(),
				config.GetPublicAPIKey(),
				config.GetPrivateAPIKey(),
				config.GetOpsManagerURL())

			if err != nil {
				return err
			}

			opts.store = s
			return opts.Run()
		},
	}
	return cmd
}
