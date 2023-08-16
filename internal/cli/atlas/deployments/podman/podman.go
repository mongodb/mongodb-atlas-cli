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
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func runPodman(debug bool, arg ...string) error {
	cmd := exec.Command("podman", arg...)

	if debug {
		fmt.Println("\n>>>  podman", strings.Join(arg, " "), "\n")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}

func CreateNetwork(debug bool, name string) error {
	return runPodman(debug, "network", "create", name)
}

func CreateVolume(debug bool, name string) error {
	return runPodman(debug, "volume", "create", name)
}

func RunContainer(debug bool, opts RunContainerOpts) error {
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

	return runPodman(debug, arg...)
}

func CopyFileToContainer(debug bool, localFile string, containerName string, filePathInContainer string) error {
	return runPodman(debug, "cp", localFile, containerName+":"+filePathInContainer)
}

func StopContainers(debug bool, names ...string) error {
	return runPodman(debug, append([]string{"stop"}, names...)...)
}

func RemoveContainers(debug bool, names ...string) error {
	return runPodman(debug, append([]string{"rm", "-f"}, names...)...)
}

func RemoveVolumes(debug bool, names ...string) error {
	return runPodman(debug, append([]string{"volume", "rm"}, names...)...)
}

func RemoveNetworks(debug bool, names ...string) error {
	return runPodman(debug, append([]string{"network", "rm"}, names...)...)
}
