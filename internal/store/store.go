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
	skipVerify    bool
	username      string
	password      string
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

var skipVerifyTransport = &http.Transport{
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
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // this is optional for some users,
}

func customCATransport(ca []byte) *http.Transport {
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

func (s *Store) httpClient(httpTransport http.RoundTripper) (*http.Client, error) {
	if s.username == "" || s.password == "" {
		client := http.DefaultClient
		client.Transport = httpTransport
		return client, nil
	}
	t := &digest.Transport{
		Username: s.username,
		Password: s.password,
	}
	t.Transport = httpTransport
	return t.Client()
}

func (s *Store) transport() (*http.Transport, error) {
	switch {
	case s.caCertificate != "":
		dat, err := ioutil.ReadFile(s.caCertificate)
		if err != nil {
			return nil, err
		}
		return customCATransport(dat), nil
	case s.skipVerify:
		return skipVerifyTransport, nil
	default:
		return defaultTransport, nil
	}
}

// Option configures a Store.
type Option func(s *Store) error

// Options turns a list of Option instances into an Option.
func Options(opts ...Option) Option {
	return func(s *Store) error {
		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}
		return nil
	}
}

// Service configures the service.
func Service(service string) Option {
	return func(s *Store) error {
		s.service = service
		return nil
	}
}

// WithBaseURL configures the base URL for the underling client.
// the url should not contain any path, to add the public API path use WithPublicPathBaseURL.
func WithBaseURL(configURL string) Option {
	return func(s *Store) error {
		s.baseURL = configURL
		return nil
	}
}

// WithPublicPathBaseURL if you use WithBaseURL and need the Store to connect to the public API.
func WithPublicPathBaseURL() Option {
	return func(s *Store) error {
		if s.service == config.CloudService {
			s.baseURL += atlas.APIPublicV1Path
			return nil
		}
		s.baseURL += opsmngr.APIPublicV1Path
		return nil
	}
}

// WithCACertificate when using a custom CA certificate.
func WithCACertificate(caCertificate string) Option {
	return func(s *Store) error {
		s.caCertificate = caCertificate
		return nil
	}
}

// SkipVerify skips CA certificate verification, use at your own risk.
func SkipVerify() Option {
	return func(s *Store) error {
		s.skipVerify = true
		return nil
	}
}

// CredentialsGetter how to get credentials when Store must be authenticated.
type CredentialsGetter interface {
	PublicAPIKey() string
	PrivateAPIKey() string
}

// WithAuthentication sets the store credentials
func WithAuthentication(c CredentialsGetter) Option {
	return func(s *Store) error {
		s.username = c.PublicAPIKey()
		s.password = c.PrivateAPIKey()
		return nil
	}
}

func withAtlasClient(client *http.Client) Option {
	return func(s *Store) error {
		opts := []atlas.ClientOpt{atlas.SetUserAgent(userAgent)}
		if s.baseURL != "" {
			opts = append(opts, atlas.SetBaseURL(s.baseURL))
		}
		c, err := atlas.New(client, opts...)
		if err != nil {
			return err
		}
		s.client = c
		return nil
	}
}

func withOpsManagerClient(client *http.Client) Option {
	return func(s *Store) error {
		opts := []opsmngr.ClientOpt{opsmngr.SetUserAgent(userAgent)}
		if s.baseURL != "" {
			opts = append(opts, opsmngr.SetBaseURL(s.baseURL))
		}
		c, err := opsmngr.New(client, opts...)
		if err != nil {
			return err
		}

		s.client = c
		return nil
	}
}

// TransportConfigGetter ops manager special transport settings.
type TransportConfigGetter interface {
	OpsManagerCACertificate() string
	OpsManagerSkipVerify() string
}

// NetworkPresets use when reading any special network settings.
func NetworkPresets(c TransportConfigGetter) Option {
	options := make([]Option, 0)
	if caCertificate := c.OpsManagerCACertificate(); caCertificate != "" {
		options = append(options, WithCACertificate(caCertificate))
	}
	if skipVerify := c.OpsManagerSkipVerify(); skipVerify != yes {
		options = append(options, SkipVerify())
	}
	return Options(options...)
}

// Config all settings for a new Store
type Config interface {
	CredentialsGetter
	TransportConfigGetter
	Service() string
	OpsManagerURL() string
}

// PublicAuthenticatedPreset default settings when connecting to the public API with authentication.
func PublicAuthenticatedPreset(c Config) Option {
	options := []Option{Service(c.Service()), WithAuthentication(c)}
	if configURL := c.OpsManagerURL(); configURL != "" {
		options = append(options, WithBaseURL(configURL), WithPublicPathBaseURL())
	}
	options = append(options, NetworkPresets(c))
	return Options(options...)
}

// PublicUnauthenticatedPreset default settings when connecting to the public API without authentication.
func PublicUnauthenticatedPreset(c Config) Option {
	options := []Option{Service(c.Service())}
	if configURL := c.OpsManagerURL(); configURL != "" {
		options = append(options, WithBaseURL(configURL), WithPublicPathBaseURL())
	}
	options = append(options, NetworkPresets(c))
	return Options(options...)
}

// PrivateAuthenticatedPreset default settings when connecting to the private API with authentication.
func PrivateAuthenticatedPreset(c Config) Option {
	options := []Option{Service(c.Service()), WithAuthentication(c)}
	if configURL := c.OpsManagerURL(); configURL != "" {
		options = append(options, WithBaseURL(configURL))
	}
	options = append(options, NetworkPresets(c))
	return Options(options...)
}

// PrivateUnauthenticatedPreset default settings when connecting to the private API without authentication.
func PrivateUnauthenticatedPreset(c Config) Option {
	options := []Option{Service(c.Service())}
	if configURL := c.OpsManagerURL(); configURL != "" {
		options = append(options, WithBaseURL(configURL))
	}
	options = append(options, NetworkPresets(c))
	return Options(options...)
}

// New return a new Store based on the given list of Option.
//
// Usage:
//
//	// get a new client for Atlas
//	store := store.New(WithService("cloud"))
//
//	// get a new client based on the config
//	store := store.New(PublicAuthenticatedPreset(config))
func New(opts ...Option) (*Store, error) {
	store := new(Store)

	// apply the list of options to Server
	for _, opt := range opts {
		if err := opt(store); err != nil {
			return nil, err
		}
	}

	httpTransport, err := store.transport()
	if err != nil {
		return nil, err
	}
	client, err := store.httpClient(httpTransport)
	if err != nil {
		return nil, err
	}

	switch store.service {
	case config.CloudService:
		err = withAtlasClient(client)(store)
	case config.CloudManagerService, config.OpsManagerService:
		err = withOpsManagerClient(client)(store)
	default:
		return nil, fmt.Errorf("unsupported service: %s", store.service)
	}
	if err != nil {
		return nil, err
	}
	return store, nil
}

type ManifestGetter interface {
	Service() string
	OpsManagerVersionManifestURL() string
}

// NewVersionManifest ets the appropriate client for the manifest version page
func NewVersionManifest(c ManifestGetter) (*Store, error) {
	s := new(Store)
	s.service = c.Service()
	if s.service != config.OpsManagerService {
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
	s.baseURL = versionManifestStaticPath
	if baseURL := c.OpsManagerVersionManifestURL(); baseURL != "" {
		s.baseURL = baseURL
	}
	if err := withOpsManagerClient(http.DefaultClient)(s); err != nil {
		return nil, err
	}

	return s, nil
}
