package teams

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const updateTemplate = "Team's roles updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store  store.TeamRolesUpdater
	teamID string
	roles  []string
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateProjectTeamRoles(opts.ConfigProjectID(), opts.teamID, opts.newTeamUpdateRoles())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newTeamUpdateRoles() *atlas.TeamUpdateRoles {
	return &atlas.TeamUpdateRoles{
		RoleNames: opts.roles,
	}
}

// mongocli iam team(s) user(s) updates teamId --projectId projectId --role role
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update <teamId>",
		Aliases: []string{"updates"},
		Args:    cobra.ExactArgs(1),
		Short:   updateTeamRoles,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.teamID = args[0]
			return opts.PreRunE(
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.TeamRole)

	cmd.Flags().StringVar(&opts.OrgID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
