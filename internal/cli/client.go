package cli

import (
	"errors"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

// newAuthenticatedClient get the appropriate client for the profile/provider selected
func newAuthenticatedClient(c *Configuration) (interface{}, error) {
	transport := digest.NewTransport(c.GetPublicAPIKey(), c.GetPrivateAPIKey())

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	switch c.GetService() {
	case "cloud":
		return atlas.New(client, atlas.SetBaseURL(CloudDefaultURL), atlas.SetUserAgent(DefaultUserAgent))
	case "cloud-manager":
		return cloudmanager.New(client, cloudmanager.SetBaseURL(cloudmanager.DefaultBaseURL), cloudmanager.SetUserAgent(DefaultUserAgent))
	case "ops-manager":
		baseURL := c.GetOpsManagerURL()
		if baseURL == "" {
			return nil, errors.New("ops manager url not set")
		}
		return cloudmanager.New(client, cloudmanager.SetBaseURL(baseURL), cloudmanager.SetUserAgent(DefaultUserAgent))
	default:
		return nil, errors.New("unsupported provider")
	}
}
