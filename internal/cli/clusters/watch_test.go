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
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

func TestWatch_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	expected := &atlasClustersPinned.AdvancedClusterDescription{StateName: pointer.Get("IDLE")}

	opts := &WatchOpts{
		name:          "test",
		store:         mockStore,
		isFlexCluster: false,
	}

	mockStore.
		EXPECT().
		AtlasCluster(opts.ProjectID, opts.name).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run(context.Background()))
}

func TestWatch_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	expected := &atlasv2.FlexClusterDescription20241113{StateName: pointer.Get("IDLE")}

	opts := &WatchOpts{
		name:          "test",
		store:         mockStore,
		isFlexCluster: true,
	}

	mockStore.
		EXPECT().
		FlexCluster(opts.ProjectID, opts.name).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run(context.Background()))
}
