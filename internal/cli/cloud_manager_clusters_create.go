package cli

import (
	"fmt"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/convert"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type cmClustersCreateOpts struct {
	*globalOpts
	filename string
	fs       afero.Fs
	store    store.AutomationStore
}

func (opts *cmClustersCreateOpts) init() error {
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

func (opts *cmClustersCreateOpts) Run() error {
	newConfig, err := convert.ReadInClusterConfig(opts.fs, opts.filename)
	if err != nil {
		return err
	}
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	for _, rs := range current.ReplicaSets {
		if rs.ID == newConfig.Name {
			return fmt.Errorf("cluster %s already exists", newConfig.Name)
		}

	}

	err = newConfig.PatchReplicaSet(current)

	if err != nil {
		return err
	}

	if _, err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Printf("Changes are being applied, please check %s/v2/%s#deployment/topology for status\n", opts.OpsManagerURL(), opts.ProjectID())

	return nil
}

// mcli cloud-manager cluster(s) create --projectId projectId --file myfile.yaml
func CloudManagerClustersCreateBuilder() *cobra.Command {
	opts := &cmClustersCreateOpts{
		globalOpts: newGlobalOpts(),
		fs:         afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Command to create a Cloud Manager cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().StringVar(&opts.filename, flags.File, "", "File with config")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	_ = cmd.MarkFlagRequired(flags.File)

	return cmd
}
