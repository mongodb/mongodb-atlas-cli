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
	"runtime"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/shirou/gopsutil/v3/host"
)

func (opts *DeploymentOpts) SelectDeployments(ctx context.Context, projectID string) (Deployment, error) {
	var atlasDeployments, localDeployments []Deployment
	var atlasErr, localErr error

	if opts.IsAtlasDeploymentType() || opts.NoDeploymentTypeSet() {
		if atlasDeployments, atlasErr = opts.AtlasDeployments(projectID); atlasErr != nil {
			if opts.IsAtlasDeploymentType() {
				return Deployment{}, atlasErr
			}
			if !errors.Is(atlasErr, ErrNotAuthenticated) {
				_, _ = log.Warningf("Warning: failed to retrieve Atlas deployments because %q\n", atlasErr.Error())
			}
		}
	}

	if opts.IsLocalDeploymentType() || opts.NoDeploymentTypeSet() {
		if localErr = opts.LocalDeploymentPreRun(ctx); localErr != nil {
			if opts.IsLocalDeploymentType() {
				return Deployment{}, localErr
			}
			_, _ = log.Warningf("Warning: failed to retrieve Local deployments because %q\n", localErr.Error())
		}

		localDeployments, localErr = opts.GetLocalDeployments(ctx)
		if localErr != nil {
			if opts.IsLocalDeploymentType() {
				return Deployment{}, localErr
			}
			_, _ = log.Warningf("Warning: failed to retrieve Local deployments because %q\n", localErr.Error())
		}
	}

	if atlasErr != nil && localErr != nil {
		return Deployment{}, errors.New("failed to retrieve atlas and local deployments")
	}

	if opts.DeploymentName == "" {
		return opts.Select(append(localDeployments, atlasDeployments...))
	}

	return opts.findDeploymentByName(localDeployments, atlasDeployments)
}

func (opts *DeploymentOpts) findDeploymentByName(localDeployments []Deployment, atlasDeployments []Deployment) (Deployment, error) {
	deployments := make([]Deployment, 0)
	for _, d := range localDeployments {
		if d.Name == opts.DeploymentName {
			deployments = append(deployments, d)
		}
	}

	for _, d := range atlasDeployments {
		if d.Name == opts.DeploymentName {
			deployments = append(deployments, d)
		}
	}

	return opts.Select(deployments)
}

func (opts *DeploymentOpts) AtlasDeployments(projectID string) ([]Deployment, error) {
	if !opts.IsCliAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	if projectID == "" {
		projectID = opts.Config.ProjectID()
	}

	if projectID == "" {
		if err := opts.DefaultSetter.AskProject(); err != nil {
			return nil, err
		}
		projectID = opts.DefaultSetter.ProjectID
	}

	listOpts := &store.ListOptions{
		PageNum:      cli.DefaultPage,
		ItemsPerPage: MaxItemsPerPage,
	}

	atlasClusters, err := opts.AtlasClusterListStore.ProjectClusters(projectID, listOpts)
	if err != nil {
		return nil, err
	}

	deployments := make([]Deployment, len(atlasClusters.GetResults()))
	for i, c := range atlasClusters.GetResults() {
		stateName := c.GetStateName()
		if c.GetPaused() {
			// for paused clusters, Atlas returns stateName IDLE and Paused=true
			stateName = PausedState
		}
		deployments[i] = Deployment{
			Type:           "ATLAS",
			Name:           c.GetName(),
			MongoDBVersion: c.GetMongoDBVersion(),
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *DeploymentOpts) LocalDeploymentPreRun(_ context.Context) error {
	if !localDeploymentSupportedByOs() {
		_, _ = log.Warningln("Local deployments are not supported on this OS, to see local deployments requirements visit https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-deploy-local/.")
	}

	return opts.ContainerEngine.Ready()
}

func localDeploymentSupportedByOs() bool {
	os := runtime.GOOS
	switch os {
	case "darwin":
		// MacOS Intel and M1 are supported
		return true
	case "windows":
		// Windows is not supported
		return false
	case "linux":
		// Depends on distro
		support, err := isLinuxDistroSupported()
		if err != nil {
			// If something went wrong in finding OS distro, then assume support
			_, _ = log.Debugln(err)
			return true
		}
		return support
	default:
		// Other unknown OS are not supported
		return false
	}
}

func isLinuxDistroSupported() (bool, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return false, err
	}

	distro := strings.ToLower(hostInfo.Platform)
	if distro == "" {
		return false, errors.New("unable to find OS distro")
	}

	_, _ = log.Debugln("Detected linux distro: ", distro)
	return strings.Contains(distro, "centos") || strings.Contains(distro, "redhat") || strings.Contains(distro, "rhel"), nil
}
