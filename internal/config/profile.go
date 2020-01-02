package config

import (
	"fmt"

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
	GetAPIPath() string
}

type Profile struct {
	Name string
}

var _ Config = &Profile{}

func New(name string) Config {
	return &Profile{Name: name}
}

// GetService get configured service
func (p *Profile) GetService() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, service))
}

// SetService set configured service
func (p *Profile) SetService(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, service), value)
}

// GetPublicAPIKey get configured public api key
func (p *Profile) GetPublicAPIKey() string {
	if viper.IsSet(publicAPIKey) {
		return viper.GetString(publicAPIKey)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, publicAPIKey))
}

// SetPublicAPIKey set configured publicAPIKey
func (p *Profile) SetPublicAPIKey(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, publicAPIKey), value)
}

// GetPrivateAPIKey get configured private api key
func (p *Profile) GetPrivateAPIKey() string {
	if viper.IsSet(privateAPIKey) {
		return viper.GetString(privateAPIKey)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, privateAPIKey))
}

// SetPrivateAPIKey set configured private api key
func (p *Profile) SetPrivateAPIKey(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, privateAPIKey), value)
}

// GetOpsManagerURL get configured ops manager base url
func (p *Profile) GetOpsManagerURL() string {
	if viper.IsSet(opsManagerURL) {
		return viper.GetString(opsManagerURL)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, opsManagerURL))
}

func (p *Profile) GetAPIPath() string {
	baseURL := p.GetOpsManagerURL()
	if baseURL != "" {
		if p.GetService() == CloudService {
			return baseURL + atlasAPIPath
		}
		return baseURL + publicAPIPath
	}
	return ""
}

// SetOpsManagerURL set configured ops manager base url
func (p *Profile) SetOpsManagerURL(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, opsManagerURL), value)
}

// Save save the configuration to disk
func Load() error {
	// Find home directory.
	configDir, err := configHome()
	if err != nil {
		return err
	}
	viper.SetConfigType(configType)
	viper.SetConfigName(Name)
	viper.AddConfigPath(configDir)

	viper.RegisterAlias("base_url", opsManagerURL)

	viper.SetEnvPrefix(Name)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return createConfigFile()
		}
		return err
	}
	return nil
}

// Save save the configuration to disk
func (p *Profile) Save() error {
	return viper.WriteConfig()
}
