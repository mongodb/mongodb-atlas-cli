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
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

func getPrivateLinkConnections() []atlasv2.StreamsPrivateLinkConnection {
	var connections []atlasv2.StreamsPrivateLinkConnection

	for i := 0; i < 5; i++ {
		conn := atlasv2.NewStreamsPrivateLinkConnection()
		conn.SetId(fmt.Sprintf("testId%d", i))
		conn.SetProvider("Azure")
		conn.SetRegion("US_EAST_2")
		conn.SetServiceEndpointId("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace")
		conn.SetDnsDomain("test-namespace.servicebus.windows.net")

		connections = append(connections, *conn)
	}

	return connections
}

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPrivateLinkLister(ctrl)

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	connections := getPrivateLinkConnections()
	expected := atlasv2.NewPaginatedApiStreamsPrivateLink()
	expected.SetResults(connections)

	mockStore.
		EXPECT().
		ListPrivateLinkEndpoints(gomock.Eq(listOpts.ConfigProjectID())).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, listTemplate, expected)
}
