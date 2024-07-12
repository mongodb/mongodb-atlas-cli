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

package store

import (
	"errors"
	"io"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_streams.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store StreamsLister,StreamsDescriber,StreamsCreator,StreamsDeleter,StreamsUpdater,StreamsDownloader,ConnectionCreator,ConnectionDeleter,ConnectionUpdater,StreamsConnectionDescriber,StreamsConnectionLister

type StreamsLister interface {
	ProjectStreams(*atlasv2.ListStreamInstancesApiParams) (*atlasv2.PaginatedApiStreamsTenant, error)
}

type StreamsDescriber interface {
	AtlasStream(string, string) (*atlasv2.StreamsTenant, error)
}

type StreamsCreator interface {
	CreateStream(string, *atlasv2.StreamsTenant) (*atlasv2.StreamsTenant, error)
}

type StreamsDeleter interface {
	DeleteStream(string, string) error
}

type StreamsUpdater interface {
	UpdateStream(string, string, *atlasv2.StreamsDataProcessRegion) (*atlasv2.StreamsTenant, error)
}

type StreamsDownloader interface {
	DownloadAuditLog(*atlasv2.DownloadStreamTenantAuditLogsApiParams) (io.ReadCloser, error)
}

type StreamsConnectionLister interface {
	StreamsConnections(string, string) (*atlasv2.PaginatedApiStreamsConnection, error)
}

type ConnectionCreator interface {
	CreateConnection(string, string, *atlasv2.StreamsConnection) (*atlasv2.StreamsConnection, error)
}

type ConnectionDeleter interface {
	DeleteConnection(string, string, string) error
}

type StreamsConnectionDescriber interface {
	StreamConnection(string, string, string) (*atlasv2.StreamsConnection, error)
}

type ConnectionUpdater interface {
	UpdateConnection(string, string, string, *atlasv2.StreamsConnection) (*atlasv2.StreamsConnection, error)
}

func (s *Store) ProjectStreams(opts *atlasv2.ListStreamInstancesApiParams) (*atlasv2.PaginatedApiStreamsTenant, error) {
	result, _, err := s.clientv2.StreamsApi.ListStreamInstancesWithParams(s.ctx, opts).Execute()
	return result, err
}

func (s *Store) AtlasStream(projectID, name string) (*atlasv2.StreamsTenant, error) {
	result, _, err := s.clientv2.StreamsApi.GetStreamInstance(s.ctx, projectID, name).Execute()
	return result, err
}

func (s *Store) CreateStream(projectID string, processor *atlasv2.StreamsTenant) (*atlasv2.StreamsTenant, error) {
	result, _, err := s.clientv2.StreamsApi.CreateStreamInstance(s.ctx, projectID, processor).Execute()
	return result, err
}

func (s *Store) DeleteStream(projectID, name string) error {
	_, _, err := s.clientv2.StreamsApi.DeleteStreamInstance(s.ctx, projectID, name).Execute()
	return err
}

func (s *Store) UpdateStream(projectID, name string, streamsDataProcessRegion *atlasv2.StreamsDataProcessRegion) (*atlasv2.StreamsTenant, error) {
	result, _, err := s.clientv2.StreamsApi.UpdateStreamInstance(s.ctx, projectID, name, streamsDataProcessRegion).Execute()
	return result, err
}

func (s *Store) DownloadAuditLog(request *atlasv2.DownloadStreamTenantAuditLogsApiParams) (io.ReadCloser, error) {
	result, _, err := s.clientv2.StreamsApi.DownloadStreamTenantAuditLogsWithParams(s.ctx, request).Execute()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("returned file is empty")
	}
	return result, nil
}

// StreamsConnections encapsulates the logic to manage different cloud providers.
func (s *Store) StreamsConnections(projectID, tenantName string) (*atlasv2.PaginatedApiStreamsConnection, error) {
	connections, _, err := s.clientv2.StreamsApi.ListStreamConnections(s.ctx, projectID, tenantName).Execute()
	return connections, err
}

// StreamConnection encapsulates the logic to manage different cloud providers.
func (s *Store) StreamConnection(projectID, tenantName, connectionName string) (*atlasv2.StreamsConnection, error) {
	result, _, err := s.clientv2.StreamsApi.GetStreamConnection(s.ctx, projectID, tenantName, connectionName).Execute()
	return result, err
}

// CreateConnection encapsulates the logic to manage different cloud providers.
func (s *Store) CreateConnection(projectID, tenantName string, opts *atlasv2.StreamsConnection) (*atlasv2.StreamsConnection, error) {
	result, _, err := s.clientv2.StreamsApi.CreateStreamConnection(s.ctx, projectID, tenantName, opts).Execute()
	return result, err
}

// UpdateConnection encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateConnection(projectID, tenantName, connectionsName string, opts *atlasv2.StreamsConnection) (*atlasv2.StreamsConnection, error) {
	result, _, err := s.clientv2.StreamsApi.UpdateStreamConnection(s.ctx, projectID, tenantName, connectionsName, opts).Execute()
	return result, err
}

// DeleteConnection encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteConnection(projectID, tenantName, connectionName string) error {
	_, _, err := s.clientv2.StreamsApi.DeleteStreamConnection(s.ctx, projectID, tenantName, connectionName).Execute()
	return err
}
