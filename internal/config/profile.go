// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

//go:generate mockgen -destination=../mocks/mock_profile.go -package=mocks github.com/mongodb/mongocli/internal/config Setter,Saver,SetSaver,Getter,Config

type profile struct {
	name      *string
	configDir string
	fs        afero.Fs
}

func Properties() []string {
	return []string{
		projectID,
		orgID,
		service,
		publicAPIKey,
		privateAPIKey,
		opsManagerURL,
		baseURL,
		opsManagerCACertificate,
		opsManagerSkipVerify,
	}
}

var p = newProfile()

type Setter interface {
	Set(string, string)
}

type Saver interface {
	Save() error
}

type SetSaver interface {
	Setter
	Saver
}

type Getter interface {
	GetString(string) string
}

type Config interface {
	Setter
	Getter
	Saver
	Service() string
	PublicAPIKey() string
	PrivateAPIKey() string
	OpsManagerURL() string
	OpsManagerCACertificate() string
	OpsManagerSkipVerify() string
}

func Default() Config {
	return p
}

func newProfile() *profile {
	configDir, err := configHome()
	if err != nil {
		log.Fatal(err)
	}
	name := "default"
	np := &profile{
		name:      &name,
		configDir: configDir,
		fs:        afero.NewOsFs(),
	}
	np.SetService(CloudService)
	return np
}

func Name() string { return p.Name() }
func (p *profile) Name() string {
	return *p.name
}

// Maybe a more intuitive way of doing this
func SetName(name *string) { p.SetName(name) }
func (p *profile) SetName(name *string) {
	p.name = name
}

func Set(name, value string) { p.Set(name, value) }
func (p *profile) Set(name, value string) {
	viper.Set(fmt.Sprintf("%s.%s", *p.name, name), value)
}

func GetString(name string) string { return p.GetString(name) }
func (p *profile) GetString(name string) string {
	if viper.IsSet(name) && viper.GetString(name) != "" {
		return viper.GetString(name)
	}
	if p.name != nil {
		return viper.GetString(fmt.Sprintf("%s.%s", *p.name, name))
	}
	return ""
}

// Service get configured service
func Service() string { return p.Service() }
func (p *profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	serviceKey := fmt.Sprintf("%s.%s", *p.name, service)
	if viper.IsSet(serviceKey) {
		return viper.GetString(serviceKey)
	}
	return CloudService
}

// SetService set configured service
func SetService(v string) { p.SetService(v) }
func (p *profile) SetService(v string) {
	p.Set(service, v)
}

// PublicAPIKey get configured public api key
func PublicAPIKey() string { return p.PublicAPIKey() }
func (p *profile) PublicAPIKey() string {
	return p.GetString(publicAPIKey)
}

// SetPublicAPIKey set configured publicAPIKey
func SetPublicAPIKey(v string) { p.SetPublicAPIKey(v) }
func (p *profile) SetPublicAPIKey(v string) {
	p.Set(publicAPIKey, v)
}

// PrivateAPIKey get configured private api key
func PrivateAPIKey() string { return p.PrivateAPIKey() }
func (p *profile) PrivateAPIKey() string {
	return p.GetString(privateAPIKey)
}

// SetPrivateAPIKey set configured private api key
func SetPrivateAPIKey(v string) { p.SetPrivateAPIKey(v) }
func (p *profile) SetPrivateAPIKey(v string) {
	p.Set(privateAPIKey, v)
}

// OpsManagerURL get configured ops manager base url
func OpsManagerURL() string { return p.OpsManagerURL() }
func (p *profile) OpsManagerURL() string {
	return p.GetString(opsManagerURL)
}

// SetOpsManagerURL set configured ops manager base url
func SetOpsManagerURL(v string) { p.SetOpsManagerURL(v) }
func (p *profile) SetOpsManagerURL(v string) {
	p.Set(opsManagerURL, v)
}

// OpsManagerCACertificate get configured ops manager CA certificate location
func OpsManagerCACertificate() string { return p.OpsManagerCACertificate() }
func (p *profile) OpsManagerCACertificate() string {
	return p.GetString(opsManagerCACertificate)
}

// SkipVerify get configured ops manager CA certificate location
func OpsManagerSkipVerify() string { return p.OpsManagerSkipVerify() }
func (p *profile) OpsManagerSkipVerify() string {
	return p.GetString(opsManagerSkipVerify)
}

// ProjectID get configured project ID
func ProjectID() string { return p.ProjectID() }
func (p *profile) ProjectID() string {
	return p.GetString(projectID)
}

// SetProjectID sets the global project ID
func SetProjectID(v string) { p.SetProjectID(v) }
func (p *profile) SetProjectID(v string) {
	p.Set(projectID, v)
}

// OrgID get configured organization ID
func OrgID() string { return p.OrgID() }
func (p *profile) OrgID() string {
	return p.GetString(orgID)
}

// SetOrgID sets the global organization ID
func SetOrgID(v string) { p.SetOrgID(v) }
func (p *profile) SetOrgID(v string) {
	p.Set(orgID, v)
}

// LoadRawConfig loads the configuration file, but without reading environment variables
func LoadRawConfig() error { return p.Load(false) }

// Load loads the configuration from disk
func Load() error { return p.Load(false) }
func (p *profile) Load(readEnvironmentVars bool) error {
	viper.SetConfigType(configType)
	viper.SetConfigName(ToolName)
	viper.SetConfigPermissions(0600)
	viper.AddConfigPath(p.configDir)

	if readEnvironmentVars {
		viper.SetEnvPrefix(EnvPrefix)
	}

	// TODO: review why this is not working as expected
	viper.RegisterAlias(baseURL, opsManagerURL)

	if readEnvironmentVars {
		viper.AutomaticEnv()
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// ignore if it doesn't exists
		var e viper.ConfigFileNotFoundError
		if errors.As(err, &e) {
			return nil
		}
		return err
	}
	return nil
}

// Save the configuration to disk
func Save() error { return p.Save() }
func (p *profile) Save() error {
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
	configFile := fmt.Sprintf("%s/%s.toml", p.configDir, ToolName)
	return viper.WriteConfigAs(configFile)
}
