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

package cli

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/spf13/cobra"
)

type atlasAlertsConfigFieldsTypeOpts struct {
	store store.MatcherFieldsLister
}

func (opts *atlasAlertsConfigFieldsTypeOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasAlertsConfigFieldsTypeOpts) Run() error {
	result, err := opts.store.MatcherFields()

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas alerts config(s) fields type
func AtlasAlertsConfigsFieldsBuilder() *cobra.Command {
	opts := &atlasAlertsConfigFieldsTypeOpts{}
	cmd := &cobra.Command{
		Use:     "type",
		Short:   description.AlertsConfigFieldsType,
		Aliases: []string{"types"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	return cmd
}
