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
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/terminal"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
)

const (
	MongodHostnamePrefix = "mongod"
	MongotHostnamePrefix = "mongot"
	CheckHostnamePrefix  = "check"
	spinnerSpeed         = 100 * time.Millisecond
	// based on https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster
	clusterNamePattern    = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
	MongotDockerImageName = "docker.io/mongodb/mongodb-atlas-search:preview-20230922T141716Z"
	PausedState           = "PAUSED"
	StoppedState          = "STOPPED"
	IdleState             = "IDLE"
	DeletingState         = "DELETING"
	RestartingState       = "RESTARTING"
	LocalCluster          = "local"
	AtlasCluster          = "atlas"
	PromptTypeMessage     = "What type of deployment would you like to work with?"
)

var (
	errInvalidDeploymentName        = errors.New("invalid cluster name")
	errDeploymentTypeNotImplemented = errors.New("deployment type not implemented")
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
	DeploymentName string
	DeploymentType string
	MdbVersion     string
	Port           int
	PodmanClient   podman.Client
	CredStore      store.CredentialsGetter
	s              *spinner.Spinner
}

type Deployment struct {
	Type           string
	Name           string
	MongoDBVersion string
	StateName      string
}

func (opts *DeploymentOpts) InitStore(podmanClient podman.Client) func() error {
	return func() error {
		opts.PodmanClient = podmanClient
		return nil
	}
}

func (opts *DeploymentOpts) LocalMongodHostname() string {
	return fmt.Sprintf("%s-%s", MongodHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongotHostname() string {
	return fmt.Sprintf("%s-%s", MongotHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalCheckHostname() string {
	return fmt.Sprintf("%s-%s", CheckHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalNetworkName() string {
	return fmt.Sprintf("mdb-local-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongotDataVolume() string {
	return fmt.Sprintf("mongot-local-data-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongodDataVolume() string {
	return fmt.Sprintf("mongod-local-data-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongoMetricsVolume() string {
	return fmt.Sprintf("mongot-local-metrics-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) MongodDockerImageName() string {
	return fmt.Sprintf("docker.io/mongodb/mongodb-enterprise-server:%s-ubi8", opts.MdbVersion)
}

func LocalDeploymentName(hostname string) string {
	return strings.TrimPrefix(hostname, fmt.Sprintf("%s-", MongodHostnamePrefix))
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

func (opts *DeploymentOpts) PostRunMessages() error {
	if !opts.IsCliAuthenticated() {
		if _, err := log.Warningln("\nTo get output for both local and Atlas deployments, run \"atlas login\" command to authenticate your Atlas account."); err != nil {
			return err
		}
	}

	if err := podman.Installed(); errors.Is(err, podman.ErrPodmanNotFound) {
		if _, err = log.Warningln("\nTo get output for both local and Atlas deployments, install Podman."); err != nil {
			return err
		}
	}
	return nil
}

func (opts *DeploymentOpts) IsCliAuthenticated() bool {
	if opts.CredStore == nil {
		opts.CredStore = config.Default()
	}
	return opts.CredStore.AuthType() != config.NotLoggedIn
}

func (opts *DeploymentOpts) GetLocalDeployments(ctx context.Context) ([]Deployment, error) {
	if err := opts.PodmanClient.Ready(ctx); err != nil {
		return nil, err
	}

	mdbContainers, err := opts.PodmanClient.ListContainers(ctx, MongodHostnamePrefix)
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

		name := strings.TrimPrefix(c.Names[0], MongodHostnamePrefix+"-")
		deployments[i] = Deployment{
			Type:           "LOCAL",
			Name:           name,
			MongoDBVersion: c.Labels["version"],
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *DeploymentOpts) PromptDeploymentType() error {
	p := &survey.Select{
		Message: PromptTypeMessage,
		Options: DeploymentTypeOptions,
		Help:    usage.DeploymentType,
		Description: func(value string, index int) string {
			return deploymentTypeDescription[value]
		},
	}

	err := telemetry.TrackAskOne(p, &opts.DeploymentType, nil)
	if err != nil {
		return err
	}

	if !strings.EqualFold(opts.DeploymentType, AtlasCluster) && !strings.EqualFold(opts.DeploymentType, LocalCluster) {
		return fmt.Errorf("%w: %s", errDeploymentTypeNotImplemented, deploymentTypeDescription[opts.DeploymentType])
	}

	return nil
}
