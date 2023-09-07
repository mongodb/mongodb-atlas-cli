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
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/terminal"
)

const (
	MongodHostnamePrefix = "mongod"
	MongotHostnamePrefix = "mongot"
	spinnerSpeed         = 100 * time.Millisecond
)

type DeploymentOpts struct {
	DeploymentName string
	DeploymentType string
	MdbVersion     string
	Port           int
	podmanClient   podman.Client
	s              *spinner.Spinner
}

func (opts *DeploymentOpts) InitStore(podmanClient podman.Client) func() error {
	return func() error {
		opts.podmanClient = podmanClient
		return nil
	}
}

func (opts *DeploymentOpts) LocalMongodHostname() string {
	return fmt.Sprintf("%s-%s", MongodHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongotHostname() string {
	return fmt.Sprintf("%s-%s", MongotHostnamePrefix, opts.DeploymentName)
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
