package store

import (
	"context"
	"fmt"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
)

type IndexCreator interface {
	CreateIndex(string, string, *atlas.IndexConfiguration) error
}

// CreateIndex encapsulate the logic to manage different cloud providers
func (s *Store) CreateIndex(projectID, clusterName string, index *atlas.IndexConfiguration) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Indexes.Create(context.Background(), projectID, clusterName, index)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
