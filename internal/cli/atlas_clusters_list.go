package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

type AtlasClustersListOpts struct {
	profile      string
	projectID    string
	pageNum      int
	itemsPerPage int
	config       config.Config
	store        store.ClusterLister
}

func (opts *AtlasClustersListOpts) Run() error {
	listOpts := opts.newListOptions()
	result, err := opts.store.ProjectClusters(opts.projectID, listOpts)

	if err != nil {
		return err
	}

	return prettyJSON(result)
}

func (opts *AtlasClustersListOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// mcli atlas cluster(s) list --projectId projectId [--page N] [--limit N]
func AtlasClustersListBuilder() *cobra.Command {
	opts := new(AtlasClustersListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Command to list Atlas clusters",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
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
	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, "Page number")
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, "Items per page")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.ProjectID)

	return cmd
}
