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

package azure

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockPrivateEndpointDescriber(ctrl)

	opts := &DescribeOpts{
		store: mockStore,
	}

	expected := &atlasv2.EndpointService{}

	mockStore.
		EXPECT().
		PrivateEndpoint(opts.ProjectID, provider, opts.id).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	require.NoError(t, err)

	test.VerifyOutputTemplate(t, describeTemplate, expected)
}
