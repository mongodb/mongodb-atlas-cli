package cli

import (
	"context"
	"fmt"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ProjectsService interface {
	GetAllProjects() (interface{}, *atlas.Response, error)
	Create(string) (interface{}, *atlas.Response, error)
}

// Projects handles
type Projects struct {
	Configuration *Configuration
}

// ListProjects encapsulate the logic to manage different cloud providers
func (s *Projects) ListProjects() (interface{}, *atlas.Response, error) {
	client, err := newAuthenticatedClient(s.Configuration)
	if err != nil {
		return nil, nil, err
	}

	service := s.Configuration.GetService()
	switch service {
	case CLoudService:
		return client.(*atlas.Client).Projects.GetAllProjects(context.Background())
	case CloudManagerService, OpsManagerService:
		return client.(*cloudmanager.Client).Projects.GetAllProjects(context.Background())
	default:
		return nil, nil, fmt.Errorf("unsupported service: %s", service)
	}
}

// CreateProject encapsulate the logic to manage different cloud providers
func (s *Projects) CreateProject(name, orgID string) (interface{}, *atlas.Response, error) {
	client, err := newAuthenticatedClient(s.Configuration)
	if err != nil {
		return nil, nil, err
	}

	service := s.Configuration.GetService()

	switch service {
	case CLoudService:
		project := &atlas.Project{Name: name, OrgID: orgID}
		return client.(*atlas.Client).Projects.Create(context.Background(), project)
	case CloudManagerService, OpsManagerService:
		project := &cloudmanager.Project{Name: name, OrgID: orgID}
		return client.(*cloudmanager.Client).Projects.Create(context.Background(), project)
	default:
		return nil, nil, fmt.Errorf("unsupported service: %s", service)
	}
}
