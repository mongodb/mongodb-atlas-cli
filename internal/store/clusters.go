// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

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

type ClusterDeleter interface {
	DeleteCluster(string, string) error
}

type ClusterUpdater interface {
	UpdateCluster(*atlas.Cluster) (*atlas.Cluster, error)
}

type ClusterStore interface {
	ClusterLister
	ClusterDescriber
	ClusterCreator
	ClusterDeleter
	ClusterUpdater
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

// UpdateCluster encapsulate the logic to manage different cloud providers
func (s *Store) UpdateCluster(cluster *atlas.Cluster) (*atlas.Cluster, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Clusters.Update(context.Background(), cluster.GroupID, cluster.Name, cluster)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteCluster encapsulate the logic to manage different cloud providers
func (s *Store) DeleteCluster(projectID, name string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Clusters.Delete(context.Background(), projectID, name)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
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
