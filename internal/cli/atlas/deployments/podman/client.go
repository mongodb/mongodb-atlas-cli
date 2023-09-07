// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package podman

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/containers/podman/v4/pkg/machine"
)

var ErrPodmanNotReady = errors.New("podman not found in your system, check requirements at http://docpage")

type RunContainerOpts struct {
	Detach   bool
	Image    string
	Name     string
	Hostname string
	// map[hostVolume, pathInContainer]
	Volumes map[string]string
	// map[hostPort, containerPort]
	Ports      map[int]int
	Network    string
	EnvVars    map[string]string
	Args       []string
	Entrypoint string
	Cmd        string
}

type Container struct {
	ID    string   `json:"ID"`
	Names []string `json:"Names"`
	State string   `json:"State"`
	Image string   `json:"Image"`
	Ports []struct {
		HostPort      int `json:"host_port"`
		ContainerPort int `json:"container_port"`
	} `json:"Ports,omitempty"`
	Labels map[string]string `json:"Labels"`
}

//go:generate mockgen -destination=../../../../mocks/mock_podman.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman Client

type Client interface {
	Ready(ctx context.Context) error
	CreateNetwork(ctx context.Context, name string) ([]byte, error)
	CreateVolume(ctx context.Context, name string) ([]byte, error)
	RunContainer(ctx context.Context, opts RunContainerOpts) ([]byte, error)
	CopyFileToContainer(ctx context.Context, localFile string, containerName string, filePathInContainer string) ([]byte, error)
	StopContainers(ctx context.Context, names ...string) ([]byte, error)
	RemoveContainers(ctx context.Context, names ...string) ([]byte, error)
	RemoveVolumes(ctx context.Context, names ...string) ([]byte, error)
	RemoveNetworks(ctx context.Context, names ...string) ([]byte, error)
	ListContainers(ctx context.Context, nameFilter string) ([]Container, error)
}

type client struct {
	debug     bool
	outWriter io.Writer
}

func (o *client) machineInit(ctx context.Context) error {
	_, err := o.machineInspect(ctx)
	if err == nil { // machine is already present
		return nil
	}

	_, err = o.runPodman(ctx, "machine", "init")
	return err
}

func (o *client) machineInspect(ctx context.Context) (*machine.InspectInfo, error) {
	var info []machine.InspectInfo
	b, err := o.runPodman(ctx, "machine", "inspect", machine.DefaultMachineName)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &info); err != nil {
		return nil, err
	}
	return &info[0], nil
}

func (o *client) machineStart(ctx context.Context) error {
	info, err := o.machineInspect(ctx)
	if err != nil {
		return err
	}
	if info.State != machine.Running {
		_, err := o.runPodman(ctx, "machine", "start", machine.DefaultMachineName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *client) Ready(ctx context.Context) error {
	if _, err := exec.LookPath("podman"); err != nil {
		return ErrPodmanNotReady
	}

	if runtime.GOOS != "windows" && runtime.GOOS != "darwin" {
		// macOs and Windows require VMs
		return nil
	}

	if err := o.machineInit(ctx); err != nil {
		return err
	}

	return o.machineStart(ctx)
}

func (o *client) runPodman(ctx context.Context, arg ...string) ([]byte, error) {
	if o.debug {
		_, _ = o.outWriter.Write([]byte(fmt.Sprintln(append([]string{"podman"}, arg...))))
	}

	cmd := exec.CommandContext(ctx, "podman", arg...)

	output, err := cmd.Output() // ignore stderr

	if o.debug {
		_, _ = o.outWriter.Write(output)
		if exitErr, ok := err.(*exec.ExitError); ok {
			_, _ = o.outWriter.Write(exitErr.Stderr)
		}
	}

	return output, err
}

func (o *client) CreateNetwork(ctx context.Context, name string) ([]byte, error) {
	return o.runPodman(ctx, "network", "create", name)
}

func (o *client) CreateVolume(ctx context.Context, name string) ([]byte, error) {
	return o.runPodman(ctx, "volume", "create", name)
}

func (o *client) RunContainer(ctx context.Context, opts RunContainerOpts) ([]byte, error) {
	arg := []string{"run",
		"--name", opts.Name,
		"--hostname", opts.Hostname,
		"--network", opts.Network,
	}

	for hostVolume, pathInContainer := range opts.Volumes {
		arg = append(arg, "-v", hostVolume+":"+pathInContainer)
	}

	for hostPort, containerPort := range opts.Ports {
		arg = append(arg, "-p", "127.0.0.1:"+strconv.Itoa(hostPort)+":"+strconv.Itoa(containerPort))
	}

	for envVar, value := range opts.EnvVars {
		arg = append(arg, "-e", envVar+"="+value)
	}

	if opts.Detach {
		arg = append(arg, "-d")
	}

	if opts.Entrypoint != "" {
		arg = append(arg, "--entrypoint", opts.Entrypoint)
	}

	arg = append(arg, opts.Image)

	if opts.Cmd != "" {
		arg = append(arg, opts.Cmd)
	}

	arg = append(arg, opts.Args...)

	return o.runPodman(ctx, arg...)
}

func (o *client) CopyFileToContainer(ctx context.Context, localFile string, containerName string, filePathInContainer string) ([]byte, error) {
	return o.runPodman(ctx, "cp", localFile, containerName+":"+filePathInContainer)
}

func (o *client) StopContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"stop"}, names...)...)
}

func (o *client) RemoveContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"rm", "-f"}, names...)...)
}

func (o *client) RemoveVolumes(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"volume", "rm", "-f"}, names...)...)
}

func (o *client) RemoveNetworks(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"network", "rm", "-f"}, names...)...)
}

func (o *client) ListContainers(ctx context.Context, nameFilter string) ([]Container, error) {
	response, err := o.runPodman(ctx, "ps", "--all", "--format", "json", "--filter", "name="+nameFilter)
	if err != nil {
		return nil, err
	}

	var containers []Container
	err = json.Unmarshal(response, &containers)
	return containers, err
}

func NewClient(debug bool, outWriter io.Writer) Client {
	return &client{
		debug:     debug,
		outWriter: outWriter,
	}
}
