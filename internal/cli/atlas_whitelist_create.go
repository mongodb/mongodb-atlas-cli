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

type AtlasWhitelistCreateOpts struct {
	profile   string
	projectID string
	entry     string
	entryType string
	comment   string
	config    config.Config
	store     store.ProjectIPWhitelistCreator
}

func (opts *AtlasWhitelistCreateOpts) Run() error {
	entry := opts.newWhitelist()
	result, err := opts.store.CreateProjectIPWhitelist(entry)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *AtlasWhitelistCreateOpts) newWhitelist() *atlas.ProjectIPWhitelist {
	projectIPWhitelist := &atlas.ProjectIPWhitelist{
		GroupID: opts.projectID,
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
	opts := new(AtlasWhitelistCreateOpts)
	cmd := &cobra.Command{
		Use:   "create [entry]",
		Short: "Command to create a cluster with Atlas",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.config = config.New(opts.profile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := store.New(opts.config)

			if err != nil {
				return err
			}

			opts.store = s
			opts.entry = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.entryType, flags.Type, "ipAddress", "Type of entry, cidrBlock, or ipAddress")
	cmd.Flags().StringVar(&opts.comment, flags.Comment, "", "Optional comment")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.ProjectID)

	return cmd
}
