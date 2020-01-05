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

type atlasDBUsersCreateOpts struct {
	*atlasOpts
	username string
	password string
	roles    []string
	store    store.DatabaseUserCreator
}

func (opts *atlasDBUsersCreateOpts) init() error {
	opts.loadConfig()

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

func (opts *atlasDBUsersCreateOpts) Run() error {
	user := opts.newDatabaseUser()
	result, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *atlasDBUsersCreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		DatabaseName: adminDB,
		Roles:        opts.buildRoles(),
		GroupID:      opts.ProjectID(),
		Username:     opts.username,
		Password:     opts.password,
	}
}

func (opts *atlasDBUsersCreateOpts) buildRoles() []atlas.Role {
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
	opts := &atlasDBUsersCreateOpts{
		atlasOpts: newAtlasOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Command to create a cluster with Atlas",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.username, flags.Username, "", "Username")
	cmd.Flags().StringVar(&opts.password, flags.Password, "", "Password")
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, "Role")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.Username)
	_ = cmd.MarkFlagRequired(flags.Password)
	_ = cmd.MarkFlagRequired(flags.Role)

	return cmd
}
