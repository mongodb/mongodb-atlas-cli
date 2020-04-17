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
	"time"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertsAcknowledgeOpts struct {
	globalOpts
	alertID string
	until   string
	comment string
	store   store.AlertAcknowledger
}

func (opts *atlasAlertsAcknowledgeOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertsAcknowledgeOpts) Run() error {

	body := opts.newAcknowledgeRequest()
	result, err := opts.store.AcknowledgeAlert(opts.ProjectID(), opts.alertID, body)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasAlertsAcknowledgeOpts) newAcknowledgeRequest() *atlas.AcknowledgeRequest {

	until := opts.until

	// To acknowledge an alert “forever”, set the field value to 100 years in the future.
	if until == "" {
		until = time.Now().AddDate(100, 1, 1).Format(time.RFC3339)
	}

	return &atlas.AcknowledgeRequest{
		AcknowledgedUntil:      until,
		AcknowledgementComment: opts.comment,
	}

}

// mongocli atlas alerts acknowledge alertID --projectId projectId
func AtlasAlertsAcknowledgeBuilder() *cobra.Command {
	opts := new(atlasAlertsAcknowledgeOpts)
	cmd := &cobra.Command{
		Use:     "acknowledge [alertId]",
		Short:   description.AcknowledgeAlerts,
		Aliases: []string{"ack", "unacknowledge", "unack"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.until, flags.Until, "", usage.Until)
	cmd.Flags().StringVar(&opts.comment, flags.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
