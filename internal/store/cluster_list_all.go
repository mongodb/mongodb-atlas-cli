package store

import (
	"context"
	"fmt"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type ListAllClusters interface {
	ListAllClusters() ([]om.AllClustersProject, error)
}

// CreateOwner encapsulate the logic to manage different cloud providers
func (s *Store) ListAllClusters() ([]om.AllClustersProject, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*om.Client).AllCusters.List(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
