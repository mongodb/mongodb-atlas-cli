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

package compass

import (
	"context"
	"errors"
	"os/exec"
	"time"
)

const waitForRunningStateDuration = 10 * time.Second

func binPath() string {
	if p, err := exec.LookPath(compassBin); err == nil {
		return p
	}
	return ""
}

func Detect() bool {
	return binPath() != ""
}

func Run(username, password, mongoURI string) error {
	args := []string{mongoURI}
	if username != "" && password != "" {
		args = append(args, "--username", username, "--password", password)
	}

	path := binPath()
	if path != compassBin {
		path += compassBin
	}

	cmd := exec.Command(path, args...)
	if err := cmd.Start(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), waitForRunningStateDuration)
	defer cancel()

	processExited := make(chan error)

	// Check if the process is still running
	go func() {
		if err := cmd.Wait(); err != nil {
			processExited <- err
		} else {
			processExited <- errors.New("MongoDB Compass process has exited")
		}
	}()

	select {
	case <-ctx.Done():
		// compass still running
		return nil
	case err := <-processExited:
		return err
	}
}
