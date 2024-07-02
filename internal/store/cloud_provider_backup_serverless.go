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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup_serverless.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ServerlessSnapshotsLister,ServerlessSnapshotsDescriber,ServerlessRestoreJobsLister,ServerlessRestoreJobsDescriber,ServerlessRestoreJobsCreator

type ServerlessSnapshotsLister interface {
	ServerlessSnapshots(string, string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error)
}

type ServerlessSnapshotsDescriber interface {
	ServerlessSnapshot(string, string, string) (*atlasv2.ServerlessBackupSnapshot, error)
}

type ServerlessRestoreJobsLister interface {
	ServerlessRestoreJobs(string, string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error)
}

type ServerlessRestoreJobsDescriber interface {
	ServerlessRestoreJob(string, string, string) (*atlasv2.ServerlessBackupRestoreJob, error)
}

type ServerlessRestoreJobsCreator interface {
	ServerlessCreateRestoreJobs(string, string, *atlasv2.ServerlessBackupRestoreJob) (*atlasv2.ServerlessBackupRestoreJob, error)
}

// ServerlessSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	res := s.clientv2.CloudBackupsApi.ListServerlessBackups(s.ctx, projectID, clusterName)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// ServerlessSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshot(projectID, instanceName, snapshotID string) (*atlasv2.ServerlessBackupSnapshot, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.CloudBackupsApi.GetServerlessBackup(s.ctx, projectID, instanceName, snapshotID).Execute()
	return result, err
}

// ServerlessRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJobs(projectID, instanceName string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	res := s.clientv2.CloudBackupsApi.ListServerlessBackupRestoreJobs(s.ctx, projectID, instanceName)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// ServerlessRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJob(projectID, instanceName string, jobID string) (*atlasv2.ServerlessBackupRestoreJob, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.CloudBackupsApi.GetServerlessBackupRestoreJob(s.ctx, projectID, instanceName, jobID).Execute()
	return result, err
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessCreateRestoreJobs(projectID, clusterName string, request *atlasv2.ServerlessBackupRestoreJob) (*atlasv2.ServerlessBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateServerlessBackupRestoreJob(s.ctx, projectID, clusterName, request).Execute()
	return result, err
}
