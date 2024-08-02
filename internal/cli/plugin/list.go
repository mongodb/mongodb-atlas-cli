// Copyright 2024 MongoDB Inc
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

package plugin

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

type ListOps struct {
	cli.OutputOpts
	Opts
}

func (opts *ListOps) Run() error {
	return opts.Print(opts.plugins)
}

const listTemplate = `NAME	DESCRIPTION	VERSION {{range valueOrEmptySlice .}}
{{.Name}}	{{.Description}}	{{.Version}}	COMMANDS: {{range valueOrEmptySlice .Commands}}
		{{.Name}}{{end}}
		{{end}}
`

func ListBuilder(plugins []*plugin.Plugin) *cobra.Command {
	opts := &ListOps{
		Opts: Opts{
			plugins: plugins,
		},
	}
	const use = "list"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Returns a list of all installed plugins.",
		PreRun: func(cmd *cobra.Command, _ []string) {
			opts.OutWriter = cmd.OutOrStdout()
			opts.Template = listTemplate
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	return cmd
}
