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

package privatelink

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	t.Run("should error when no connectionID is provided", func(t *testing.T) {
		describeOpts := &DescribeOpts{}

		require.ErrorContains(t, describeOpts.Run(), "connectionID is missing")
	})

	t.Run("should call the store get privateLink method with the correct parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := NewMockDescriber(ctrl)

		connectionID := "123456789012"
		describeOpts := &DescribeOpts{
			store:        mockStore,
			connectionID: connectionID,
		}

		mockStore.
			EXPECT().
			DescribePrivateLinkEndpoint(gomock.Eq(describeOpts.ConfigProjectID()), gomock.Eq(connectionID)).
			Times(1)

		require.NoError(t, describeOpts.Run())
	})

	t.Run("should print the result", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := NewMockDescriber(ctrl)

		buf := new(bytes.Buffer)
		describeOpts := &DescribeOpts{
			store:        mockStore,
			connectionID: "123456789012",
			OutputOpts: cli.OutputOpts{
				Template:  describeTemplate,
				OutWriter: buf,
			},
		}

		expected := atlasv2.NewStreamsPrivateLinkConnection("AZURE")
		expected.SetId(describeOpts.connectionID)
		expected.SetInterfaceEndpointId("vpce-123456789012345678")
		expected.SetServiceEndpointId("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace")
		expected.SetDnsDomain("test-namespace.servicebus.windows.net")
		expected.SetProvider("Azure")
		expected.SetRegion("US_EAST_2")

		mockStore.
			EXPECT().
			// This test does not assert the parameters passed to the store method
			DescribePrivateLinkEndpoint(gomock.Any(), gomock.Any()).
			Return(expected, nil).
			Times(1)

		require.NoError(t, describeOpts.Run())
		test.VerifyOutputTemplate(t, describeTemplate, expected)
	})
}
