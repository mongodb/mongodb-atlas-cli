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

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_backup_config.go -package=mocks github.com/mongodb/mongocli/internal/store BackupConfigGetter,BackupConfigLister,BackupConfigUpdater

type BackupConfigGetter interface {
	GetBackupConfig(string, string) (*opsmngr.BackupConfig, error)
}

type BackupConfigLister interface {
	ListBackupConfigs(string, *atlas.ListOptions) (*opsmngr.BackupConfigs, error)
}

type BackupConfigUpdater interface {
	UpdateBackupConfig(*opsmngr.BackupConfig) (*opsmngr.BackupConfig, error)
}

// GetBackupConfig encapsulates the logic to manage different cloud providers
func (s *Store) GetBackupConfig(projectID, clusterID string) (*opsmngr.BackupConfig, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).BackupConfigs.Get(context.Background(), projectID, clusterID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ListBackupConfigs encapsulates the logic to manage different cloud providers
func (s *Store) ListBackupConfigs(projectID string, options *atlas.ListOptions) (*opsmngr.BackupConfigs, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).BackupConfigs.List(context.Background(), projectID, options)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UpdateBackupConfig encapsulates the logic to manage different cloud providers
func (s *Store) UpdateBackupConfig(backupConfig *opsmngr.BackupConfig) (*opsmngr.BackupConfig, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).BackupConfigs.Update(context.Background(), backupConfig.GroupID, backupConfig.ClusterID, backupConfig)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
