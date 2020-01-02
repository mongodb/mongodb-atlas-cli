package store

import (
	"context"
	"fmt"

	"github.com/10gen/mcli/internal/config"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type DatabaseUserCreator interface {
	CreateDatabaseUser(*atlas.DatabaseUser) (*atlas.DatabaseUser, error)
}

type DatabaseUserStore interface {
	DatabaseUserCreator
}

// CreateDatabaseUser encapsulate the logic to manage different cloud providers
func (s *Store) CreateDatabaseUser(user *atlas.DatabaseUser) (*atlas.DatabaseUser, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DatabaseUsers.Create(context.Background(), user.GroupID, user)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
