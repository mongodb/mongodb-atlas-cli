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

package policies

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=policies . Describer

type Describer interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts

	store Describer
}

var describePoliciesTemplate = `ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range valueOrEmptySlice .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	raw, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(raw)
}

func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Return the individual policy items of the backup compliance policy for your project.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describePoliciesTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
