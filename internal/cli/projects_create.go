package cli

import (
	"errors"

	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/cobra"
)

type CreateProjectOpts struct {
	orgID string
	name  string
	store store.ProjectCreator
}

func (opts *CreateProjectOpts) Run() error {
	projects, err := opts.store.CreateProject(opts.name, opts.orgID)

	if err != nil {
		return err
	}

	err = prettyJSON(projects)

	if err != nil {
		return err
	}

	return nil
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
			config := NewConfig()
			s, err := store.New(config.GetService(),
				config.GetPublicAPIKey(),
				config.GetPrivateAPIKey(),
				config.GetOpsManagerURL())

			if err != nil {
				return err
			}

			opts.store = s
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.orgID, "orgId", "", "Organization ID for the project")
	return cmd
}
