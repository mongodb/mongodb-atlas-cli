package store

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_global_deployment.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store GlobalDeploymentDescriber

type GlobalDeploymentDescriber interface {
	GlobalDeployment(string, string) (*atlas.GlobalCluster, error)
}

func (s *Store) GlobalDeployment(projectID, instanceName string) (*atlas.GlobalCluster, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).GlobalClusters.Get(s.ctx, projectID, instanceName)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
