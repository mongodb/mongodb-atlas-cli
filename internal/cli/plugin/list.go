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
	"github.com/spf13/cobra"
)

type ListOps struct {
	cli.OutputOpts
	Opts
}

func (opts *ListOps) Run() error {
	return opts.Print(opts.plugins.GetValidAndInvalidPlugins())
}

const listTemplate = `NAME	DESCRIPTION	VERSION COMMANDS {{range valueOrEmptySlice .}}
{{.Name}}	{{.Description}}	{{.Version}}	{{- range $i, $cmd := .Commands}}{{if $i}}
			{{else}}	{{end}}{{$cmd.Name}}{{formatAliases $cmd.Aliases}}{{end}}
		{{end}}
`

func ListBuilder(pluginOpts *Opts) *cobra.Command {
	opts := &ListOps{}
	opts.Opts = *pluginOpts

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
