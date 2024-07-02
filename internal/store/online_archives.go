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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_online_archives.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store OnlineArchiveLister,OnlineArchiveDescriber,OnlineArchiveCreator,OnlineArchiveUpdater,OnlineArchiveDeleter

type OnlineArchiveLister interface {
	OnlineArchives(string, string, *atlas.ListOptions) (*atlasv2.PaginatedOnlineArchive, error)
}

type OnlineArchiveDescriber interface {
	OnlineArchive(string, string, string) (*atlasv2.BackupOnlineArchive, error)
}

type OnlineArchiveCreator interface {
	CreateOnlineArchive(string, string, *atlasv2.BackupOnlineArchiveCreate) (*atlasv2.BackupOnlineArchive, error)
}

type OnlineArchiveUpdater interface {
	UpdateOnlineArchive(string, string, *atlasv2.BackupOnlineArchive) (*atlasv2.BackupOnlineArchive, error)
}

type OnlineArchiveDeleter interface {
	DeleteOnlineArchive(string, string, string) error
}

// OnlineArchives encapsulate the logic to manage different cloud providers.
func (s *Store) OnlineArchives(projectID, clusterName string, lstOpt *atlas.ListOptions) (*atlasv2.PaginatedOnlineArchive, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.OnlineArchiveApi.ListOnlineArchives(s.ctx, projectID, clusterName).
		PageNum(lstOpt.PageNum).ItemsPerPage(lstOpt.ItemsPerPage).IncludeCount(lstOpt.IncludeCount).Execute()
	return result, err
}

// OnlineArchive encapsulate the logic to manage different cloud providers.
func (s *Store) OnlineArchive(projectID, clusterName, archiveID string) (*atlasv2.BackupOnlineArchive, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.OnlineArchiveApi.GetOnlineArchive(s.ctx, projectID, archiveID, clusterName).Execute()
	return result, err
}

// CreateOnlineArchive encapsulate the logic to manage different cloud providers.
func (s *Store) CreateOnlineArchive(projectID, clusterName string, archive *atlasv2.BackupOnlineArchiveCreate) (*atlasv2.BackupOnlineArchive, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.OnlineArchiveApi.CreateOnlineArchive(s.ctx, projectID, clusterName, archive).Execute()
	return result, err
}

// UpdateOnlineArchive encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateOnlineArchive(projectID, clusterName string, archive *atlasv2.BackupOnlineArchive) (*atlasv2.BackupOnlineArchive, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.OnlineArchiveApi.UpdateOnlineArchive(s.ctx, projectID, archive.GetId(), clusterName, archive).Execute()
	return result, err
}

// DeleteOnlineArchive encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOnlineArchive(projectID, clusterName, archiveID string) error {
	if s.service == config.CloudGovService {
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	_, _, err := s.clientv2.OnlineArchiveApi.DeleteOnlineArchive(s.ctx, projectID, archiveID, clusterName).Execute()
	return err
}
