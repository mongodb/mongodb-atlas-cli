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
	"context"
	"fmt"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_continuous_jobs.go -package=mocks github.com/mongodb/mongocli/internal/store ContinuousJobLister,ContinuousJobCreator

type ContinuousJobLister interface {
	ContinuousRestoreJobs(string, string, *atlas.ListOptions) (*atlas.ContinuousJobs, error)
}

type ContinuousJobCreator interface {
	CreateContinuousRestoreJob(string, string, *atlas.ContinuousJobRequest) (*atlas.ContinuousJobs, error)
}

func (s *Store) ContinuousRestoreJobs(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.ContinuousJobs, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ContinuousRestoreJobs.List(context.Background(), projectID, clusterID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).ContinuousRestoreJobs.List(context.Background(), projectID, clusterID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateContinuousRestoreJob encapsulate the logic to manage different cloud providers
func (s *Store) CreateContinuousRestoreJob(projectID, clusterID string, request *atlas.ContinuousJobRequest) (*atlas.ContinuousJobs, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ContinuousRestoreJobs.Create(context.Background(), projectID, clusterID, request)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).ContinuousRestoreJobs.Create(context.Background(), projectID, clusterID, request)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
