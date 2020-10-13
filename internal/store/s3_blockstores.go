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

//go:generate mockgen -destination=../mocks/mock_backup_s3_blockstores.go -package=mocks github.com/mongodb/mongocli/internal/store S3BlockstoresLister,S3BlockstoresDeleter,S3BlockstoresCreator,S3BlockstoresUpdater,S3BlockstoresDescriber

type S3BlockstoresLister interface {
	ListS3Blockstores(*atlas.ListOptions) (*opsmngr.S3Blockstores, error)
}

type S3BlockstoresDeleter interface {
	DeleteS3Blockstore(string) error
}

type S3BlockstoresCreator interface {
	CreateS3Blockstores(*opsmngr.S3Blockstore) (*opsmngr.S3Blockstore, error)
}

type S3BlockstoresUpdater interface {
	UpdateS3Blockstores(string, *opsmngr.S3Blockstore) (*opsmngr.S3Blockstore, error)
}

type S3BlockstoresDescriber interface {
	GetS3Blockstore(string) (*opsmngr.S3Blockstore, error)
}

// ListS3Blockstores encapsulates the logic to manage different cloud providers
func (s *Store) ListS3Blockstores(options *atlas.ListOptions) (*opsmngr.S3Blockstores, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).S3BlockstoreConfig.List(context.Background(), options)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteS3Blockstores encapsulates the logic to manage different cloud providers
func (s *Store) DeleteS3Blockstore(blockstoreID string) error {
	switch s.service {
	case config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).S3BlockstoreConfig.Delete(context.Background(), blockstoreID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateS3Blockstores encapsulates the logic to manage different cloud providers
func (s *Store) CreateS3Blockstores(blockstore *opsmngr.S3Blockstore) (*opsmngr.S3Blockstore, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).S3BlockstoreConfig.Create(context.Background(), blockstore)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UpdateS3Blockstores encapsulates the logic to manage different cloud providers
func (s *Store) UpdateS3Blockstores(blockstoreID string, blockstore *opsmngr.S3Blockstore) (*opsmngr.S3Blockstore, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).S3BlockstoreConfig.Update(context.Background(), blockstoreID, blockstore)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ListS3Blockstores encapsulates the logic to manage different cloud providers
func (s *Store) GetS3Blockstore(blockstoreID string) (*opsmngr.S3Blockstore, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).S3BlockstoreConfig.Get(context.Background(), blockstoreID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
