// Copyright 2023 MongoDB Inc
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

package alerts

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type UnacknowledgeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	alertID string
	comment string
	store   store.AlertAcknowledger
}

func (opts *UnacknowledgeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var unackTemplate = "Alert '{{.Id}}' unacknowledged\n"

func (opts *UnacknowledgeOpts) Run() error {
	body := opts.newUnacknowledgeRequest()
	params := &atlasv2.AcknowledgeAlertApiParams{
		GroupId:          opts.ConfigProjectID(),
		AlertId:          opts.alertID,
		AcknowledgeAlert: body,
	}
	r, err := opts.store.AcknowledgeAlert(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UnacknowledgeOpts) newUnacknowledgeRequest() *atlasv2.AcknowledgeAlert {
	unacknowledge := true
	return &atlasv2.AcknowledgeAlert{
		AcknowledgedUntil:      nil,
		AcknowledgementComment: &opts.comment,
		UnacknowledgeAlert:     &unacknowledge,
	}
}

// atlas alerts unacknowledge <ID> --projectId projectId --comment comment.
func UnacknowledgeBuilder() *cobra.Command {
	opts := new(UnacknowledgeOpts)
	cmd := &cobra.Command{
		Use:     "unacknowledge <alertId>",
		Short:   "Unacknowledge the specified alert for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Aliases: []string{"unack"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"alertIdDesc": "Unique ID of the alert you want to unacknowledge.",
			"output":      unackTemplate,
		},
		Example: `  # Unacknowledge the alert with the ID 5d1113b25a115342acc2d1aa in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas alerts unacknowledge 5d1113b25a115342acc2d1aa --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), unackTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
