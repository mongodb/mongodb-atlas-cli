package cli

import (
	"strings"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

const (
	adminDB = "admin"
	roleSep = "@"
)

type AtlasDBUsersCreateOpts struct {
	profile   string
	projectID string
	username  string
	password  string
	roles     []string
	config    config.Config
	store     store.DatabaseUserCreator
}

func (opts *AtlasDBUsersCreateOpts) Run() error {
	user := opts.newDatabaseUser()
	result, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *AtlasDBUsersCreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		DatabaseName: adminDB,
		Roles:        opts.buildRoles(),
		GroupID:      opts.projectID,
		Username:     opts.username,
		Password:     opts.password,
	}
}

func (opts *AtlasDBUsersCreateOpts) buildRoles() []atlas.Role {
	rolesLen := len(opts.roles)
	roles := make([]atlas.Role, rolesLen)
	for i, roleP := range opts.roles {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := adminDB
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}

// mcli atlas dbuser(s) create --username username --password password --role roleName@dbName [--projectId projectId]
func AtlasDBUsersCreateBuilder() *cobra.Command {
	opts := new(AtlasDBUsersCreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Command to create a cluster with Atlas",
		Args:  cobra.ExactArgs(0),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.config = config.New(opts.profile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := store.New(opts.config)

			if err != nil {
				return err
			}

			opts.store = s
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.username, flags.Username, "", "Username")
	cmd.Flags().StringVar(&opts.password, flags.Password, "", "Password")
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, "Role")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.ProjectID)
	_ = cmd.MarkFlagRequired(flags.Username)
	_ = cmd.MarkFlagRequired(flags.Password)
	_ = cmd.MarkFlagRequired(flags.Role)

	return cmd
}
