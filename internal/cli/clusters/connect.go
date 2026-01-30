// Copyright 2026 MongoDB Inc
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

package clusters

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/connect"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/vscode"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

var (
	errNetworkPeeringConnectionNotConfigured = errors.New("network peering connection is not configured for this cluster")
	errConnectionError                       = errors.New("could not connect")
	errClusterNotStarted                     = errors.New("cluster not started")
	errNoClusters                            = errors.New("currently there are no clusters in your project")
	errClusterRequiredOnPipe                 = errors.New("cluster name is required when piping the output of the command")
	promptSelectCluster                      = "Select a cluster"
	promptStartClusterConfirm                = "Cluster seems stopped, would you like to start it?"
	promptConnectWithFormat                  = "How would you like to connect to %s?"
	promptDBUsername                         = "Username for authenticating to MongoDB cluster"
	promptDBUserPassword                     = "Password for authenticating to MongoDB cluster"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=connect_mock_test.go -package=clusters . ConnectClusterStore

type ConnectClusterStore interface {
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	LatestAtlasCluster(string, string) (*atlasv2.ClusterDescription20240805, error)
	StartCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	StartClusterLatest(string, string) (*atlasv2.ClusterDescription20240805, error)
	LatestProjectClusters(string, *store.ListOptions) (*atlasv2.PaginatedClusterDescription20240805, error)
}

func Run(ctx context.Context, opts *ConnectOpts) error {
	return opts.Connect(ctx)
}

// atlas cluster connect [clusterName].
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}
	cmd := &cobra.Command{
		Use:   "connect [clusterName]",
		Short: "Connect to an Atlas cluster. If the cluster is paused, run atlas cluster start first.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster that you want to connect to.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.name = args[0]
			}
			preRunE := []prerun.CmdOpt{
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.resolveClusterName(cmd.Context()),
			}
			if opts.autoScalingMode != "" {
				preRunE = append(preRunE, validate.AutoScalingMode(opts.autoScalingMode))
			}
			return opts.PreRunE(preRunE...)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.ConnectWith, flag.ConnectWith, "", usage.ConnectWithConnect)
	cmd.Flags().StringVar(&opts.autoScalingMode, flag.AutoScalingMode, "", usage.AutoScalingMode)
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.ConnectionStringType, flag.ConnectionStringType, connect.ConnectionStringTypeStandard, usage.ConnectionStringType)

	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return connect.ConnectWithOptions, cobra.ShellCompDirectiveDefault
	})

	opts.AddOutputOptFlags(cmd)

	return cmd
}

type ConnectOpts struct {
	cli.OutputOpts
	cli.ProjectOpts
	cli.InputOpts
	name                 string
	autoScalingMode      string
	ConnectWith          string
	DBUsername           string
	DBUserPassword       string
	ConnectionStringType string
	store                ConnectClusterStore
}

func (opts *ConnectOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ConnectOpts) resolveClusterName(ctx context.Context) func() error {
	return func() error {
		stat, _ := os.Stdout.Stat()
		if (stat.Mode()&os.ModeCharDevice) == 0 && opts.name == "" {
			return errClusterRequiredOnPipe
		}
		if opts.name != "" {
			return nil
		}
		return opts.selectCluster(ctx)
	}
}

func (opts *ConnectOpts) selectCluster(_ context.Context) error {
	listOpts := &store.ListOptions{
		PageNum:      cli.DefaultPage,
		ItemsPerPage: connect.MaxItemsPerPage,
	}
	clusters, err := opts.store.LatestProjectClusters(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}
	results := clusters.GetResults()
	if len(results) == 0 {
		return errNoClusters
	}
	if len(results) == 1 {
		opts.name = results[0].GetName()
		return nil
	}
	names := make([]string, len(results))
	for i, c := range results {
		names[i] = c.GetName()
	}
	return telemetry.TrackAskOne(&survey.Select{
		Message: promptSelectCluster,
		Options: names,
		Help:    usage.ClusterName,
	}, &opts.name, survey.WithValidator(survey.Required))
}

func (opts *ConnectOpts) startCluster(_ context.Context) error {
	ok, err := opts.promptStartCluster()
	if err != nil {
		return err
	}
	if !ok {
		return errClusterNotStarted
	}
	_, err = opts.store.StartCluster(opts.ConfigProjectID(), opts.name)
	return err
}

func (*ConnectOpts) promptStartCluster() (bool, error) {
	var result bool
	p := &survey.Confirm{
		Message: promptStartClusterConfirm,
		Default: true,
	}
	err := telemetry.TrackAskOne(p, &result)
	return result, err
}

func (opts *ConnectOpts) Connect(ctx context.Context) error {
	if opts.autoScalingMode != "" && isIndependentShardScaling(opts.autoScalingMode) {
		return opts.connectLatest(ctx)
	}

	r, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	stateName := r.GetStateName()
	if r.GetPaused() {
		stateName = connect.PausedState
	}
	if stateName == connect.StoppedState || stateName == connect.PausedState {
		if err := opts.startCluster(ctx); err != nil {
			return err
		}
	}

	if err := opts.askConnectWith(); err != nil {
		return err
	}
	if err := opts.validateAndPromptAtlasOpts(); err != nil {
		return err
	}
	return opts.connectToAtlas(r)
}

func (opts *ConnectOpts) connectLatest(_ context.Context) error {
	r, err := opts.store.LatestAtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	stateName := r.GetStateName()
	if r.GetPaused() {
		stateName = connect.PausedState
	}
	if stateName == connect.StoppedState || stateName == connect.PausedState {
		if _, err := opts.store.StartClusterLatest(opts.ConfigProjectID(), opts.name); err != nil {
			return err
		}
	}

	if err := opts.askConnectWith(); err != nil {
		return err
	}
	if err := opts.validateAndPromptAtlasOpts(); err != nil {
		return err
	}
	return opts.connectToAtlasLatest(r)
}

func (opts *ConnectOpts) askConnectWith() error {
	if opts.ConnectWith == "" {
		var err error
		if opts.ConnectWith, err = opts.promptConnectWith(); err != nil {
			return err
		}
	}
	return connect.ValidateConnectWith(opts.ConnectWith)
}

func (opts *ConnectOpts) connectToCluster(connectionString string) error {
	switch opts.ConnectWith {
	case connect.ConnectWithConnectionString:
		return opts.Print(connectionString)
	case connect.ConnectWithCompass:
		if !compass.Detect() {
			return compass.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case connect.ConnectWithMongosh:
		if !mongosh.Detect() {
			return mongosh.ErrMongoshNotInstalled
		}
		return mongosh.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case connect.ConnectWithVsCode:
		if !vscode.Detect() {
			return vscode.ErrVsCodeCliNotInstalled
		}
		if _, err := log.Warningln("Launching VsCode..."); err != nil {
			return err
		}
		return vscode.SaveConnection(connectionString, opts.name, connect.AtlasCluster)
	}
	return nil
}

func (opts *ConnectOpts) promptConnectWith() (string, error) {
	p := &survey.Select{
		Message: fmt.Sprintf(promptConnectWithFormat, opts.name),
		Options: connect.ConnectWithOptions,
		Description: func(value string, _ int) string {
			return map[string]string{
				connect.ConnectWithConnectionString: "Connection String",
				connect.ConnectWithMongosh:          "MongoDB Shell",
				connect.ConnectWithCompass:          "MongoDB Compass",
				connect.ConnectWithVsCode:           "MongoDB for VsCode",
			}[value]
		},
	}
	var response string
	err := telemetry.TrackAskOne(p, &response, nil)
	return response, err
}

func (opts *ConnectOpts) validateAndPromptAtlasOpts() error {
	requiresAuth := opts.ConnectWith == connect.ConnectWithMongosh || opts.ConnectWith == connect.ConnectWithCompass
	if requiresAuth && opts.DBUsername == "" {
		if err := opts.promptDBUsername(); err != nil {
			return err
		}
	}
	if requiresAuth && opts.DBUserPassword == "" {
		if err := opts.promptDBUserPassword(); err != nil {
			return err
		}
	}
	return opts.validateAndPromptConnectionStringType()
}

func (opts *ConnectOpts) promptDBUsername() error {
	p := &survey.Input{
		Message: promptDBUsername,
	}
	return telemetry.TrackAskOne(p, &opts.DBUsername)
}

func (opts *ConnectOpts) promptDBUserPassword() error {
	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.DBUserPassword)
		return err
	}

	p := &survey.Password{
		Message: promptDBUserPassword,
	}
	return telemetry.TrackAskOne(p, &opts.DBUserPassword)
}

func (opts *ConnectOpts) validateAndPromptConnectionStringType() error {
	if opts.ConnectionStringType == "" {
		p := &survey.Select{
			Message: connect.PromptConnectionStringType,
			Options: connect.ConnectionStringTypeOptions,
			Help:    usage.ConnectionStringType,
		}

		if err := telemetry.TrackAskOne(p, &opts.ConnectionStringType, nil); err != nil {
			return err
		}
	}
	return connect.ValidateConnectionStringType(opts.ConnectionStringType)
}

func (opts *ConnectOpts) connectToAtlas(r *atlasClustersPinned.AdvancedClusterDescription) error {
	if r.GetStateName() != idle {
		return fmt.Errorf("%w: cluster is not in an idle state yet, try again in a few moments", errConnectionError)
	}

	if r.ConnectionStrings == nil {
		return fmt.Errorf("%w: server did not return connectionstrings", errConnectionError)
	}

	if opts.ConnectionStringType == connect.ConnectionStringTypePrivate {
		if r.GetConnectionStrings().PrivateSrv == nil {
			return errNetworkPeeringConnectionNotConfigured
		}
		return opts.connectToCluster(*r.GetConnectionStrings().PrivateSrv)
	}

	if r.ConnectionStrings.StandardSrv == nil {
		return fmt.Errorf("%w: server did not return connectionstring", errConnectionError)
	}
	return opts.connectToCluster(*r.ConnectionStrings.StandardSrv)
}

func (opts *ConnectOpts) connectToAtlasLatest(r *atlasv2.ClusterDescription20240805) error {
	if r.GetStateName() != idle {
		return fmt.Errorf("%w: cluster is not in an idle state yet, try again in a few moments", errConnectionError)
	}

	if r.ConnectionStrings == nil {
		return fmt.Errorf("%w: server did not return connectionstrings", errConnectionError)
	}

	if opts.ConnectionStringType == connect.ConnectionStringTypePrivate {
		if r.GetConnectionStrings().PrivateSrv == nil {
			return errNetworkPeeringConnectionNotConfigured
		}
		return opts.connectToCluster(*r.GetConnectionStrings().PrivateSrv)
	}

	if r.ConnectionStrings.StandardSrv == nil {
		return fmt.Errorf("%w: server did not return connectionstring", errConnectionError)
	}
	return opts.connectToCluster(*r.ConnectionStrings.StandardSrv)
}
