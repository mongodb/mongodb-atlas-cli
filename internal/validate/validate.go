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

package validate

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
)

const (
	minPasswordLength       = 10
	clusterWideScaling      = "clusterWideScaling"
	independentShardScaling = "independentShardScaling"
	clientIDLength          = 34
)

var (
	ErrAlreadyAuthenticatedAPIKeys = errors.New("already authenticated with an API key")
	ErrAlreadyAuthenticatedToken   = errors.New("already authenticated with an account")
	ErrInvalidPath                 = errors.New("invalid path")
	ErrInvalidClusterName          = errors.New("invalid cluster name")
	ErrInvalidDBUsername           = errors.New("invalid db username")
	ErrWeakPassword                = errors.New("the password provided is too common")
	ErrShortPassword               = errors.New("the password provided is too short")

	minimumPluginVersions = map[string]string{
		"atlas-cli-plugin-kubernetes": "v1.1.7",
		"atlas-cli-plugin-gsa":        "v0.0.2",
	}
)

// toString tries to cast an interface to string.
func toString(val any) (string, error) {
	var u string
	var ok bool
	if u, ok = val.(string); !ok {
		return "", fmt.Errorf("'%v' is not valid", val)
	}
	return u, nil
}

// URL validates a value is a valid URL for the cli store.
func URL(val any) error {
	s, err := toString(val)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(s, "/") {
		return fmt.Errorf("'%s' must have a trailing slash", s)
	}
	_, err = url.ParseRequestURI(s)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid URL", s)
	}

	return nil
}

// OptionalURL validates a value is a valid URL for the cli store.
func OptionalURL(val any) error {
	if val == nil {
		return nil
	}
	s, err := toString(val)
	if err != nil {
		return err
	}
	if s == "" {
		return nil
	}

	return URL(val)
}

// OptionalObjectID validates a value is a valid ObjectID.
func OptionalObjectID(val any) error {
	if val == nil {
		return nil
	}
	s, err := toString(val)
	if err != nil || s == "" {
		return err
	}
	return ObjectID(s)
}

// ObjectID validates a value is a valid ObjectID.
func ObjectID(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", s)
	}
	return nil
}

func ObjectIDByType(name, s string) error {
	if strings.HasPrefix(name, "client_") {
		if len(s) != clientIDLength {
			return fmt.Errorf("the provided value '%s' is not a valid client ID", s)
		}
		return nil
	}

	return ObjectID(s)
}

// Credentials validates public and private API keys have been set.
func Credentials() error {
	if t, err := config.Token(); t != nil {
		return err
	}
	if config.PrivateAPIKey() != "" && config.PublicAPIKey() != "" {
		return nil
	}

	return commonerrors.ErrUnauthorized
}

// NoAPIKeys there are no API keys in the profile, used for login/register/setup commands.
func NoAPIKeys() error {
	if config.PrivateAPIKey() == "" && config.PublicAPIKey() == "" {
		return nil
	}
	return fmt.Errorf(`%w (%s)

To authenticate using your Atlas username and password on a new profile, run: atlas auth login --profile <profile_name>`,
		ErrAlreadyAuthenticatedAPIKeys,
		config.PublicAPIKey(),
	)
}

func AutoScalingMode(autoScalingMode string) func() error {
	return func() error {
		if autoScalingMode == "" {
			return nil
		}

		if !strings.EqualFold(autoScalingMode, clusterWideScaling) && !strings.EqualFold(autoScalingMode, independentShardScaling) {
			return fmt.Errorf("invalid auto scaling mode: %s. Valid values are %s or %s", autoScalingMode, clusterWideScaling, independentShardScaling)
		}

		if strings.EqualFold(autoScalingMode, independentShardScaling) {
			fmt.Fprintf(os.Stderr, "'independentShardScaling' autoscaling cluster(s) detected, use --autoScalingMode independentShardScaling for clusters-related commands when interacting with this cluster\n")
		}
		return nil
	}
}

// NoAccessToken there is no access token in the profile, used for login/register/setup commands.
func NoAccessToken() error {
	if config.AccessToken() == "" {
		return nil
	}
	subject, _ := config.AccessTokenSubject()
	return fmt.Errorf(`%w (%s)

To log out, run: atlas auth logout`,
		ErrAlreadyAuthenticatedToken,
		subject,
	)
}

func FlagInSlice(value, flag string, validValues []string) error {
	if slices.Contains(validValues, value) {
		return nil
	}

	return fmt.Errorf(`invalid value for "%s", allowed values: "%s"`, flag, strings.Join(validValues, `", "`))
}

func ConditionalFlagNotInSlice(conditionalFlag string, conditionalFlagValue string, flag string, invalidFlags []string) error {
	if !slices.Contains(invalidFlags, flag) {
		return nil
	}

	return fmt.Errorf(`invalid flag "%s" in combination with "%s=%s", not allowed values: "%s"`, flag, conditionalFlag, conditionalFlagValue, strings.Join(invalidFlags, `", "`))
}

func Path(val any) error {
	path, ok := val.(string)
	if !ok {
		return fmt.Errorf("%w: %v", ErrInvalidPath, val)
	}

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidPath, path)
	}

	return nil
}

func OptionalPath(val any) error {
	if val == nil {
		return nil
	}
	s, err := toString(val)
	if err != nil || s == "" {
		return err
	}
	return Path(val)
}

func ClusterName(val any) error {
	name, ok := val.(string)
	if !ok {
		return fmt.Errorf("%w: %v", ErrInvalidClusterName, val)
	}
	match, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9-]*$", name)
	if match {
		return nil
	}

	return fmt.Errorf("%w. Cluster names can only contain ASCII letters, numbers, and hyphens: %s", ErrInvalidClusterName, name)
}

func DBUsername(val any) error {
	name, ok := val.(string)
	if !ok {
		return fmt.Errorf("%w: %v", ErrInvalidDBUsername, val)
	}
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+[a-zA-Z0-9-_]*$", name)
	if match {
		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidDBUsername, name)
}

func WeakPassword(val any) error {
	password, ok := val.(string)
	if !ok {
		return ErrWeakPassword
	}

	if len(password) < minPasswordLength {
		return fmt.Errorf("%w min: %d", ErrShortPassword, minPasswordLength)
	}

	if commonPasswords[strings.ToLower(password)] {
		return ErrWeakPassword
	}

	return nil
}

// PluginVersion validates the version of a plugin against the minimum required version.
// If a plugin is not listed in the minimumPluginVersions map, it is considered valid.
func PluginVersion(name string, version *semver.Version) error {
	minVersionStr, exists := minimumPluginVersions[name]
	if !exists {
		return nil // No version requirement for this plugin
	}

	minVersion, err := semver.NewVersion(minVersionStr)
	if err != nil {
		return err
	}

	if version.LessThan(minVersion) {
		return fmt.Errorf("plugin %s version v%s is below minimum required version %s for this version of AtlasCLI. Please update the plugin using 'atlas plugin update %s'",
			name, version.String(), minVersionStr, name)
	}

	return nil
}
