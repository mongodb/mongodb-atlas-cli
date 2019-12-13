package store

import (
	"context"
	"fmt"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ProjectLister interface {
	GetAllProjects() (interface{}, error)
}

type ProjectCreator interface {
	CreateProject(string, string) (interface{}, error)
}

type ProjectStore interface {
	ProjectLister
	ProjectCreator
}

// GetAllProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetAllProjects() (interface{}, error) {

	switch s.service {
	case CLoudService:
		result, _, err := s.client.(*atlas.Client).Projects.GetAllProjects(context.Background())
		return result, err
	case CloudManagerService, OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).Projects.GetAllProjects(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateProject encapsulate the logic to manage different cloud providers
func (s *Store) CreateProject(name, orgID string) (interface{}, error) {

	switch s.service {
	case CLoudService:
		project := &atlas.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*atlas.Client).Projects.Create(context.Background(), project)
		return result, err
	case CloudManagerService, OpsManagerService:
		project := &cloudmanager.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*cloudmanager.Client).Projects.Create(context.Background(), project)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
