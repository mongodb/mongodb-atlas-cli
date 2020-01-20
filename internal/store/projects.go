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
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ProjectLister interface {
	GetAllProjects() (interface{}, error)
	GetOrgProjects(string) (interface{}, error)
}

type OrgProjectLister interface {
	GetOrgProjects(string) (interface{}, error)
}

type ProjectCreator interface {
	CreateProject(string, string) (interface{}, error)
}

type ProjectDeleter interface {
	DeleteProject(string) error
}

type ProjectStore interface {
	ProjectLister
	ProjectCreator
	ProjectDeleter
}

// GetAllProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetAllProjects() (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Projects.GetAllProjects(context.Background())
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).Projects.GetAllProjects(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// GetOrgProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetOrgProjects(orgID string) (interface{}, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).Organizations.GetProjects(context.Background(), orgID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateProject encapsulate the logic to manage different cloud providers
func (s *Store) CreateProject(name, orgID string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		project := &atlas.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*atlas.Client).Projects.Create(context.Background(), project)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		project := &cloudmanager.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*cloudmanager.Client).Projects.Create(context.Background(), project)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteProject encapsulate the logic to manage different cloud providers
func (s *Store) DeleteProject(projectID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Projects.Delete(context.Background(), projectID)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*cloudmanager.Client).Projects.Delete(context.Background(), projectID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
