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
	"encoding/json"
	"io"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/containers/podman/v4/pkg/machine"
)

type RunContainerOpts struct {
	Detach   bool
	Image    string
	Name     string
	Hostname string
	// map[hostVolume, pathInContainer]
	Volumes map[string]string
	// map[hostPort, containerPort]
	Ports   map[int]int
	Network string
	EnvVars map[string]string
	Args    []string
}

type Container struct {
	Names []string `json:"Names"`
	State string   `json:"State"`
	Image string   `json:"Image"`
	Ports []struct {
		HostPort      int `json:"host_port"`
		ContainerPort int `json:"container_port"`
	} `json:"Ports,omitempty"`
}

type Client interface {
	Ready() bool
	Setup() error
	CreateNetwork(name string) ([]byte, error)
	CreateVolume(name string) ([]byte, error)
	RunContainer(opts RunContainerOpts) ([]byte, error)
	CopyFileToContainer(localFile string, containerName string, filePathInContainer string) ([]byte, error)
	StopContainers(names ...string) ([]byte, error)
	RemoveContainers(names ...string) ([]byte, error)
	RemoveVolumes(names ...string) ([]byte, error)
	RemoveNetworks(names ...string) ([]byte, error)
	ListContainers(nameFilter string) ([]Container, error)
}

type client struct {
	debug     bool
	outWriter io.Writer
}

func (o *client) Ready() bool {
	_, err := exec.LookPath("podman")
	return err == nil
}

func (o *client) machineInit() error {
	_, err := o.machineInspect()
	if err == nil { // machine is already present
		return nil
	}

	_, err = o.runPodman("machine", "init")
	return err
}

func (o *client) machineInspect() (*machine.InspectInfo, error) {
	var info []machine.InspectInfo
	b, err := o.runPodman("machine", "inspect", machine.DefaultMachineName)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &info); err != nil {
		return nil, err
	}
	return &info[0], nil
}

func (o *client) machineStart() error {
	info, err := o.machineInspect()
	if err != nil {
		return err
	}
	if info.State != machine.Running {
		_, err := o.runPodman("machine", "start", machine.DefaultMachineName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *client) Setup() error {
	if runtime.GOOS != "windows" && runtime.GOOS != "darwin" {
		// macOs and Windows require VMs
		return nil
	}

	if err := o.machineInit(); err != nil {
		return err
	}

	return o.machineStart()
}

func (o *client) runPodman(arg ...string) ([]byte, error) {
	cmd := exec.Command("podman", arg...)

	output, err := cmd.CombinedOutput()

	if o.debug {
		_, _ = o.outWriter.Write(output)
	}

	return output, err
}

func (o *client) CreateNetwork(name string) ([]byte, error) {
	return o.runPodman("network", "create", name)
}

func (o *client) CreateVolume(name string) ([]byte, error) {
	return o.runPodman("volume", "create", name)
}

func (o *client) RunContainer(opts RunContainerOpts) ([]byte, error) {
	arg := []string{"run",
		"--name", opts.Name,
		"--hostname", opts.Hostname,
		"--network", opts.Network,
	}

	for hostVolume, pathInContainer := range opts.Volumes {
		arg = append(arg, "-v", hostVolume+":"+pathInContainer)
	}

	for hostPort, containerPort := range opts.Ports {
		arg = append(arg, "-p", strconv.Itoa(hostPort)+":"+strconv.Itoa(containerPort))
	}

	for envVar, value := range opts.EnvVars {
		arg = append(arg, "-e", envVar+"="+value)
	}

	if opts.Detach {
		arg = append(arg, "-d")
	}

	arg = append(arg, opts.Image)

	arg = append(arg, opts.Args...)

	return o.runPodman(arg...)
}

func (o *client) CopyFileToContainer(localFile string, containerName string, filePathInContainer string) ([]byte, error) {
	return o.runPodman("cp", localFile, containerName+":"+filePathInContainer)
}

func (o *client) StopContainers(names ...string) ([]byte, error) {
	return o.runPodman(append([]string{"stop"}, names...)...)
}

func (o *client) RemoveContainers(names ...string) ([]byte, error) {
	return o.runPodman(append([]string{"rm", "-f"}, names...)...)
}

func (o *client) RemoveVolumes(names ...string) ([]byte, error) {
	return o.runPodman(append([]string{"volume", "rm"}, names...)...)
}

func (o *client) RemoveNetworks(names ...string) ([]byte, error) {
	return o.runPodman(append([]string{"network", "rm"}, names...)...)
}

func (o *client) ListContainers(nameFilter string) ([]Container, error) {
	response, err := o.runPodman("ps", "--all", "--format", "json", "--filter", "name="+nameFilter)
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
