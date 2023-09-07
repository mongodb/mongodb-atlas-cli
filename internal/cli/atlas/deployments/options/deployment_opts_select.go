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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
)

var errEmptyDeployments = errors.New("currently there are no deployment in your local system")
var errDeploymentNotFound = errors.New("deployment not found")

func (opts *DeploymentOpts) CheckIfDeploymentExists(ctx context.Context) error {
	containers, err := opts.podmanClient.ListContainers(ctx, MongodHostnamePrefix)
	if err != nil {
		return err
	}

	found := false
	for _, c := range containers {
		for _, n := range c.Names {
			if n == opts.LocalMongodHostname() {
				found = true
			}
		}
	}

	if !found {
		return fmt.Errorf("%w: %s", errDeploymentNotFound, opts.DeploymentName)

	}

	return nil
}

func (opts *DeploymentOpts) Select(ctx context.Context) error {
	containers, err := opts.podmanClient.ListContainers(ctx, MongodHostnamePrefix)
	if err != nil {
		return err
	}

	if len(containers) == 0 {
		return errEmptyDeployments
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
			return "local"
		},
	}, &opts.DeploymentName, survey.WithValidator(survey.Required))
}
