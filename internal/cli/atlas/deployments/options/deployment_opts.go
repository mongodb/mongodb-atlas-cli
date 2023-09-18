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
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/terminal"
)

const (
	MongodHostnamePrefix = "mongod"
	MongotHostnamePrefix = "mongot"
	CheckHostnamePrefix  = "check"
	spinnerSpeed         = 100 * time.Millisecond
	// based on https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster
	clusterNamePattern    = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
	MongotDockerImageName = "docker.io/mongodb/mongodb-atlas-search:preview"
)

var (
	errInvalidDeploymentName = errors.New("invalid cluster name")
)

type DeploymentOpts struct {
	DeploymentName string
	DeploymentType string
	MdbVersion     string
	Port           int
	PodmanClient   podman.Client
	CredStore      store.CredentialsGetter
	s              *spinner.Spinner
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
		_, err := log.Warningln("To get output for both local and Atlas clusters, run \"atlas login\" command to authenticate your Atlas account.")
		if err != nil {
			return err
		}
	}

	if err := podman.Installed(); err == podman.ErrPodmanNotFound {
		_, err = log.Warningln("To get output for both local and Atlas clusters, install Podman.")
		if err != nil {
			return err
		}
	}
	return nil
}

func (opts *DeploymentOpts) IsCliAuthenticated() bool {
	return opts.CredStore.AuthType() != config.NotLoggedIn
}
