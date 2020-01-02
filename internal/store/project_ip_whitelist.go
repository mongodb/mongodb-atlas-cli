package store

import (
	"context"
	"fmt"

	"github.com/10gen/mcli/internal/config"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ProjectIPWhitelistCreator interface {
	CreateProjectIPWhitelist(*atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error)
}

type ProjectIPWhitelistStore interface {
	ProjectIPWhitelistCreator
}

// CreateProjectIPWhitelist encapsulate the logic to manage different cloud providers
func (s *Store) CreateProjectIPWhitelist(whitelist *atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPWhitelist.Create(context.Background(), whitelist.GroupID, []*atlas.ProjectIPWhitelist{whitelist})
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
