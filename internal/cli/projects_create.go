package cli

import (
	"errors"

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
	opts := CreateProjectOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("requires a name argument")
			}
			return nil
		},
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
