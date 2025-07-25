// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate go tool go.uber.org/mock/mockgen -destination=./mocks.go -package=config github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config Store

package config

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// ViperConfigStore implements the config.Store interface
type ViperConfigStore struct {
	viper     viper.Viper
	configDir string
	fs        afero.Fs
}

// ViperConfigStore specific methods
func NewViperStore(fs afero.Fs) (*ViperConfigStore, error) {
	configDir, err := CLIConfigHome()
	if err != nil {
		return nil, err
	}

	v := viper.New()

	v.SetConfigName("config")

	if hasMongoCLIEnvVars() {
		v.SetEnvKeyReplacer(strings.NewReplacer(AtlasCLIEnvPrefix, MongoCLIEnvPrefix))
	}

	v.SetConfigType(configType)
	v.SetConfigPermissions(configPerm)
	v.AddConfigPath(configDir)
	v.SetFs(fs)

	v.SetEnvPrefix(AtlasCLIEnvPrefix)
	v.AutomaticEnv()

	// aliases only work for a config file, this won't work for env variables
	v.RegisterAlias(baseURL, OpsManagerURLField)

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err != nil {
		// ignore if it doesn't exists
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return nil, err
		}
	}
	return &ViperConfigStore{
		viper:     *v,
		configDir: configDir,
		fs:        fs,
	}, nil
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

func ViperConfigStoreFilename(configDir string) string {
	return filepath.Join(configDir, "config.toml")
}

func (s *ViperConfigStore) Filename() string {
	return ViperConfigStoreFilename(s.configDir)
}

// ConfigStore implementation

func (s *ViperConfigStore) Save() error {
	exists, err := afero.DirExists(s.fs, s.configDir)
	if err != nil {
		return err
	}
	if !exists {
		if err := s.fs.MkdirAll(s.configDir, defaultPermissions); err != nil {
			return err
		}
	}

	return s.viper.WriteConfigAs(s.Filename())
}

func (s *ViperConfigStore) GetProfileNames() []string {
	allKeys := s.viper.AllSettings()

	profileNames := make([]string, 0, len(allKeys))
	for key := range allKeys {
		if !slices.Contains(AllProperties(), key) {
			profileNames = append(profileNames, key)
		}
	}
	// keys in maps are non-deterministic, trying to give users a consistent output
	sort.Strings(profileNames)
	return profileNames
}

func (s *ViperConfigStore) RenameProfile(oldProfileName string, newProfileName string) error {
	if err := validateName(newProfileName); err != nil {
		return err
	}

	// Configuration needs to be deleted from toml, as viper doesn't support this yet.
	// FIXME :: change when https://github.com/spf13/viper/pull/519 is merged.
	configurationAfterDelete := s.viper.AllSettings()

	t, err := toml.TreeFromMap(configurationAfterDelete)
	if err != nil {
		return err
	}

	t.Set(newProfileName, t.Get(oldProfileName))

	err = t.Delete(oldProfileName)
	if err != nil {
		return err
	}

	tomlString := t.String()

	f, err := s.fs.OpenFile(s.Filename(), fileFlags, configPerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(tomlString); err != nil {
		return err
	}

	return nil
}

func (s *ViperConfigStore) DeleteProfile(profileName string) error {
	// Configuration needs to be deleted from toml, as viper doesn't support this yet.
	// FIXME :: change when https://github.com/spf13/viper/pull/519 is merged.
	settings := viper.AllSettings()

	t, err := toml.TreeFromMap(settings)
	if err != nil {
		return err
	}

	// Delete from the toml manually
	err = t.Delete(profileName)
	if err != nil {
		return err
	}

	tomlString := t.String()

	f, err := s.fs.OpenFile(s.Filename(), fileFlags, configPerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(tomlString)
	return err
}

func (s *ViperConfigStore) GetHierarchicalValue(profileName string, propertyName string) any {
	if s.viper.IsSet(propertyName) && s.viper.Get(propertyName) != "" {
		return s.viper.Get(propertyName)
	}
	settings := s.viper.GetStringMap(profileName)
	return settings[propertyName]
}

func (s *ViperConfigStore) SetProfileValue(profileName string, propertyName string, value any) {
	settings := s.viper.GetStringMap(profileName)
	settings[propertyName] = value
	s.viper.Set(profileName, settings)
}

func (s *ViperConfigStore) GetProfileValue(profileName string, propertyName string) any {
	settings := s.viper.GetStringMap(profileName)
	return settings[propertyName]
}

func (s *ViperConfigStore) GetProfileStringMap(profileName string) map[string]string {
	return s.viper.GetStringMapString(profileName)
}

func (s *ViperConfigStore) SetGlobalValue(propertyName string, value any) {
	s.viper.Set(propertyName, value)
}

func (s *ViperConfigStore) GetGlobalValue(propertyName string) any {
	return s.viper.Get(propertyName)
}

func (s *ViperConfigStore) IsSetGlobal(propertyName string) bool {
	return s.viper.IsSet(propertyName)
}
