package store

import (
	"errors"
	"net/http"
	"net/url"
	"runtime"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/version"
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	// DefaultUserAgent to be submitted by the client
	DefaultUserAgent = config.Name + "/" + version.Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
)

type Store struct {
	service   string
	baseURL   *url.URL
	transport *http.Client
	client    interface{}
}

// New get the appropriate client for the profile/service selected
func New(c config.Config) (*Store, error) {
	s := &Store{service: c.GetService()}
	s.transport, _ = digest.NewTransport(c.GetPublicAPIKey(), c.GetPrivateAPIKey()).Client()

	if c.GetAPIPath() != "" {
		s.baseURL, _ = url.Parse(c.GetAPIPath())
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

func (s *Store) atlas() *atlas.Client {
	atlasClient := atlas.NewClient(s.transport)
	if s.baseURL != nil {
		atlasClient.BaseURL = s.baseURL
	}
	atlasClient.UserAgent = DefaultUserAgent

	return atlasClient
}

func (s *Store) cloudManager() *cloudmanager.Client {
	cloudManagerClient := cloudmanager.NewClient(s.transport)
	if s.baseURL != nil {
		cloudManagerClient.BaseURL = s.baseURL
	}
	cloudManagerClient.UserAgent = DefaultUserAgent

	return cloudManagerClient
}

func (s *Store) opsManager() *cloudmanager.Client {
	opsManagerClient := cloudmanager.NewClient(s.transport)
	opsManagerClient.BaseURL = s.baseURL
	opsManagerClient.UserAgent = DefaultUserAgent

	return opsManagerClient
}
