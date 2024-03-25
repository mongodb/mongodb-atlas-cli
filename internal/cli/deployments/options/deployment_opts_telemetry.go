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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
)

//go:generate mockgen -destination=../../../mocks/mock_deployment_opts_telemetry.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options DeploymentTelemetry
type DeploymentTelemetry interface {
	AppendDeploymentType()
}

func NewDeploymentTypeTelemetry(opts *DeploymentOpts) DeploymentTelemetry {
	return opts
}

func (opts *DeploymentOpts) AppendDeploymentType() {
	var deploymentType string
	if opts.IsLocalDeploymentType() {
		deploymentType = LocalCluster
	} else if opts.IsAtlasDeploymentType() {
		deploymentType = AtlasCluster
	}
	if deploymentType != "" {
		telemetry.AppendOption(telemetry.WithDeploymentType(deploymentType))
	}
}
