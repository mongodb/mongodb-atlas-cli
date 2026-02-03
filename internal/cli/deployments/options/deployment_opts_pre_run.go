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
	"runtime"
	"slices"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

var errDeploymentUnexpectedState = errors.New("deployment is in unexpected state")

func (opts *DeploymentOpts) listDeployments(ctx context.Context, projectID string) ([]Deployment, error) {
	var atlasDeployments, localDeployments []Deployment
	var atlasErr, localErr error

	if opts.IsAtlasDeploymentType() || opts.NoDeploymentTypeSet() {
		if atlasDeployments, atlasErr = opts.AtlasDeployments(projectID); atlasErr != nil {
			if opts.IsAtlasDeploymentType() {
				return nil, atlasErr
			}
			if !isUnauthenticatedErr(atlasErr) {
				_, _ = log.Warningf("Warning: failed to retrieve Atlas deployments because %q\n", atlasErr.Error())
			}
		}
	}

	if opts.IsLocalDeploymentType() || opts.NoDeploymentTypeSet() {
		if localErr = opts.LocalDeploymentPreRun(ctx); localErr != nil {
			if opts.IsLocalDeploymentType() {
				return nil, localErr
			}
			_, _ = log.Debugf("Warning: failed to retrieve Local deployments because %q\n", localErr.Error())
		}

		localDeployments, localErr = opts.GetLocalDeployments(ctx)
		if localErr != nil {
			if opts.IsLocalDeploymentType() {
				return nil, localErr
			}
			_, _ = log.Warningf("Warning: failed to retrieve Local deployments because %q\n", localErr.Error())
		}
	}

	if atlasErr != nil && localErr != nil {
		return nil, errors.New("failed to retrieve atlas and local deployments")
	}

	return append(localDeployments, atlasDeployments...), nil
}

func (opts *DeploymentOpts) SelectDeployments(ctx context.Context, projectID string, states ...string) (Deployment, error) {
	deployments, err := opts.listDeployments(ctx, projectID)
	if err != nil {
		return Deployment{}, errors.New("failed to retrieve atlas and local deployments")
	}

	var ok bool
	deployments, ok = opts.filterDeploymentByName(deployments...)
	if !ok {
		deployments, _ = opts.filterDeploymentByStateLocal(deployments, states...)
	}

	d, err := opts.Select(deployments)
	if err != nil {
		return Deployment{}, err
	}

	if d.Type == "LOCAL" && len(states) > 0 && !slices.Contains(states, d.StateName) {
		return Deployment{}, fmt.Errorf("%w: %s", errDeploymentUnexpectedState, d.StateName)
	}

	return d, nil
}

func isUnauthenticatedErr(err error) bool {
	if errors.Is(err, ErrNotAuthenticated) {
		return true
	}

	target, ok := atlasv2.AsError(err)
	return ok && target.GetReason() == "Unauthorized"
}

func (opts *DeploymentOpts) filterDeploymentByName(deployments ...Deployment) ([]Deployment, bool) {
	if opts.DeploymentName == "" {
		return deployments, false
	}

	filteredDeployments := []Deployment{}

	for _, d := range deployments {
		if d.Name == opts.DeploymentName {
			filteredDeployments = append(filteredDeployments, d)
		}
	}

	return filteredDeployments, true
}

func (*DeploymentOpts) filterDeploymentByStateLocal(deployments []Deployment, states ...string) ([]Deployment, bool) {
	if len(states) == 0 {
		return deployments, false
	}

	filteredDeployments := []Deployment{}

	for _, d := range deployments {
		if d.Type == "ATLAS" || slices.Contains(states, d.StateName) {
			filteredDeployments = append(filteredDeployments, d)
		}
	}

	return filteredDeployments, true
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

	atlasClusters, err := opts.AtlasClusterListStore.LatestProjectClusters(projectID, listOpts)
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

func (opts *DeploymentOpts) LocalDeploymentPreRun(ctx context.Context) error {
	if !localDeploymentSupportedByOs() {
		_, _ = log.Warningln("Local deployments are not supported on this OS, to see local deployments requirements visit https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-deploy-local/.")
	}

	if err := opts.ContainerEngine.Ready(); err != nil {
		return err
	}

	return opts.ContainerEngine.VerifyVersion(ctx)
}

func localDeploymentSupportedByOs() bool {
	os := runtime.GOOS
	switch os {
	case "darwin":
		// MacOS Intel and M1 are supported
		return true
	case "windows":
		// Windows is supported
		return true
	case "linux":
		// Linux is supported
		return true
	default:
		// Other unknown OS are not supported
		return false
	}
}
