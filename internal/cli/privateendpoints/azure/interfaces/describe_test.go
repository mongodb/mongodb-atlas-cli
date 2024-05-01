// Copyright 2021 MongoDB Inc
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

package interfaces

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115013/admin"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockInterfaceEndpointDescriber(ctrl)

	opts := &DescribeOpts{
		store:                    mockStore,
		privateEndpointServiceID: "privateEndpointID",
		privateEndpointID:        "id",
	}

	expected := &atlasv2.PrivateLinkEndpoint{}

	mockStore.
		EXPECT().
		InterfaceEndpoint(opts.ProjectID, provider, opts.privateEndpointServiceID, opts.privateEndpointID).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	require.NoError(t, err)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.EndpointServiceID},
	)
}
