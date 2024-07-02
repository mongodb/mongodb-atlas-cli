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

package search

import (
	"strings"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func StringInSliceFold(a []string, x string) bool {
	for _, b := range a {
		if strings.EqualFold(b, x) {
			return true
		}
	}
	return false
}

// DefaultRegion returns the index of the default region.
func DefaultRegion(regions []atlasv2.AvailableCloudProviderRegion) int {
	for i, v := range regions {
		if v.GetDefault() {
			return i
		}
	}
	return -1
}
