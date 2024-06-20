// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var ErrDockerNotFound = errors.New("podman not found in your system, check requirements at https://dochub.mongodb.org/core/atlas-cli-deploy-local-reqs")

const filterBy = "name"

type dockerImpl struct {
}

func newDockerEngine() Engine {
	return &dockerImpl{}
}

func (*dockerImpl) Name() string {
	return "docker"
}

func (*dockerImpl) Ready(context.Context) error {
	if _, err := exec.LookPath("docker"); err != nil {
		return ErrDockerNotFound
	}
	return nil
}

func (*dockerImpl) run(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "docker", args...)
	return cmd.Output()
}

func (e *dockerImpl) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	buf, err := e.run(ctx, "container", "logs", name)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(buf), "\n"), nil
}

func portsFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := []string{}
	if flags.Ports != nil {
		host := "127.0.0.1:"
		if flags.BindIPAll != nil && *flags.BindIPAll {
			host = ""
		}
		for _, mapping := range flags.Ports {
			proto := ""
			if mapping.ContainerProtocol != "" {
				proto = "/" + mapping.ContainerProtocol
			}
			args = append(args, "-p", fmt.Sprintf("%s%d:%d%s", host, mapping.HostPort, mapping.ContainerPort, proto))
		}
	}
	return args
}

func runFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := portsFlags(flags)

	if flags.Detach != nil && *flags.Detach {
		args = append(args, "--detach")
	}

	if flags.Remove != nil && *flags.Remove {
		args = append(args, "--rm")
	}

	if flags.Hostname != nil {
		args = append(args, "--hostname", *flags.Hostname)
	}

	if flags.Env != nil {
		for key, value := range flags.Env {
			args = append(args, "-e", fmt.Sprintf("%s:%s", key, value))
		}
	}

	if flags.Network != nil {
		args = append(args, "--network", *flags.Network)
	}

	if flags.IP != nil {
		args = append(args, "--ip", *flags.IP)
	}

	if flags.Entrypoint != nil {
		args = append(args, "--entrypoint", *flags.Entrypoint)
	}

	return args
}

func (e *dockerImpl) ContainerRun(ctx context.Context, image string, flags *RunFlags) (string, error) {
	args := []string{"run"}
	args = append(args, runFlags(flags)...)
	args = append(args, image)
	if flags != nil && flags.Cmd != nil {
		args = append(args, *flags.Cmd)
	}

	if flags != nil && flags.Args != nil {
		args = append(args, flags.Args...)
	}

	buf, err := e.run(ctx, args...)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (e *dockerImpl) ContainerList(ctx context.Context, names ...string) ([]Container, error) {
	args := []string{"container", "ls", "--format", "json"}

	if len(names) > 0 {
		for _, name := range names {
			args = append(args, "-f", filterBy+"="+name)
		}
	}
	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	result := []Container{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *dockerImpl) ContainerRm(ctx context.Context, names ...string) error {
	args := []string{"container", "rm", "-v"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerStart(ctx context.Context, names ...string) error {
	args := []string{"container", "start"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerStop(ctx context.Context, names ...string) error {
	args := []string{"container", "stop"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerUnpause(ctx context.Context, names ...string) error {
	args := []string{"container", "unpause"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerInspect(ctx context.Context, names ...string) ([]*InspectData, error) {
	args := []string{"container", "inspect", "--format", "json"}
	args = append(args, names...)

	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	result := []*InspectData{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *dockerImpl) ImageList(ctx context.Context, names ...string) ([]Image, error) {
	args := []string{"image", "ls", "--format", "json"}

	if len(names) > 0 {
		for _, name := range names {
			args = append(args, "-f", filterBy+"="+name)
		}
	}
	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	result := []Image{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *dockerImpl) ImagePull(ctx context.Context, name string) error {
	_, err := e.run(ctx, "image", "pull", name)
	return err
}

func (e *dockerImpl) Version(ctx context.Context) (map[string]any, error) {
	buf, err := e.run(ctx, "version", "--format", "json")
	if err != nil {
		return nil, err
	}

	result := map[string]any{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}
