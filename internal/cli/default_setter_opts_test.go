// Copyright 2022 MongoDB Inc
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

package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312024/admin"
	"go.uber.org/mock/gomock"
)

func TestDefaultOpts_DefaultQuestions(t *testing.T) {
	type fields struct {
		Service string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "cloud",
			fields: fields{
				Service: "cloud",
			},
			want: 1,
		},
		{
			name: "cloud gov",
			fields: fields{
				Service: "cloudgov",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &DefaultSetterOpts{
				Service: tt.fields.Service,
			}
			assert.Len(t, opts.DefaultQuestions(), tt.want)
		})
	}
}

func TestDefaultOpts_Projects(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	opts := &DefaultSetterOpts{
		Service: "cloud",
		Store:   mockStore,
	}
	t.Run("empty", func(t *testing.T) {
		expectedProjects := &atlasv2.PaginatedAtlasGroup{}
		mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(1)
		_, _, err := opts.projects()
		require.Error(t, err)
	})
	t.Run("with one project", func(t *testing.T) {
		expectedProjects := &atlasv2.PaginatedAtlasGroup{
			Results: []atlasv2.Group{
				{
					Id:   pointer.Get("1"),
					Name: "Project 1",
				},
			},
			TotalCount: pointer.Get(1),
		}
		mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(1)
		gotIDs, gotNames, err := opts.projects()
		require.NoError(t, err)
		assert.Equal(t, []string{"1"}, gotIDs)
		assert.Equal(t, []string{"Project 1"}, gotNames)
	})
	t.Run("too many projects", func(t *testing.T) {
		expectedProjects := &atlasv2.PaginatedAtlasGroup{
			Results:    []atlasv2.Group{},
			TotalCount: pointer.Get(resultsLimit + 1),
		}
		mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(1)
		_, _, err := opts.projects()
		require.ErrorIs(t, err, errTooManyResults)
	})
}

// expectOrgs stubs a successful listOrgs response carrying the given payload.
func expectOrgs(t *testing.T, executor *api.MockCommandExecutor, payload any) {
	t.Helper()

	body, err := json.Marshal(payload)
	require.NoError(t, err)

	executor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(&api.CommandResponse{
			IsSuccess: true,
			HTTPCode:  http.StatusOK,
			Output:    io.NopCloser(bytes.NewReader(body)),
		}, nil).
		Times(1)
}

func TestDefaultOpts_Orgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutor := api.NewMockCommandExecutor(ctrl)

	opts := &DefaultSetterOpts{
		Service:     "cloud",
		OrgExecutor: mockExecutor,
	}
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		expectOrgs(t, mockExecutor, &atlasv2.PaginatedOrganization{})
		_, err := opts.orgs(ctx, "")
		require.Error(t, err)
	})
	t.Run("with one org", func(t *testing.T) {
		expectedOrgs := &atlasv2.PaginatedOrganization{
			Results: []atlasv2.AtlasOrganization{
				{
					Id:   pointer.Get("1"),
					Name: "Org 1",
				},
			},
			TotalCount: pointer.Get(1),
		}
		expectOrgs(t, mockExecutor, expectedOrgs)
		gotOrgs, err := opts.orgs(ctx, "")
		require.NoError(t, err)
		assert.Equal(t, expectedOrgs.GetResults(), gotOrgs)
	})

	t.Run("with no org", func(t *testing.T) {
		expectOrgs(t, mockExecutor, &atlasv2.PaginatedOrganization{
			Results: []atlasv2.AtlasOrganization{},
		})
		_, err := opts.orgs(ctx, "")
		require.Error(t, err)
		require.EqualError(t, err, errNoResults.Error())
	})

	t.Run("with null body", func(t *testing.T) {
		expectOrgs(t, mockExecutor, nil)
		_, err := opts.orgs(ctx, "")
		require.Error(t, err)
		require.EqualError(t, err, errNoResults.Error())
	})

	t.Run("404 maps to no results", func(t *testing.T) {
		mockExecutor.EXPECT().
			ExecuteCommand(gomock.Any(), gomock.Any()).
			Return(&api.CommandResponse{
				IsSuccess: false,
				HTTPCode:  http.StatusNotFound,
				Output:    io.NopCloser(strings.NewReader(`{}`)),
			}, nil).
			Times(1)
		_, err := opts.orgs(ctx, "")
		require.ErrorIs(t, err, errNoResults)
	})

	t.Run("too many orgs", func(t *testing.T) {
		expectOrgs(t, mockExecutor, &atlasv2.PaginatedOrganization{
			Results:    []atlasv2.AtlasOrganization{},
			TotalCount: pointer.Get(resultsLimit + 1),
		})
		_, err := opts.orgs(ctx, "")
		require.ErrorIs(t, err, errTooManyResults)
	})
}

// The server defaults includeGlobal to true, so a regression here degrades silently.
func TestDefaultOpts_Orgs_SuppressesGlobalOrgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutor := api.NewMockCommandExecutor(ctrl)

	var got api.CommandRequest
	mockExecutor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, commandRequest api.CommandRequest) (*api.CommandResponse, error) {
			got = commandRequest
			return &api.CommandResponse{
				IsSuccess: true,
				HTTPCode:  http.StatusOK,
				Output:    io.NopCloser(strings.NewReader(`{"totalCount":0,"results":[]}`)),
			}, nil
		}).
		Times(1)

	opts := &DefaultSetterOpts{Service: "cloud", OrgExecutor: mockExecutor}
	_, _ = opts.orgs(context.Background(), "myFilter")

	httpRequest, err := api.ConvertToHTTPRequest("https://cloud.mongodb.com", got)
	require.NoError(t, err)

	query := httpRequest.URL.Query()
	assert.Equal(t, "false", query.Get("includeGlobal"))
	assert.Equal(t, "myFilter", query.Get("name"))
	assert.Equal(t, strconv.Itoa(resultsLimit), query.Get("itemsPerPage"))
}
