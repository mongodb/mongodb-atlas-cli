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
package namespaces

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const listTemplate = `NAMESPACE	TYPE{{range .Namespaces}}
{{.Namespace}}	{{.Type}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store       store.PerformanceAdvisorNamespacesLister
	processName string
	hostID      string
	since       int64
	duration    int64
}

func (opts *ListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	host, err := opts.host()
	if err != nil {
		return err
	}
	r, err := opts.store.PerformanceAdvisorNamespaces(opts.ConfigProjectID(), host, opts.newNamespaceOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newNamespaceOptions() *atlas.NamespaceOptions {
	return &atlas.NamespaceOptions{
		Since:    opts.since,
		Duration: opts.duration,
	}
}

func (opts *ListOpts) validateProcessName() error {
	const length = 2
	process := strings.Split(opts.processName, ":")
	if len(process) != length {
		return fmt.Errorf("'%v' is not valid", opts.processName)
	}
	return nil
}

func (opts *ListOpts) host() (string, error) {
	if opts.processName == "" {
		return opts.hostID, nil
	}
	err := opts.validateProcessName()
	if err != nil {
		return "", err
	}
	return opts.processName, nil
}

func (opts *ListOpts) markRequired(cmd *cobra.Command) func() error {
	return func() error {
		if config.Service() == config.CloudService {
			return cmd.MarkFlagRequired(flag.ProcessName)
		}
		return cmd.MarkFlagRequired(flag.HostID)
	}
}

// mongocli atlas performanceAdvisor namespace(s) list  --processName processName --since since --duration duration  --projectId projectId
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Short:   list,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.markRequired(cmd),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.hostID, flag.HostID, "", usage.HostID)
	cmd.Flags().StringVar(&opts.processName, flag.ProcessName, "", usage.ProcessName)
	cmd.Flags().Int64Var(&opts.since, flag.Since, 0, usage.Since)
	cmd.Flags().Int64Var(&opts.duration, flag.Duration, 0, usage.Duration)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	return cmd
}
