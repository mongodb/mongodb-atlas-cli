// Copyright 2026 MongoDB Inc
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

package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

// commandInfo describes the `atlas api` command that implements an operation.
type commandInfo struct {
	group                 string
	name                  string
	params                map[string]shared_api.Parameter
	versions              map[string]bool
	watcherGetOperationID string
	hasWatcher            bool
}

func buildCommandIndex() map[string]commandInfo {
	index := make(map[string]commandInfo)

	for _, group := range api.Commands {
		groupName := strcase.ToLowerCamel(group.Name)
		for _, command := range group.Commands {
			if _, exists := index[command.OperationID]; exists {
				continue
			}

			name := strcase.ToLowerCamel(command.OperationID)
			if command.ShortOperationID != "" {
				name = strcase.ToLowerCamel(command.ShortOperationID)
			}

			params := make(map[string]shared_api.Parameter)
			for _, parameter := range append(command.RequestParameters.URLParameters, command.RequestParameters.QueryParameters...) {
				params[parameter.Name] = parameter
			}

			versions := make(map[string]bool)
			for _, version := range command.Versions {
				versions[version.Version.String()] = true
			}

			info := commandInfo{
				group:    groupName,
				name:     name,
				params:   params,
				versions: versions,
			}
			if command.Watcher != nil && command.Watcher.Get.OperationID != "" {
				info.hasWatcher = true
				info.watcherGetOperationID = command.Watcher.Get.OperationID
			}

			index[command.OperationID] = info
		}
	}

	return index
}

func (c commandInfo) parameter(name string) (shared_api.Parameter, error) {
	parameter, ok := c.params[name]
	if !ok {
		return shared_api.Parameter{}, fmt.Errorf("parameter %q is not declared on the command", name)
	}

	return parameter, nil
}

func (c commandInfo) supportsVersion(version string) bool {
	return c.versions[version]
}
