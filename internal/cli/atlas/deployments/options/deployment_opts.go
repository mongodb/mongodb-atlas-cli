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
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
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
	PausedState           = "PAUSED"
	StoppedState          = "STOPPED"
	IdleState             = "IDLE"
	DeletingState         = "DELETING"
	RestartingState       = "RESTARTING"
)

var (
	errInvalidDeploymentName = errors.New("invalid cluster name")
)

var localStateMap = map[string]string{
	"running":  IdleState,
	"removing": DeletingState,
	// a "created" container is ready to be started but is currently stopped,
	// which for a local deployment is equivalent to being paused.
	"created":    PausedState,
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
	Type            string
	Name            string
	MongoDBVersion  string
	StateName       string
	MongoDContainer *podman.Container
	MongoTContainer *podman.Container
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
			Type:            "LOCAL",
			Name:            name,
			MongoDBVersion:  c.Labels["version"],
			StateName:       stateName,
			MongoDContainer: c,
		}
	}

	return deployments, nil
}

func (opts *DeploymentOpts) GetLocalDeploymentsWithContainers(ctx context.Context) (map[string]Deployment, error) {
	if err := opts.PodmanClient.Ready(ctx); err != nil {
		return nil, err
	}

	mdbContainers, err := opts.PodmanClient.ListContainers(ctx, MongotHostnamePrefix)
	if err != nil {
		return nil, err
	}

	sort.Slice(mdbContainers, func(i, j int) bool {
		return mdbContainers[i].Names[0] < mdbContainers[j].Names[0]
	})

	deployments := make(map[string]Deployment)
	for _, c := range mdbContainers {
		stateName, found := localStateMap[c.State]
		if !found {
			stateName = strings.ToUpper(c.State)
		}

		name := strings.TrimPrefix(c.Names[0], MongotHostnamePrefix+"-")
		deployments[name] = Deployment{
			Type:            "LOCAL",
			Name:            name,
			MongoDBVersion:  c.Labels["version"],
			StateName:       stateName,
			MongoTContainer: c,
		}
	}

	mdbContainers, err = opts.PodmanClient.ListContainers(ctx, MongodHostnamePrefix)
	if err != nil {
		return nil, err
	}

	sort.Slice(mdbContainers, func(i, j int) bool {
		return mdbContainers[i].Names[0] < mdbContainers[j].Names[0]
	})

	for _, c := range mdbContainers {
		name := strings.TrimPrefix(c.Names[0], MongodHostnamePrefix+"-")
		if v, ok := deployments[name]; ok {
			v.MongoDContainer = c
			deployments[name] = v
		}
	}

	print("\nANDREA\n")
	res2B, _ := json.Marshal(deployments)
	fmt.Println(string(res2B))
	print("\n\n")
	return deployments, nil
}

//func (opts *DeploymentOpts) GetLocalDeploymentsWithContainers(ctx context.Context) ([]Deployment, error) {
//	deployments, err := opts.GetLocalDeployments(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	mdtContainers, err := opts.PodmanClient.ListContainers(ctx, "")
//	if err != nil {
//		return nil, err
//	}
//
//	sort.Slice(mdtContainers, func(i, j int) bool {
//		return mdtContainers[i].Names[0] < mdtContainers[j].Names[0]
//	})
//
//	for i, c := range mdtContainers {
//		deployments[i].MongoTContainer = c
//	}
//
//	return deployments, nil
//}
