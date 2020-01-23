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

var userAgent = fmt.Sprintf("%s/%s (%s;%s)", config.Name, version.Version, runtime.GOOS, runtime.GOARCH)

const atlasAPIPath = "/api/atlas/v1.0/"

type Store struct {
	service string
	baseURL *url.URL
	client  interface{}
}

type Config interface {
	Service() string
	PublicAPIKey() string
	PrivateAPIKey() string
	OpsManagerURL() string
}

// New get the appropriate client for the profile/service selected
func New(c Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()
	client, err := digest.NewTransport(c.PublicAPIKey(), c.PrivateAPIKey()).Client()

	if err != nil {
		return nil, err
	}

	if c.OpsManagerURL() != "" {
		baseURL, err := url.Parse(s.apiPath(c.OpsManagerURL()))
		if err != nil {
			return nil, err
		}
		s.baseURL = baseURL
	}

	switch s.service {
	case config.CloudService:
		s.setAtlasClient(client)
	case config.CloudManagerService, config.OpsManagerService:
		s.setCloudManagerClient(client)
	default:
		return nil, errors.New("unsupported service")
	}

	return s, nil
}

func (s *Store) setAtlasClient(client *http.Client) {
	atlasClient := atlas.NewClient(client)
	if s.baseURL != nil {
		atlasClient.BaseURL = s.baseURL
	}
	atlasClient.UserAgent = userAgent

	s.client = atlasClient
}

func (s *Store) setCloudManagerClient(client *http.Client) {
	cmClient := cloudmanager.NewClient(client)
	if s.baseURL != nil {
		cmClient.BaseURL = s.baseURL
	}
	cmClient.UserAgent = userAgent

	s.client = cmClient
}

func (s *Store) apiPath(baseURL string) string {
	if s.service == config.CloudService {
		return baseURL + atlasAPIPath
	}
	return baseURL + cloudmanager.APIPublicV1Path
}
