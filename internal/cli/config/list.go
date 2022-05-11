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

package config

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/prompt"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var listTemplate = `PROFILE NAME{{range .}}
{{.}}{{end}}
`

type listOpts struct {
	cli.OutputOpts
}

func (opts *listOpts) Run() error {
	p := prompt.NewConfirm("Are you sure?")
	r := false

	err := telemetry.TrackAskOne(p, &r)
	if err != nil {
		return err
	}
	if !r {
		return errors.New("you did not confirm")
	}
	return opts.Print(config.List())
}

func ListBuilder() *cobra.Command {
	o := &listOpts{}
	o.Template = listTemplate
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Return a list of available profiles by name.",
		Long:    "If you did not specify a name for your profile, it displays as the default profile.",
		Example: fmt.Sprintf("  $ %s config ls", config.BinName()),
		PreRun: func(cmd *cobra.Command, args []string) {
			o.OutWriter = cmd.OutOrStdout()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	cmd.Flags().StringVarP(&o.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
