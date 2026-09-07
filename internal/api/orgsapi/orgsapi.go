// Copyright 2025 MongoDB Inc
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

// Package orgsapi lists organizations through the generated API command definitions so that
// the hidden includeGlobal query parameter can be sent.
//
// includeGlobal is declared in MMS as @Parameter(hidden = true) on listOrgs, so it is absent
// from the OpenAPI spec and therefore from both atlas-sdk-go and api.Commands. It defaults to
// true and only gates a global role check, making it a no-op for customers. See CLOUDP-432111.
package orgsapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312024/admin"
)

// Neither identifier is compiler checked: listOrgsOperationID must match api.Commands and
// includeGlobalParam must match MMS. Both are covered by tests.
const (
	listOrgsOperationID = "listOrgs"
	includeGlobalParam  = "includeGlobal"
	jsonFormat          = "json"
)

var (
	// ErrNotFound is returned when the API responds 404.
	ErrNotFound = errors.New("no organizations found")

	ErrCommandNotFound = errors.New("api command not found")
	ErrNoVersions      = errors.New("api command has no versions")
	ErrMissingExecutor = errors.New("command executor is nil")
)

// ListOrgsOptions are the supported listOrgs query parameters.
type ListOrgsOptions struct {
	Name         string
	PageNum      int
	ItemsPerPage int
	IncludeCount bool
	// IncludeGlobal is only ever sent as false; when true the parameter is omitted so the
	// server default applies.
	IncludeGlobal bool
}

// findCommand looks up a command by its operation ID in the generated definitions.
func findCommand(operationID string) (shared_api.Command, error) {
	for _, group := range api.Commands {
		for _, command := range group.Commands {
			if command.OperationID == operationID {
				return command, nil
			}
		}
	}

	return shared_api.Command{}, fmt.Errorf("%w: %q", ErrCommandNotFound, operationID)
}

// listOrgsCommand returns a copy of the generated listOrgs command with the includeGlobal
// parameter appended. api.Commands is global state backing the cobra tree, and copying a
// Command only copies the slice header, so the parameters are cloned before appending.
func listOrgsCommand() (shared_api.Command, error) {
	command, err := findCommand(listOrgsOperationID)
	if err != nil {
		return shared_api.Command{}, err
	}

	queryParameters := slices.Clone(command.RequestParameters.QueryParameters)
	queryParameters = append(queryParameters, shared_api.Parameter{
		Name:        includeGlobalParam,
		Description: "Flag that indicates whether organizations visible through a global role are returned.",
		Type:        shared_api.ParameterType{Type: shared_api.TypeBool},
	})
	command.RequestParameters.QueryParameters = queryParameters

	return command, nil
}

// latestVersion returns the most recent version of a command. Versions are sorted ascending
// by the generation tool.
func latestVersion(command shared_api.Command) (shared_api.Version, error) {
	if len(command.Versions) == 0 {
		return nil, fmt.Errorf("%w: %q", ErrNoVersions, command.OperationID)
	}

	return command.Versions[len(command.Versions)-1].Version, nil
}

func (o ListOrgsOptions) parameters() map[string][]string {
	parameters := map[string][]string{}

	if o.Name != "" {
		parameters["name"] = []string{o.Name}
	}
	if o.PageNum > 0 {
		parameters["pageNum"] = []string{strconv.Itoa(o.PageNum)}
	}
	if o.ItemsPerPage > 0 {
		parameters["itemsPerPage"] = []string{strconv.Itoa(o.ItemsPerPage)}
	}
	if o.IncludeCount {
		parameters["includeCount"] = []string{"true"}
	}
	if !o.IncludeGlobal {
		parameters[includeGlobalParam] = []string{"false"}
	}

	return parameters
}

// NewCommandRequest builds the listOrgs request.
func NewCommandRequest(o ListOrgsOptions) (api.CommandRequest, error) {
	command, err := listOrgsCommand()
	if err != nil {
		return api.CommandRequest{}, err
	}

	version, err := latestVersion(command)
	if err != nil {
		return api.CommandRequest{}, err
	}

	return api.CommandRequest{
		Command: command,
		Format:  jsonFormat,
		// Executor.SetContentType would derive this, but it is needed for direct conversion.
		ContentType: jsonFormat,
		Parameters:  o.parameters(),
		Version:     version,
	}, nil
}

// ListOrgs returns the organizations the caller has access to.
func ListOrgs(ctx context.Context, executor api.CommandExecutor, o ListOrgsOptions) (*atlasv2.PaginatedOrganization, error) {
	if executor == nil {
		return nil, ErrMissingExecutor
	}

	commandRequest, err := NewCommandRequest(o)
	if err != nil {
		return nil, err
	}

	response, err := executor.ExecuteCommand(ctx, commandRequest)
	if err != nil {
		return nil, err
	}
	defer response.Output.Close()

	if !response.IsSuccess {
		if response.HTTPCode == http.StatusNotFound {
			return nil, ErrNotFound
		}

		body, _ := io.ReadAll(response.Output)
		return nil, fmt.Errorf("failed to list organizations (HTTP %d): %s", response.HTTPCode, body)
	}

	var result atlasv2.PaginatedOrganization
	if err := json.NewDecoder(response.Output).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
