// Copyright 2024 MongoDB Inc
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

package api

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
)

type ProfileFlagValueProvider struct {
	profile *config.Profile
}

func NewProfileFlagValueProvider(profile *config.Profile) *ProfileFlagValueProvider {
	return &ProfileFlagValueProvider{
		profile: profile,
	}
}

func NewProfileFlagValueProviderForDefaultProfile() *ProfileFlagValueProvider {
	return NewProfileFlagValueProvider(config.Default())
}

func (p *ProfileFlagValueProvider) ValueForFlag(flagName string) (*string, error) {
	switch flagName {
	case flag.GroupID:
		return noneEmptyStringPointer(p.profile.ProjectID()), nil
	case flag.OrgID:
		return noneEmptyStringPointer(p.profile.OrgID()), nil
	case flag.Version:
		return noneEmptyStringPointer(p.profile.APIVersion()), nil
	}

	return nil, nil
}

func noneEmptyStringPointer(value string) *string {
	if value != "" {
		return pointer.Get(value)
	}

	return nil
}
