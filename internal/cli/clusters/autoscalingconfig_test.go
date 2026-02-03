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

package clusters

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
	"go.uber.org/mock/gomock"
)

func TestGetAutoscalingConfigOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterAutoscalingConfigGetter(ctrl)

	tests := []struct {
		name string
		mode string
	}{
		{name: "independentShardScaling", mode: independentShardScalingResponse},
		{name: "clusterWideScaling", mode: clusterWideScalingResponse},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			expected := &atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
				AutoScalingMode: pointer.Get(test.mode),
			}

			opts := &GetAutoscalingConfigOpts{
				store: mockStore,
				name:  "ProjectBar",
				OutputOpts: cli.OutputOpts{
					Template:  "{{.AutoScalingMode}}",
					OutWriter: buf,
				},
			}

			mockStore.
				EXPECT().
				GetClusterAutoScalingConfig(opts.ConfigProjectID(), opts.name).Return(expected, nil).Times(1)

			require.NoError(t, opts.Run())

			require.Equal(t, test.mode, buf.String())
		})
	}
}
