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

// This code was autogenerated at 2023-06-21T10:33:04+01:00. Note: Manual updates are allowed, but may be overwritten.

package store

import (
	"io"

	"go.mongodb.org/atlas-sdk/v20240530005/admin"
)

//go:generate mockgen -destination=../mocks/mock_data_federation.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store DataFederationLister,DataFederationDescriber,DataFederationStore,DataFederationCreator,DataFederationUpdater,DataFederationDeleter,DataFederationLogDownloader

type DataFederationStore interface {
	DataFederationLister
	DataFederationDescriber
}

type DataFederationLister interface {
	DataFederationList(string) ([]admin.DataLakeTenant, error)
}

type DataFederationCreator interface {
	CreateDataFederation(string, *admin.DataLakeTenant) (*admin.DataLakeTenant, error)
}

type DataFederationDeleter interface {
	DeleteDataFederation(string, string) error
}

type DataFederationDescriber interface {
	DataFederation(string, string) (*admin.DataLakeTenant, error)
}

type DataFederationUpdater interface {
	UpdateDataFederation(string, string, *admin.DataLakeTenant) (*admin.DataLakeTenant, error)
}

type DataFederationLogDownloader interface {
	DataFederationLogs(string, string, int64, int64) (io.ReadCloser, error)
}

// DataFederationList encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationList(projectID string) ([]admin.DataLakeTenant, error) {
	req := s.clientv2.DataFederationApi.ListFederatedDatabases(s.ctx, projectID)
	result, _, err := req.Execute()
	return result, err
}

// DataFederation encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederation(projectID, id string) (*admin.DataLakeTenant, error) {
	result, _, err := s.clientv2.DataFederationApi.GetFederatedDatabase(s.ctx, projectID, id).Execute()
	return result, err
}

// CreateDataFederation encapsulates the logic to manage different cloud providers.
func (s *Store) CreateDataFederation(projectID string, opts *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	result, _, err := s.clientv2.DataFederationApi.CreateFederatedDatabase(s.ctx, projectID, opts).SkipRoleValidation(false).Execute()
	return result, err
}

// UpdateDataFederation encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateDataFederation(projectID, id string, opts *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	result, _, err := s.clientv2.DataFederationApi.UpdateFederatedDatabase(s.ctx, projectID, id, opts).SkipRoleValidation(false).Execute()
	return result, err
}

// DeleteDataFederation encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteDataFederation(projectID, id string) error {
	_, _, err := s.clientv2.DataFederationApi.DeleteFederatedDatabase(s.ctx, projectID, id).Execute()
	return err
}

// DataFederationLogs encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationLogs(projectID, id string, startDate, endDate int64) (io.ReadCloser, error) {
	req := s.clientv2.DataFederationApi.DownloadFederatedDatabaseQueryLogs(s.ctx, projectID, id)
	if startDate != 0 {
		req = req.StartDate(startDate)
	}
	if endDate != 0 {
		req = req.EndDate(endDate)
	}
	result, _, err := req.Execute()
	return result, err
}
