package cli

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Profile string
}

// GetService get configured service
func (c *Configuration) GetService() string {
	return viper.GetString(fmt.Sprintf("%s.service", c.Profile))
}

// SetService set configured service
func (c *Configuration) SetService(service string) {
	viper.Set(fmt.Sprintf("%s.service", c.Profile), service)
}

// GetPublicAPIKey get configured public api key
func (c *Configuration) GetPublicAPIKey() string {
	return viper.GetString(fmt.Sprintf("%s.public_api_key", c.Profile))
}

// SetPublicAPIKey set configured publicAPIKey
func (c *Configuration) SetPublicAPIKey(publicAPIKey string) {
	viper.Set(fmt.Sprintf("%s.public_api_key", c.Profile), publicAPIKey)
}

// GetPrivateAPIKey get configured private api key
func (c *Configuration) GetPrivateAPIKey() string {
	return viper.GetString(fmt.Sprintf("%s.private_api_key", c.Profile))
}

// SetPrivateAPIKey set configured private api key
func (c *Configuration) SetPrivateAPIKey(privateAPIKey string) {
	viper.Set(fmt.Sprintf("%s.private_api_key", c.Profile), privateAPIKey)
}

// GetOpsManagerURL get configured ops manager base url
func (c *Configuration) GetOpsManagerURL() string {
	return viper.GetString(fmt.Sprintf("%s.ops_manager_url", c.Profile)) + "/api/public/v1.0/"
}

// SetOpsManagerURL set configured ops manager base url
func (c *Configuration) SetOpsManagerURL(opsManagerURL string) {
	viper.Set(fmt.Sprintf("%s.ops_manager_url", c.Profile), opsManagerURL)
}

// Save save the configuration to disk
func (c *Configuration) Save() error {
	return viper.WriteConfig()
}
