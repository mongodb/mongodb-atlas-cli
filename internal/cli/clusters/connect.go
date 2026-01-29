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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/vscode"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

const (
	connectWithConnectionString = "connectionString"
	connectWithMongosh          = "mongosh"
	connectWithCompass          = "compass"
	connectWithVsCode           = "vscode"
	pausedState                 = "PAUSED"
	stoppedState                = "STOPPED"
	maxItemsPerPage             = 500
	atlasDeploymentType         = "atlas"
)

var (
	connectWithOptions                       = []string{connectWithMongosh, connectWithCompass, connectWithVsCode, connectWithConnectionString}
	connectionStringTypeStandard             = "standard"
	connectionStringTypePrivate              = "private"
	connectionStringTypeOptions              = []string{connectionStringTypeStandard, connectionStringTypePrivate}
	errConnectionStringTypeNotImplemented    = errors.New("connection string type not implemented")
	errInvalidConnectWith                    = errors.New("invalid --connectWith option")
	errNetworkPeeringConnectionNotConfigured = errors.New("network peering connection is not configured for this cluster")
	errConnectionError                       = errors.New("could not connect")
	errClusterNotStarted                     = errors.New("cluster not started")
	errNoClusters                            = errors.New("currently there are no clusters in your project")
	errClusterRequiredOnPipe                 = errors.New("cluster name is required when piping the output of the command")
	promptConnectionStringType               = "What type of connection string type would you like to use?"
	promptSelectCluster                      = "Select a cluster"
	promptStartClusterConfirm                = "Cluster seems stopped, would you like to start it?"
	promptConnectWithFormat                  = "How would you like to connect to %s?"
	promptDBUsername                         = "Username for authenticating to MongoDB cluster"
	promptDBUserPassword                     = "Password for authenticating to MongoDB cluster"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=connect_mock_test.go -package=clusters . ConnectClusterStore

type ConnectClusterStore interface {
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	StartCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
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
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.resolveClusterName(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.ConnectWith, flag.ConnectWith, "", usage.ConnectWithConnect)
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.ConnectionStringType, flag.ConnectionStringType, connectionStringTypeStandard, usage.ConnectionStringType)

	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return connectWithOptions, cobra.ShellCompDirectiveDefault
	})

	opts.AddOutputOptFlags(cmd)

	return cmd
}

type ConnectOpts struct {
	cli.OutputOpts
	cli.ProjectOpts
	cli.InputOpts
	name                 string
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
		ItemsPerPage: maxItemsPerPage,
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
	r, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	stateName := r.GetStateName()
	if r.GetPaused() {
		stateName = pausedState
	}
	if stateName == stoppedState || stateName == pausedState {
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

func (opts *ConnectOpts) askConnectWith() error {
	if opts.ConnectWith == "" {
		var err error
		if opts.ConnectWith, err = opts.promptConnectWith(); err != nil {
			return err
		}
	}
	return validateConnectWith(opts.ConnectWith)
}

func validateConnectWith(s string) error {
	if !search.StringInSliceFold(connectWithOptions, s) {
		return fmt.Errorf("%w: %s", errInvalidConnectWith, s)
	}
	return nil
}

func (opts *ConnectOpts) connectToCluster(connectionString string) error {
	switch opts.ConnectWith {
	case connectWithConnectionString:
		return opts.Print(connectionString)
	case connectWithCompass:
		if !compass.Detect() {
			return compass.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case connectWithMongosh:
		if !mongosh.Detect() {
			return mongosh.ErrMongoshNotInstalled
		}
		return mongosh.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case connectWithVsCode:
		if !vscode.Detect() {
			return vscode.ErrVsCodeCliNotInstalled
		}
		if _, err := log.Warningln("Launching VsCode..."); err != nil {
			return err
		}
		return vscode.SaveConnection(connectionString, opts.name, atlasDeploymentType)
	}
	return nil
}

func (opts *ConnectOpts) promptConnectWith() (string, error) {
	p := &survey.Select{
		Message: fmt.Sprintf(promptConnectWithFormat, opts.name),
		Options: connectWithOptions,
		Description: func(value string, _ int) string {
			return map[string]string{
				connectWithConnectionString: "Connection String",
				connectWithMongosh:          "MongoDB Shell",
				connectWithCompass:          "MongoDB Compass",
				connectWithVsCode:           "MongoDB for VsCode",
			}[value]
		},
	}
	var response string
	err := telemetry.TrackAskOne(p, &response, nil)
	return response, err
}

func (opts *ConnectOpts) validateAndPromptAtlasOpts() error {
	requiresAuth := opts.ConnectWith == connectWithMongosh || opts.ConnectWith == connectWithCompass
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
			Message: promptConnectionStringType,
			Options: connectionStringTypeOptions,
			Help:    usage.ConnectionStringType,
		}

		if err := telemetry.TrackAskOne(p, &opts.ConnectionStringType, nil); err != nil {
			return err
		}
	}
	if !search.StringInSliceFold(connectionStringTypeOptions, opts.ConnectionStringType) {
		return fmt.Errorf("%w: %s", errConnectionStringTypeNotImplemented, opts.ConnectionStringType)
	}
	return nil
}

func (opts *ConnectOpts) connectToAtlas(r *atlasClustersPinned.AdvancedClusterDescription) error {
	if r.GetStateName() != "IDLE" {
		return fmt.Errorf("%w: cluster is not in an idle state yet, try again in a few moments", errConnectionError)
	}

	if r.ConnectionStrings == nil {
		return fmt.Errorf("%w: server did not return connectionstrings", errConnectionError)
	}

	if opts.ConnectionStringType == connectionStringTypePrivate {
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
