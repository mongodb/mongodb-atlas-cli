package store

import (
	"context"
	"fmt"

	"github.com/10gen/mcli/internal/config"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ClusterLister interface {
	ProjectClusters(string, *atlas.ListOptions) ([]atlas.Cluster, error)
}

type ClusterDescriber interface {
	Cluster(string, string) (*atlas.Cluster, error)
}

type ClusterCreator interface {
	CreateCluster(*atlas.Cluster) (*atlas.Cluster, error)
}

type ClusterStore interface {
	ClusterLister
	ClusterDescriber
	ClusterCreator
}

// CreateCluster encapsulate the logic to manage different cloud providers
func (s *Store) CreateCluster(cluster *atlas.Cluster) (*atlas.Cluster, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Clusters.Create(context.Background(), cluster.GroupID, cluster)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ProjectClusters encapsulate the logic to manage different cloud providers
func (s *Store) ProjectClusters(projectID string, opts *atlas.ListOptions) ([]atlas.Cluster, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Clusters.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Cluster encapsulate the logic to manage different cloud providers
func (s *Store) Cluster(projectID string, name string) (*atlas.Cluster, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Clusters.Get(context.Background(), projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
