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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTmpl = "Updating advanced configuration settings of your cluster'.\n"

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
	oplogSizeMB                      int
	sampleRefreshIntervalBIConnector int
	sampleSizeBIConnector            int
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

func (opts *UpdateOpts) newProcessArgs() *atlasv2.ClusterDescriptionProcessArgs {
	args := &atlasv2.ClusterDescriptionProcessArgs{}

	if opts.defaultReadConcern != "" {
		args.DefaultReadConcern = &opts.defaultReadConcern
	}

	if opts.defaultWriteConcern != "" {
		args.DefaultWriteConcern = &opts.defaultWriteConcern
	}

	if opts.minimumEnabledTLSProtocol != "" {
		args.MinimumEnabledTlsProtocol = &opts.minimumEnabledTLSProtocol
	}

	if opts.sampleSizeBIConnector != -1 {
		args.SampleSizeBIConnector = &opts.sampleSizeBIConnector
	}

	if opts.sampleRefreshIntervalBIConnector != -1 {
		args.SampleRefreshIntervalBIConnector = &opts.sampleRefreshIntervalBIConnector
	}

	if opts.disableTableScan {
		args.NoTableScan = &opts.disableTableScan
	}

	if opts.enableTableScan {
		noTableScan := !opts.enableTableScan
		args.NoTableScan = &noTableScan
	}

	if opts.disableJavascript || opts.enableJavascript {
		args.JavascriptEnabled = &opts.enableJavascript
	}

	if opts.disableFailIndexKeyTooLong || opts.enableFailIndexKeyTooLong {
		args.FailIndexKeyTooLong = &opts.enableFailIndexKeyTooLong
	}

	if opts.oplogSizeMB != 0 {
		args.OplogSizeMB = &opts.oplogSizeMB
	}

	if opts.oplogMinRetentionHours != 0 {
		args.OplogMinRetentionHours = &opts.oplogMinRetentionHours
	}

	return args
}

// atlas cluster(s) advancedSettings update <clusterName> --projectId projectId [--readConcern readConcern]
// [--writeConcern writeConcern] [--disableFailIndexKeyTooLong true/fale] [--enableFailIndexKeyTooLong true/fale] [--disableJavascript true/false] [--enableJavascript true/false]
// [--minimumEnabledTLSProtocol minimumEnabledTLSProtocol] [--disableTableScan true/false] [--enableTableScan true/false] [--oplogMinRetentionHours oplogMinRetentionHours]
// [--oplogSizeMB oplogSizeMB] [--sampleRefreshIntervalBIConnector sampleRefreshIntervalBIConnector] [--sampleSizeBIConnector sampleSizeBIConnector].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <clusterName>",
		Short: "Update advanced configuration settings for one cluster.",
		Long: `Updates the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings.
Atlas supports this command only for M10+ clusters.
`,
		Example: `  # Update the minimum oplog size for a cluster:
  atlas cluster advancedSettings update <clusterName> --projectId <projectId> --oplogSizeMB 1000

  # Update the minimum TLS protocol version for a cluster:
  atlas cluster advancedSettings update <clusterName> --projectId <projectId> --minimumEnabledTLSProtocol "TLS1_2"`,
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
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to update.",
			"output":          updateTmpl,
		},
	}

	cmd.Flags().StringVar(&opts.defaultReadConcern, flag.ReadConcern, "", usage.ReadConcern)
	cmd.Flags().StringVar(&opts.defaultWriteConcern, flag.WriteConcern, "", usage.WriteConcernAdvancedSettings)
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
	cmd.Flags().IntVar(&opts.oplogSizeMB, flag.OplogSizeMB, 0, usage.OplogSizeMB)
	cmd.Flags().IntVar(&opts.sampleRefreshIntervalBIConnector, flag.SampleRefreshIntervalBIConnector, -1, usage.SampleRefreshIntervalBIConnector)
	cmd.Flags().IntVar(&opts.sampleSizeBIConnector, flag.SampleSizeBIConnector, -1, usage.SampleSizeBIConnector)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
