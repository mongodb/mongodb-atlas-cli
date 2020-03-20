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
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertsDescribeOpts struct {
	*globalOpts
	alertID string
	store   store.AlertDescriber
}

func (opts *atlasAlertsDescribeOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertsDescribeOpts) Run() error {
	result, err := opts.store.Alert(opts.ProjectID(), opts.alertID)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas alerts describe alertID --projectId projectId
func AtlasAlertsDescribeBuilder() *cobra.Command {
	opts := &atlasAlertsDescribeOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "describe [name]",
		Short: "Describe an Atlas Alert.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
