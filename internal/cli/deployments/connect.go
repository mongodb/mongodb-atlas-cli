// Copyright 2023 MongoDB Inc
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

package deployments

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
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
)

var (
	ConnectionStringTypeStandard             = "standard"
	connectionStringTypePrivate              = "private"
	connectionStringTypeOptions              = []string{ConnectionStringTypeStandard, connectionStringTypePrivate}
	errConnectionStringTypeNotImplemented    = errors.New("connection string type not implemented")
	errNetworkPeeringConnectionNotConfigured = errors.New("network peering connection is not configured for this deployment")
	errConnectionError                       = errors.New("could not connect")
	promptConnectionStringType               = "What type of connection string type would you like to use?"
)

func Run(ctx context.Context, opts *ConnectOpts) error {
	return opts.Connect(ctx)
}

func PostRun(opts *ConnectOpts) {
	opts.DeploymentTelemetry.AppendDeploymentType()
	opts.DeploymentTelemetry.AppendDeploymentUUID()
}

// atlas deployments connect [clusterName].
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}
	cmd := &cobra.Command{
		Use:     "connect [deploymentName]",
		Short:   "Connect to a deployment that is running locally or in Atlas. If the deployment is paused, make sure to run atlas deployments start first.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment that you want to connect to.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitAtlasStore(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(cmd.Context(), opts)
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			PostRun(opts)
		},
	}

	cmd.Flags().StringVar(&opts.ConnectWith, flag.ConnectWith, "", usage.ConnectWithConnect)
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.ConnectionStringType, flag.ConnectionStringType, ConnectionStringTypeStandard, usage.ConnectionStringType)

	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return options.ConnectWithOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})

	return cmd
}

type ConnectOpts struct {
	cli.OutputOpts
	options.DeploymentOpts
	ConnectWith string
	ConnectToAtlasOpts
}

var errDeploymentNotStarted = errors.New("deployment not started")

func (opts *ConnectOpts) startDeployment(ctx context.Context, deployment options.Deployment) error {
	ok, err := opts.promptStartDeployment()
	if err != nil {
		return err
	}
	if !ok {
		return errDeploymentNotStarted
	}

	return opts.Spin(func() error {
		if opts.DeploymentType == options.AtlasCluster {
			_, err := opts.Store.StartCluster(opts.ProjectID, deployment.Name)

			return err
		}

		return opts.StartLocal(ctx, deployment)
	})
}

func (opts *ConnectOpts) Connect(ctx context.Context) error {
	d, err := opts.SelectDeployments(ctx, opts.ConfigProjectID(), options.IdleState, options.StoppedState, options.PausedState)
	if err != nil {
		return err
	}

	if d.StateName == options.StoppedState || d.StateName == options.PausedState {
		if err := opts.startDeployment(ctx, d); err != nil {
			return err
		}
	}

	if err := opts.askConnectWith(); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() {
		if err := opts.validateAndPromptAtlasOpts(); err != nil {
			return err
		}

		return opts.connectToAtlas()
	}

	return opts.connectToLocal(ctx)
}

func (opts *ConnectOpts) askConnectWith() error {
	if opts.ConnectWith == "" {
		var err error
		if opts.ConnectWith, err = opts.PromptConnectWith(); err != nil {
			return err
		}
	}

	return options.ValidateConnectWith(opts.ConnectWith)
}

func (opts *ConnectOpts) connectToDeployment(connectionString string) error {
	switch opts.ConnectWith {
	case options.ConnectWithConnectionString:
		return opts.Print(connectionString)
	case options.CompassConnect:
		if !compass.Detect() {
			return compass.ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case options.MongoshConnect:
		if !mongosh.Detect() {
			return mongosh.ErrMongoshNotInstalled
		}
		return mongosh.Run(opts.DBUsername, opts.DBUserPassword, connectionString)
	case options.VsCodeConnect:
		if !vscode.Detect() {
			return vscode.ErrVsCodeCliNotInstalled
		}
		if _, err := log.Warningln("Launching VsCode..."); err != nil {
			return err
		}
		return vscode.SaveConnection(connectionString, opts.DeploymentName, opts.DeploymentType)
	}

	return nil
}

func (opts *ConnectOpts) promptDBUsername() error {
	p := &survey.Input{
		Message: "Username for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DBUsername)
}

func (*ConnectOpts) promptStartDeployment() (bool, error) {
	var result bool
	p := &survey.Confirm{
		Message: "Deployment seems stopped, would you like to start it?",
		Default: true,
	}

	err := telemetry.TrackAskOne(p, &result)
	return result, err
}

func (opts *ConnectOpts) promptDBUserPassword() error {
	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.DBUserPassword)
		return err
	}

	p := &survey.Password{
		Message: "Password for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DBUserPassword)
}

func (opts *ConnectOpts) connectToLocal(ctx context.Context) error {
	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	return opts.connectToDeployment(connectionString)
}

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=connect_mock_test.go -package=deployments . ClusterDescriberStarter

type ClusterDescriberStarter interface {
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	StartCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
}

type ConnectToAtlasOpts struct {
	cli.ProjectOpts
	cli.InputOpts
	ConnectionStringType string
	Store                ClusterDescriberStarter
}

func (opts *ConnectToAtlasOpts) InitAtlasStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.Store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ConnectOpts) validateAndPromptAtlasOpts() error {
	requiresAuth := opts.ConnectWith == options.MongoshConnect || opts.ConnectWith == options.CompassConnect
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

func (opts *ConnectToAtlasOpts) validateAndPromptConnectionStringType() error {
	if opts.ConnectionStringType == "" {
		p := &survey.Select{
			Message: promptConnectionStringType,
			Options: connectionStringTypeOptions,
			Help:    usage.ConnectionStringType,
		}

		err := telemetry.TrackAskOne(p, &opts.ConnectionStringType, nil)
		if err != nil {
			return err
		}
	}

	if !search.StringInSliceFold(connectionStringTypeOptions, opts.ConnectionStringType) {
		return fmt.Errorf("%w: %s", errConnectionStringTypeNotImplemented, opts.ConnectionStringType)
	}

	return nil
}

func (opts *ConnectOpts) connectToAtlas() error {
	r, err := opts.Store.AtlasCluster(opts.ConfigProjectID(), opts.DeploymentName)
	if err != nil {
		return err
	}

	// Connectionstrings are empty when the server is not in IDLE
	// r.ConnectionStrings.PrivateSrv == nil and r.ConnectionStrings.StandardSrv == nil
	if r.GetStateName() != "IDLE" {
		return fmt.Errorf("%w: cluster is not in an idle state yet, try again in a few moments", errConnectionError)
	}

	// This field is optional, if not set, throw an error
	if r.ConnectionStrings == nil {
		return fmt.Errorf("%w: server did not return connectionstrings", errConnectionError)
	}

	if opts.ConnectionStringType == connectionStringTypePrivate {
		if r.GetConnectionStrings().PrivateSrv == nil {
			return errNetworkPeeringConnectionNotConfigured
		}
		return opts.connectToDeployment(*r.GetConnectionStrings().PrivateSrv)
	}

	// Make sure the string pointer is not nil before dereferencing
	if r.ConnectionStrings.StandardSrv == nil {
		return fmt.Errorf("%w: server did not return connectionstring", errConnectionError)
	}
	return opts.connectToDeployment(*r.ConnectionStrings.StandardSrv)
}
