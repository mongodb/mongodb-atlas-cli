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

package atlas

//go:generate mockgen -destination=../../mocks/atlas/store.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CredentialsGetter

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlasauth "go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	yes                   = "yes"
	responseHeaderTimeout = 1 * time.Minute
	telemetryTimeout      = 1 * time.Second
	tlsHandshakeTimeout   = 5 * time.Second
	timeout               = 5 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 5
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
	cloudGovServiceURL    = "https://cloud.mongodbgov.com/"
)

type Store struct {
	// Cloud or Gov
	service     string
	baseURL     string
	telemetry   bool
	username    string
	password    string
	accessToken *atlasauth.Token
	client      *atlas.Client
	clientv2    *atlasv2.APIClient
	ctx         context.Context
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

var telemetryTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   telemetryTimeout,
		KeepAlive: keepAlive,
	}).DialContext,
	MaxIdleConns:          maxIdleConns,
	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	Proxy:                 http.ProxyFromEnvironment,
	IdleConnTimeout:       idleConnTimeout,
	ExpectContinueTimeout: expectContinueTimeout,
}

func (s *Store) httpClient(httpTransport http.RoundTripper) (*http.Client, error) {
	if s.username == "" && s.password == "" && s.accessToken == nil {
		return &http.Client{Transport: httpTransport}, nil
	}
	if s.username != "" && s.password != "" {
		t := &digest.Transport{
			Username: s.username,
			Password: s.password,
		}
		t.Transport = httpTransport
		return t.Client()
	}
	tr := &Transport{
		token: s.accessToken,
		base:  httpTransport,
	}

	return &http.Client{Transport: tr}, nil
}

type Transport struct {
	token *atlasauth.Token
	base  http.RoundTripper
}

func (tr *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.token.SetAuthHeader(req)
	return tr.base.RoundTrip(req)
}

func (s *Store) transport() *http.Transport {
	switch {
	case s.telemetry:
		return telemetryTransport
	default:
		return defaultTransport
	}
}

// Option is any configuration for Store.
// New will take a list of Option and process them sequentially.
// The store package provides a list of pointers and preset set of Option you can use
// but you can implement your own.
type Option func(s *Store) error

// Options turns a list of Option instances into a single Option.
// This is a helper when combining multiple Option.
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

// WithBaseURL sets the base URL for the underling HTTP client.
// the url should not contain any path, to add the public API path use WithPublicPathBaseURL.
func WithBaseURL(configURL string) Option {
	return func(s *Store) error {
		s.baseURL = configURL
		return nil
	}
}

func Telemetry() Option {
	return func(s *Store) error {
		s.telemetry = true
		return nil
	}
}

// CredentialsGetter interface for how to get credentials when Store must be authenticated.
type CredentialsGetter interface {
	PublicAPIKey() string
	PrivateAPIKey() string
	Token() (*atlasauth.Token, error)
}

// WithAuthentication sets the store credentials.
func WithAuthentication(c CredentialsGetter) Option {
	return func(s *Store) error {
		s.username = c.PublicAPIKey()
		s.password = c.PrivateAPIKey()

		if s.username == "" && s.password == "" {
			t, err := c.Token()
			if err != nil {
				return err
			}
			s.accessToken = t
		}
		return nil
	}
}

// WithContext sets the store context.
func WithContext(ctx context.Context) Option {
	return func(s *Store) error {
		s.ctx = ctx
		return nil
	}
}

// setAtlasClient sets the internal client to use an Atlas client and methods.
func (s *Store) setAtlasClient(client *http.Client) error {
	opts := []atlas.ClientOpt{atlas.SetUserAgent(config.UserAgent)}
	if s.baseURL != "" {
		opts = append(opts, atlas.SetBaseURL(s.baseURL))
	}
	if log.IsDebugLevel() {
		opts = append(opts, atlas.SetWithRaw())
	}
	c, err := atlas.New(client, opts...)
	if err != nil {
		return err
	}

	err = s.createV2Client(client)
	if err != nil {
		return err
	}

	c.OnResponseProcessed(func(resp *atlas.Response) {
		respHeaders := ""
		for key, value := range resp.Header {
			respHeaders += fmt.Sprintf("%v: %v\n", key, strings.Join(value, " "))
		}

		_, _ = log.Debugf(`request:
%v %v
response:
%v %v
%v
%v
`, resp.Request.Method, resp.Request.URL.String(), resp.Proto, resp.Status, respHeaders, string(resp.Raw))
	})
	s.client = c
	return nil
}

/**
* Creates client for v2 generated API.
 */
func (s *Store) createV2Client(client *http.Client) error {
	opts := []atlasv2.ClientModifier{
		atlasv2.UseHTTPClient(client),
		atlasv2.UseUserAgent(config.UserAgent),
		atlasv2.UseDebug(log.IsDebugLevel())}

	if s.baseURL != "" {
		opts = append(opts, atlasv2.UseBaseURL(s.baseURL))
	}
	c, err := atlasv2.NewClient(opts...)
	if err != nil {
		return err
	}
	s.clientv2 = c
	return nil
}

// AuthenticatedConfig an interface of the methods needed to set up a Store.
type AuthenticatedConfig interface {
	CredentialsGetter
	ServiceGetter
}

type ServiceGetter interface {
	Service() string
	// Name of the config value for custom service URL. Named opsmannager for legacy reasons.
	OpsManagerURL() string
}

func baseURLOption(c ServiceGetter) Option {
	if configURL := c.OpsManagerURL(); configURL != "" {
		return WithBaseURL(configURL)
	} else if c.Service() == config.CloudGovService {
		return WithBaseURL(cloudGovServiceURL)
	}
	return nil
}

func Service(service string) Option {
	return func(s *Store) error {
		s.service = service
		return nil
	}
}

// AuthenticatedPreset is the default Option when connecting to the public API with authentication.
func AuthenticatedPreset(c AuthenticatedConfig) Option {
	options := []Option{Service(c.Service()), WithAuthentication(c)}
	if baseURLOpt := baseURLOption(c); baseURLOpt != nil {
		options = append(options, baseURLOpt)
	}
	return Options(options...)
}

// UnauthenticatedPreset is the default Option when connecting to the public API without authentication.
func UnauthenticatedPreset(c ServiceGetter) Option {
	options := []Option{Service(c.Service())}
	if option := baseURLOption(c); option != nil {
		options = append(options, option)
	}
	return Options(options...)
}

// New returns a new Store based on the given list of Option.
//
// Usage:
//
//	// get a new Store for Atlas
//	store := store.New(Service("cloud"))
//
//	// get a new Store for the public API based on a Config interface
//	store := store.New(AuthenticatedPreset(config))
//
//	// get a new Store for the private API based on a Config interface
//	store := store.New(PrivateAuthenticatedPreset(config))
func New(opts ...Option) (*Store, error) {
	store := new(Store)

	// apply the list of options to Server
	for _, opt := range opts {
		if err := opt(store); err != nil {
			return nil, err
		}
	}

	httpTransport := store.transport()
	client, err := store.httpClient(httpTransport)
	if err != nil {
		return nil, err
	}

	err = store.setAtlasClient(client)
	if err != nil {
		return nil, err
	}

	if store.ctx == nil {
		store.ctx = context.Background()
	}

	return store, nil
}
