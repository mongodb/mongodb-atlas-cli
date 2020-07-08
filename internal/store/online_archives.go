package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_online_archives.go -package=mocks github.com/mongodb/mongocli/internal/store OnlineArchiveLister

type OnlineArchiveLister interface {
	OnlineArchives(string, string) ([]*atlas.OnlineArchive, error)
}

// Containers encapsulate the logic to manage different cloud providers
func (s *Store) OnlineArchives(projectID, clusterName string) ([]*atlas.OnlineArchive, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).OnlineArchives.List(context.Background(), projectID, clusterName)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
