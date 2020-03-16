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
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
)

func Agent() *om.Agent {
	return &om.Agent{
		TypeName:  "AUTOMATION",
		Hostname:  "example",
		ConfCount: 59,
		LastConf:  "2015-06-18T14:21:42Z",
		StateName: "ACTIVE",
		PingCount: 6,
		IsManaged: true,
		LastPing:  "2015-06-18T14:21:42Z",
	}
}

func Agents() *om.Agents {
	return &om.Agents{
		Links:      []*atlas.Link{},
		Results:    []*om.Agent{Agent()},
		TotalCount: 1,
	}
}
