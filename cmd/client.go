package cmd

import (
	"errors"
	"fmt"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/viper"
)

// newAuthenticatedClient get the appropriate client for the profile/provider selected
func newAuthenticatedClient(profile string) (interface{}, error) {
	// setup a transport to handle digest
	publicKey := viper.GetString(fmt.Sprintf("%s.public_key", profile))
	privateKey := viper.GetString(fmt.Sprintf("%s.private_key", profile))
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	provider := viper.GetString(fmt.Sprintf("%s.service", profile))
	switch provider {
	case "cloud":
		return atlas.New(client, atlas.SetBaseURL(CloudDefaultURL), atlas.SetUserAgent(DefaultUserAgent))
	case "cloud-manager":
		return cloudmanager.New(client, cloudmanager.SetBaseURL(cloudmanager.DefaultBaseURL), cloudmanager.SetUserAgent(DefaultUserAgent))
	case "ops-manager":
		baseURL := viper.GetString(fmt.Sprintf("%s.base_url", profile))
		return cloudmanager.New(client, cloudmanager.SetBaseURL(baseURL+"/api/public/v1.0/"), cloudmanager.SetUserAgent(DefaultUserAgent))
	default:
		return nil, errors.New("unsupported provider")
	}
}
