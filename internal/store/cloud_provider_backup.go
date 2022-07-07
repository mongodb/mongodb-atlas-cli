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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store RestoreJobsLister,RestoreJobsDescriber,RestoreJobsCreator,SnapshotsLister,SnapshotsCreator,SnapshotsDescriber,SnapshotsDeleter,ExportJobsLister,ExportBucketsLister,ExportBucketsCreator,ExportBucketsDescriber

type RestoreJobsLister interface {
	RestoreJobs(string, string, *atlas.ListOptions) (*atlas.CloudProviderSnapshotRestoreJobs, error)
}

type RestoreJobsDescriber interface {
	RestoreJob(string, string, string) (*atlas.CloudProviderSnapshotRestoreJob, error)
}

type RestoreJobsCreator interface {
	CreateRestoreJobs(string, string, *atlas.CloudProviderSnapshotRestoreJob) (*atlas.CloudProviderSnapshotRestoreJob, error)
}

type SnapshotsLister interface {
	Snapshots(string, string, *atlas.ListOptions) (*atlas.CloudProviderSnapshots, error)
}

type SnapshotsDescriber interface {
	Snapshot(string, string, string) (*atlas.CloudProviderSnapshot, error)
}

type SnapshotsCreator interface {
	CreateSnapshot(string, string, *atlas.CloudProviderSnapshot) (*atlas.CloudProviderSnapshot, error)
}

type SnapshotsDeleter interface {
	DeleteSnapshot(string, string, string) error
}

type ExportJobsLister interface {
	ExportJobs(string, string, *atlas.ListOptions) (*atlas.CloudProviderSnapshotExportJobs, error)
}

type ExportBucketsLister interface {
	ExportBuckets(string, *atlas.ListOptions) (*atlas.CloudProviderSnapshotExportBuckets, error)
}

type ExportBucketsCreator interface {
	CreateExportBucket(string, *atlas.CloudProviderSnapshotExportBucket) (*atlas.CloudProviderSnapshotExportBucket, error)
}

type ExportBucketsDescriber interface {
	DescribeExportBucket(string, string) (*atlas.CloudProviderSnapshotExportBucket, error)
}

// RestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshotRestoreJobs, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.List(s.ctx, o, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// RestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) RestoreJob(projectID, clusterName, jobID string) (*atlas.CloudProviderSnapshotRestoreJob, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		JobID:       jobID,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.Get(s.ctx, o)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRestoreJobs(projectID, clusterName string, request *atlas.CloudProviderSnapshotRestoreJob) (*atlas.CloudProviderSnapshotRestoreJob, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.Create(s.ctx, o, request)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) CreateSnapshot(projectID, clusterName string, request *atlas.CloudProviderSnapshot) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.Create(s.ctx, o, request)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Snapshots encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshots, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetAllCloudProviderSnapshots(s.ctx, o, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Snapshot encapsulates the logic to manage different cloud providers.
func (s *Store) Snapshot(projectID, clusterName, snapshotID string) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetOneCloudProviderSnapshot(s.ctx, o)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteSnapshot(projectID, clusterName, snapshotID string) error {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		SnapshotID:  snapshotID,
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.client.(*atlas.Client).CloudProviderSnapshots.Delete(s.ctx, o)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ExportJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ExportJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshotExportJobs, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotExportJobs.List(s.ctx, projectID, clusterName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ExportBuckets encapsulates the logic to manage different cloud providers.
func (s *Store) ExportBuckets(projectID string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshotExportBuckets, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotExportBuckets.List(s.ctx, projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportBucket(projectID string, bucket *atlas.CloudProviderSnapshotExportBucket) (*atlas.CloudProviderSnapshotExportBucket, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotExportBuckets.Create(s.ctx, projectID, bucket)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DescribeExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeExportBucket(projectID, bucketID string) (*atlas.CloudProviderSnapshotExportBucket, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotExportBuckets.Get(s.ctx, projectID, bucketID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
