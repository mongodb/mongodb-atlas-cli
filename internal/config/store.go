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

package config

import (
	"slices"
	"sort"

	"github.com/spf13/viper"
)

type Store interface {
	Save() error

	GetProfileNames() []string
	RenameProfile(oldProfileName string, newProfileName string) error
	DeleteProfile(profileName string) error

	GetHierarchicalValue(profileName string, propertyName string) any

	SetProfileValue(profileName string, propertyName string, value any)
	GetProfileValue(profileName string, propertyName string) any
	GetProfileStringMap(profileName string) map[string]string

	SetGlobalValue(propertyName string, value any)
	GetGlobalValue(propertyName string) any
	IsSetGlobal(propertyName string) bool
}

// Temporary InMemoryStore to mimick legacy behavior
// Will be removed when we get rid of static references in the profile
type InMemoryStore struct {
	v *viper.Viper
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		v: viper.New(),
	}
}

func (*InMemoryStore) Save() error {
	return nil
}

func (s *InMemoryStore) GetProfileNames() []string {
	allKeys := s.v.AllKeys()

	profileNames := make([]string, 0, len(allKeys))
	for _, key := range allKeys {
		if !slices.Contains(GlobalProperties(), key) {
			profileNames = append(profileNames, key)
		}
	}
	// keys in maps are non-deterministic, trying to give users a consistent output
	sort.Strings(profileNames)
	return profileNames
}

func (*InMemoryStore) RenameProfile(_, _ string) error {
	panic("not implemented")
}

func (*InMemoryStore) DeleteProfile(_ string) error {
	panic("not implemented")
}

func (s *InMemoryStore) GetHierarchicalValue(profileName string, propertyName string) any {
	if s.v.IsSet(propertyName) && s.v.Get(propertyName) != "" {
		return s.v.Get(propertyName)
	}
	settings := s.v.GetStringMap(profileName)
	return settings[propertyName]
}

func (s *InMemoryStore) SetProfileValue(profileName string, propertyName string, value any) {
	settings := s.v.GetStringMap(profileName)
	settings[propertyName] = value
	s.v.Set(profileName, settings)
}

func (s *InMemoryStore) GetProfileValue(profileName string, propertyName string) any {
	settings := s.v.GetStringMap(profileName)
	return settings[propertyName]
}

func (s *InMemoryStore) GetProfileStringMap(profileName string) map[string]string {
	return s.v.GetStringMapString(profileName)
}

func (s *InMemoryStore) SetGlobalValue(propertyName string, value any) {
	s.v.Set(propertyName, value)
}

func (s *InMemoryStore) GetGlobalValue(propertyName string) any {
	return s.v.Get(propertyName)
}

func (s *InMemoryStore) IsSetGlobal(propertyName string) bool {
	return s.v.IsSet(propertyName)
}
