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
	"fmt"
	"time"
)

const waitForRunningStateDuration = 10 * time.Second

var errCompassExited = errors.New("MongoDB Compass process has exited")

func Detect() bool {
	return binPath() != ""
}

func Run(username, password, mongoURI string) error {
	path := binPath()

	cmd := compassCmd(path, username, password, mongoURI)
	if err := cmd.Start(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), waitForRunningStateDuration)
	defer cancel()

	processExited := make(chan error)

	// Check if the process is still running
	go func() {
		if err := cmd.Wait(); err != nil {
			processExited <- fmt.Errorf("MongoDB Compass failed to start: %w", err)
		} else {
			processExited <- errCompassExited
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
