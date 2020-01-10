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
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/version"
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type Store struct {
	service   string
	baseURL   *url.URL
	transport *http.Client
	client    interface{}
}

// New get the appropriate client for the profile/service selected
func New(c config.Config) (*Store, error) {
	s := &Store{service: c.Service()}
	s.transport, _ = digest.NewTransport(c.PublicAPIKey(), c.PrivateAPIKey()).Client()

	if c.APIPath() != "" {
		s.baseURL, _ = url.Parse(c.APIPath())
	}

	// fmt.Println("s.baseURL", s.baseURL)
	switch s.service {
	case config.CloudService:
		s.client = s.atlas()
	case config.CloudManagerService:
		s.client = s.cloudManager()
	case config.OpsManagerService:
		s.client = s.opsManager()
	default:
		return nil, errors.New("unsupported service")
	}

	return s, nil
}

func (s *Store) userAgent() string {
	return fmt.Sprintf("%s/%s (%s;%s)", config.Name, version.Version, runtime.GOOS, runtime.GOARCH)
}

func (s *Store) atlas() *atlas.Client {
	atlasClient := atlas.NewClient(s.transport)
	if s.baseURL != nil {
		atlasClient.BaseURL = s.baseURL
	}
	atlasClient.UserAgent = s.userAgent()

	return atlasClient
}

func (s *Store) cloudManager() *cloudmanager.Client {
	cloudManagerClient := cloudmanager.NewClient(s.transport)
	if s.baseURL != nil {
		cloudManagerClient.BaseURL = s.baseURL
	}
	cloudManagerClient.UserAgent = s.userAgent()

	return cloudManagerClient
}

func (s *Store) opsManager() *cloudmanager.Client {
	opsManagerClient := cloudmanager.NewClient(s.transport)
	opsManagerClient.BaseURL = s.baseURL
	opsManagerClient.UserAgent = s.userAgent()

	return opsManagerClient
}
