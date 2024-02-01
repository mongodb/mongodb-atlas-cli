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

package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/prerun"
	store "github.com/andreangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

const describeTemplate = `ID	FIRST NAME	LAST NAME	USERNAME	EMAIL
{{.Id}}	{{.FirstName}}	{{.LastName}}	{{.Username}}	{{.EmailAddress}}
`

type DescribeOpts struct {
	cli.OutputOpts
	store    store.UserDescriber
	username string
	id       string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	var r interface{}
	var err error

	if opts.username != "" {
		r, err = opts.store.UserByName(opts.username)
	}

	if opts.id != "" {
		r, err = opts.store.UserByID(opts.id)
	}

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) validate() error {
	if opts.id == "" && opts.username == "" {
		return errors.New("must supply one of 'id' or 'username'")
	}
	return nil
}

// DescribeBuilder atlas user(s) describe [--id id|--username USERNAME].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Annotations: map[string]string{"output": describeTemplate},
		Example: fmt.Sprintf(`  # Return the JSON-formatted details for the MongoDB user with the ID 5dd56c847a3e5a1f363d424d:
  %[1]s users describe --id 5dd56c847a3e5a1f363d424d --output json
  
  # Return the JSON-formatted details for the MongoDB user with the username myUser:
  %[1]s users describe --username myUser --output json`, cli.ExampleAtlasEntryPoint()),
		Short: "Return the details for the specified MongoDB user.",
		Long: `You can specify either the unique 24-digit ID that identifies the MongoDB user or the username for the MongoDB user.
		
User accounts and API keys with any role can run this command.`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return prerun.ExecuteE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flag.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.id, flag.ID, "", usage.UserID)
	cmd.MarkFlagsMutuallyExclusive(flag.Username, flag.ID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
