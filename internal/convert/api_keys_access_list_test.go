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

// +build unit

package convert

import (
	"testing"

	"github.com/magiconair/properties/assert"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestFromWhitelistAPIKeysToAccessListAPIKeys(t *testing.T) {
	type test struct {
		input *atlas.WhitelistAPIKeys
		want  *atlas.AccessListAPIKeys
	}

	tests := []test{
		{
			input: nil,
			want:  nil,
		},
		{
			input: &atlas.WhitelistAPIKeys{
				Results: []*atlas.WhitelistAPIKey{
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
				},
				TotalCount: 1,
			},
			want: &atlas.AccessListAPIKeys{
				Results: []*atlas.AccessListAPIKey{
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
				},
				TotalCount: 1,
			},
		},
		{
			input: &atlas.WhitelistAPIKeys{
				Results: []*atlas.WhitelistAPIKey{
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
				},
				TotalCount: 2,
			},
			want: &atlas.AccessListAPIKeys{
				Results: []*atlas.AccessListAPIKey{
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
					{
						CidrBlock:       "1",
						Count:           1,
						Created:         "1",
						IPAddress:       "1",
						LastUsed:        "1",
						LastUsedAddress: "1",
						Links:           nil,
					},
				},
				TotalCount: 2,
			},
		},
	}

	for _, tc := range tests {
		got := FromWhitelistAPIKeysToAccessListAPIKeys(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func TestFromAccessListAPIKeysReqToWhitelistAPIKeysReq(t *testing.T) {
	type test struct {
		input []*atlas.AccessListAPIKeysReq
		want  []*atlas.WhitelistAPIKeysReq
	}

	tests := []test{
		{
			input: nil,
			want:  nil,
		},
		{
			input: []*atlas.AccessListAPIKeysReq{
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
			},
			want: []*atlas.WhitelistAPIKeysReq{
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
			},
		},
		{
			input: []*atlas.AccessListAPIKeysReq{
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
			},
			want: []*atlas.WhitelistAPIKeysReq{
				{
					IPAddress: "1",
					CidrBlock: "1",
				},
			},
		},
	}

	for _, tc := range tests {
		got := FromAccessListAPIKeysReqToWhitelistAPIKeysReq(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
