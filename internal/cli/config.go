package cli

import (
	"fmt"

	"github.com/10gen/mcli/internal/store"
	"github.com/spf13/viper"
)

type Config interface {
	GetService() string
	SetService(string)
	GetPublicAPIKey() string
	SetPublicAPIKey(string)
	GetPrivateAPIKey() string
	SetPrivateAPIKey(string)
	GetOpsManagerURL() string
	SetOpsManagerURL(string)
	IsOpsManager() bool
}

type VConfig struct {
	Profile string
}

func NewConfig() *VConfig {
	profile := viper.GetString("profile")

	return &VConfig{Profile: profile}
}

// GetService get configured service
func (c *VConfig) GetService() string {
	return viper.GetString(fmt.Sprintf("%s.service", c.Profile))
}

// SetService set configured service
func (c *VConfig) SetService(service string) {
	viper.Set(fmt.Sprintf("%s.service", c.Profile), service)
}

// IsOpsManager check if Ops Manager
func (c *VConfig) IsOpsManager() bool {
	return c.GetService() == store.OpsManagerService
}

// GetPublicAPIKey get configured public api key
func (c *VConfig) GetPublicAPIKey() string {
	return viper.GetString(fmt.Sprintf("%s.public_api_key", c.Profile))
}

// SetPublicAPIKey set configured publicAPIKey
func (c *VConfig) SetPublicAPIKey(publicAPIKey string) {
	viper.Set(fmt.Sprintf("%s.public_api_key", c.Profile), publicAPIKey)
}

// GetPrivateAPIKey get configured private api key
func (c *VConfig) GetPrivateAPIKey() string {
	return viper.GetString(fmt.Sprintf("%s.private_api_key", c.Profile))
}

// SetPrivateAPIKey set configured private api key
func (c *VConfig) SetPrivateAPIKey(privateAPIKey string) {
	viper.Set(fmt.Sprintf("%s.private_api_key", c.Profile), privateAPIKey)
}

// GetOpsManagerURL get configured ops manager base url
func (c *VConfig) GetOpsManagerURL() string {
	return viper.GetString(fmt.Sprintf("%s.ops_manager_url", c.Profile)) + "/api/public/v1.0/"
}

// SetOpsManagerURL set configured ops manager base url
func (c *VConfig) SetOpsManagerURL(opsManagerURL string) {
	viper.Set(fmt.Sprintf("%s.ops_manager_url", c.Profile), opsManagerURL)
}

// Save save the configuration to disk
func (c *VConfig) Save() error {
	return viper.WriteConfig()
}
