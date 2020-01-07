package config

import (
	"fmt"

	"github.com/spf13/afero"
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
	ProjectID() string
	SetProjectID(string)
	APIPath() string
	Save() error
}

type Profile struct {
	Name      string
	configDir string
	fs        afero.Fs
}

var _ Config = new(Profile)

func New(name string) (Config, error) {
	configDir, err := configHome()
	if err != nil {
		return nil, err
	}

	return &Profile{
		Name:      name,
		configDir: configDir,
		fs:        afero.NewOsFs(),
	}, nil
}

// Service get configured service
func (p *Profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	serviceKey := fmt.Sprintf("%s.%s", p.Name, service)
	if viper.IsSet(serviceKey) {
		return viper.GetString(serviceKey)
	}
	return CloudService
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

// ProjectID get configured project ID
func (p *Profile) ProjectID() string {
	if viper.IsSet(projectID) {
		return viper.GetString(projectID)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.Name, projectID))
}

func (p *Profile) SetProjectID(value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.Name, projectID), value)
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
	viper.SetConfigPermissions(0600)
	viper.AddConfigPath(configDir)

	viper.SetEnvPrefix(Name)
	// TODO: review why this is not working as expected
	viper.RegisterAlias(baseURL, opsManagerURL)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// ignore if it doesn't exists
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return err
	}
	return nil
}

// Save the configuration to disk
func (p *Profile) Save() error {
	exists, err := afero.DirExists(p.fs, p.configDir)
	if err != nil {
		return err
	}
	if !exists {
		err := p.fs.MkdirAll(p.configDir, 0700)
		if err != nil {
			return err
		}
	}
	// TODO: We can now read but not write, see https://github.com/spf13/viper/pull/813
	configFile := fmt.Sprintf("%s/%s.toml", p.configDir, Name)
	return viper.WriteConfigAs(configFile)
}
