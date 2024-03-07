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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type AcknowledgeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	alertID string
	until   string
	comment string
	forever bool
	store   store.AlertAcknowledger
}

func (opts *AcknowledgeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var ackTemplate = "Alert '{{.ID}}' acknowledged until {{.AcknowledgedUntil}}\n"

func (opts *AcknowledgeOpts) Run() error {
	body := opts.newAcknowledgeRequest()
	r, err := opts.store.AcknowledgeAlert(opts.ConfigProjectID(), opts.alertID, body)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *AcknowledgeOpts) newAcknowledgeRequest() *atlas.AcknowledgeRequest {
	if opts.forever {
		// To acknowledge an alert “forever”, set the field value to 100 years in the future.
		const years = 100
		opts.until = time.Now().AddDate(years, 1, 1).Format(time.RFC3339)
	}

	return &atlas.AcknowledgeRequest{
		AcknowledgedUntil:      &opts.until,
		AcknowledgementComment: opts.comment,
	}
}

// mongocli atlas alerts acknowledge <ID> --projectId projectId --forever --comment comment --until until.
func AcknowledgeBuilder() *cobra.Command {
	opts := new(AcknowledgeOpts)
	opts.Template = ackTemplate
	cmd := &cobra.Command{
		Use:     "acknowledge <alertId>",
		Short:   "Acknowledges the specified alert for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Aliases: []string{"ack"},
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if opts.forever && opts.until != "" {
				return fmt.Errorf("--%s and --%s are exclusive", flag.Forever, flag.Until)
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ackTemplate),
			)
		},
		Annotations: map[string]string{
			"alertIdDesc": "ID of the alert you want to acknowledge or unacknowledge.",
		},
		Example: fmt.Sprintf(`  # Acknowledge an alert with the ID 5d1113b25a115342acc2d1aa in the project with the ID 5e2211c17a3e5a48f5497de3 until January 1 2028:
  %s alerts acknowledge 5d1113b25a115342acc2d1aa --until 2028-01-01T20:24:26Z --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}
	cmd.OutOrStdout()
	cmd.Flags().BoolVarP(&opts.forever, flag.Forever, flag.ForeverShort, false, usage.Forever)
	cmd.Flags().StringVar(&opts.until, flag.Until, "", usage.Until)
	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
