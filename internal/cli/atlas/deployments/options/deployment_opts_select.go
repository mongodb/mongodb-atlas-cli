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
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
)

var errEmptyLocalDeployments = errors.New("currently there are no deployment in your local system")
var errNoDeployments = errors.New("currently there are no deployments")
var ErrDeploymentNotFound = errors.New("deployment not found")
var errDeploymentRequiredOnPipe = fmt.Errorf("deployment name is required  when piping the output of the command")

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
