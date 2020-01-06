package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

type atlasClustersListOpts struct {
	*globalOpts
	pageNum      int
	itemsPerPage int
	store        store.ClusterLister
}

func (opts *atlasClustersListOpts) init() error {
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

func (opts *atlasClustersListOpts) Run() error {
	listOpts := opts.newListOptions()
	result, err := opts.store.ProjectClusters(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *atlasClustersListOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// mcli atlas cluster(s) list --projectId projectId [--page N] [--limit N]
func AtlasClustersListBuilder() *cobra.Command {
	opts := &atlasClustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Command to list Atlas clusters",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, "Page number")
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, "Items per page")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
