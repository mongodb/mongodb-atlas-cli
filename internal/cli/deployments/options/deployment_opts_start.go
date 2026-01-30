// Copyright 2024 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/connect"
)

var ErrDeploymentIsDeleting = errors.New("deployment state is DELETING")

func (opts *DeploymentOpts) StartLocal(ctx context.Context, deployment Deployment) error {
	if deployment.StateName == IdleState || deployment.StateName == RestartingState {
		return nil
	}

	if deployment.StateName == connect.StoppedState {
		return opts.ContainerEngine.ContainerStart(ctx, opts.LocalMongodHostname())
	}

	if deployment.StateName == connect.PausedState {
		return opts.ContainerEngine.ContainerUnpause(ctx, opts.LocalMongodHostname())
	}

	return ErrDeploymentIsDeleting
}
