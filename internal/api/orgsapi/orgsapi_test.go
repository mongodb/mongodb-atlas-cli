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

package orgsapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const testBaseURL = "https://cloud.mongodb.com"

func TestFindCommand_ListOrgsExists(t *testing.T) {
	command, err := findCommand(listOrgsOperationID)
	require.NoError(t, err)
	assert.Equal(t, listOrgsOperationID, command.OperationID)
	assert.NotEmpty(t, command.Versions)
}

func TestFindCommand_Unknown(t *testing.T) {
	_, err := findCommand("thisOperationDoesNotExist")
	require.ErrorIs(t, err, ErrCommandNotFound)
}

// The generated definitions currently have len == cap, so append would reallocate even
// without the clone. Give the input spare capacity so the clone is load bearing.
func TestWithIncludeGlobal_DoesNotWriteIntoCallersBackingArray(t *testing.T) {
	backing := make([]shared_api.Parameter, 2, 4)
	backing[0] = shared_api.Parameter{Name: "name"}
	backing[1] = shared_api.Parameter{Name: "sentinel"}

	command := shared_api.Command{
		OperationID: listOrgsOperationID,
		RequestParameters: shared_api.RequestParameters{
			// Length 1 with spare capacity, so an uncloned append would overwrite backing[1].
			QueryParameters: backing[:1],
		},
	}

	got := withIncludeGlobal(command)

	require.Len(t, got.RequestParameters.QueryParameters, 2)
	assert.Equal(t, includeGlobalParam, got.RequestParameters.QueryParameters[1].Name)
	assert.Equal(t, "sentinel", backing[1].Name, "includeGlobal was written into the caller's backing array")
}

func TestListOrgsCommand_DoesNotMutateGlobalCommands(t *testing.T) {
	before, err := findCommand(listOrgsOperationID)
	require.NoError(t, err)
	originalCount := len(before.RequestParameters.QueryParameters)

	command, err := listOrgsCommand()
	require.NoError(t, err)
	assert.Len(t, command.RequestParameters.QueryParameters, originalCount+1)

	after, err := findCommand(listOrgsOperationID)
	require.NoError(t, err)
	assert.Len(t, after.RequestParameters.QueryParameters, originalCount)
	for _, parameter := range after.RequestParameters.QueryParameters {
		assert.NotEqual(t, includeGlobalParam, parameter.Name, "includeGlobal leaked into api.Commands")
	}
}

// The server defaults includeGlobal to true, so a regression here degrades silently.
func TestNewCommandRequest_QueryString(t *testing.T) {
	tests := []struct {
		name              string
		options           ListOrgsOptions
		wantIncludeGlobal string
		wantOmitted       bool
	}{
		{
			name:              "suppressed sends false",
			options:           ListOrgsOptions{IncludeGlobal: false},
			wantIncludeGlobal: "false",
		},
		{
			name:        "default omits the parameter entirely",
			options:     ListOrgsOptions{IncludeGlobal: true},
			wantOmitted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandRequest, err := NewCommandRequest(tt.options)
			require.NoError(t, err)

			httpRequest, err := api.ConvertToHTTPRequest(testBaseURL, commandRequest)
			require.NoError(t, err)

			query := httpRequest.URL.Query()
			if tt.wantOmitted {
				assert.False(t, query.Has(includeGlobalParam), "expected no includeGlobal in %q", httpRequest.URL.RawQuery)
				return
			}
			assert.Equal(t, tt.wantIncludeGlobal, query.Get(includeGlobalParam))
		})
	}
}

func TestNewCommandRequest_PassesThroughPagination(t *testing.T) {
	commandRequest, err := NewCommandRequest(ListOrgsOptions{
		Name:         "myOrg",
		PageNum:      2,
		ItemsPerPage: 50,
		IncludeCount: true,
	})
	require.NoError(t, err)

	httpRequest, err := api.ConvertToHTTPRequest(testBaseURL, commandRequest)
	require.NoError(t, err)

	query := httpRequest.URL.Query()
	assert.Equal(t, "myOrg", query.Get("name"))
	assert.Equal(t, "2", query.Get("pageNum"))
	assert.Equal(t, "50", query.Get("itemsPerPage"))
	assert.Equal(t, "true", query.Get("includeCount"))
}

func TestNewCommandRequest_OmitsEmptyValues(t *testing.T) {
	commandRequest, err := NewCommandRequest(ListOrgsOptions{})
	require.NoError(t, err)

	httpRequest, err := api.ConvertToHTTPRequest(testBaseURL, commandRequest)
	require.NoError(t, err)

	query := httpRequest.URL.Query()
	assert.False(t, query.Has("name"))
	assert.False(t, query.Has("pageNum"))
	assert.False(t, query.Has("itemsPerPage"))
	assert.False(t, query.Has("includeCount"))
}

func TestListOrgs(t *testing.T) {
	body := `{"totalCount":1,"results":[{"id":"5e2211c17a3e5a48f5497de3","name":"myOrg"}]}`

	ctrl := gomock.NewController(t)
	executor := api.NewMockCommandExecutor(ctrl)

	var got api.CommandRequest
	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, commandRequest api.CommandRequest) (*api.CommandResponse, error) {
			got = commandRequest
			return &api.CommandResponse{
				IsSuccess: true,
				HTTPCode:  http.StatusOK,
				Output:    io.NopCloser(strings.NewReader(body)),
			}, nil
		}).
		Times(1)

	result, err := ListOrgs(context.Background(), executor, ListOrgsOptions{Name: "myOrg"})
	require.NoError(t, err)
	assert.Equal(t, 1, result.GetTotalCount())
	require.Len(t, result.GetResults(), 1)
	assert.Equal(t, "myOrg", result.GetResults()[0].GetName())

	// Assert what ListOrgs itself put on the wire, not just what NewCommandRequest builds.
	httpRequest, err := api.ConvertToHTTPRequest(testBaseURL, got)
	require.NoError(t, err)
	query := httpRequest.URL.Query()
	assert.Equal(t, "false", query.Get(includeGlobalParam))
	assert.Equal(t, "myOrg", query.Get("name"))
}

func TestListOrgs_NilExecutor(t *testing.T) {
	_, err := ListOrgs(context.Background(), nil, ListOrgsOptions{})
	require.ErrorIs(t, err, ErrMissingExecutor)
}

func TestListOrgs_MalformedBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	executor := api.NewMockCommandExecutor(ctrl)
	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(&api.CommandResponse{
			IsSuccess: true,
			HTTPCode:  http.StatusOK,
			Output:    io.NopCloser(strings.NewReader(`{"totalCount":`)),
		}, nil).
		Times(1)

	_, err := ListOrgs(context.Background(), executor, ListOrgsOptions{})
	require.Error(t, err)
}

func TestLatestVersion_NoVersions(t *testing.T) {
	_, err := latestVersion(shared_api.Command{OperationID: listOrgsOperationID})
	require.ErrorIs(t, err, ErrNoVersions)
}

func TestListOrgs_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	executor := api.NewMockCommandExecutor(ctrl)
	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(&api.CommandResponse{
			IsSuccess: false,
			HTTPCode:  http.StatusNotFound,
			Output:    io.NopCloser(strings.NewReader(`{"error":404}`)),
		}, nil).
		Times(1)

	_, err := ListOrgs(context.Background(), executor, ListOrgsOptions{})
	require.ErrorIs(t, err, ErrNotFound)
}

func TestListOrgs_ErrorResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	executor := api.NewMockCommandExecutor(ctrl)
	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(&api.CommandResponse{
			IsSuccess: false,
			HTTPCode:  http.StatusUnauthorized,
			Output:    io.NopCloser(strings.NewReader(`{"detail":"nope"}`)),
		}, nil).
		Times(1)

	_, err := ListOrgs(context.Background(), executor, ListOrgsOptions{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "401")
	assert.Contains(t, err.Error(), "nope")
}

func TestListOrgs_ExecutorError(t *testing.T) {
	wantErr := errors.New("boom")

	ctrl := gomock.NewController(t)
	executor := api.NewMockCommandExecutor(ctrl)
	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(nil, wantErr).
		Times(1)

	_, err := ListOrgs(context.Background(), executor, ListOrgsOptions{})
	require.ErrorIs(t, err, wantErr)
}
