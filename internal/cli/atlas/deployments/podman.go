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

package local

import (
	"os"
	"os/exec"
)

func runPodman(debug bool, arg ...string) error {
	cmd := exec.Command("podman", arg...)
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}

func createNetwork(debug bool, name string) error {
	return runPodman(debug, "network", "create", name)
}

func createVolume(debug bool, name string) error {
	return runPodman(debug, "volume", "create", name)
}

func runContainer(debug bool, arg ...string) error {
	return runPodman(debug, append([]string{"run"}, arg...)...)
}

func copyFileToContainer(debug bool, localFile string, containerName string, filePathInContainer string) error {
	return runPodman(debug, "cp", "create", localFile, containerName+":"+filePathInContainer)
}
