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

//go:build unit

package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
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
			name: "cloud manager",
			fields: fields{
				Service: "cloud-manager",
			},
			want: 1,
		},
		{
			name: "ops manager",
			fields: fields{
				Service: "ops-manager",
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
		Service: "cloud-manager",
		Store:   mockStore,
	}
	t.Run("empty", func(t *testing.T) {
		expectedProjects := &opsmngr.Projects{}
		mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(1)
		_, _, err := opts.projects()
		require.Error(t, err)
	})
	t.Run("with one project", func(t *testing.T) {
		expectedProjects := &opsmngr.Projects{
			Results: []*opsmngr.Project{
				{
					ID:   "1",
					Name: "Project 1",
				},
			},
			TotalCount: 1,
		}
		mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(1)
		gotIDs, gotNames, err := opts.projects()
		require.NoError(t, err)
		assert.Equal(t, []string{"1"}, gotIDs)
		assert.Equal(t, []string{"Project 1"}, gotNames)
	})
}

func TestDefaultOpts_Orgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	opts := &DefaultSetterOpts{
		Service: "cloud-manager",
		Store:   mockStore,
	}
	t.Run("empty", func(t *testing.T) {
		expectedOrgs := &opsmngr.Organizations{}
		mockStore.EXPECT().Organizations(gomock.Any()).Return(expectedOrgs, nil).Times(1)
		_, err := opts.orgs("")
		require.Error(t, err)
	})
	t.Run("with one org", func(t *testing.T) {
		expectedOrgs := &opsmngr.Organizations{
			Results: []*opsmngr.Organization{
				{
					ID:   "1",
					Name: "Org 1",
				},
			},
			TotalCount: 1,
		}
		mockStore.EXPECT().Organizations(gomock.Any()).Return(expectedOrgs, nil).Times(1)
		gotOrgs, err := opts.orgs("")
		require.NoError(t, err)
		assert.Equal(t, expectedOrgs.Results, gotOrgs)
	})

	t.Run("with no org", func(t *testing.T) {
		expectedOrgs := &opsmngr.Organizations{
			Results: []*opsmngr.Organization{},
		}
		mockStore.EXPECT().Organizations(gomock.Any()).Return(expectedOrgs, nil).Times(1)
		_, err := opts.orgs("")
		require.Error(t, err)
		require.EqualError(t, err, errNoResults.Error())
	})

	t.Run("with nil org", func(t *testing.T) {
		mockStore.EXPECT().Organizations(gomock.Any()).Return(nil, errNoResults).Times(1)
		_, err := opts.orgs("")
		require.Error(t, err)
		require.EqualError(t, err, errNoResults.Error())
	})
}
