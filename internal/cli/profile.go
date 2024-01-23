package cli

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
)

func InitProfile(profile string) error {
	if profile != "" {
		return config.SetName(profile)
	} else if profile = config.GetString(flag.Profile); profile != "" {
		return config.SetName(profile)
	} else if availableProfiles := config.List(); len(availableProfiles) == 1 {
		return config.SetName(availableProfiles[0])
	}

	return nil
}
