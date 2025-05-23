// Copyright 2022 MongoDB Inc
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

package advancedsettings

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

var describeTemplate = `MINIMUM TLS	JAVASCRIPT ENABLED	OPLOG SIZE MB
{{.MinimumEnabledTlsProtocol}}	{{.JavascriptEnabled}}	{{if .GetOplogSizeMB}}{{.GetOplogSizeMB}} {{else}}N/A{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=advancedsettings . AtlasClusterConfigurationOptionsDescriber

type AtlasClusterConfigurationOptionsDescriber interface {
	AtlasClusterConfigurationOptions(string, string) (*atlasClustersPinned.ClusterDescriptionProcessArgs, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name  string
	store AtlasClusterConfigurationOptionsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.AtlasClusterConfigurationOptions(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas cluster(s) advancedSettings describe <clusterName>  --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe <clusterName>",
		Aliases: []string{"get"},
		Short:   "Retrieve advanced configuration settings for one cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster for which you want to retrieve configuration settings.",
		},
		Example: `  atlas clusters advancedSettings describe Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
