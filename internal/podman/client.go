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
	"strings"

	"github.com/containers/podman/v4/pkg/machine"
)

var (
	ErrPodmanNotFound  = errors.New("podman not found in your system, check requirements at https://dochub.mongodb.org/core/atlas-cli-deploy-local-reqs")
	ErrNetworkNotFound = errors.New("network ip range was not found")
)

type Diagnostic struct {
	Installed    bool
	MachineFound bool
	MachineState string
	MachineInfo  *machine.InspectInfo
	Version      *Version
	Images       []string
	Errors       []string
}

const PodmanRunningState = machine.Running

type RunContainerOpts struct {
	Detach   bool
	Remove   bool
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
	IP         string
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

type Image struct {
	ID          string   `json:"ID"`
	RepoTags    string   `json:"RepoTags"`
	RepoDigests []string `json:"RepoDigests"`
	Created     int      `json:"Created"`
	CreatedAt   string   `json:"CreatedAt"`
	Size        int      `json:"Size"`
	SharedSize  int      `json:"SharedSize"`
	VirtualSize int      `json:"VirtualSize"`
	Labels      struct {
		Architecture string `json:"architecture"`
		BuildDate    string `json:"build-date"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Version      string `json:"version"`
	} `json:"Labels"`
	Containers int `json:"Containers"`
	Names      []string
}

type Version struct {
	Client struct {
		APIVersion string `json:"APIVersion"`
		Version    string `json:"Version"`
		GoVersion  string `json:"GoVersion"`
		GitCommit  string `json:"GitCommit"`
		BuiltTime  string `json:"BuiltTime"`
		Built      int    `json:"Built"`
		OsArch     string `json:"OsArch"`
		Os         string `json:"Os"`
	} `json:"Client"`

	Server struct {
		APIVersion string `json:"APIVersion"`
		Version    string `json:"Version"`
		GoVersion  string `json:"GoVersion"`
		GitCommit  string `json:"GitCommit"`
		BuiltTime  string `json:"BuiltTime"`
		Built      int    `json:"Built"`
		OsArch     string `json:"OsArch"`
		Os         string `json:"Os"`
	} `json:"Server"`
}

//go:generate mockgen -destination=../mocks/mock_podman.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/podman Client

type Client interface {
	Ready(ctx context.Context) error
	Diagnostics(ctx context.Context) *Diagnostic
	CreateNetwork(ctx context.Context, name string) ([]byte, error)
	CreateVolume(ctx context.Context, name string) ([]byte, error)
	RunContainer(ctx context.Context, opts RunContainerOpts) ([]byte, error)
	CopyFileToContainer(ctx context.Context, localFile string, containerName string, filePathInContainer string) ([]byte, error)
	StopContainers(ctx context.Context, names ...string) ([]byte, error)
	RemoveContainers(ctx context.Context, names ...string) ([]byte, error)
	RemoveVolumes(ctx context.Context, names ...string) ([]byte, error)
	RemoveNetworks(ctx context.Context, names ...string) ([]byte, error)
	ListContainers(ctx context.Context, nameFilter string) ([]*Container, error)
	ListImages(ctx context.Context, nameFilter string) ([]*Image, error)
	PullImage(ctx context.Context, name string) ([]byte, error)
	Version(ctx context.Context) (*Version, error)
	Logs(ctx context.Context) ([]interface{}, error)
	ContainerLogs(ctx context.Context, name string) ([]string, error)
	Network(ctx context.Context, name string) (*Network, error)
	Exec(ctx context.Context, name string, args ...string) error
}

type Network struct {
	ID         string `json:"ID"`
	Name       string `json:"Name"`
	DNSEnabled string `json:"DNSEnabled"`
	Subnets    []struct {
		Subnet  string `json:"Subnet"`
		Gateway string `json:"gateway"`
	} `json:"Subnets"`
}

type client struct {
	debug     bool
	outWriter io.Writer
}

func (o *client) Diagnostics(ctx context.Context) *Diagnostic {
	d := &Diagnostic{
		Installed:    true,
		MachineFound: true,
	}

	err := Installed()
	if err != nil {
		d.Installed = false
		d.Errors = append(d.Errors, fmt.Errorf("failed to detect podman installed: %w", err).Error())
	}

	d.Version, err = o.Version(ctx)
	if err != nil {
		d.Errors = append(d.Errors, fmt.Errorf("failed to collect podman version: %w", err).Error())
	}

	info, err := o.machineInspect(ctx)
	if err != nil {
		d.MachineFound = false
		d.Errors = append(d.Errors, fmt.Errorf("failed to detect podman machine: %w", err).Error())
	} else {
		d.MachineInfo = info
		d.MachineState = info.State
	}

	images, err := o.ListImages(ctx, "")
	if err != nil {
		d.Errors = append(d.Errors, fmt.Errorf("failed to list podman images: %w", err).Error())
	} else {
		d.Images = make([]string, 0, len(images))
		for _, img := range images {
			d.Images = append(d.Images, img.Names...)
		}
	}
	return d
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
	b, err := o.runPodman(ctx, "machine", "inspect", machine.DefaultMachineName)
	if err != nil {
		return nil, err
	}
	var info []machine.InspectInfo
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
	if info.State != PodmanRunningState {
		_, err := o.runPodman(ctx, "machine", "start")
		if err != nil {
			return err
		}
	}

	return nil
}

func Installed() error {
	if _, err := exec.LookPath("podman"); err != nil {
		return ErrPodmanNotFound
	}
	return nil
}

func (o *client) Ready(ctx context.Context) error {
	if err := Installed(); err != nil {
		return err
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

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
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

	if opts.IP != "" {
		arg = append(arg, "--ip", opts.IP)
	}

	if opts.Detach {
		arg = append(arg, "-d")
	}

	if opts.Remove {
		arg = append(arg, "--rm")
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

func (o *client) ListContainers(ctx context.Context, nameFilter string) ([]*Container, error) {
	args := []string{"ps", "--all", "--format", "json"}
	if nameFilter != "" {
		args = append(args, "--filter", "name="+nameFilter)
	}

	response, err := o.runPodman(ctx, args...)
	if err != nil {
		return nil, err
	}

	var containers []*Container
	err = json.Unmarshal(response, &containers)
	return containers, err
}

func (o *client) ListImages(ctx context.Context, nameFilter string) ([]*Image, error) {
	args := []string{"image", "list", "--format", "json"}

	if nameFilter != "" {
		args = append(args, "--filter", "reference="+nameFilter)
	}

	response, err := o.runPodman(ctx, args...)
	if err != nil {
		return nil, err
	}

	var images []*Image
	if err = json.Unmarshal(response, &images); err != nil {
		return nil, err
	}
	return images, err
}

func (o *client) PullImage(ctx context.Context, name string) ([]byte, error) {
	return o.runPodman(ctx, "pull", name)
}

func (o *client) Version(ctx context.Context) (version *Version, err error) {
	output, err := o.runPodman(ctx, "version", "--format", "json")
	if err != nil {
		return nil, err
	}

	var v *Version
	if err = json.Unmarshal(output, &v); err != nil {
		return nil, err
	}
	return v, err
}

func (o *client) Logs(ctx context.Context) ([]interface{}, error) {
	output, err := o.runPodman(ctx, "ps", "--format", "json")
	if err != nil {
		return nil, err
	}

	var podmanLogs []interface{}

	err = json.Unmarshal(output, &podmanLogs)
	if err != nil {
		return podmanLogs, err
	}

	return podmanLogs, nil
}

func (o *client) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	output, err := o.runPodman(ctx, "container", "logs", name)
	if err != nil {
		return []string{""}, err
	}

	logs := strings.Split(string(output), "\n")
	return logs, nil
}

func (o *client) Network(ctx context.Context, name string) (*Network, error) {
	output, err := o.runPodman(ctx, "network", "inspect", name, "--format", "json")
	if err != nil {
		return nil, err
	}

	var n []*Network
	if err = json.Unmarshal(output, &n); err != nil {
		return nil, err
	}

	if len(n) == 0 {
		return nil, ErrNetworkNotFound
	}

	return n[0], err
}

func (o *client) Exec(ctx context.Context, name string, args ...string) error {
	_, err := o.runPodman(ctx, append([]string{"exec", name}, args...)...)
	return err
}

func NewClient(debug bool, outWriter io.Writer) Client {
	return &client{
		debug:     debug,
		outWriter: outWriter,
	}
}
