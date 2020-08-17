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
)

//go:generate mockgen -destination=../mocks/mock_peeringConnections.go -package=mocks github.com/mongodb/mongocli/internal/store PeeringConnectionLister,PeeringConnectionDeleter

type PeeringConnectionLister interface {
	PeeringConnections(string, *atlas.ListOptions) ([]atlas.Peer, error)
}

type PeeringConnectionDeleter interface {
	DeletePeeringConnection(string, string) error
}

// PeeringConnections encapsulates the logic to manage different cloud providers
func (s *Store) PeeringConnections(projectID string, opts *atlas.ListOptions) ([]atlas.Peer, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Peers.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers
func (s *Store) DeletePeeringConnection(projectID, peerID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Peers.Delete(context.Background(), projectID, peerID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
