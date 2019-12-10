package cmd

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/viper"
)

// Version for client
// DefaultBaseURL API default base URL
// DefaultUserAgent To be submitted by the client
const (
	Version          = "0.1"
	DefaultUserAgent = "mcli/" + Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
)

//Config ...
type Config struct {
	PublicKey  string
	PrivateKey string
}

//NewClient ...
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
	if provider == "cloud" {
		//Initialize the MongoDB Atlas API Client.
		return mongodbatlas.NewClient(client), nil
	}
	if provider == "cloud-manager" {
		return cloudmanager.NewClient(client), nil
	}
	if provider == "ops-manager" {
		baseURL := viper.GetString(fmt.Sprintf("%s.base_url", profile))
		return cloudmanager.New(client, cloudmanager.SetBaseURL(baseURL+"/api/public/v1.0/"), cloudmanager.SetUserAgent(DefaultUserAgent))
	}
	return nil, errors.New("unsupported provider")
}
