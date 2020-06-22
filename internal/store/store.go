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
	"net"
	"net/http"
	"time"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/version"
	"go.mongodb.org/ops-manager/opsmngr"
)

var userAgent = fmt.Sprintf("%s/%s", config.ToolName, version.Version)

const (
	atlasAPIPath = "api/atlas/v1.0/"
	yes          = "yes"
)

type Store struct {
	service       string
	baseURL       string
	caCertificate string
	skipVerify    string
	client        interface{}
}

func newTransport(c config.Config) *digest.Transport {
	t := &digest.Transport{
		Username: c.PublicAPIKey(),
		Password: c.PrivateAPIKey(),
	}
	t.Transport = &http.Transport{
		ResponseHeaderTimeout: 10 * time.Minute,
		TLSHandshakeTimeout:   10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   4,
		Proxy:                 http.ProxyFromEnvironment,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return t
}

// New get the appropriate client for the profile/service selected
func New(c config.Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()

	client, err := newTransport(c).Client()

	if err != nil {
		return nil, err
	}

	if configURL := c.OpsManagerURL(); configURL != "" {
		s.baseURL = s.apiPath(configURL)
	}
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		s.caCertificate = caCertificate
	}
	if skipVerify := c.OpsManagerSkipVerify(); skipVerify != yes {
		s.skipVerify = skipVerify
	}

	switch s.service {
	case config.CloudService:
		err = s.setAtlasClient(client)
	case config.CloudManagerService, config.OpsManagerService:
		err = s.setOpsManagerClient(client)
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}

	return s, err
}

func NewUnauthenticated(c config.Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()

	if configURL := c.OpsManagerURL(); configURL != "" {
		s.baseURL = s.apiPath(configURL)
	}
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		s.caCertificate = caCertificate
	}
	if skipVerify := c.OpsManagerSkipVerify(); skipVerify != yes {
		s.skipVerify = skipVerify
	}

	err := s.setOpsManagerClient(nil)

	if s.service != config.OpsManagerService {
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}

	return s, err
}

func (s *Store) setAtlasClient(client *http.Client) error {
	opts := make([]atlas.ClientOpt, 0)
	opts = append(opts, atlas.SetUserAgent(userAgent))
	if s.baseURL != "" {
		opts = append(opts, atlas.SetBaseURL(s.baseURL))
	}
	c, err := atlas.New(client, opts...)

	s.client = c
	return err
}

func (s *Store) setOpsManagerClient(client *http.Client) error {
	opts := make([]opsmngr.ClientOpt, 0)
	opts = append(opts, opsmngr.SetUserAgent(userAgent))
	if s.baseURL != "" {
		opts = append(opts, opsmngr.SetBaseURL(s.baseURL))
	}
	if s.caCertificate != "" {
		opts = append(opts, opsmngr.OptionCAValidate(s.caCertificate))
	}
	if s.skipVerify == yes {
		opts = append(opts, opsmngr.OptionSkipVerify())
	}
	c, err := opsmngr.New(client, opts...)

	s.client = c
	return err
}

func (s *Store) apiPath(baseURL string) string {
	if s.service == config.CloudService {
		return baseURL + atlasAPIPath
	}
	return baseURL + opsmngr.APIPublicV1Path
}
