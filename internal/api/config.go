package api

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"go.mongodb.org/atlas/mongodbatlas"
)

type AuthenticatedConfigWrapper struct {
	authenticatedConfig store.AuthenticatedConfig
}

func NewAuthenticatedConfigWrapper(authenticatedConfig store.AuthenticatedConfig) *AuthenticatedConfigWrapper {
	return &AuthenticatedConfigWrapper{
		authenticatedConfig: authenticatedConfig,
	}
}

func (c *AuthenticatedConfigWrapper) GetAccessToken() (string, error) {
	token, err := c.authenticatedConfig.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (c *AuthenticatedConfigWrapper) GetBaseURL() (string, error) {
	// If the profile has overwritten the URL, use that one
	if configURL := c.authenticatedConfig.OpsManagerURL(); configURL != "" {
		return configURL, nil
	}

	// If the service is cloud gov, use the cloud gov base url
	if c.authenticatedConfig.Service() == config.CloudGovService {
		return store.CloudGovServiceURL, nil
	}

	// By default, return the default base URL
	return mongodbatlas.CloudURL, nil
}
