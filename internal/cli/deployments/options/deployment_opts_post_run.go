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

package options

import (
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

func (opts *DeploymentOpts) PostRunMessages() error {
	if !opts.IsCliAuthenticated() {
		if _, err := log.Warningln("\nTo list both local and cloud Atlas deployments, authenticate to your Atlas account using the \"atlas login\" command."); err != nil {
			return err
		}
	}

	if err := opts.ContainerEngine.Ready(); errors.Is(err, container.ErrContainerEngineNotFound) {
		if _, err = log.Warningln("\nTo get output for both local and Atlas deployments, install " + opts.ContainerEngine.Name()); err != nil {
			return err
		}
	}
	return nil
}
