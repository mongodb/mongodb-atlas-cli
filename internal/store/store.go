// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/version"
	"go.mongodb.org/ops-manager/opsmngr"
)

var userAgent = fmt.Sprintf("%s/%s (%s;%s)", config.ToolName, version.Version, runtime.GOOS, runtime.GOARCH)

const atlasAPIPath = "api/atlas/v1.0/"

type Store struct {
	service string
	baseURL *url.URL
	client  interface{}
}

// New get the appropriate client for the profile/service selected
func New(c config.Config) (*Store, error) {
	s := new(Store)
	s.service = config.Service()
	client, err := digest.NewTransport(c.PublicAPIKey(), c.PrivateAPIKey()).Client()

	if err != nil {
		return nil, err
	}

	configURL := c.OpsManagerURL()
	if configURL != "" {
		apiPath := s.apiPath(configURL)
		baseURL, err := url.Parse(apiPath)
		if err != nil {
			return nil, err
		}
		s.baseURL = baseURL
	}

	switch s.service {
	case config.CloudService:
		s.setAtlasClient(client)
	case config.CloudManagerService, config.OpsManagerService:
		s.setOpsManagerClient(client)
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}

	return s, nil
}

func NewUnauthenticated(c config.Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()

	configURL := c.OpsManagerURL()
	if configURL != "" {
		apiPath := s.apiPath(configURL)
		baseURL, err := url.Parse(apiPath)
		if err != nil {
			return nil, err
		}
		s.baseURL = baseURL
	}

	switch s.service {
	case config.OpsManagerService:
		s.setOpsManagerClient(nil)
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
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

func (s *Store) setOpsManagerClient(client *http.Client) {
	cmClient := opsmngr.NewClient(client)
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
	return baseURL + opsmngr.APIPublicV1Path
}
