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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

// UpdateMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateMaintenanceWindow(projectID string, maintenanceWindow *atlasv2.GroupMaintenanceWindow) error {
	_, err := s.clientv2.MaintenanceWindowsApi.UpdateMaintenanceWindow(s.ctx, projectID, maintenanceWindow).Execute()
	return err
}

// ClearMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) ClearMaintenanceWindow(projectID string) error {
	_, err := s.clientv2.MaintenanceWindowsApi.ResetMaintenanceWindow(s.ctx, projectID).Execute()
	return err
}

// DeferMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) DeferMaintenanceWindow(projectID string) error {
	_, err := s.clientv2.MaintenanceWindowsApi.DeferMaintenanceWindow(s.ctx, projectID).Execute()
	return err
}

// MaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) MaintenanceWindow(projectID string) (*atlasv2.GroupMaintenanceWindow, error) {
	resp, _, err := s.clientv2.MaintenanceWindowsApi.GetMaintenanceWindow(s.ctx, projectID).Execute()
	return resp, err
}
