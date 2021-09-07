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

package accesslists

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func Builder() *cobra.Command {
	const use = "accessLists"
	deprecated := append([]string{"whitelists"}, cli.GenerateAliases("whitelists")...)
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Manage the IP access list for your API Key.",
		Aliases: cli.GenerateAliases(use, deprecated...),
	}

	cmd.AddCommand(
		ListBuilder(),
		CreateBuilder(),
		DeleteBuilder(),
	)

	return cmd
}

// parseServiceVersion parses service version into semver.Version.
func parseServiceVersion(s store.ServiceVersionDescriber) (*semver.Version, error) {
	v, err := s.ServiceVersion()
	if err != nil {
		return nil, err
	}

	versionParts := strings.Split(v.Version, ".")

	const maxVersionParts = 2

	if len(versionParts) > maxVersionParts {
		versionParts = versionParts[0:maxVersionParts] // ops manager versions are not semantic this is converting it to x.y
	}

	newVersion := strings.Join(versionParts, ".")

	return semver.NewVersion(newVersion)
}

// shouldUseAccessList returns true when service is Cloud, CloudGov, Cloud Manager or Ops Manager (version 5+)
// and returns false when Ops Manager 4 or below.
func shouldUseAccessList(s store.ServiceVersionDescriber) (bool, error) {
	if config.Service() != config.OpsManagerService {
		return true, nil
	}

	sv, err := parseServiceVersion(s)
	if err != nil {
		return false, err
	}

	constrain, _ := semver.NewConstraint(">= 5.0")
	return constrain.Check(sv), nil
}

// fromWhitelistAPIKeysToAccessListAPIKeys convert from atlas.WhitelistAPIKeys format to atlas.AccessListAPIKeys
// We use this function with whitelist endpoints to keep supporting OM 4.2 and OM 4.4.
func fromWhitelistAPIKeysToAccessListAPIKeys(in *atlas.WhitelistAPIKeys) *atlas.AccessListAPIKeys {
	if in == nil {
		return nil
	}

	out := &atlas.AccessListAPIKeys{
		TotalCount: in.TotalCount,
		Links:      in.Links,
	}

	results := make([]*atlas.AccessListAPIKey, len(in.Results))
	for i, element := range in.Results {
		results[i] = fromWhitelistAPIKeyToAccessListAPIKey(element)
	}

	out.Results = results
	return out
}

// fromWhitelistAPIKeyToAccessListAPIKey convert from atlas.WhitelistAPIKey format to atlas.AccessListAPIKey
// We use this function with whitelist endpoints to keep supporting OM 4.2 and OM 4.4.
func fromWhitelistAPIKeyToAccessListAPIKey(in *atlas.WhitelistAPIKey) *atlas.AccessListAPIKey {
	return &atlas.AccessListAPIKey{
		CidrBlock:       in.CidrBlock,
		Count:           in.Count,
		Created:         in.Created,
		IPAddress:       in.IPAddress,
		LastUsed:        in.LastUsed,
		LastUsedAddress: in.LastUsedAddress,
		Links:           in.Links,
	}
}
