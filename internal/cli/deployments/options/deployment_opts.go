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
	"io"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/setup"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
)

const (
	spinnerSpeed = 100 * time.Millisecond
	// based on https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster
	clusterNamePattern = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
	PausedState        = "PAUSED"
	StoppedState       = "STOPPED"
	IdleState          = "IDLE"
	DeletingState      = "DELETING"
	RestartingState    = "RESTARTING"
	LocalCluster       = "local"
	AtlasCluster       = "atlas"
	CompassConnect     = "compass"
	MongoshConnect     = "mongosh"
	PromptTypeMessage  = "What type of deployment would you like to work with?"
	MaxItemsPerPage    = 500
	ContainerFilter    = "mongodb-atlas-local=container"
)

var (
	errInvalidDeploymentName        = errors.New("invalid cluster name")
	errDeploymentTypeNotImplemented = errors.New("deployment type not implemented")
	ErrNotAuthenticated             = errors.New("you are not authenticated. Please, run atlas auth login")
	ErrCompassNotInstalled          = errors.New("did not find MongoDB Compass, install: https://dochub.mongodb.org/core/install-compass")
	ErrMongoshNotInstalled          = errors.New("did not find mongosh, install: https://dochub.mongodb.org/core/install-mongosh")
	DeploymentTypeOptions           = []string{LocalCluster, AtlasCluster}
	deploymentTypeDescription       = map[string]string{
		LocalCluster: "Local Database",
		AtlasCluster: "Atlas Database",
	}
)

var localStateMap = map[string]string{
	"running":  IdleState,
	"removing": DeletingState,
	// a "created" container is ready to be started but is currently stopped
	"created":    StoppedState,
	"paused":     PausedState,
	"restarting": RestartingState,
	"exited":     StoppedState,
	"dead":       StoppedState,
}

type DeploymentOpts struct {
	DeploymentName        string
	DeploymentType        string
	MdbVersion            string
	Port                  int
	DBUsername            string
	DBUserPassword        string
	ContainerEngine       container.Engine
	CredStore             store.CredentialsGetter
	s                     *spinner.Spinner
	DefaultSetter         cli.DefaultSetterOpts
	AtlasClusterListStore store.ClusterLister
	Config                setup.ProfileReader
	DeploymentTelemetry   DeploymentTelemetry
}

type Deployment struct {
	Type           string
	Name           string
	MongoDBVersion string
	StateName      string
}

func (opts *DeploymentOpts) InitStore(ctx context.Context, writer io.Writer) func() error {
	return func() error {
		var err error
		opts.ContainerEngine = container.New()
		opts.Config = config.Default()
		opts.CredStore = config.Default()
		if opts.AtlasClusterListStore, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx)); err != nil {
			return err
		}
		opts.DefaultSetter.OutWriter = writer
		opts.DeploymentTelemetry = NewDeploymentTypeTelemetry(opts)
		opts.UpdateDeploymentTelemetry()
		return opts.DefaultSetter.InitStore(ctx)
	}
}

func (opts *DeploymentOpts) LocalMongodHostname() string {
	return opts.DeploymentName
}

func (opts *DeploymentOpts) MongodDockerImageName() string {
	return "docker.io/mongodb/mongodb-atlas-local:" + opts.MdbVersion
}

func (opts *DeploymentOpts) StartSpinner() {
	if terminal.IsTerminal(log.Writer()) {
		opts.s = spinner.New(spinner.CharSets[9], spinnerSpeed)
		opts.s.Start()
	}
}

func (opts *DeploymentOpts) StopSpinner() {
	if terminal.IsTerminal(log.Writer()) {
		opts.s.Stop()
	}
}

func ValidateDeploymentName(n string) error {
	if matched, _ := regexp.MatchString(clusterNamePattern, n); !matched {
		return fmt.Errorf("%w: %s", errInvalidDeploymentName, n)
	}
	return nil
}

func (opts *DeploymentOpts) IsCliAuthenticated() bool {
	if opts.CredStore == nil {
		opts.CredStore = config.Default()
	}
	return opts.CredStore.AuthType() != config.NotLoggedIn
}

func (opts *DeploymentOpts) GetLocalContainers(ctx context.Context) ([]container.Container, error) {
	return opts.ContainerEngine.ContainerList(ctx, ContainerFilter)
}

func (opts *DeploymentOpts) GetLocalDeployments(ctx context.Context) ([]Deployment, error) {
	mdbContainers, err := opts.GetLocalContainers(ctx)
	if err != nil {
		return nil, err
	}
	sort.Slice(mdbContainers, func(i, j int) bool {
		return mdbContainers[i].Names[0] < mdbContainers[j].Names[0]
	})

	deployments := make([]Deployment, len(mdbContainers))
	for i, c := range mdbContainers {
		stateName, found := localStateMap[c.State]
		if !found {
			stateName = strings.ToUpper(c.State)
		}

		name := c.Names[0]
		deployments[i] = Deployment{
			Type:           "LOCAL",
			Name:           name,
			MongoDBVersion: c.Labels["version"],
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *DeploymentOpts) promptDeploymentType() error {
	p := &survey.Select{
		Message: PromptTypeMessage,
		Options: DeploymentTypeOptions,
		Help:    usage.DeploymentType,
		Description: func(value string, _ int) string {
			return deploymentTypeDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.DeploymentType, nil)
}

func validateDeploymentType(s string) error {
	if !search.StringInSliceFold(DeploymentTypeOptions, s) {
		return fmt.Errorf("%w: %s", errDeploymentTypeNotImplemented, s)
	}
	return nil
}

func (opts *DeploymentOpts) ValidateAndPromptDeploymentType() error {
	if opts.DeploymentType == "" {
		if err := opts.promptDeploymentType(); err != nil {
			return err
		}
	} else if err := validateDeploymentType(opts.DeploymentType); err != nil {
		return err
	}
	return nil
}

func (opts *DeploymentOpts) IsAtlasDeploymentType() bool {
	return strings.EqualFold(opts.DeploymentType, AtlasCluster)
}

func (opts *DeploymentOpts) IsLocalDeploymentType() bool {
	return strings.EqualFold(opts.DeploymentType, LocalCluster)
}

func (opts *DeploymentOpts) NoDeploymentTypeSet() bool {
	return strings.EqualFold(opts.DeploymentType, "")
}

func (opts *DeploymentOpts) IsAuthEnabled() bool {
	return opts.DBUsername != ""
}

func (opts *DeploymentOpts) UpdateDeploymentTelemetry() {
	if opts.DeploymentTelemetry == nil {
		opts.DeploymentTelemetry = NewDeploymentTypeTelemetry(opts)
	}
	opts.DeploymentTelemetry.AppendDeploymentType()
}
