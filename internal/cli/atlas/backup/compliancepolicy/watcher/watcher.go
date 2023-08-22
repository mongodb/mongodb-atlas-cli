// Copyright 2023 MongoDB Inc
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

package watcher

import (
	"errors"

	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	active = "ACTIVE"
)

var errInvalidStateField = errors.New("invalid State field for watching")

// CompliancePolicyWatcherFactory allows to init a compliance policy watcher in a functional way.
//
// Parameters
//
//   - projectID: The ID of the project for which the compliance policy details need to be fetched.
//   - store: An implementation of the CompliancePolicyDescriber interface, which is used to describe/fetch the compliance policy.
//   - policy: A pointer to a DataProtectionSettings object which will be updated with the fetched details.
func CompliancePolicyWatcherFactory(projectID string, store store.CompliancePolicyDescriber, policy *atlasv2.DataProtectionSettings) func() (bool, error) {
	return func() (bool, error) {
		res, err := store.DescribeCompliancePolicy(projectID)
		if err != nil {
			return false, err
		}
		*policy = *res
		if res.GetState() == "" {
			return false, errInvalidStateField
		}
		return res.GetState() == active, nil
	}
}
