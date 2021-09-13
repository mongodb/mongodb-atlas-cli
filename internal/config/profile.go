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
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mongodb/mongocli/internal/search"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

//go:generate mockgen -destination=../mocks/mock_profile.go -package=mocks github.com/mongodb/mongocli/internal/config SetSaver

const (
	ToolName                     = "mongocli"      // ToolName of the CLI
	EnvPrefix                    = "mcli"          // EnvPrefix prefix for ENV variables
	DefaultProfile               = "default"       // DefaultProfile default
	CloudService                 = "cloud"         // CloudService setting when using Atlas API
	CloudGovService              = "cloudgov"      // CloudGovService setting when using Atlas API for Government
	CloudManagerService          = "cloud-manager" // CloudManagerService settings when using CLoud Manager API
	OpsManagerService            = "ops-manager"   // OpsManagerService settings when using Ops Manager API
	JSON                         = "json"          // JSON output format as json
	projectID                    = "project_id"
	orgID                        = "org_id"
	mongoShellPath               = "mongosh_path"
	configType                   = "toml"
	service                      = "service"
	publicAPIKey                 = "public_api_key"
	privateAPIKey                = "private_api_key"
	opsManagerURL                = "ops_manager_url"
	baseURL                      = "base_url"
	opsManagerCACertificate      = "ops_manager_ca_certificate"
	opsManagerSkipVerify         = "ops_manager_skip_verify"
	opsManagerVersionManifestURL = "ops_manager_version_manifest_url"
	output                       = "output"
	fileFlags                    = os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	configPerm                   = 0600
	skipUpdateCheck              = "skip_update_check"
)

type Setter interface {
	Set(string, interface{})
}

type GlobalSetter interface {
	SetGlobal(string, interface{})
}

type Saver interface {
	Save() error
}

type SetSaver interface {
	Setter
	Saver
	GlobalSetter
}

type Profile struct {
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
		output,
		opsManagerURL,
		baseURL,
		opsManagerCACertificate,
		opsManagerSkipVerify,
		mongoShellPath,
		skipUpdateCheck,
	}
}

func BooleanProperties() []string {
	return []string{
		skipUpdateCheck,
	}
}

func GlobalProperties() []string {
	return []string{
		skipUpdateCheck,
	}
}

func IsTrue(s string) bool {
	return search.StringInSlice([]string{"true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES"}, s)
}

var p = newProfile()

func Default() *Profile {
	return p
}

// List returns the names of available profiles.
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
	return search.StringInSlice(List(), name)
}

func newProfile() *Profile {
	configDir, err := configHome()
	if err != nil {
		log.Fatal(err)
	}
	np := &Profile{
		name:      DefaultProfile,
		configDir: configDir,
		fs:        afero.NewOsFs(),
	}
	return np
}

func Name() string { return p.Name() }
func (p *Profile) Name() string {
	return p.name
}

func SetName(name string) { p.SetName(name) }
func (p *Profile) SetName(name string) {
	p.name = strings.ToLower(name)
}

func Set(name string, value interface{}) { p.Set(name, value) }
func (p *Profile) Set(name string, value interface{}) {
	settings := viper.GetStringMap(p.Name())
	settings[name] = value
	viper.Set(p.name, settings)
}

func SetGlobal(name string, value interface{}) { viper.Set(name, value) }
func (p *Profile) SetGlobal(name string, value interface{}) {
	SetGlobal(name, value)
}

func Get(name string) interface{} { return p.Get(name) }
func (p *Profile) Get(name string) interface{} {
	if viper.IsSet(name) && viper.Get(name) != "" {
		return viper.Get(name)
	}
	settings := viper.GetStringMap(p.Name())
	return settings[name]
}

func GetString(name string) string { return p.GetString(name) }
func (p *Profile) GetString(name string) string {
	value := p.Get(name)
	if value == nil {
		return ""
	}
	return value.(string)
}

func GetBool(name string) bool { return p.GetBool(name) }
func (p *Profile) GetBool(name string) bool {
	value := p.Get(name)
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return IsTrue(v)
	default:
		return false
	}
}

// Service get configured service.
func Service() string { return p.Service() }
func (p *Profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}
	settings := viper.GetStringMapString(p.Name())
	return settings[service]
}

func IsCloud() bool {
	return p.Service() == "" || p.Service() == CloudService || p.Service() == CloudGovService
}

// SetService set configured service.
func SetService(v string) { p.SetService(v) }
func (p *Profile) SetService(v string) {
	p.Set(service, v)
}

// PublicAPIKey get configured public api key.
func PublicAPIKey() string { return p.PublicAPIKey() }
func (p *Profile) PublicAPIKey() string {
	return p.GetString(publicAPIKey)
}

// SetPublicAPIKey set configured publicAPIKey.
func SetPublicAPIKey(v string) { p.SetPublicAPIKey(v) }
func (p *Profile) SetPublicAPIKey(v string) {
	p.Set(publicAPIKey, v)
}

// PrivateAPIKey get configured private api key.
func PrivateAPIKey() string { return p.PrivateAPIKey() }
func (p *Profile) PrivateAPIKey() string {
	return p.GetString(privateAPIKey)
}

// SetPrivateAPIKey set configured private api key.
func SetPrivateAPIKey(v string) { p.SetPrivateAPIKey(v) }
func (p *Profile) SetPrivateAPIKey(v string) {
	p.Set(privateAPIKey, v)
}

// OpsManagerURL get configured ops manager base url.
func OpsManagerURL() string { return p.OpsManagerURL() }
func (p *Profile) OpsManagerURL() string {
	return p.GetString(opsManagerURL)
}

// SetOpsManagerURL set configured ops manager base url.
func SetOpsManagerURL(v string) { p.SetOpsManagerURL(v) }
func (p *Profile) SetOpsManagerURL(v string) {
	p.Set(opsManagerURL, v)
}

// OpsManagerCACertificate get configured ops manager CA certificate location.
func OpsManagerCACertificate() string { return p.OpsManagerCACertificate() }
func (p *Profile) OpsManagerCACertificate() string {
	return p.GetString(opsManagerCACertificate)
}

// OpsManagerSkipVerify get configured if transport should skip CA verification.
func OpsManagerSkipVerify() string { return p.OpsManagerSkipVerify() }
func (p *Profile) OpsManagerSkipVerify() string {
	return p.GetString(opsManagerSkipVerify)
}

// OpsManagerVersionManifestURL get configured ops manager version manifest base url.
func OpsManagerVersionManifestURL() string { return p.OpsManagerVersionManifestURL() }
func (p *Profile) OpsManagerVersionManifestURL() string {
	return p.GetString(opsManagerVersionManifestURL)
}

// ProjectID get configured project ID.
func ProjectID() string { return p.ProjectID() }
func (p *Profile) ProjectID() string {
	return p.GetString(projectID)
}

// SetProjectID sets the global project ID.
func SetProjectID(v string) { p.SetProjectID(v) }
func (p *Profile) SetProjectID(v string) {
	p.Set(projectID, v)
}

// OrgID get configured organization ID.
func OrgID() string { return p.OrgID() }
func (p *Profile) OrgID() string {
	return p.GetString(orgID)
}

// SetOrgID sets the global organization ID.
func SetOrgID(v string) { p.SetOrgID(v) }
func (p *Profile) SetOrgID(v string) {
	p.Set(orgID, v)
}

// MongoShellPath get the configured MongoDB Shell path.
func MongoShellPath() string { return p.MongoShellPath() }
func (p *Profile) MongoShellPath() string {
	return p.GetString(mongoShellPath)
}

// SetMongoShellPath sets the global MongoDB Shell path.
func SetMongoShellPath(v string) { p.SetMongoShellPath(v) }
func (p *Profile) SetMongoShellPath(v string) {
	SetGlobal(mongoShellPath, v)
}

// SkipUpdateCheck get the global skip update check.
func SkipUpdateCheck() bool { return p.SkipUpdateCheck() }
func (p *Profile) SkipUpdateCheck() bool {
	return p.GetBool(skipUpdateCheck)
}

// SetSkipUpdateCheck sets the global skip update check.
func SetSkipUpdateCheck(v bool) { p.SetSkipUpdateCheck(v) }
func (p *Profile) SetSkipUpdateCheck(v bool) {
	SetGlobal(skipUpdateCheck, v)
}

// Output get configured output format.
func Output() string { return p.Output() }
func (p *Profile) Output() string {
	return p.GetString(output)
}

// SetOutput sets the global output format.
func SetOutput(v string) { p.SetOutput(v) }
func (p *Profile) SetOutput(v string) {
	p.Set(output, v)
}

// IsAccessSet return true if API keys have been set up.
// For Ops Manager we also check for the base URL.
func IsAccessSet() bool { return p.IsAccessSet() }
func (p *Profile) IsAccessSet() bool {
	isSet := p.PublicAPIKey() != "" && p.PrivateAPIKey() != ""
	if p.Service() == OpsManagerService {
		isSet = isSet && p.OpsManagerURL() != ""
	}

	return isSet
}

// Map returns a map describing the configuration.
func Map() map[string]string { return p.Map() }
func (p *Profile) Map() map[string]string {
	settings := viper.GetStringMapString(p.Name())
	profileSettings := make(map[string]string, len(settings)+1)
	if p.MongoShellPath() != "" {
		profileSettings[mongoShellPath] = p.MongoShellPath()
	}
	for k, v := range settings {
		if k == privateAPIKey || k == publicAPIKey {
			profileSettings[k] = "redacted"
		} else {
			profileSettings[k] = v
		}
	}

	return profileSettings
}

// SortedKeys returns the properties of the Profile sorted.
func SortedKeys() []string { return p.SortedKeys() }
func (p *Profile) SortedKeys() []string {
	config := p.Map()
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
func (p *Profile) Delete() error {
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

	f, err := p.fs.OpenFile(p.Filename(), fileFlags, configPerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(s)
	return err
}

func (p *Profile) Filename() string {
	return filepath.Join(p.configDir, ToolName+".toml")
}

// Rename replaces the Profile to a new Profile name, overwriting any Profile that existed before.
func Rename(newProfileName string) error { return p.Rename(newProfileName) }
func (p *Profile) Rename(newProfileName string) error {
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

	f, err := p.fs.OpenFile(p.Filename(), fileFlags, configPerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(s); err != nil {
		return err
	}

	return nil
}

// Load loads the configuration from disk.
func Load() error { return p.Load(true) }
func (p *Profile) Load(readEnvironmentVars bool) error {
	viper.SetConfigType(configType)
	viper.SetConfigName(ToolName)
	viper.SetConfigPermissions(configPerm)
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

// Save the configuration to disk.
func Save() error { return p.Save() }
func (p *Profile) Save() error {
	exists, err := afero.DirExists(p.fs, p.configDir)
	if err != nil {
		return err
	}
	if !exists {
		const defaultPermissions = 0700
		if err := p.fs.MkdirAll(p.configDir, defaultPermissions); err != nil {
			return err
		}
	}

	return viper.WriteConfigAs(p.Filename())
}

func configHome() (string, error) {
	if home := os.Getenv("XDG_CONFIG_HOME"); home != "" {
		return home, nil
	}
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config", home), nil
}
