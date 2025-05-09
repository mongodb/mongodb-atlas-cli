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

package interfaces

import (
	"testing"

	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
	"go.uber.org/mock/gomock"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockInterfaceEndpointCreator(ctrl)

	createOpts := &CreateOpts{
		store:                    mockStore,
		privateEndpointServiceID: "privateEndpointServiceID",
		privateEndpointGroupID:   "privateEndpointGroupID",
	}

	expected := &atlasv2.PrivateLinkEndpoint{}
	mockStore.
		EXPECT().
		CreateInterfaceEndpoint(createOpts.ProjectID, provider, createOpts.privateEndpointServiceID, createOpts.newInterfaceEndpointConnection()).
		Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	require.NoError(t, err)
}

func TestCreateOpts_ValidateEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		wantErr  bool
	}{
		{
			name:     "correct",
			endpoint: "foo@127.0.0.1",
			wantErr:  false,
		},
		{
			name:     "only endpoint name",
			endpoint: "foo",
			wantErr:  true,
		},
		{
			name:     "only IP address",
			endpoint: "127.0.0.1",
			wantErr:  true,
		},
		{
			name:     "no endpoint name",
			endpoint: "@127.0.0.1",
			wantErr:  true,
		},
		{
			name:     "no IP address",
			endpoint: "foo@",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		endpoints := make([]string, 1)
		endpoints[0] = tt.endpoint
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			opts := &CreateOpts{
				Endpoints: endpoints,
			}
			if err := opts.validateEndpoints(); (err != nil) != wantErr {
				t.Errorf("ValidateEndpoints() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}
