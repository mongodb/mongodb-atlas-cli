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

package search

import (
	"fmt"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDefaultRegion(t *testing.T) {
	tests := []struct {
		input []atlasv2.AvailableCloudProviderRegion
		want  int
	}{
		{
			input: []atlasv2.AvailableCloudProviderRegion{},
			want:  -1,
		},
		{
			input: []atlasv2.AvailableCloudProviderRegion{
				{
					Name:    pointer.Get("test"),
					Default: pointer.Get(false),
				},
			},
			want: -1,
		},
		{
			input: []atlasv2.AvailableCloudProviderRegion{
				{
					Name:    pointer.Get("test"),
					Default: pointer.Get(true),
				},
			},
			want: 0,
		},
		{
			input: []atlasv2.AvailableCloudProviderRegion{
				{
					Name:    pointer.Get("test"),
					Default: pointer.Get(false),
				},
				{
					Name:    pointer.Get("test2"),
					Default: pointer.Get(true),
				},
				{
					Name:    pointer.Get("test1"),
					Default: pointer.Get(false),
				},
			},
			want: 1,
		},
		{
			input: []atlasv2.AvailableCloudProviderRegion{
				{
					Name:    pointer.Get("test"),
					Default: pointer.Get(false),
				},
				{
					Name:    pointer.Get("test2"),
					Default: pointer.Get(false),
				},
				{
					Name:    pointer.Get("test1"),
					Default: pointer.Get(false),
				},
			},
			want: -1,
		},
	}

	for i, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			out := DefaultRegion(tt.input)
			assert.Equal(t, tt.want, out)
		})
	}
}
