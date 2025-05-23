// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slowoperationthreshold

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const EnableTemplate = `Atlas management of the slow operation enabled
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=enable_mock_test.go -package=slowoperationthreshold . PerformanceAdvisorSlowOperationThresholdEnabler

type PerformanceAdvisorSlowOperationThresholdEnabler interface {
	EnablePerformanceAdvisorSlowOperationThreshold(string) error
}

type EnableOpts struct {
	cli.ProjectOpts
	store PerformanceAdvisorSlowOperationThresholdEnabler
}

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *EnableOpts) Run() error {
	if err := opts.store.EnablePerformanceAdvisorSlowOperationThreshold(opts.ConfigProjectID()); err != nil {
		return err
	}

	fmt.Print(EnableTemplate)
	return nil
}

// atlas performanceAdvisor sot enable  [--projectId projectId].
func EnableBuilder() *cobra.Command {
	opts := new(EnableOpts)
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable the application-managed slow operation threshold for your project.",
		Long: `The slow operation threshold determines which operations are flagged by the Performance Advisor and Query Profiler. When enabled, the application uses the average execution time for operations on your cluster to determine slow-running queries. As a result, the threshold is more pertinent to your cluster workload. Application-managed slow operation threshold is enabled by default for dedicated clusters (M10+).

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": EnableTemplate,
		},
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)

	return cmd
}
