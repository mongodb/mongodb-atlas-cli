// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package processes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func Test_autoCompleteOpts_tierSuggestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessLister(ctrl)

	au := &AutoCompleteOpts{
		store: mockStore,
	}
	expected := &atlasv2.PaginatedHostViewAtlas{
		Results: &[]atlasv2.ApiHostViewAtlas{
			{
				Id:       pointer.Get("test.com:27017"),
				Hostname: pointer.Get("test"),
				Port:     pointer.Get(27017),
			},
		},
	}
	mockStore.
		EXPECT().
		Processes(&atlasv2.ListAtlasProcessesApiParams{}).
		Return(expected, nil).
		Times(1)

	res, err := au.processSuggestion("")
	require.NoError(t, err)
	require.Equal(t, []string{"test.com:27017"}, res)
}
