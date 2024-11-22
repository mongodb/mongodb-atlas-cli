package api

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
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
	case "groupId":
		return pointer.Get(p.profile.ProjectID()), nil
	case "orgId":
		return pointer.Get(p.profile.OrgID()), nil
	}

	return nil, nil
}
