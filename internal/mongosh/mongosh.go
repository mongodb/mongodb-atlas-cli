// Copyright 2020 MongoDB Inc
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

package mongosh

import (
	"os"
	"os/exec"
	"syscall"
)

func Detect() bool {
	return binPath() != ""
}

func binPath() string {
	if p, err := exec.LookPath(mongoshBin); err == nil {
		return p
	}

	return ""
}

func execCommand(args ...string) error {
	a := append([]string{mongoshBin}, args...)
	env := os.Environ()
	return syscall.Exec(binPath(), a, env) //nolint:gosec // false positive, this path won't be tampered
}

func SetTelemetry(enable bool) error {
	cmd := "disableTelemetry()"
	if enable {
		cmd = "enableTelemetry()"
	}
	return execCommand("--nodb", "--eval", cmd)
}

func Run(username, password, mongoURI string) error {
	return execCommand("-u", username, "-p", password, mongoURI)
}

func Exec(debug bool, args ...string) error {
	cmd := exec.Command(mongoshBin, args...)
	cmd.Env = os.Environ()
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}
