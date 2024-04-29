// Copyright 2023 MongoDB Inc
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

package clusters

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

const storeErrMsg = "store error: "

type autoCompleteOpts struct {
	cli.GlobalOpts
	providers []string
	tier      string
	store     store.CloudProviderRegionsLister
}

func (opts *autoCompleteOpts) autocompleteTier() cli.AutoFunc {
	return func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if err := opts.parseFlags(cmd); err != nil {
			cobra.CompErrorln(fmt.Sprintf("failed to parse flags: %v", err))
			return nil, cobra.ShellCompDirectiveError
		}

		if err := validate.Credentials(); err != nil {
			cobra.CompErrorln("no credentials")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.ValidateProjectID(); err != nil {
			cobra.CompErrorln("no project ID")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.initStore(cmd.Context()); err != nil {
			cobra.CompErrorln(storeErrMsg + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}

		suggestions, err := opts.tierSuggestions(toComplete)
		if err != nil {
			cobra.CompErrorln("error fetching tier: " + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}
		return suggestions, cobra.ShellCompDirectiveDefault
	}
}

func (opts *autoCompleteOpts) tierSuggestions(toComplete string) ([]string, error) {
	result, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), "", opts.providers)
	if err != nil {
		return nil, err
	}
	availableTiers := map[string]bool{}
	for _, p := range result.GetResults() {
		for _, i := range p.GetInstanceSizes() {
			if _, ok := availableTiers[i.GetName()]; !ok && strings.HasPrefix(i.GetName(), strings.ToUpper(toComplete)) {
				availableTiers[i.GetName()] = true
			}
		}
	}
	suggestion := make([]string, len(availableTiers))
	i := 0
	for k := range availableTiers {
		suggestion[i] = k
		i++
	}
	sort.Strings(suggestion)
	return suggestion, nil
}

func (opts *autoCompleteOpts) autocompleteRegion() cli.AutoFunc {
	return func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if err := opts.parseFlags(cmd); err != nil {
			cobra.CompErrorln(fmt.Sprintf("failed to parse flags: %v", err))
			return nil, cobra.ShellCompDirectiveError
		}

		if err := validate.Credentials(); err != nil {
			cobra.CompErrorln("no credentials")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.ValidateProjectID(); err != nil {
			cobra.CompErrorln("no project ID")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.initStore(cmd.Context()); err != nil {
			cobra.CompErrorln(storeErrMsg + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}
		suggestions, err := opts.regionSuggestions(toComplete)
		if err != nil {
			cobra.CompErrorln("error fetching regions: " + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}
		return suggestions, cobra.ShellCompDirectiveDefault
	}
}

func (opts *autoCompleteOpts) regionSuggestions(toComplete string) ([]string, error) {
	result, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), opts.tier, opts.providers)
	if err != nil {
		return nil, err
	}
	availableRegions := map[string]bool{}
	for _, p := range result.GetResults() {
		for _, i := range p.GetInstanceSizes() {
			for _, r := range i.GetAvailableRegions() {
				if _, ok := availableRegions[r.GetName()]; !ok && strings.HasPrefix(r.GetName(), strings.ToUpper(toComplete)) {
					availableRegions[r.GetName()] = true
				}
			}
		}
	}
	suggestion := make([]string, len(availableRegions))
	i := 0
	for k := range availableRegions {
		suggestion[i] = k
		i++
	}
	sort.Strings(suggestion)
	return suggestion, nil
}

func (opts *autoCompleteOpts) initStore(ctx context.Context) error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
	return err
}

func (opts *autoCompleteOpts) parseFlags(cmd *cobra.Command) error {
	profile := cmd.Flag(flag.Profile).Value.String()

	if err := cli.InitProfile(profile); err != nil {
		return err
	}

	if project := cmd.Flag(flag.ProjectID).Value.String(); project != "" {
		opts.ProjectID = project
	}
	opts.providers = make([]string, 0, 1)
	if provider := cmd.Flag(flag.Provider).Value.String(); provider != "" {
		opts.providers = append(opts.providers, provider)
	}

	if tier := cmd.Flag(flag.Tier).Value.String(); tier != "" {
		opts.tier = tier
	}

	return nil
}
