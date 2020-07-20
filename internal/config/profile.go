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
	"sort"
	"strings"

	"github.com/mongodb/mongocli/internal/search"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

//go:generate mockgen -destination=../mocks/mock_profile.go -package=mocks github.com/mongodb/mongocli/internal/config Setter,Saver,SetSaver,Getter,Config

type profile struct {
	name      string
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

// List returns the names of available profiles
func List() []string {
	m := viper.AllSettings()

	keys := make([]string, 0, len(m))
	for k := range m {
		if !search.StringInSlice(Properties(), k) {
			keys = append(keys, k)
		}
	}
	// keys in maps are non deterministic, trying to give users a consistent output
	sort.Strings(keys)
	return keys
}

// Exists returns true if there are any set settings for the profile name.
func Exists(name string) bool {
	return len(viper.GetStringMap(name)) > 0
}

func newProfile() *profile {
	configDir, err := configHome()
	if err != nil {
		log.Fatal(err)
	}
	np := &profile{
		name:      DefaultProfile,
		configDir: configDir,
		fs:        afero.NewOsFs(),
	}
	return np
}

func Name() string { return p.Name() }
func (p *profile) Name() string {
	return p.name
}

func SetName(name string) { p.SetName(name) }
func (p *profile) SetName(name string) {
	p.name = strings.ToLower(name)
}

func Set(name, value string) { p.Set(name, value) }
func (p *profile) Set(name, value string) {
	viper.Set(fmt.Sprintf("%s.%s", p.name, name), value)
}

func GetString(name string) string { return p.GetString(name) }
func (p *profile) GetString(name string) string {
	if viper.IsSet(name) && viper.GetString(name) != "" {
		return viper.GetString(name)
	}
	return viper.GetString(fmt.Sprintf("%s.%s", p.name, name))
}

// Service get configured service
func Service() string { return p.Service() }
func (p *profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	serviceKey := fmt.Sprintf("%s.%s", p.name, service)
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

// IsAccessSet return true if API keys have been set up.
// For Ops Manager we also check for the base URL.
func IsAccessSet() bool { return p.IsAccessSet() }
func (p *profile) IsAccessSet() bool {
	isSet := p.PublicAPIKey() != "" && p.PrivateAPIKey() != ""
	if p.Service() == OpsManagerService {
		isSet = isSet && p.OpsManagerURL() != ""
	}

	return isSet
}

// Get returns a map describing the configuration.
func Get() map[string]string { return p.Get() }
func (p *profile) Get() map[string]string {
	settings := viper.GetStringMapString(p.Name())
	newSettings := make(map[string]string, len(settings))

	for k, v := range settings {
		if k == privateAPIKey || k == publicAPIKey {
			newSettings[k] = "redacted"
		} else {
			newSettings[k] = v
		}
	}

	return newSettings
}

// SortedKeys returns the properties of the profile sorted.
func SortedKeys() []string { return p.SortedKeys() }
func (p *profile) SortedKeys() []string {
	config := p.Get()
	keys := make([]string, 0, len(config))
	for k := range config {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Delete deletes an existing configuration. The profiles are reloaded afterwards, as
// this edits the file directly.
func Delete() error { return p.Delete() }
func (p *profile) Delete() error {
	// Configuration needs to be deleted from toml, as viper doesn't support this yet.
	// FIXME :: change when https://github.com/spf13/viper/pull/519 is merged.
	settings := viper.AllSettings()

	t, err := toml.TreeFromMap(settings)
	if err != nil {
		return err
	}

	// Delete from the toml manually
	err = t.Delete(p.Name())
	if err != nil {
		return err
	}

	s := t.String()

	f, err := p.fs.OpenFile(fmt.Sprintf("%s/%s.toml", p.configDir, ToolName), fileFlags, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(s); err != nil {
		return err
	}

	// Force reload, so that viper has the new configuration
	return p.Load(true)
}

// Rename replaces the profile to a new profile name, overwriting any profile that existed before.
func Rename(newProfileName string) error { return p.Rename(newProfileName) }
func (p *profile) Rename(newProfileName string) error {
	// Configuration needs to be deleted from toml, as viper doesn't support this yet.
	// FIXME :: change when https://github.com/spf13/viper/pull/519 is merged.
	configurationAfterDelete := viper.AllSettings()

	t, err := toml.TreeFromMap(configurationAfterDelete)
	if err != nil {
		return err
	}

	t.Set(newProfileName, t.Get(p.Name()))

	err = t.Delete(p.Name())
	if err != nil {
		return err
	}

	s := t.String()

	f, err := p.fs.OpenFile(fmt.Sprintf("%s/%s.toml", p.configDir, ToolName), fileFlags, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(s); err != nil {
		return err
	}

	// Force reload, so that viper has the new configuration
	return p.Load(true)
}

// Load loads the configuration from disk
func Load() error { return p.Load(true) }
func (p *profile) Load(readEnvironmentVars bool) error {
	viper.SetConfigType(configType)
	viper.SetConfigName(ToolName)
	viper.SetConfigPermissions(0600)
	viper.AddConfigPath(p.configDir)
	viper.SetFs(p.fs)

	if readEnvironmentVars {
		viper.SetEnvPrefix(EnvPrefix)
		viper.AutomaticEnv()
	}

	// aliases only work for a config file, this won't work for env variables
	viper.RegisterAlias(baseURL, opsManagerURL)

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
