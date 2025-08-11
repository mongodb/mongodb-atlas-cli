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
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	AuthTypeField            = "auth_type"
	publicAPIKey             = "public_api_key"
	privateAPIKey            = "private_api_key"
	AccessTokenField         = "access_token"
	RefreshTokenField        = "refresh_token"
	ClientIDField            = "client_id"
	ClientSecretField        = "client_secret"
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
	Version                  = "version"                // versionField is the key for the configuration version
)

// Workaround to keep existing code working
// We cannot set the profile immediately because of a race condition which breaks all the unit tests
//
// The goal is to get rid of this, but we will need to do this gradually, since it's a large change that affects almost every command
func SetProfile(profile *Profile) {
	defaultProfile = profile
}

var (
	defaultProfile = &Profile{
		name:        DefaultProfile,
		configStore: NewInMemoryStore(),
	}
	profileContextKey = profileKey{}
)

type Profile struct {
	name        string
	configStore Store
}

func NewProfile(name string, configStore Store) *Profile {
	return &Profile{
		name:        name,
		configStore: configStore,
	}
}

type profileKey struct{}

// Setting a value
func WithProfile(ctx context.Context, profile *Profile) context.Context {
	return context.WithValue(ctx, profileContextKey, profile)
}

// Getting a value
func ProfileFromContext(ctx context.Context) (*Profile, bool) {
	if ctx == nil {
		return nil, false
	}

	profile, ok := ctx.Value(profileContextKey).(*Profile)
	return profile, ok
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
		Version,
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

func SetDefaultProfile(profile *Profile) {
	defaultProfile = profile
}

// List returns the names of available profiles.
func List() []string { return Default().List() }
func (p *Profile) List() []string {
	return p.configStore.GetProfileNames()
}

// Exists returns true if a profile with the give name exists.
func Exists(name string) bool {
	return slices.Contains(List(), name)
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
	p.configStore.SetProfileValue(p.Name(), name, value)
}

func SetGlobal(name string, value any) { Default().SetGlobal(name, value) }
func (p *Profile) SetGlobal(name string, value any) {
	p.configStore.SetGlobalValue(name, value)
}

func Get(name string) any { return Default().Get(name) }
func (p *Profile) Get(name string) any {
	return p.configStore.GetHierarchicalValue(p.Name(), name)
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

func GetInt64(name string) int64 { return Default().GetInt64(name) }
func (p *Profile) GetInt64(name string) int64 {
	value := p.Get(name)
	if value == nil {
		return 0
	}
	return value.(int64)
}

// Service get configured service.
func Service() string { return Default().Service() }
func (p *Profile) Service() string {
	if p.configStore.IsSetGlobal(service) {
		serviceValue, _ := p.configStore.GetGlobalValue(service).(string)
		return serviceValue
	}

	serviceValue, _ := p.configStore.GetProfileValue(p.Name(), service).(string)
	return serviceValue
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

type AuthMechanism string

const (
	APIKeys        AuthMechanism = "api_keys"
	UserAccount    AuthMechanism = "user_account"
	ServiceAccount AuthMechanism = "service_account"
	NoAuth         AuthMechanism = "no_auth"
)

// AuthType gets the configured auth type.
func AuthType() AuthMechanism { return Default().AuthType() }
func (p *Profile) AuthType() AuthMechanism {
	return AuthMechanism(p.GetString(AuthTypeField))
}

// SetAuthType sets the configured auth type.
func SetAuthType(v AuthMechanism) { Default().SetAuthType(v) }
func (p *Profile) SetAuthType(v AuthMechanism) {
	p.Set(AuthTypeField, string(v))
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

// ClientID get configured client ID.
func ClientID() string { return Default().ClientID() }
func (p *Profile) ClientID() string {
	return p.GetString(ClientIDField)
}

// SetClientID set configured client ID.
func SetClientID(v string) { Default().SetClientID(v) }
func (p *Profile) SetClientID(v string) {
	p.Set(ClientIDField, v)
}

// ClientSecret get configured client secret.
func ClientSecret() string { return Default().ClientSecret() }
func (p *Profile) ClientSecret() string {
	return p.GetString(ClientSecretField)
}

// SetClientSecret set configured client secret.
func SetClientSecret(v string) { Default().SetClientSecret(v) }
func (p *Profile) SetClientSecret(v string) {
	p.Set(ClientSecretField, v)
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
func (p *Profile) IsTelemetryEnabledSet() bool {
	return p.configStore.IsSetGlobal(TelemetryEnabledProperty)
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

// IsAccessSet return true if Service Account or API Keys credentials have been set up.
func IsAccessSet() bool { return Default().IsAccessSet() }
func (p *Profile) IsAccessSet() bool {
	isSet := p.PublicAPIKey() != "" && p.PrivateAPIKey() != "" ||
		p.ClientID() != "" && p.ClientSecret() != ""

	return isSet
}

// Map returns a map describing the configuration.
func Map() map[string]string { return Default().Map() }
func (p *Profile) Map() map[string]string {
	settings := p.configStore.GetProfileStringMap(p.Name())
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
	return p.configStore.DeleteProfile(p.Name())
}

// Rename replaces the Profile to a new Profile name, overwriting any Profile that existed before.
func Rename(newProfileName string) error { return Default().Rename(newProfileName) }
func (p *Profile) Rename(newProfileName string) error {
	if err := validateName(newProfileName); err != nil {
		return err
	}

	return p.configStore.RenameProfile(p.Name(), newProfileName)
}

// Save the configuration to disk.
func Save() error { return Default().Save() }
func (p *Profile) Save() error {
	return p.configStore.Save()
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
