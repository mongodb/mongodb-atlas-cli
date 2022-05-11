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

package config

import (
	"fmt"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
	exec "golang.org/x/sys/execabs"
)

type editOpts struct {
}

func (*editOpts) Run() error {
	editor := defaultEditor
	if v := os.Getenv("VISUAL"); v != "" {
		editor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		editor = e
	}
	cmd := exec.Command(editor, config.Filename()) //nolint:gosec // it's ok to let users do this
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func EditBuilder() *cobra.Command {
	opt := &editOpts{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Opens the the config with the default text editor.",
		Long:  `Will use the default editor to open the config file. You can use EDITOR or VISUAL envs to change the default.`,
		Example: fmt.Sprintf(`
  To open the config
  $ %s config edit
`, config.BinName()),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opt.Run()
		},
		Annotations: map[string]string{
			"toc": "true",
		},
		Args: require.NoArgs,
	}

	return cmd
}
