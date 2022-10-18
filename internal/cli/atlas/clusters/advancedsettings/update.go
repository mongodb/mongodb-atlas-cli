// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package advancedsettings

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/openlyinc/pointy"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	updateTmpl                   = "Updating advanced configuration settings of your cluster'.\n"
	defaultReadConcern           = "available"
	defaultWriteConcern          = "1"
	defaultSampleSizeBIConnector = 1000
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name                             string
	defaultReadConcern               string
	defaultWriteConcern              string
	minimumEnabledTLSProtocol        string
	disableTableScan                 bool
	enableTableScan                  bool
	disableFailIndexKeyTooLong       bool
	enableFailIndexKeyTooLong        bool
	disableJavascript                bool
	enableJavascript                 bool
	oplogMinRetentionHours           float64
	oplogSizeMB                      int64
	sampleRefreshIntervalBIConnector int64
	sampleSizeBIConnector            int64
	store                            store.AtlasClusterConfigurationOptionsUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateAtlasClusterConfigurationOptions(opts.ConfigProjectID(), opts.name, opts.newProcessArgs())
	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newProcessArgs() *atlas.ProcessArgs {
	args := &atlas.ProcessArgs{}

	if opts.defaultReadConcern != defaultReadConcern {
		args.DefaultReadConcern = opts.defaultReadConcern
	}

	if opts.defaultWriteConcern != defaultWriteConcern {
		args.DefaultWriteConcern = opts.defaultWriteConcern
	}

	if opts.minimumEnabledTLSProtocol != "" {
		args.MinimumEnabledTLSProtocol = opts.minimumEnabledTLSProtocol
	}

	if opts.sampleSizeBIConnector != defaultSampleSizeBIConnector {
		args.SampleSizeBIConnector = pointy.Int64(opts.sampleSizeBIConnector)
	}

	if opts.sampleRefreshIntervalBIConnector != 0 {
		args.SampleRefreshIntervalBIConnector = pointy.Int64(opts.sampleRefreshIntervalBIConnector)
	}

	if opts.disableTableScan {
		args.NoTableScan = pointy.Bool(opts.disableTableScan)
	}

	if opts.enableTableScan {
		args.NoTableScan = pointy.Bool(!opts.enableTableScan)
	}

	if opts.disableJavascript {
		args.JavascriptEnabled = pointy.Bool(false)
	}

	if opts.enableJavascript {
		args.JavascriptEnabled = pointy.Bool(true)
	}

	if opts.disableFailIndexKeyTooLong {
		args.FailIndexKeyTooLong = pointy.Bool(false)
	}

	if opts.enableFailIndexKeyTooLong {
		args.FailIndexKeyTooLong = pointy.Bool(true)
	}

	if opts.oplogSizeMB != 0 {
		args.OplogSizeMB = pointy.Int64(opts.oplogSizeMB)
	}

	if opts.oplogMinRetentionHours != 0 {
		args.OplogMinRetentionHours = pointy.Float64(opts.oplogMinRetentionHours)
	}

	return args
}

// atlas cluster(s) advancedSettings update <clusterName> --projectId projectId [--readConcern readConcern]
// [--writeConcern writeConcern] [--disableFailIndexKeyTooLong true/fale] [--disableJavascript true/false]
// [--minimumEnabledTLSProtocol minimumEnabledTLSProtocol] [--disableTableScan true/false] [--oplogMinRetentionHours oplogMinRetentionHours]
// [--oplogSizeMB oplogSizeMB] [--sampleRefreshIntervalBIConnector sampleRefreshIntervalBIConnector] [--sampleSizeBIConnector sampleSizeBIConnector].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <clusterName>",
		Short: "Update advanced configuration settings for one cluster.",
		Example: fmt.Sprintf(`  Update oplog minimum size for a cluster:
  $ %[1]s cluster advancedSettings update <clusterName> --projectId <projectId> --oplogSizeMB 1000

  Update minimum TLS protocol for a cluster:
  $ %[1]s cluster advancedSettings update <clusterName> --projectId <projectId> --minimumEnabledTLSProtocol "TLS1_0"`,
			cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				opts.name = args[0]
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTmpl),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to update.",
		},
	}

	cmd.Flags().StringVar(&opts.defaultReadConcern, flag.ReadConcern, defaultReadConcern, usage.ReadConcern)
	cmd.Flags().StringVar(&opts.defaultWriteConcern, flag.WriteConcern, defaultWriteConcern, usage.WriteConcernAdvancedSettings)
	cmd.Flags().StringVar(&opts.minimumEnabledTLSProtocol, flag.TLSProtocol, "", usage.TLSProtocol)

	cmd.Flags().BoolVar(&opts.disableTableScan, flag.DisableTableScan, false, usage.DisableTableScan)
	cmd.Flags().BoolVar(&opts.enableTableScan, flag.EnableTableScan, false, usage.EnableTableScan)
	cmd.MarkFlagsMutuallyExclusive(flag.DisableTableScan, flag.EnableTableScan)

	cmd.Flags().BoolVar(&opts.disableFailIndexKeyTooLong, flag.DisableFailIndexKeyTooLong, false, usage.DisableFailIndexKeyTooLong)
	cmd.Flags().BoolVar(&opts.enableFailIndexKeyTooLong, flag.EnableFailIndexKeyTooLong, false, usage.EnableFailIndexKeyTooLong)
	cmd.MarkFlagsMutuallyExclusive(flag.DisableFailIndexKeyTooLong, flag.EnableFailIndexKeyTooLong)

	cmd.Flags().BoolVar(&opts.disableJavascript, flag.DisableJavascript, false, usage.DisableJavascript)
	cmd.Flags().BoolVar(&opts.enableJavascript, flag.EnableJavascript, false, usage.EnableJavascript)
	cmd.MarkFlagsMutuallyExclusive(flag.DisableJavascript, flag.EnableJavascript)

	cmd.Flags().Float64Var(&opts.oplogMinRetentionHours, flag.OplogMinRetentionHours, 0, usage.OplogMinRetentionHours)
	cmd.Flags().Int64Var(&opts.oplogSizeMB, flag.OplogSizeMB, 0, usage.OplogSizeMB)
	cmd.Flags().Int64Var(&opts.sampleRefreshIntervalBIConnector, flag.SampleRefreshIntervalBIConnector, 0, usage.SampleRefreshIntervalBIConnector)
	cmd.Flags().Int64Var(&opts.sampleSizeBIConnector, flag.SampleSizeBIConnector, defaultSampleSizeBIConnector, usage.SampleSizeBIConnector)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
