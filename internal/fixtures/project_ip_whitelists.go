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

package fixtures

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func ProjectIPWhitelist() []atlas.ProjectIPWhitelist {
	return []atlas.ProjectIPWhitelist{*IPWhiteList()}
}

func IPWhiteList() *atlas.ProjectIPWhitelist {
	return &atlas.ProjectIPWhitelist{
		Comment:   "test",
		GroupID:   "5def8d5dce4bd936ac99ae9c",
		CIDRBlock: "37.228.254.100/32",
		IPAddress: "37.228.254.100",
	}
}
