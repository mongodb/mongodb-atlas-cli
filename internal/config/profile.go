package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config interface {
	Service() string
	SetService(string)
	PublicAPIKey() string
	SetPublicAPIKey(string)
	PrivateAPIKey() string
	SetPrivateAPIKey(string)
	OpsManagerURL() string
	SetOpsManagerURL(string)
	APIPath() string
}

type Profile struct {
	Name string
}

var _ Config = &Profile{}

func New(name string) Config {
	return &Profile{Name: name}
}

// Service get configured service
func (p *Profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, service))
}

// SetService set configured service
func (p *Profile) SetService(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, service), value)
}

// PublicAPIKey get configured public api key
func (p *Profile) PublicAPIKey() string {
	if viper.IsSet(publicAPIKey) {
		return viper.GetString(publicAPIKey)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, publicAPIKey))
}

// SetPublicAPIKey set configured publicAPIKey
func (p *Profile) SetPublicAPIKey(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, publicAPIKey), value)
}

// PrivateAPIKey get configured private api key
func (p *Profile) PrivateAPIKey() string {
	if viper.IsSet(privateAPIKey) {
		return viper.GetString(privateAPIKey)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, privateAPIKey))
}

// SetPrivateAPIKey set configured private api key
func (p *Profile) SetPrivateAPIKey(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, privateAPIKey), value)
}

// OpsManagerURL get configured ops manager base url
func (p *Profile) OpsManagerURL() string {
	if viper.IsSet(opsManagerURL) {
		return viper.GetString(opsManagerURL)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, opsManagerURL))
}

func (p *Profile) APIPath() string {
	baseURL := p.OpsManagerURL()
	if baseURL != "" {
		if p.Service() == CloudService {
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
