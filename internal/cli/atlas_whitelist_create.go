package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

const (
	cidrBlock = "cidrBlock"
	ipAddress = "ipAddress"
)

type atlasWhitelistCreateOpts struct {
	*globalOpts
	entry     string
	entryType string
	comment   string
	store     store.ProjectIPWhitelistCreator
}

func (opts *atlasWhitelistCreateOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

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

func (opts *atlasWhitelistCreateOpts) Run() error {
	entry := opts.newWhitelist()
	result, err := opts.store.CreateProjectIPWhitelist(entry)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *atlasWhitelistCreateOpts) newWhitelist() *atlas.ProjectIPWhitelist {
	projectIPWhitelist := &atlas.ProjectIPWhitelist{
		GroupID: opts.ProjectID(),
		Comment: opts.comment,
	}
	switch opts.entryType {
	case cidrBlock:
		projectIPWhitelist.CIDRBlock = opts.entry
	case ipAddress:
		projectIPWhitelist.IPAddress = opts.entry
	}
	return projectIPWhitelist
}

// mcli atlas whitelist(s) create value --type cidrBlock|ipAddress [--comment comment] [--projectId projectId]
func AtlasWhitelistCreateBuilder() *cobra.Command {
	opts := &atlasWhitelistCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create [entry]",
		Short: "Command to create a cluster with Atlas",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.entry = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.entryType, flags.Type, ipAddress, "Type of entry, cidrBlock, or ipAddress")
	cmd.Flags().StringVar(&opts.comment, flags.Comment, "", "Optional comment")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
