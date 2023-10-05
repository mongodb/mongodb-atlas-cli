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

	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
)

func (opts *ConnectOpts) validateAndPromptLocalOpts(ctx context.Context) error {
	if opts.DeploymentName == "" {
		if err := opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	} else if err := opts.DeploymentOpts.CheckIfDeploymentExists(ctx); err != nil {
		return err
	}

	return nil
}

func (opts *ConnectOpts) connectToLocal(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	telemetry.AppendOption(telemetry.WithDeploymentType(LocalCluster))

	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	return opts.connectToDeployment(connectionString)
}
