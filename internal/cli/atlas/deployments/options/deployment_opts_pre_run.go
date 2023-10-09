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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/shirou/gopsutil/v3/host"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

func (opts *DeploymentOpts) SelectDeployments(ctx context.Context, projectId string) error {
	if opts.IsLocalDeploymentType() {
		return opts.displayLocalDeployments(ctx)
	}

	if opts.IsAtlasDeploymentType() {
		return opts.displayAtlasDeployments(ctx, projectId)
	}

	if !opts.IsCliAuthenticated() {
		return opts.displayLocalDeployments(ctx)
	}

	var atlasDeployments []Deployment
	var atlasErr error
	if atlasDeployments, atlasErr = opts.GetAtlasDeployments(projectId); atlasErr != nil {
		defer log.Warningf("Failed to retrieve Atlas deployments with: %s", atlasErr.Error())
		return opts.displayLocalDeployments(ctx)
	}

	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	localDeployments, err := opts.GetLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return err
	}

	if opts.DeploymentName == "" {
		return opts.Select(append(localDeployments, atlasDeployments...))
	}

	return opts.findDeployment(localDeployments, atlasDeployments)
}

func (opts *DeploymentOpts) findDeployment(localDeployments []Deployment, atlasDeployments []Deployment) error {
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

func (opts *DeploymentOpts) displayLocalDeployments(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	localDeployments, err := opts.GetLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return err
	}

	if opts.DeploymentName != "" {
		for _, d := range localDeployments {
			if d.Name == opts.DeploymentName {
				opts.DeploymentType = LocalCluster
				opts.DeploymentName = d.Name
				return nil
			}
		}
	}

	return opts.Select(localDeployments)
}

func (opts *DeploymentOpts) displayAtlasDeployments(ctx context.Context, projectID string) error {
	if !opts.IsCliAuthenticated() {
		return ErrNotAuthenticated
	}

	var atlasDeployments []Deployment
	var atlasErr error
	if atlasDeployments, atlasErr = opts.GetAtlasDeployments(projectID); atlasErr != nil {
		log.Warningf("Displaying only local deployments, failed to retrieve Atlas deployments with: %s\n", atlasErr.Error())
		return opts.displayLocalDeployments(ctx)
	}

	if opts.DeploymentName != "" {
		for _, d := range atlasDeployments {
			if d.Name == opts.DeploymentName {
				opts.DeploymentType = AtlasCluster
				opts.DeploymentName = d.Name
				return nil
			}
		}
	}

	return opts.Select(atlasDeployments)
}

func (opts *DeploymentOpts) GetAtlasDeployments(projectID string) ([]Deployment, error) {
	if projectID == "" {
		projectID = opts.Config.ProjectID()
	}

	if projectID == "" {
		if err := opts.DefaultSetter.AskProject(); err != nil {
			return nil, err
		}
		projectID = opts.DefaultSetter.ProjectID
	}

	listOpts := &mongodbatlas.ListOptions{
		PageNum:      cli.DefaultPage,
		ItemsPerPage: MaxItemsPerPage,
	}

	projectClusters, err := opts.AtlasClusterListStore.ProjectClusters(projectID, listOpts)
	if err != nil {
		return nil, err
	}
	atlasClusters := projectClusters.(*admin.PaginatedAdvancedClusterDescription)

	deployments := make([]Deployment, len(atlasClusters.Results))
	for i, c := range atlasClusters.Results {
		stateName := *c.StateName
		if *c.Paused {
			// for paused clusters, Atlas returns stateName IDLE and Paused=true
			stateName = PausedState
		}
		deployments[i] = Deployment{
			Type:           "ATLAS",
			Name:           *c.Name,
			MongoDBVersion: *c.MongoDBVersion,
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *DeploymentOpts) LocalDeploymentPreRun(ctx context.Context) error {
	if localDeploymentSupportedByOs() {
		_, _ = log.Warningln("Local deployments are not supported on this OS, to see local deployments requirements visit https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-deploy-local/.")
	}

	return opts.PodmanClient.Ready(ctx)
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
