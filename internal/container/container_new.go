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

//go:build !linux
// +build !linux

package container

import (
	"context"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

func New() Engine {
	// Try Docker first
	docker := newDockerEngine()
	if isEngineReady(docker) {
		_, _ = log.Debug("Using Docker engine")
		return docker
	}

	// Try containerd second
	containerd := newContainerdEngine()
	if isEngineReady(containerd) {
		_, _ = log.Debug("Using containerd engine")
		return containerd
	}

	// If neither is ready, return Docker as the default (it will show appropriate error messages)
	_, _ = log.Debug("No engines ready, defaulting to Docker engine")
	return docker
}

// isEngineReady checks if an engine is both available and functional
func isEngineReady(engine Engine) bool {
	// First check if the binary is available
	if err := engine.Ready(); err != nil {
		return false
	}

	// Then check if the daemon/service is actually working
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := engine.VerifyVersion(ctx); err != nil {
		return false
	}

	return true
}
