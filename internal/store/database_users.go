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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_database_users.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store DatabaseUserLister,DatabaseUserCreator,DatabaseUserDeleter,DatabaseUserUpdater,DatabaseUserDescriber,DBUserCertificateLister,DBUserCertificateCreator

type DatabaseUserLister interface {
	DatabaseUsers(groupID string, opts *ListOptions) (*atlasv2.PaginatedApiAtlasDatabaseUser, error)
}

type DatabaseUserCreator interface {
	CreateDatabaseUser(*atlasv2.CloudDatabaseUser) (*atlasv2.CloudDatabaseUser, error)
}

type DatabaseUserDeleter interface {
	DeleteDatabaseUser(string, string, string) error
}

type DatabaseUserUpdater interface {
	UpdateDatabaseUser(*atlasv2.UpdateDatabaseUserApiParams) (*atlasv2.CloudDatabaseUser, error)
}

type DatabaseUserDescriber interface {
	DatabaseUser(string, string, string) (*atlasv2.CloudDatabaseUser, error)
}

type DBUserCertificateLister interface {
	DBUserCertificates(string, string, *atlas.ListOptions) (*atlasv2.PaginatedUserCert, error)
}

type DBUserCertificateCreator interface {
	CreateDBUserCertificate(string, string, int) (string, error)
}

// CreateDatabaseUser encapsulate the logic to manage different cloud providers.
func (s *Store) CreateDatabaseUser(user *atlasv2.CloudDatabaseUser) (*atlasv2.CloudDatabaseUser, error) {
	result, _, err := s.clientv2.DatabaseUsersApi.CreateDatabaseUser(s.ctx, user.GroupId, user).Execute()
	return result, err
}

func (s *Store) DeleteDatabaseUser(authDB, groupID, username string) error {
	_, _, err := s.clientv2.DatabaseUsersApi.DeleteDatabaseUser(s.ctx, groupID, authDB, username).Execute()
	return err
}

func (s *Store) DatabaseUsers(projectID string, opts *ListOptions) (*atlasv2.PaginatedApiAtlasDatabaseUser, error) {
	res := s.clientv2.DatabaseUsersApi.ListDatabaseUsers(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

func (s *Store) UpdateDatabaseUser(params *atlasv2.UpdateDatabaseUserApiParams) (*atlasv2.CloudDatabaseUser, error) {
	result, _, err := s.clientv2.DatabaseUsersApi.UpdateDatabaseUserWithParams(s.ctx, params).Execute()
	return result, err
}

func (s *Store) DatabaseUser(authDB, groupID, username string) (*atlasv2.CloudDatabaseUser, error) {
	result, _, err := s.clientv2.DatabaseUsersApi.GetDatabaseUser(s.ctx, groupID, authDB, username).Execute()
	return result, err
}

// DBUserCertificates retrieves the current Atlas managed certificates for a database user.
func (s *Store) DBUserCertificates(projectID, username string, opts *atlas.ListOptions) (*atlasv2.PaginatedUserCert, error) {
	res := s.clientv2.X509AuthenticationApi.ListDatabaseUserCertificates(s.ctx, projectID, username)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// CreateDBUserCertificate creates a new Atlas managed certificates for a database user.
func (s *Store) CreateDBUserCertificate(projectID, username string, monthsUntilExpiration int) (string, error) {
	userCert := &atlasv2.UserCert{
		MonthsUntilExpiration: pointer.Get(monthsUntilExpiration),
	}
	cert, _, err := s.clientv2.X509AuthenticationApi.CreateDatabaseUserCertificate(s.ctx, projectID, username, userCert).Execute()
	return cert, err
}
