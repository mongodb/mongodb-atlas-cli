package cli

import (
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/spf13/cobra"
)

type opsManagerClusterListAll struct {
	store        store.ListAllClusters
}


func (opts *opsManagerClusterListAll) init() error {
	var err error
	opts.store, err = store.New()
	return err
}


func (opts *opsManagerClusterListAll) Run() error {
	result, err := opts.store.ListAllClusters()
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

func OpsManagerListAllClustersBuilder() *cobra.Command {
	opts := new(opsManagerClusterListAll)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List All clusters",
		Args:  cobra.OnlyValidArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	return cmd
}