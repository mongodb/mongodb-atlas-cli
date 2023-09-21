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
	"fmt"
	"strconv"

	"github.com/containers/podman/v4/libpod/define"
)

const deploymentTypeLocal = "local"

func (opts *DeploymentOpts) updateFields(c *define.InspectContainerData) {
	opts.DeploymentType = deploymentTypeLocal
	opts.MdbVersion = c.Config.Labels["version"]
	if len(c.HostConfig.PortBindings) > 0 {
		for _, port := range c.HostConfig.PortBindings {
			opts.Port, _ = strconv.Atoi(port[0].HostPort)
		}
	}
}

func (opts *DeploymentOpts) ConnectionString(ctx context.Context) (string, error) {
	if opts.Port == 0 {
		c, err := opts.findMongoDContainer(ctx)
		if err != nil {
			return "", err
		}
		opts.updateFields(c)
	}
	return fmt.Sprintf("mongodb://localhost:%d/?directConnection=true", opts.Port), nil
}
