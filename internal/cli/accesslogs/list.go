// Copyright 2021 MongoDB Inc
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

package accesslogs

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	success      = "success"
	fail         = "fail"
	listTemplate = `HOSTNAME	AUTH RESULT	LOG LINE {{range valueOrEmptySlice .AccessLogs}}
{{if .Hostname}}{{.Hostname}} {{else}}N/A{{end}}{{.Hostname}}	{{.AuthResult}}	{{.LogLine}}{{end}}
`
	missingClusterNameHostnameErrorMessage = "one between --%s and --%s must be set"
	invalidValueAuthResultErrorMessage     = `--%s must be set to "%s" or "%s"`
)

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	hostname    string
	clusterName string
	start       string
	end         string
	nLogs       int
	ipAddresses string
	authResult  string
	store       store.AccessLogsLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	if opts.clusterName != "" {
		r, err := opts.store.AccessLogsByClusterName(opts.ConfigProjectID(), opts.clusterName, opts.newAccessLogOptions())
		if err != nil {
			return err
		}

		return opts.Print(r)
	}

	r, err := opts.store.AccessLogsByHostname(opts.ConfigProjectID(), opts.hostname, opts.newAccessLogOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newAccessLogOptions() *atlas.AccessLogOptions {
	var authResult *bool
	if opts.authResult != "" {
		isSuccess := strings.EqualFold(opts.authResult, success)
		authResult = &isSuccess
	}

	return &atlas.AccessLogOptions{
		Start:      opts.start,
		End:        opts.end,
		NLogs:      opts.nLogs,
		IPAddress:  opts.ipAddresses,
		AuthResult: authResult,
	}
}

func (opts *ListOpts) ValidateInput() error {
	if err := opts.ValidateProjectID(); err != nil {
		return err
	}

	if opts.clusterName == "" && opts.hostname == "" {
		return fmt.Errorf(missingClusterNameHostnameErrorMessage, flag.ClusterName, flag.Hostname)
	}

	if opts.authResult != "" && !strings.EqualFold(opts.authResult, success) && !strings.EqualFold(opts.authResult, fail) {
		return fmt.Errorf(invalidValueAuthResultErrorMessage, flag.AuthResult, success, fail)
	}

	return nil
}

// atlas accessLogs(s) list|ls  [--projectId projectId] [--clusterName clusterName] [--start start] [--end end] [--nLogs nLogs] [--ipAddress ipAddress] [--authResult success|fail].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Retrieve the access logs of a cluster identified by the cluster's name or hostname.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Monitoring Admin"),
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of all authentication requests made against the cluster named Cluster0 for the project with ID 618d48e05277a606ed2496fe:		
  atlas accesslogs list --output json --projectId 618d48e05277a606ed2496fe --clusterName Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateInput,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.hostname, flag.Hostname, "", usage.Hostname)
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.start, flag.Start, "", usage.AccessLogStartDate)
	cmd.Flags().StringVar(&opts.end, flag.End, "", usage.AccessLogEndDate)
	cmd.Flags().IntVar(&opts.nLogs, flag.NLog, 0, usage.NLog)
	cmd.Flags().StringVar(&opts.ipAddresses, flag.IP, "", usage.AccessLogIP)
	cmd.Flags().StringVar(&opts.authResult, flag.AuthResult, "", usage.AuthResult)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
