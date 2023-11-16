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

package processes

import (
	"context"
	"sort"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

type AutoCompleteOpts struct {
	cli.GlobalOpts
	store store.ProcessLister
}

func (opts *AutoCompleteOpts) AutocompleteProcesses() cli.AutoFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		opts.parseFlags(cmd)
		if err := validate.Credentials(); err != nil {
			cobra.CompErrorln("no credentials")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.ValidateProjectID(); err != nil {
			cobra.CompErrorln("no project ID")
			return nil, cobra.ShellCompDirectiveError
		}
		if err := opts.initStore(cmd.Context()); err != nil {
			cobra.CompErrorln("store error: " + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}
		suggestions, err := opts.processSuggestion(toComplete)
		if err != nil {
			cobra.CompErrorln("error fetching: " + err.Error())
			return nil, cobra.ShellCompDirectiveError
		}
		return suggestions, cobra.ShellCompDirectiveDefault
	}
}

func (opts *AutoCompleteOpts) processSuggestion(toComplete string) ([]string, error) {
	processesList := &atlasv2.ListAtlasProcessesApiParams{
		GroupId: opts.ConfigProjectID(),
	}
	result, err := opts.store.Processes(processesList)
	if err != nil {
		return nil, err
	}
	suggestion := make([]string, 0, len(result.Results))
	for _, p := range result.Results {
		if !strings.HasPrefix(p.GetId(), toComplete) {
			continue
		}
		suggestion = append(suggestion, p.GetId())
	}
	sort.Strings(suggestion)
	return suggestion, nil
}

func (opts *AutoCompleteOpts) initStore(ctx context.Context) error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
	return err
}

func (opts *AutoCompleteOpts) parseFlags(cmd *cobra.Command) {
	profile := cmd.Flag(flag.Profile).Value.String()
	if profile != "" {
		config.SetName(profile)
	} else if profile = config.GetString(flag.Profile); profile != "" {
		config.SetName(profile)
	} else if availableProfiles := config.List(); len(availableProfiles) == 1 {
		config.SetName(availableProfiles[0])
	}
	if project := cmd.Flag(flag.ProjectID).Value.String(); project != "" {
		opts.ProjectID = project
	}
}
