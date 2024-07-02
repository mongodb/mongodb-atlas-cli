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

package serverless

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

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	instanceName                      string
	enableServerlessContinuousBackup  bool
	disableServerlessContinuousBackup bool
	disableTerminationProtection      bool
	enableTerminationProtection       bool
	tag                               map[string]string
	store                             store.ServerlessInstanceUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Serverless instance {{.Name}} updated.\n"

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateServerlessInstance(opts.ConfigProjectID(), opts.instanceName, opts.newServerlessUpdateRequestParams())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *UpdateOpts) newServerlessUpdateRequestParams() *atlasv2.ServerlessInstanceDescriptionUpdate {
	params := &atlasv2.ServerlessInstanceDescriptionUpdate{
		TerminationProtectionEnabled: cli.ReturnValueForSetting(opts.enableTerminationProtection, opts.disableTerminationProtection),
	}

	serverlessContinuousBackupEnabled := cli.ReturnValueForSetting(opts.enableServerlessContinuousBackup, opts.disableServerlessContinuousBackup)
	if serverlessContinuousBackupEnabled != nil {
		params.ServerlessBackupOptions = &atlasv2.ClusterServerlessBackupOptions{
			ServerlessContinuousBackupEnabled: serverlessContinuousBackupEnabled,
		}
	}

	if len(opts.tag) > 0 {
		var tags []atlasv2.ResourceTag
		for k, v := range opts.tag {
			if k != "" && v != "" {
				tags = append(tags, atlasv2.ResourceTag{Key: k, Value: v})
			}
		}
		params.Tags = &tags
	}

	return params
}

// atlas serverless|sl update <instanceName> [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <instanceName>",
		Short: "Updates one serverless instance in the specified project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"instanceNameDesc": "Human-readable label that identifies your serverless instance.",
			"output":           updateTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.instanceName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.enableTerminationProtection, flag.EnableTerminationProtection, false, usage.EnableTerminationProtection)
	cmd.Flags().BoolVar(&opts.disableTerminationProtection, flag.DisableTerminationProtection, false, usage.DisableTerminationProtection)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableTerminationProtection, flag.DisableTerminationProtection)

	cmd.Flags().BoolVar(&opts.enableServerlessContinuousBackup, flag.EnableServerlessContinuousBackup, false, usage.EnableServerlessContinuousBackup)
	cmd.Flags().BoolVar(&opts.disableServerlessContinuousBackup, flag.DisableServerlessContinuousBackup, false, usage.DisableServerlessContinuousBackup)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableServerlessContinuousBackup, flag.DisableServerlessContinuousBackup)

	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.ServerlessTag+usage.UpdateWarning)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
