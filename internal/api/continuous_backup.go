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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/api/mock_continuous_backup.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/api CheckpointsLister,ContinuousJobLister,ContinuousJobCreator,ContinuousSnapshotsLister

type CheckpointsLister interface {
	Checkpoints(string, string, *atlas.ListOptions) (*atlas.Checkpoints, error)
}

type ContinuousJobLister interface {
	ContinuousRestoreJobs(string, string, *atlas.ListOptions) (*atlas.ContinuousJobs, error)
}

type ContinuousJobCreator interface {
	CreateContinuousRestoreJob(string, string, *atlas.ContinuousJobRequest) (*atlas.ContinuousJobs, error)
}

type ContinuousSnapshotsLister interface {
	ContinuousSnapshots(string, string, *atlas.ListOptions) (*atlas.ContinuousSnapshots, error)
}

// Checkpoints encapsulate the logic to manage different cloud providers.
func (s *Store) Checkpoints(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.Checkpoints, error) {
	result, _, err := s.client.(*atlas.Client).Checkpoints.List(s.ctx, projectID, clusterID, opts)
	return result, err
}

// ContinuousRestoreJobs encapsulate the logic to manage different cloud providers.
func (s *Store) ContinuousRestoreJobs(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.ContinuousJobs, error) {
	result, _, err := s.client.(*atlas.Client).ContinuousRestoreJobs.List(s.ctx, projectID, clusterID, opts)
	return result, err
}

// CreateContinuousRestoreJob encapsulate the logic to manage different cloud providers.
func (s *Store) CreateContinuousRestoreJob(projectID, clusterID string, request *atlas.ContinuousJobRequest) (*atlas.ContinuousJobs, error) {
	result, _, err := s.client.(*atlas.Client).ContinuousRestoreJobs.Create(s.ctx, projectID, clusterID, request)
	return result, err
}

// ContinuousSnapshots encapsulate the logic to manage different cloud providers.
func (s *Store) ContinuousSnapshots(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.ContinuousSnapshots, error) {
	result, _, err := s.client.(*atlas.Client).ContinuousSnapshots.List(s.ctx, projectID, clusterID, opts)
	return result, err
}
