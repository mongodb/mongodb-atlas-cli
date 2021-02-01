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
	yes                       = "yes"
	responseHeaderTimeout     = 1 * time.Minute
	tlsHandshakeTimeout       = 5 * time.Second
	timeout                   = 5 * time.Second
	keepAlive                 = 30 * time.Second
	maxIdleConns              = 5
	maxIdleConnsPerHost       = 4
	idleConnTimeout           = 30 * time.Second
	expectContinueTimeout     = 1 * time.Second
	versionManifestStaticPath = "https://opsmanager.mongodb.com/"
)

type Store struct {
	service       string
	baseURL       string
	caCertificate string
	skipVerify    string
	client        interface{}
}

var defaultTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepAlive,
	}).DialContext,
	MaxIdleConns:          maxIdleConns,
	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	Proxy:                 http.ProxyFromEnvironment,
	IdleConnTimeout:       idleConnTimeout,
	ExpectContinueTimeout: expectContinueTimeout,
}

func customCATransport(ca []byte) http.RoundTripper {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)
	tlsClientConfig := &tls.Config{ //nolint:gosec // we let users set custom certificates
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

type Config interface {
	Service() string
	PublicAPIKey() string
	PrivateAPIKey() string
	OpsManagerURL() string
	OpsManagerCACertificate() string
	OpsManagerSkipVerify() string
	OpsManagerVersionManifestURL() string
}

func authenticatedClient(c Config) (*http.Client, error) {
	t := &digest.Transport{
		Username: c.PublicAPIKey(),
		Password: c.PrivateAPIKey(),
	}
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		dat, err := ioutil.ReadFile(caCertificate)
		if err != nil {
			return nil, err
		}
		t.Transport = customCATransport(dat)
	} else if skipVerify := c.OpsManagerSkipVerify(); skipVerify == yes {
		t.Transport = skipVerifyTransport()
	} else {
		t.Transport = http.DefaultTransport
	}

	return t.Client()
}

func defaultClient(c Config) (*http.Client, error) {
	client := http.DefaultClient
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		dat, err := ioutil.ReadFile(caCertificate)
		if err != nil {
			return nil, err
		}
		client.Transport = customCATransport(dat)
	} else if skipVerify := c.OpsManagerSkipVerify(); skipVerify == yes {
		client.Transport = skipVerifyTransport()
	} else {
		client.Transport = defaultTransport
	}

	return client, nil
}

// New get the appropriate client for the profile/service selected
func New(c Config) (*Store, error) {
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

// NewUnauthenticated a client to interact with the Ops Manager APIs that don't require authentication
func NewUnauthenticated(c Config) (*Store, error) {
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
	if err := s.setOpsManagerClient(client); err != nil {
		return nil, err
	}
	return s, nil
}

// NewVersionManifest ets the appropriate client for the manifest version page
func NewVersionManifest(c Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()
	if s.service != config.OpsManagerService {
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
	s.baseURL = versionManifestStaticPath

	if baseURL := c.OpsManagerVersionManifestURL(); baseURL != "" {
		s.baseURL = baseURL
	}

	client, err := defaultClient(c)
	if err != nil {
		return nil, err
	}

	if err := s.setOpsManagerClient(client); err != nil {
		return nil, err
	}

	return s, nil
}

// NewPrivateUnauth gets the appropriate client for the atlas private api
func NewPrivateUnauth(c Config) (*Store, error) {
	s := new(Store)
	s.service = c.Service()
	if s.service != config.CloudService {
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
	s.baseURL = atlas.CloudURL

	if configURL := c.OpsManagerURL(); configURL != "" {
		s.baseURL = s.apiPath(configURL)
	}

	client, err := defaultClient(c)
	if err != nil {
		return nil, err
	}

	if err2 := s.setAtlasClient(client); err2 != nil {
		return nil, err2
	}

	return s, err
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
		return baseURL + atlas.APIPublicV1Path
	}
	return baseURL + opsmngr.APIPublicV1Path
}
