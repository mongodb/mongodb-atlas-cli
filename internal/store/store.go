package store

import (
	"errors"
	"net/http"
	"net/url"
	"runtime"

	"github.com/10gen/mcli/internal/version"
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	// DefaultUserAgent to be submitted by the client
	DefaultUserAgent = "mcli/" + version.Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	// CloudDefaultURL Atlas default URL
	CloudDefaultURL = "https://cloud-qa.mongodb.com/api/atlas/v1.0/"
	// CLoudService setting when using Atlas API
	CLoudService = "cloud"
	// CloudManagerService settings when using CLoud Manager API
	CloudManagerService = "cloud-manager"
	// OpsManagerService settings when using Ops Manager API
	OpsManagerService = "ops-manager"
)

type Store struct {
	service   string
	baseURL   *url.URL
	transport *http.Client
	client    interface{}
}

// New get the appropriate client for the profile/service selected
func New(service, publicAPIKey, privateAPIKey, baseURL string) (*Store, error) {
	c := &Store{}
	c.transport, _ = digest.NewTransport(publicAPIKey, privateAPIKey).Client()
	c.service = service
	c.baseURL, _ = url.Parse(baseURL)
	switch c.service {
	case CLoudService:
		c.client = c.atlas()
	case CloudManagerService:
		c.client = c.cloudManager()
	case OpsManagerService:
		c.client = c.opsManager()
	default:
		return nil, errors.New("unsupported service")
	}

	return c, nil
}

func (s *Store) atlas() *atlas.Client {
	atlasClient := atlas.NewClient(s.transport)
	atlasClient.BaseURL = s.baseURL
	atlasClient.UserAgent = DefaultUserAgent

	return atlasClient
}

func (s *Store) cloudManager() *cloudmanager.Client {
	cloudManagerClient := cloudmanager.NewClient(s.transport)
	cloudManagerClient.BaseURL = s.baseURL
	cloudManagerClient.UserAgent = DefaultUserAgent

	return cloudManagerClient
}

func (s *Store) opsManager() *cloudmanager.Client {
	opsManagerClient := cloudmanager.NewClient(s.transport)
	opsManagerClient.BaseURL = s.baseURL
	opsManagerClient.UserAgent = DefaultUserAgent

	return opsManagerClient
}
