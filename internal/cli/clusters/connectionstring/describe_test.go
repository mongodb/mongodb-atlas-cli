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

package connectionstring

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDescribe_Run_StandardConnectionString(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	expected := &admin.AdvancedClusterDescription{
		ConnectionStrings: &admin.ClusterConnectionStrings{
			StandardSrv: pointer.Get("test"),
		},
	}

	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	err := describeOpts.Run()
	require.NoError(t, err)
	test.VerifyOutputTemplate(t, describeTemplateStandard, expected.ConnectionStrings)
}

func TestDescribe_Run_PrivateConnectionString(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	expected := &admin.AdvancedClusterDescription{
		ConnectionStrings: &admin.ClusterConnectionStrings{
			StandardSrv: pointer.Get("test"),
			PrivateSrv:  pointer.Get("test"),
		},
	}

	describeOpts := &DescribeOpts{
		name:   "test",
		store:  mockStore,
		csType: "private",
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	err := describeOpts.Run()
	require.NoError(t, err)
	test.VerifyOutputTemplate(t, describeTemplatePrivate, expected.ConnectionStrings)
}

func TestDescribe_Run_PrivateEndpointsConnectionString(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	expected := &admin.AdvancedClusterDescription{
		ConnectionStrings: &admin.ClusterConnectionStrings{
			StandardSrv: pointer.Get("test"),
			PrivateSrv:  pointer.Get("test"),
			PrivateEndpoint: &[]admin.ClusterDescriptionConnectionStringsPrivateEndpoint{
				{
					SrvShardOptimizedConnectionString: pointer.Get("test"),
				},
			},
		},
	}

	describeOpts := &DescribeOpts{
		name:   "test",
		store:  mockStore,
		csType: "privateEndpoints",
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	err := describeOpts.Run()
	require.NoError(t, err)
	test.VerifyOutputTemplate(t, describeTemplateShardOptimized, expected.ConnectionStrings)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
