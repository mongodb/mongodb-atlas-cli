// Copyright 2020 MongoDB Inc
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

package clusters

import (
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/mocks"
	"github.com/andreangiolillo/mongocli-test/internal/test/fixture"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudManagerClustersDescriber(ctrl)

	t.Run("describe cluster simplified", func(t *testing.T) {
		descOpts := &DescribeOpts{
			store: mockStore,
			name:  "myReplicaSet",
		}
		expected := &opsmngr.Cluster{}
		mockStore.
			EXPECT().
			OpsManagerCluster(descOpts.ProjectID, descOpts.name).
			Return(expected, nil).
			Times(1)

		require.NoError(t, descOpts.Run())
	})

	t.Run("describe cluster for JSON", func(t *testing.T) {
		descOpts := &DescribeOpts{
			store:      mockStore,
			OutputOpts: cli.OutputOpts{Output: "json"},
			name:       "myReplicaSet",
		}

		expected := fixture.AutomationConfig()
		mockStore.
			EXPECT().
			GetAutomationConfig(descOpts.ProjectID).
			Return(expected, nil).
			Times(1)

		require.NoError(t, descOpts.Run())
	})
}

func TestDescribeOpts_validateArg(t *testing.T) {
	type fields struct {
		OutputOpts cli.OutputOpts
		name       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "no output valid id",
			fields: fields{
				OutputOpts: cli.OutputOpts{},
				name:       "62bdbde098ecb26e563edf84",
			},
			wantErr: require.NoError,
		},
		{
			name: "no output invalid id",
			fields: fields{
				OutputOpts: cli.OutputOpts{},
				name:       "test",
			},
			wantErr: require.Error,
		},
		{
			name: "json output valid id",
			fields: fields{
				OutputOpts: cli.OutputOpts{Output: "json"},
				name:       "test",
			},
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &DescribeOpts{
				OutputOpts: tt.fields.OutputOpts,
				name:       tt.fields.name,
			}
			tt.wantErr(t, opts.validateArg())
		})
	}
}
