// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/compass"
	"github.com/andreaangiolillo/mongocli-test/internal/log"
	"github.com/andreaangiolillo/mongocli-test/internal/mongosh"
	"github.com/andreaangiolillo/mongocli-test/internal/telemetry"
)

var (
	ConnectionStringTypeStandard             = "standard"
	connectionStringTypePrivate              = "private"
	connectionStringTypeOptions              = []string{ConnectionStringTypeStandard, connectionStringTypePrivate}
	errConnectionStringTypeNotImplemented    = errors.New("connection string type not implemented")
	errNetworkPeeringConnectionNotConfigured = errors.New("network peering connection is not configured for this deployment")
	promptConnectionStringType               = "What type of connection string type would you like to use?"
)

type ConnectOpts struct {
	cli.OutputOpts
	DeploymentOpts
	ConnectWith string
	ConnectToAtlasOpts
}

func (opts *ConnectOpts) Connect(ctx context.Context) error {
	if _, err := opts.SelectDeployments(ctx, opts.ConfigProjectID()); err != nil {
		return err
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
		if opts.ConnectWith, err = opts.DeploymentOpts.PromptConnectWith(); err != nil {
			return err
		}
	}

	return ValidateConnectWith(opts.ConnectWith)
}

func (opts *ConnectOpts) connectToDeployment(connectionString string) error {
	switch opts.ConnectWith {
	case ConnectWithConnectionString:
		return opts.Print(connectionString)
	case CompassConnect:
		if !compass.Detect() {
			return ErrCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run(opts.DeploymentOpts.DBUsername, opts.DeploymentOpts.DBUserPassword, connectionString)
	case MongoshConnect:
		if !mongosh.Detect() {
			return ErrMongoshNotInstalled
		}
		return mongosh.Run(opts.DeploymentOpts.DBUsername, opts.DeploymentOpts.DBUserPassword, connectionString)
	}

	return nil
}

func (opts *ConnectOpts) promptDBUsername() error {
	p := &survey.Input{
		Message: "Username for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DeploymentOpts.DBUsername)
}

func (opts *ConnectOpts) promptDBUserPassword() error {
	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.DeploymentOpts.DBUserPassword)
		return err
	}

	p := &survey.Password{
		Message: "Password for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DeploymentOpts.DBUserPassword)
}
