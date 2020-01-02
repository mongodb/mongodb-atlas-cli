package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type CreateProjectOpts struct {
	profile string
	orgID   string
	name    string
	store   store.ProjectCreator
}

func (opts *CreateProjectOpts) Run() error {
	projects, err := opts.store.CreateProject(opts.name, opts.orgID)

	if err != nil {
		return err
	}

	return prettyJSON(projects)
}

// createCmd represents the create command
func ProjectCreateBuilder() *cobra.Command {
	opts := new(CreateProjectOpts)
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := config.New(opts.profile)
			s, err := store.New(conf)

			if err != nil {
				return err
			}

			opts.store = s
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.orgID, flags.OrgID, "", "Organization ID for the project")
	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
