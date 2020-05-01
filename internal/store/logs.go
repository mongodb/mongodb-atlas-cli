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
	"io"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

type LogsDownloader interface {
	DownloadLog(string, string, string, io.Writer, *atlas.DateRangetOptions) error
}

type LogJobsDownloader interface {
	DownloadLogJob(string, string, io.Writer) error
}

type LogCollector interface {
	Collect(string, *opsmngr.LogCollectionJob) (*opsmngr.LogCollectionJob, error)
}

type LogJobLister interface {
	LogCollectionJobs(string, *opsmngr.LogListOptions) (*opsmngr.LogCollectionJobs, error)
}

type LogJobDeleter interface {
	DeleteCollectionJob(string, string) error
}

// LogCollectionJobs encapsulate the logic to manage different cloud providers
func (s *Store) LogCollectionJobs(groupID string, opts *opsmngr.LogListOptions) (*opsmngr.LogCollectionJobs, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).LogCollections.List(context.Background(), groupID, opts)
		return log, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteCollectionJob encapsulate the logic to manage different cloud providers
func (s *Store) DeleteCollectionJob(groupID, logID string) error {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		_, err := s.client.(*opsmngr.Client).LogCollections.Delete(context.Background(), groupID, logID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Collect encapsulate the logic to manage different cloud providers
func (s *Store) Collect(groupID string, newLog *opsmngr.LogCollectionJob) (*opsmngr.LogCollectionJob, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).LogCollections.Create(context.Background(), groupID, newLog)
		return log, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ProcessDisks encapsulate the logic to manage different cloud providers
func (s *Store) DownloadLog(groupID, host, name string, out io.Writer, opts *atlas.DateRangetOptions) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Logs.Get(context.Background(), groupID, host, name, out, opts)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DownloadLogJob encapsulate the logic to manage different cloud providers
func (s *Store) DownloadLogJob(groupID, jobID string, out io.Writer) error {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		_, err := s.client.(*opsmngr.Client).Logs.Download(context.Background(), groupID, jobID, out)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
