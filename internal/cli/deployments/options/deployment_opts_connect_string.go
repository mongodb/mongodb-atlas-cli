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
	"net/url"
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
)

const deploymentTypeLocal = "local"

func (opts *DeploymentOpts) updateFields(c *container.InspectData) {
	opts.DeploymentType = deploymentTypeLocal
	opts.MdbVersion = c.Config.Labels["version"]
	portBind, ok := c.HostConfig.PortBindings["27017/tcp"]
	if ok && len(portBind) > 0 {
		opts.Port, _ = strconv.Atoi(portBind[0].HostPort)
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
	if opts.IsAuthEnabled() {
		return fmt.Sprintf("mongodb://%s:%s@localhost:%d/?directConnection=true",
			url.QueryEscape(opts.DBUsername),
			url.QueryEscape(opts.DBUserPassword),
			opts.Port,
		), nil
	}

	return fmt.Sprintf("mongodb://localhost:%d/?directConnection=true", opts.Port), nil
}
