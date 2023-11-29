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
	"os"
	"os/exec"
)

func (opts *DeploymentOpts) RemoveLocal(ctx context.Context) error {
	buf, err := ComposeDefinition(&ComposeDefinitionOptions{
		Port:          "27017",
		MongodVersion: "7.0",
		BindIp:        "127.0.0.1",
	})
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "compose", "-f", "/dev/stdin", "down", "-v")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = buf
	cmd.Env = append(os.Environ(), "COMPOSE_PROJECT_NAME="+opts.DeploymentName, "KEY_FILE=keyfile")
	return cmd.Run()
}
