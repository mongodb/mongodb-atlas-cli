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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_database_roles.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store DatabaseRoleLister,DatabaseRoleCreator,DatabaseRoleDeleter,DatabaseRoleUpdater,DatabaseRoleDescriber

type DatabaseRoleLister interface {
	DatabaseRoles(string) ([]atlasv2.UserCustomDBRole, error)
}

type DatabaseRoleCreator interface {
	CreateDatabaseRole(string, *atlasv2.UserCustomDBRole) (*atlasv2.UserCustomDBRole, error)
}

type DatabaseRoleDeleter interface {
	DeleteDatabaseRole(string, string) error
}

type DatabaseRoleUpdater interface {
	UpdateDatabaseRole(string, string, *atlasv2.UserCustomDBRole) (*atlasv2.UserCustomDBRole, error)
	DatabaseRole(string, string) (*atlasv2.UserCustomDBRole, error)
}

type DatabaseRoleDescriber interface {
	DatabaseRole(string, string) (*atlasv2.UserCustomDBRole, error)
}

// CreateDatabaseRole encapsulate the logic to manage different cloud providers.
func (s *Store) CreateDatabaseRole(groupID string, role *atlasv2.UserCustomDBRole) (*atlasv2.UserCustomDBRole, error) {
	result, _, err := s.clientv2.CustomDatabaseRolesApi.CreateCustomDatabaseRole(s.ctx, groupID, role).Execute()
	return result, err
}

// DeleteDatabaseRole encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteDatabaseRole(groupID, roleName string) error {
	_, err := s.clientv2.CustomDatabaseRolesApi.DeleteCustomDatabaseRole(s.ctx, groupID, roleName).Execute()
	return err
}

// DatabaseRoles encapsulate the logic to manage different cloud providers.
func (s *Store) DatabaseRoles(projectID string) ([]atlasv2.UserCustomDBRole, error) {
	result, _, err := s.clientv2.CustomDatabaseRolesApi.ListCustomDatabaseRoles(s.ctx, projectID).Execute()
	return result, err
}

// UpdateDatabaseRole encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateDatabaseRole(groupID, roleName string, role *atlasv2.UserCustomDBRole) (*atlasv2.UserCustomDBRole, error) {
	dbRole := atlasv2.UpdateCustomDBRole{
		Actions:        role.Actions,
		InheritedRoles: role.InheritedRoles,
	}
	result, _, err := s.clientv2.CustomDatabaseRolesApi.UpdateCustomDatabaseRole(s.ctx, groupID, roleName, &dbRole).Execute()
	return result, err
}

// DatabaseRole encapsulate the logic to manage different cloud providers.
func (s *Store) DatabaseRole(groupID, roleName string) (*atlasv2.UserCustomDBRole, error) {
	result, _, err := s.clientv2.CustomDatabaseRolesApi.GetCustomDatabaseRole(s.ctx, groupID, roleName).Execute()
	return result, err
}
