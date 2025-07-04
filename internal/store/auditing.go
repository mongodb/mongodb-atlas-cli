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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

func (s *Store) Auditing(projectID string) (*atlasv2.AuditLog, error) {
	result, _, err := s.clientv2.AuditingApi.GetAuditingConfiguration(s.ctx, projectID).Execute()
	return result, err
}

func (s *Store) UpdateAuditingConfig(projectID string, r *atlasv2.AuditLog) (*atlasv2.AuditLog, error) {
	result, _, err := s.clientv2.AuditingApi.UpdateAuditingConfiguration(s.ctx, projectID, r).Execute()
	return result, err
}
