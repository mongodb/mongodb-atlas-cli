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
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.mongodb.org/atlas/auth"
)

const (
	MongoCLIEnvPrefix        = "MCLI"          // MongoCLIEnvPrefix prefix for MongoCLI ENV variables
	AtlasCLIEnvPrefix        = "MONGODB_ATLAS" // AtlasCLIEnvPrefix prefix for AtlasCLI ENV variables
	DefaultProfile           = "default"       // DefaultProfile default
	CloudService             = "cloud"         // CloudService setting when using Atlas API
	CloudGovService          = "cloudgov"      // CloudGovService setting when using Atlas API for Government
	projectID                = "project_id"
	orgID                    = "org_id"
	mongoShellPath           = "mongosh_path"
	configType               = "toml"
	service                  = "service"
	publicAPIKey             = "public_api_key"
	privateAPIKey            = "private_api_key"
	AccessTokenField         = "access_token"
	RefreshTokenField        = "refresh_token"
	ClientIDField            = "client_id"
	OpsManagerURLField       = "ops_manager_url"
	AccountURLField          = "account_url"
	baseURL                  = "base_url"
	apiVersion               = "api_version"
	output                   = "output"
	fileFlags                = os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	configPerm               = 0600
	defaultPermissions       = 0700
	skipUpdateCheck          = "skip_update_check"
	TelemetryEnabledProperty = "telemetry_enabled"
	AtlasCLI                 = "atlascli"
	ContainerizedHostNameEnv = "MONGODB_ATLAS_IS_CONTAINERIZED"
	GitHubActionsHostNameEnv = "GITHUB_ACTIONS"
	AtlasActionHostNameEnv   = "ATLAS_GITHUB_ACTION"
	CLIUserTypeEnv           = "CLI_USER_TYPE" // CLIUserTypeEnv is used to separate MongoDB University users from default users
	DefaultUser              = "default"       // Users that do NOT use ATLAS CLI with MongoDB University
	UniversityUser           = "university"    // Users that uses ATLAS CLI with MongoDB University
	NativeHostName           = "native"
	DockerContainerHostName  = "container"
	GitHubActionsHostName    = "all_github_actions"
	AtlasActionHostName      = "atlascli_github_action"
	LocalDeploymentImage     = "local_deployment_image" // LocalDeploymentImage is the config key for the MongoDB Local Dev Docker image
)

var (
	HostName       = getConfigHostnameFromEnvs()
	UserAgent      = fmt.Sprintf("%s/%s (%s;%s;%s)", AtlasCLI, version.Version, runtime.GOOS, runtime.GOARCH, HostName)
	CLIUserType    = newCLIUserTypeFromEnvs()
	defaultProfile = newProfile()
)

type Profile struct {
	name      string
	configDir string
	fs        afero.Fs
	err       error
}

func AllProperties() []string {
	return append(ProfileProperties(), GlobalProperties()...)
}

func BooleanProperties() []string {
	return []string{
		skipUpdateCheck,
		TelemetryEnabledProperty,
	}
}

func ProfileProperties() []string {
	return []string{
		AccessTokenField,
		apiVersion,
		baseURL,
		OpsManagerURLField,
		orgID,
		output,
		privateAPIKey,
		projectID,
		publicAPIKey,
		RefreshTokenField,
		service,
	}
}

func GlobalProperties() []string {
	return []string{
		LocalDeploymentImage,
		mongoShellPath,
		skipUpdateCheck,
		TelemetryEnabledProperty,
	}
}

func IsTrue(s string) bool {
	switch s {
	case "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES", "1":
		return true
	default:
		return false
	}
}

func Default() *Profile {
	return defaultProfile
}

// List returns the names of available profiles.
func List() []string {
	m := viper.AllSettings()

	keys := make([]string, 0, len(m))
	for k := range m {
		if !slices.Contains(AllProperties(), k) {
			keys = append(keys, k)
		}
	}
	// keys in maps are non-deterministic, trying to give users a consistent output
	sort.Strings(keys)
	return keys
}

// Exists returns true if there are any set settings for the profile name.
func Exists(name string) bool {
	return slices.Contains(List(), name)
}

// getConfigHostnameFromEnvs patches the agent hostname based on set env vars.
func getConfigHostnameFromEnvs() string {
	var builder strings.Builder

	envVars := []struct {
		envName  string
		hostName string
	}{
		{AtlasActionHostNameEnv, AtlasActionHostName},
		{GitHubActionsHostNameEnv, GitHubActionsHostName},
		{ContainerizedHostNameEnv, DockerContainerHostName},
	}

	for _, envVar := range envVars {
		if envIsTrue(envVar.envName) {
			appendToHostName(&builder, envVar.hostName)
		} else {
			appendToHostName(&builder, "-")
		}
	}
	configHostName := builder.String()

	if isDefaultHostName(configHostName) {
		return NativeHostName
	}
	return configHostName
}

// newCLIUserTypeFromEnvs patches the user type information based on set env vars.
func newCLIUserTypeFromEnvs() string {
	if value, ok := os.LookupEnv(CLIUserTypeEnv); ok {
		return value
	}

	return DefaultUser
}

func envIsTrue(env string) bool {
	return IsTrue(os.Getenv(env))
}

func appendToHostName(builder *strings.Builder, configVal string) {
	if builder.Len() > 0 {
		builder.WriteString("|")
	}
	builder.WriteString(configVal)
}

// isDefaultHostName checks if the hostname is the default placeholder.
func isDefaultHostName(hostname string) bool {
	// Using strings.Count for a more dynamic approach.
	return strings.Count(hostname, "-") == strings.Count(hostname, "|")+1
}

func newProfile() *Profile {
	configDir, err := CLIConfigHome()
	np := &Profile{
		name:      DefaultProfile,
		configDir: configDir,
		fs:        afero.NewOsFs(),
		err:       err,
	}
	return np
}

func Name() string { return Default().Name() }
func (p *Profile) Name() string {
	return p.name
}

var ErrProfileNameHasDots = errors.New("profile should not contain '.'")

func validateName(name string) error {
	if strings.Contains(name, ".") {
		return fmt.Errorf("%w: %q", ErrProfileNameHasDots, name)
	}

	return nil
}

func SetName(name string) error { return Default().SetName(name) }
func (p *Profile) SetName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	p.name = strings.ToLower(name)

	return nil
}

func Set(name string, value any) { Default().Set(name, value) }
func (p *Profile) Set(name string, value any) {
	settings := viper.GetStringMap(p.Name())
	settings[name] = value
	viper.Set(p.name, settings)
}

func SetGlobal(name string, value any) { viper.Set(name, value) }
func (*Profile) SetGlobal(name string, value any) {
	SetGlobal(name, value)
}

func Get(name string) any { return Default().Get(name) }
func (p *Profile) Get(name string) any {
	if viper.IsSet(name) && viper.Get(name) != "" {
		return viper.Get(name)
	}
	settings := viper.GetStringMap(p.Name())
	return settings[name]
}

func GetString(name string) string { return Default().GetString(name) }
func (p *Profile) GetString(name string) string {
	value := p.Get(name)
	if value == nil {
		return ""
	}
	return value.(string)
}

func GetBool(name string) bool { return Default().GetBool(name) }
func (p *Profile) GetBool(name string) bool {
	return p.GetBoolWithDefault(name, false)
}
func (p *Profile) GetBoolWithDefault(name string, defaultValue bool) bool {
	value := p.Get(name)
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return IsTrue(v)
	default:
		return defaultValue
	}
}

// Service get configured service.
func Service() string { return Default().Service() }
func (p *Profile) Service() string {
	if viper.IsSet(service) {
		return viper.GetString(service)
	}

	settings := viper.GetStringMapString(p.Name())
	return settings[service]
}

func IsCloud() bool {
	profile := Default()
	return profile.Service() == "" || profile.Service() == CloudService || profile.Service() == CloudGovService
}

// SetService set configured service.
func SetService(v string) { Default().SetService(v) }
func (p *Profile) SetService(v string) {
	p.Set(service, v)
}

// PublicAPIKey get configured public api key.
func PublicAPIKey() string { return Default().PublicAPIKey() }
func (p *Profile) PublicAPIKey() string {
	return p.GetString(publicAPIKey)
}

// SetPublicAPIKey set configured publicAPIKey.
func SetPublicAPIKey(v string) { Default().SetPublicAPIKey(v) }
func (p *Profile) SetPublicAPIKey(v string) {
	p.Set(publicAPIKey, v)
}

// PrivateAPIKey get configured private api key.
func PrivateAPIKey() string { return Default().PrivateAPIKey() }
func (p *Profile) PrivateAPIKey() string {
	return p.GetString(privateAPIKey)
}

// SetPrivateAPIKey set configured private api key.
func SetPrivateAPIKey(v string) { Default().SetPrivateAPIKey(v) }
func (p *Profile) SetPrivateAPIKey(v string) {
	p.Set(privateAPIKey, v)
}

// AccessToken get configured access token.
func AccessToken() string { return Default().AccessToken() }
func (p *Profile) AccessToken() string {
	return p.GetString(AccessTokenField)
}

// SetAccessToken set configured access token.
func SetAccessToken(v string) { Default().SetAccessToken(v) }
func (p *Profile) SetAccessToken(v string) {
	p.Set(AccessTokenField, v)
}

// RefreshToken get configured refresh token.
func RefreshToken() string { return Default().RefreshToken() }
func (p *Profile) RefreshToken() string {
	return p.GetString(RefreshTokenField)
}

// SetRefreshToken set configured refresh token.
func SetRefreshToken(v string) { Default().SetRefreshToken(v) }
func (p *Profile) SetRefreshToken(v string) {
	p.Set(RefreshTokenField, v)
}

type AuthMechanism int

const (
	APIKeys AuthMechanism = iota

	OAuth
	NotLoggedIn
)

// AuthType returns the type of authentication used in the profile.
func AuthType() AuthMechanism { return Default().AuthType() }
func (p *Profile) AuthType() AuthMechanism {
	if p.PublicAPIKey() != "" && p.PrivateAPIKey() != "" {
		return APIKeys
	}
	if p.AccessToken() != "" {
		return OAuth
	}
	return NotLoggedIn
}

// Token gets configured auth.Token.
func Token() (*auth.Token, error) { return Default().Token() }
func (p *Profile) Token() (*auth.Token, error) {
	if p.AccessToken() == "" || p.RefreshToken() == "" {
		return nil, nil
	}
	c, err := p.tokenClaims()
	if err != nil {
		return nil, err
	}
	var e time.Time
	if c.ExpiresAt != nil {
		e = c.ExpiresAt.Time
	}
	t := &auth.Token{
		AccessToken:  p.AccessToken(),
		RefreshToken: p.RefreshToken(),
		TokenType:    "Bearer",
		Expiry:       e,
	}
	return t, nil
}

// AccessTokenSubject will return the encoded subject in a JWT.
// This method won't verify the token signature, it's only safe to use to get the token claims.
func AccessTokenSubject() (string, error) { return Default().AccessTokenSubject() }
func (p *Profile) AccessTokenSubject() (string, error) {
	c, err := p.tokenClaims()
	if err != nil {
		return "", err
	}
	return c.Subject, err
}

func (p *Profile) tokenClaims() (jwt.RegisteredClaims, error) {
	c := jwt.RegisteredClaims{}
	// ParseUnverified is ok here, only want to make sure is a JWT and to get the claims for a Subject
	_, _, err := new(jwt.Parser).ParseUnverified(p.AccessToken(), &c)
	return c, err
}

// APIVersion get the default API version.
func APIVersion() string { return Default().APIVersion() }
func (p *Profile) APIVersion() string {
	return p.GetString(apiVersion)
}

// SetAPIVersion sets the default API version.
func SetAPIVersion(v string) { Default().SetAPIVersion(v) }
func (p *Profile) SetAPIVersion(v string) {
	p.Set(apiVersion, v)
}

// OpsManagerURL get configured ops manager base url.
func OpsManagerURL() string { return Default().OpsManagerURL() }
func (p *Profile) OpsManagerURL() string {
	return p.GetString(OpsManagerURLField)
}

// SetOpsManagerURL set configured ops manager base url.
func SetOpsManagerURL(v string) { Default().SetOpsManagerURL(v) }
func (p *Profile) SetOpsManagerURL(v string) {
	p.Set(OpsManagerURLField, v)
}

// AccountURL gets the configured account base url.
func AccountURL() string { return Default().AccountURL() }
func (p *Profile) AccountURL() string {
	return p.GetString(AccountURLField)
}

// ProjectID get configured project ID.
func ProjectID() string { return Default().ProjectID() }
func (p *Profile) ProjectID() string {
	return p.GetString(projectID)
}

// SetProjectID sets the global project ID.
func SetProjectID(v string) { Default().SetProjectID(v) }
func (p *Profile) SetProjectID(v string) {
	p.Set(projectID, v)
}

// OrgID get configured organization ID.
func OrgID() string { return Default().OrgID() }
func (p *Profile) OrgID() string {
	return p.GetString(orgID)
}

// SetOrgID sets the global organization ID.
func SetOrgID(v string) { Default().SetOrgID(v) }
func (p *Profile) SetOrgID(v string) {
	p.Set(orgID, v)
}

// SkipUpdateCheck get the global skip update check.
func SkipUpdateCheck() bool { return Default().SkipUpdateCheck() }
func (p *Profile) SkipUpdateCheck() bool {
	return p.GetBool(skipUpdateCheck)
}

// SetSkipUpdateCheck sets the global skip update check.
func SetSkipUpdateCheck(v bool) { Default().SetSkipUpdateCheck(v) }
func (*Profile) SetSkipUpdateCheck(v bool) {
	SetGlobal(skipUpdateCheck, v)
}

// IsTelemetryEnabledSet return true if telemetry_enabled has been set.
func IsTelemetryEnabledSet() bool { return Default().IsTelemetryEnabledSet() }
func (*Profile) IsTelemetryEnabledSet() bool {
	return viper.IsSet(TelemetryEnabledProperty)
}

// TelemetryEnabled get the configured telemetry enabled value.
func TelemetryEnabled() bool { return Default().TelemetryEnabled() }
func (p *Profile) TelemetryEnabled() bool {
	return isTelemetryFeatureAllowed() && p.GetBoolWithDefault(TelemetryEnabledProperty, true)
}

// SetTelemetryEnabled sets the telemetry enabled value.
func SetTelemetryEnabled(v bool) { Default().SetTelemetryEnabled(v) }

func (*Profile) SetTelemetryEnabled(v bool) {
	if !isTelemetryFeatureAllowed() {
		return
	}
	SetGlobal(TelemetryEnabledProperty, v)
}

func boolEnv(key string) bool {
	value, ok := os.LookupEnv(key)
	return ok && IsTrue(value)
}

func isTelemetryFeatureAllowed() bool {
	doNotTrack := boolEnv("DO_NOT_TRACK")
	return !doNotTrack
}

// Output get configured output format.
func Output() string { return Default().Output() }
func (p *Profile) Output() string {
	return p.GetString(output)
}

// SetOutput sets the global output format.
func SetOutput(v string) { Default().SetOutput(v) }
func (p *Profile) SetOutput(v string) {
	p.Set(output, v)
}

// ClientID get configured output format.
func ClientID() string { return Default().ClientID() }
func (p *Profile) ClientID() string {
	return p.GetString(ClientIDField)
}

// IsAccessSet return true if API keys have been set up.
// For Ops Manager we also check for the base URL.
func IsAccessSet() bool { return Default().IsAccessSet() }
func (p *Profile) IsAccessSet() bool {
	isSet := p.PublicAPIKey() != "" && p.PrivateAPIKey() != ""

	return isSet
}

// Map returns a map describing the configuration.
func Map() map[string]string { return Default().Map() }
func (p *Profile) Map() map[string]string {
	settings := viper.GetStringMapString(p.Name())
	profileSettings := make(map[string]string, len(settings)+1)
	for k, v := range settings {
		if k == privateAPIKey || k == AccessTokenField || k == RefreshTokenField {
			profileSettings[k] = "redacted"
		} else {
			profileSettings[k] = v
		}
	}

	return profileSettings
}

// SortedKeys returns the properties of the Profile sorted.
func SortedKeys() []string { return Default().SortedKeys() }
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
func Delete() error { return Default().Delete() }
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
	return filepath.Join(p.configDir, "config.toml")
}

func Filename() string {
	return Default().Filename()
}

// Rename replaces the Profile to a new Profile name, overwriting any Profile that existed before.
func Rename(newProfileName string) error { return Default().Rename(newProfileName) }
func (p *Profile) Rename(newProfileName string) error {
	if err := validateName(newProfileName); err != nil {
		return err
	}

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

func LoadAtlasCLIConfig() error { return Default().LoadAtlasCLIConfig(true) }
func (p *Profile) LoadAtlasCLIConfig(readEnvironmentVars bool) error {
	if p.err != nil {
		return p.err
	}

	viper.SetConfigName("config")

	if hasMongoCLIEnvVars() {
		viper.SetEnvKeyReplacer(strings.NewReplacer(AtlasCLIEnvPrefix, MongoCLIEnvPrefix))
	}

	return p.load(readEnvironmentVars, AtlasCLIEnvPrefix)
}

func hasMongoCLIEnvVars() bool {
	envVars := os.Environ()
	for _, v := range envVars {
		if strings.HasPrefix(v, MongoCLIEnvPrefix) {
			return true
		}
	}

	return false
}

func (p *Profile) load(readEnvironmentVars bool, envPrefix string) error {
	viper.SetConfigType(configType)
	viper.SetConfigPermissions(configPerm)
	viper.AddConfigPath(p.configDir)
	viper.SetFs(p.fs)

	if readEnvironmentVars {
		viper.SetEnvPrefix(envPrefix)
		viper.AutomaticEnv()
	}

	// aliases only work for a config file, this won't work for env variables
	viper.RegisterAlias(baseURL, OpsManagerURLField)

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
func Save() error { return Default().Save() }
func (p *Profile) Save() error {
	exists, err := afero.DirExists(p.fs, p.configDir)
	if err != nil {
		return err
	}
	if !exists {
		if err := p.fs.MkdirAll(p.configDir, defaultPermissions); err != nil {
			return err
		}
	}

	return viper.WriteConfigAs(p.Filename())
}

// CLIConfigHome retrieves configHome path.
func CLIConfigHome() (string, error) {
	home, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, "atlascli"), nil
}

func Path(f string) (string, error) {
	var p bytes.Buffer

	h, err := CLIConfigHome()
	if err != nil {
		return "", err
	}

	p.WriteString(h)
	p.WriteString(f)
	return p.String(), nil
}

// GetLocalDeploymentImage returns the configured MongoDB Docker image URL.
func GetLocalDeploymentImage() string { return Default().GetLocalDeploymentImage() }
func (p *Profile) GetLocalDeploymentImage() string {
	return p.GetString(LocalDeploymentImage)
}

// SetLocalDeploymentImage sets the MongoDB Docker image URL.
func SetLocalDeploymentImage(v string) { Default().SetLocalDeploymentImage(v) }
func (*Profile) SetLocalDeploymentImage(v string) {
	SetGlobal(LocalDeploymentImage, v)
}
