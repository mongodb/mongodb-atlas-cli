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
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store RestoreJobsLister,RestoreJobsDescriber,RestoreJobsCreator,SnapshotsLister,SnapshotsCreator,SnapshotsDescriber,SnapshotsDeleter,ExportJobsLister,ExportJobsDescriber,ExportJobsCreator,ExportBucketsLister,ExportBucketsCreator,ExportBucketsDeleter,ExportBucketsDescriber,ScheduleDescriber,ScheduleDescriberUpdater,ScheduleDeleter,SnapshotsDownloader

type RestoreJobsLister interface {
	RestoreJobs(string, string, *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupRestoreJob, error)
	RestoreFlexClusterJobs(args *atlasv2.ListFlexBackupRestoreJobsApiParams) (*atlasv2.PaginatedApiAtlasFlexBackupRestoreJob20241113, error)
}

type RestoreJobsDescriber interface {
	RestoreJob(string, string, string) (*atlasv2.DiskBackupSnapshotRestoreJob, error)
	RestoreFlexClusterJob(string, string, string) (*atlasv2.FlexBackupRestoreJob20241113, error)
}

type RestoreJobsCreator interface {
	CreateRestoreJobs(string, string, *atlasv2.DiskBackupSnapshotRestoreJob) (*atlasv2.DiskBackupSnapshotRestoreJob, error)
	CreateRestoreFlexClusterJobs(string, string, *atlasv2.FlexBackupRestoreJobCreate20241113) (*atlasv2.FlexBackupRestoreJob20241113, error)
}

type SnapshotsLister interface {
	Snapshots(string, string, *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupReplicaSet, error)
	FlexClusterSnapshots(*atlasv2.ListFlexBackupsApiParams) (*atlasv2.PaginatedApiAtlasFlexBackupSnapshot20241113, error)
}

type SnapshotsDescriber interface {
	Snapshot(string, string, string) (*atlasv2.DiskBackupReplicaSet, error)
	FlexClusterSnapshot(string, string, string) (*atlasv2.FlexBackupSnapshot20241113, error)
}

type SnapshotsCreator interface {
	CreateSnapshot(string, string, *atlasv2.DiskBackupOnDemandSnapshotRequest) (*atlasv2.DiskBackupSnapshot, error)
}

type SnapshotsDownloader interface {
	DownloadFlexClusterSnapshot(string, string, *atlasv2.FlexBackupSnapshotDownloadCreate20241113) (*atlasv2.FlexBackupRestoreJob20241113, error)
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
	ExportBuckets(string, *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshotExportBuckets, error)
}

type ExportJobsDescriber interface {
	ExportJob(string, string, string) (*atlasv2.DiskBackupExportJob, error)
}

type ExportBucketsCreator interface {
	CreateExportBucket(string, *atlasv2.DiskBackupSnapshotExportBucketRequest) (*atlasv2.DiskBackupSnapshotExportBucketResponse, error)
}

type ExportBucketsDeleter interface {
	DeleteExportBucket(string, string) error
}

type ExportBucketsDescriber interface {
	DescribeExportBucket(string, string) (*atlasv2.DiskBackupSnapshotExportBucketResponse, error)
}

type ScheduleDescriber interface {
	DescribeSchedule(string, string) (*atlasClustersPinned.DiskBackupSnapshotSchedule, error)
}

type ScheduleDescriberUpdater interface {
	DescribeSchedule(string, string) (*atlasClustersPinned.DiskBackupSnapshotSchedule, error)
	UpdateSchedule(string, string, *atlasClustersPinned.DiskBackupSnapshotSchedule) (*atlasClustersPinned.DiskBackupSnapshotSchedule, error)
}

type ScheduleDeleter interface {
	DeleteSchedule(string, string) error
}

// RestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupRestoreJob, error) {
	res := s.clientv2.CloudBackupsApi.ListBackupRestoreJobs(s.ctx, projectID, clusterName)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// RestoreFlexClusterJob encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreFlexClusterJob(projectID, clusterName, restoreJobID string) (*atlasv2.FlexBackupRestoreJob20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexRestoreJobsApi.GetFlexBackupRestoreJob(s.ctx, projectID, clusterName, restoreJobID).Execute()
	return result, err
}

// RestoreFlexClusterJobs encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreFlexClusterJobs(args *atlasv2.ListFlexBackupRestoreJobsApiParams) (*atlasv2.PaginatedApiAtlasFlexBackupRestoreJob20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexRestoreJobsApi.ListFlexBackupRestoreJobsWithParams(s.ctx, args).Execute()
	return result, err
}

// RestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJob(projectID, clusterName, jobID string) (*atlasv2.DiskBackupSnapshotRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetBackupRestoreJob(s.ctx, projectID, clusterName, jobID).Execute()
	return result, err
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRestoreJobs(projectID, clusterName string, request *atlasv2.DiskBackupSnapshotRestoreJob) (*atlasv2.DiskBackupSnapshotRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateBackupRestoreJob(s.ctx, projectID, clusterName, request).Execute()
	return result, err
}

// CreateRestoreFlexClusterJobs encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRestoreFlexClusterJobs(projectID, clusterName string, request *atlasv2.FlexBackupRestoreJobCreate20241113) (*atlasv2.FlexBackupRestoreJob20241113, error) {
	result, _, err := s.clientv2.FlexRestoreJobsApi.CreateFlexBackupRestoreJob(s.ctx, projectID, clusterName, request).Execute()
	return result, err
}

// CreateSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) CreateSnapshot(projectID, clusterName string, request *atlasv2.DiskBackupOnDemandSnapshotRequest) (*atlasv2.DiskBackupSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.TakeSnapshot(s.ctx, projectID, clusterName, request).Execute()
	return result, err
}

// Snapshots encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedCloudBackupReplicaSet, error) {
	res := s.clientv2.CloudBackupsApi.ListReplicaSetBackups(s.ctx, projectID, clusterName)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// FlexClusterSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) FlexClusterSnapshots(opts *atlasv2.ListFlexBackupsApiParams) (*atlasv2.PaginatedApiAtlasFlexBackupSnapshot20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexSnapshotsApi.ListFlexBackupsWithParams(s.ctx, opts).Execute()
	return result, err
}

// DownloadFlexClusterSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) DownloadFlexClusterSnapshots(groupID, name string, flexBackupSnapshotDownloadCreate20241113 *atlasv2.FlexBackupSnapshotDownloadCreate20241113) (*atlasv2.FlexBackupRestoreJob20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexSnapshotsApi.DownloadFlexBackup(s.ctx, name, groupID, flexBackupSnapshotDownloadCreate20241113).Execute()
	return result, err
}

// FlexClusterSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) FlexClusterSnapshot(groupID, name, snapshotID string) (*atlasv2.FlexBackupSnapshot20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexSnapshotsApi.GetFlexBackup(s.ctx, groupID, name, snapshotID).Execute()
	return result, err
}

// Snapshot encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshot(projectID, clusterName, snapshotID string) (*atlasv2.DiskBackupReplicaSet, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetReplicaSetBackup(s.ctx, projectID, clusterName, snapshotID).Execute()
	return result, err
}

// DeleteSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteSnapshot(projectID, clusterName, snapshotID string) error {
	_, _, err := s.clientv2.CloudBackupsApi.DeleteReplicaSetBackup(s.ctx, projectID, clusterName, snapshotID).Execute()
	return err
}

// ExportJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ExportJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasDiskBackupExportJob, error) {
	res := s.clientv2.CloudBackupsApi.ListBackupExportJobs(s.ctx, projectID, clusterName)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// ExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) ExportJob(projectID, clusterName, bucketID string) (*atlasv2.DiskBackupExportJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetBackupExportJob(s.ctx, projectID, clusterName, bucketID).Execute()
	return result, err
}

// CreateExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportJob(projectID, clusterName string, job *atlasv2.DiskBackupExportJobRequest) (*atlasv2.DiskBackupExportJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateBackupExportJob(s.ctx, projectID, clusterName, job).Execute()
	return result, err
}

// ExportBuckets encapsulates the logic to manage different cloud providers.
func (s *Store) ExportBuckets(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshotExportBuckets, error) {
	res := s.clientv2.CloudBackupsApi.ListExportBuckets(s.ctx, projectID)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// CreateExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportBucket(projectID string, bucket *atlasv2.DiskBackupSnapshotExportBucketRequest) (*atlasv2.DiskBackupSnapshotExportBucketResponse, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateExportBucket(s.ctx, projectID, bucket).Execute()
	return result, err
}

// DeleteExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteExportBucket(projectID, bucketID string) error {
	_, _, err := s.clientv2.CloudBackupsApi.DeleteExportBucket(s.ctx, projectID, bucketID).Execute()
	return err
}

// DescribeExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeExportBucket(projectID, bucketID string) (*atlasv2.DiskBackupSnapshotExportBucketResponse, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetExportBucket(s.ctx, projectID, bucketID).Execute()
	return result, err
}

// DescribeSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeSchedule(projectID, clusterName string) (*atlasClustersPinned.DiskBackupSnapshotSchedule, error) {
	result, _, err := s.clientClusters.CloudBackupsApi.GetBackupSchedule(s.ctx, projectID, clusterName).Execute()
	return result, err
}

// UpdateSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateSchedule(projectID, clusterName string, policy *atlasClustersPinned.DiskBackupSnapshotSchedule) (*atlasClustersPinned.DiskBackupSnapshotSchedule, error) {
	result, _, err := s.clientClusters.CloudBackupsApi.UpdateBackupSchedule(s.ctx, projectID, clusterName, policy).Execute()
	return result, err
}

// DeleteSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteSchedule(projectID, clusterName string) error {
	_, _, err := s.clientv2.CloudBackupsApi.DeleteAllBackupSchedules(s.ctx, projectID, clusterName).Execute()
	return err
}
