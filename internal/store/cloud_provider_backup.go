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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store RestoreJobsLister,RestoreJobsDescriber,RestoreJobsCreator,SnapshotsLister,SnapshotsCreator,SnapshotsDescriber,SnapshotsDeleter,ExportJobsLister,ExportJobsDescriber,ExportJobsCreator,ExportBucketsLister,ExportBucketsCreator,ExportBucketsDeleter,ExportBucketsDescriber,ScheduleDescriber,ScheduleDescriberUpdater,ScheduleDeleter

type RestoreJobsLister interface {
	RestoreJobs(string, string, *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupRestoreJob, error)
}

type RestoreJobsDescriber interface {
	RestoreJob(string, string, string) (*atlasv2.DiskBackupSnapshotRestoreJob, error)
}

type RestoreJobsCreator interface {
	CreateRestoreJobs(string, string, *atlasv2.DiskBackupSnapshotRestoreJob) (*atlasv2.DiskBackupSnapshotRestoreJob, error)
}

type SnapshotsLister interface {
	Snapshots(string, string, *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupReplicaSet, error)
}

type SnapshotsDescriber interface {
	Snapshot(string, string, string) (*atlasv2.DiskBackupReplicaSet, error)
}

type SnapshotsCreator interface {
	CreateSnapshot(string, string, *atlasv2.DiskBackupOnDemandSnapshotRequest) (*atlasv2.DiskBackupSnapshot, error)
}

type SnapshotsDeleter interface {
	DeleteSnapshot(string, string, string) error
}

type ExportJobsLister interface {
	ExportJobs(string, string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasDiskBackupExportJob, error)
}

type ExportJobsCreator interface {
	CreateExportJob(string, string, *atlasv2.DiskBackupExportJobRequest) (*atlasv2.DiskBackupExportJob, error)
}

type ExportBucketsLister interface {
	ExportBuckets(string, *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshotExportBucket, error)
}

type ExportJobsDescriber interface {
	ExportJob(string, string, string) (*atlasv2.DiskBackupExportJob, error)
}

type ExportBucketsCreator interface {
	CreateExportBucket(string, *atlasv2.DiskBackupSnapshotAWSExportBucket) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error)
}

type ExportBucketsDeleter interface {
	DeleteExportBucket(string, string) error
}

type ExportBucketsDescriber interface {
	DescribeExportBucket(string, string) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error)
}

type ScheduleDescriber interface {
	DescribeSchedule(string, string) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

type ScheduleDescriberUpdater interface {
	DescribeSchedule(string, string) (*atlasv2.DiskBackupSnapshotSchedule, error)
	UpdateSchedule(string, string, *atlasv2.DiskBackupSnapshotSchedule) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

type ScheduleDeleter interface {
	DeleteSchedule(string, string) error
}

// RestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupRestoreJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.CloudBackupsApi.ListBackupRestoreJobs(s.ctx, projectID, clusterName)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// RestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJob(projectID, clusterName, jobID string) (*atlasv2.DiskBackupSnapshotRestoreJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetBackupRestoreJob(s.ctx, projectID, clusterName, jobID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRestoreJobs(projectID, clusterName string, request *atlasv2.DiskBackupSnapshotRestoreJob) (*atlasv2.DiskBackupSnapshotRestoreJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.CreateBackupRestoreJob(s.ctx, projectID, clusterName, request).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) CreateSnapshot(projectID, clusterName string, request *atlasv2.DiskBackupOnDemandSnapshotRequest) (*atlasv2.DiskBackupSnapshot, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.TakeSnapshot(s.ctx, projectID, clusterName, request).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Snapshots encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupReplicaSet, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.CloudBackupsApi.ListReplicaSetBackups(s.ctx, projectID, clusterName)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Snapshot encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshot(projectID, clusterName, snapshotID string) (*atlasv2.DiskBackupReplicaSet, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetReplicaSetBackup(s.ctx, projectID, clusterName, snapshotID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteSnapshot(projectID, clusterName, snapshotID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.CloudBackupsApi.DeleteReplicaSetBackup(s.ctx, projectID, clusterName, snapshotID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ExportJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ExportJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasDiskBackupExportJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.CloudBackupsApi.ListBackupExportJobs(s.ctx, projectID, clusterName)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) ExportJob(projectID, clusterName, bucketID string) (*atlasv2.DiskBackupExportJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetBackupExportJob(s.ctx, projectID, clusterName, bucketID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportJob(projectID, clusterName string, job *atlasv2.DiskBackupExportJobRequest) (*atlasv2.DiskBackupExportJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.CreateBackupExportJob(s.ctx, projectID, clusterName, job).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ExportBuckets encapsulates the logic to manage different cloud providers.
func (s *Store) ExportBuckets(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshotExportBucket, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.CloudBackupsApi.ListExportBuckets(s.ctx, projectID)
		if opts != nil {
			res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportBucket(projectID string, bucket *atlasv2.DiskBackupSnapshotAWSExportBucket) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.CreateExportBucket(s.ctx, projectID, bucket).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteExportBucket(projectID, bucketID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.CloudBackupsApi.DeleteExportBucket(s.ctx, projectID, bucketID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DescribeExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeExportBucket(projectID, bucketID string) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetExportBucket(s.ctx, projectID, bucketID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DescribeSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeSchedule(projectID, clusterName string) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetBackupSchedule(s.ctx, projectID, clusterName).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateSchedule(projectID, clusterName string, policy *atlasv2.DiskBackupSnapshotSchedule) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.UpdateBackupSchedule(s.ctx, projectID, clusterName, policy).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteSchedule(projectID, clusterName string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.CloudBackupsApi.DeleteAllBackupSchedules(s.ctx, projectID, clusterName).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
