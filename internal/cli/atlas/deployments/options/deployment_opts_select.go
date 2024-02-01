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
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/andreangiolillo/mongocli-test/internal/telemetry"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/containers/podman/v4/libpod/define"
)

var errEmptyLocalDeployments = errors.New("currently there are no deployment in your local system")
var errNoDeployments = errors.New("currently there are no deployments")
var ErrDeploymentNotFound = errors.New("deployment not found")
var errDeploymentRequiredOnPipe = fmt.Errorf("deployment name is required  when piping the output of the command")

func (opts *DeploymentOpts) findMongoDContainer(ctx context.Context) (*define.InspectContainerData, error) {
	containers, err := opts.PodmanClient.ContainerInspect(ctx, opts.LocalMongodHostname())
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDeploymentNotFound, opts.DeploymentName)
	}

	return containers[0], nil
}

func (opts *DeploymentOpts) CheckIfDeploymentExists(ctx context.Context) error {
	c, err := opts.findMongoDContainer(ctx)
	if err != nil {
		return err
	}

	opts.updateFields(c)
	return nil
}

func (opts *DeploymentOpts) DetectLocalDeploymentName(ctx context.Context) error {
	// before asking for deployment name, check if we are piping the output
	stat, _ := os.Stdout.Stat()
	if (stat.Mode()&os.ModeCharDevice) == 0 && opts.DeploymentName == "" {
		return errDeploymentRequiredOnPipe
	}

	if opts.DeploymentName != "" {
		return opts.CheckIfDeploymentExists(ctx)
	}
	return opts.SelectLocal(ctx)
}

func (opts *DeploymentOpts) SelectLocal(ctx context.Context) error {
	containers, err := opts.PodmanClient.ListContainers(ctx, MongodHostnamePrefix)
	if err != nil {
		return err
	}

	if len(containers) == 0 {
		return errEmptyLocalDeployments
	}

	if len(containers) == 1 {
		opts.DeploymentName = LocalDeploymentName(containers[0].Names[0])
		return nil
	}

	names := make([]string, 0, len(containers))
	for _, c := range containers {
		name := LocalDeploymentName(c.Names[0])
		names = append(names, name)
	}

	return telemetry.TrackAskOne(&survey.Select{
		Message: "Select a deployment",
		Options: names,
		Help:    usage.ClusterName,
		Description: func(value string, index int) string {
			return deploymentTypeLocal
		},
	}, &opts.DeploymentName, survey.WithValidator(survey.Required))
}

func (opts *DeploymentOpts) Select(deployments []Deployment) (Deployment, error) {
	if len(deployments) == 0 {
		return Deployment{}, errNoDeployments
	}

	if len(deployments) == 1 {
		opts.DeploymentName = deployments[0].Name
		opts.DeploymentType = deployments[0].Type

		telemetry.AppendOption(telemetry.WithDeploymentType(opts.DeploymentType))
		return deployments[0], nil
	}

	displayNames := make([]string, 0, len(deployments))
	deploymentsByDisplayName := map[string]Deployment{}

	for _, d := range deployments {
		displayType := strings.ToUpper(d.Type[:1]) + strings.ToLower(d.Type[1:])
		displayName := fmt.Sprintf("%s (%s)", d.Name, displayType)
		displayNames = append(displayNames, displayName)
		deploymentsByDisplayName[displayName] = d
	}

	var displayName string
	err := telemetry.TrackAskOne(&survey.Select{
		Message: "Select a deployment",
		Options: displayNames,
		Help:    usage.ClusterName,
	}, &displayName, survey.WithValidator(survey.Required))
	if err != nil {
		return Deployment{}, err
	}

	deployment := deploymentsByDisplayName[displayName]
	opts.DeploymentName = deployment.Name
	opts.DeploymentType = deployment.Type
	telemetry.AppendOption(telemetry.WithDeploymentType(opts.DeploymentType))
	return deployment, nil
}
