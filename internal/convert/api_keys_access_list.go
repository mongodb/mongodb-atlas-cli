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

package convert

import atlas "go.mongodb.org/atlas/mongodbatlas"

func FromWhitelistAPIKeysToAccessListAPIKeys(in *atlas.WhitelistAPIKeys) *atlas.AccessListAPIKeys {
	out := &atlas.AccessListAPIKeys{
		TotalCount: in.TotalCount,
		Links:      in.Links,
	}

	var results []*atlas.AccessListAPIKey
	for _, element := range in.Results {
		results = append(results, fromWhitelistAPIKeyToAccessListAPIKey(element))
	}

	out.Results = results
	return out
}

func fromWhitelistAPIKeyToAccessListAPIKey(in *atlas.WhitelistAPIKey) *atlas.AccessListAPIKey {
	return &atlas.AccessListAPIKey{
		CidrBlock:       in.CidrBlock,
		Count:           in.Count,
		Created:         in.Created,
		IPAddress:       in.IPAddress,
		LastUsed:        in.LastUsed,
		LastUsedAddress: in.LastUsedAddress,
		Links:           in.Links,
	}
}

func FromAccessListAPIKeysReqToWhitelistAPIKeysReq(in []*atlas.AccessListAPIKeysReq) []*atlas.WhitelistAPIKeysReq {
	var out []*atlas.WhitelistAPIKeysReq

	for _, element := range in {
		accessListElement := &atlas.WhitelistAPIKeysReq{
			IPAddress: element.IPAddress,
			CidrBlock: element.CidrBlock,
		}
		out = append(out, accessListElement)
	}
	return out
}
