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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/version"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

var userAgent = fmt.Sprintf("%s/%s (%s;%s)", config.ToolName, version.Version, runtime.GOOS, runtime.GOARCH)

const (
	atlasAPIPath          = "api/atlas/v1.0/"
	yes                   = "yes"
	responseHeaderTimeout = 10 * time.Minute
	tlsHandshakeTimeout   = 10 * time.Second
	timeout               = 10 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 100
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 90 * time.Second
	expectContinueTimeout = 1 * time.Second
)

type Store struct {
	service       string
	baseURL       string
	caCertificate string
	skipVerify    string
	client        interface{}
}

func customCATransport(ca []byte) http.RoundTripper {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            caCertPool,
	}
	return &http.Transport{
		ResponseHeaderTimeout: responseHeaderTimeout,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		Proxy:                 http.ProxyFromEnvironment,
		IdleConnTimeout:       idleConnTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
		TLSClientConfig:       tlsClientConfig,
	}
}

func skipVerifyTransport() http.RoundTripper {
	tlsClientConfig := &tls.Config{InsecureSkipVerify: true} //nolint:gosec // this is optional for some users
	return &http.Transport{
		ResponseHeaderTimeout: responseHeaderTimeout,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		Proxy:                 http.ProxyFromEnvironment,
		IdleConnTimeout:       idleConnTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
		TLSClientConfig:       tlsClientConfig,
	}
}

func authenticatedClient(c config.Config) (*http.Client, error) {
	t := &digest.Transport{
		Username: c.PublicAPIKey(),
		Password: c.PrivateAPIKey(),
	}
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		dat, err := ioutil.ReadFile(caCertificate)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(dat))
		t.Transport = customCATransport(dat)
	} else if skipVerify := c.OpsManagerSkipVerify(); skipVerify == yes {
		t.Transport = skipVerifyTransport()
	} else {
		t.Transport = http.DefaultTransport
	}

	return t.Client()
}

func defaultClient(c config.Config) (*http.Client, error) {
	client := http.DefaultClient
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		dat, err := ioutil.ReadFile(caCertificate)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(dat))
		client.Transport = customCATransport(dat)
	} else if skipVerify := c.OpsManagerSkipVerify(); skipVerify == yes {
		client.Transport = skipVerifyTransport()
	} else {
		client.Transport = http.DefaultTransport
	}

	return client, nil
}

// New get the appropriate client for the profile/service selected
func New(c config.Config) (*Store, error) {
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
	client, err := authenticatedClient(c)
	if err != nil {
		return nil, err
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

	if s.service != config.OpsManagerService {
		return nil, fmt.Errorf("unsupported service: %s", s.service)
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
	client, err := defaultClient(c)
	if err != nil {
		return nil, err
	}
	err = s.setOpsManagerClient(client)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) setAtlasClient(client *http.Client) error {
	opts := make([]atlas.ClientOpt, 0)
	if s.baseURL != "" {
		opts = append(opts, atlas.SetBaseURL(s.baseURL))
	}
	c, err := atlas.New(client, opts...)
	if err != nil {
		return err
	}
	c.UserAgent = userAgent
	s.client = c
	return nil
}

func (s *Store) setOpsManagerClient(client *http.Client) error {
	opts := make([]opsmngr.ClientOpt, 0)
	if s.baseURL != "" {
		opts = append(opts, opsmngr.SetBaseURL(s.baseURL))
	}
	c, err := opsmngr.New(client, opts...)
	if err != nil {
		return err
	}
	c.UserAgent = userAgent
	s.client = c
	return nil
}

func (s *Store) apiPath(baseURL string) string {
	if s.service == config.CloudService {
		return baseURL + atlasAPIPath
	}
	return baseURL + opsmngr.APIPublicV1Path
}
