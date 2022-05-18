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

package logs

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/decryption"
	"github.com/spf13/cobra"
)

// MongoCLIBuilder is to split "mongocli atlas logs" and "atlas logs".
func MongoCLIBuilder() *cobra.Command {
	const use = "logs"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Download host logs for your project.",
	}
	cmd.AddCommand(DownloadBuilder())

	return cmd
}

// Builder is the up-to-date builder used by atlascli.
func Builder() *cobra.Command {
	const use = "logs"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Download host logs for your project.",
	}

	keyProvidersCmd := decryption.KeyProvidersBuilder()
	keyProvidersCmd.Hidden = true
	decryptCmd := DecryptBuilder()
	decryptCmd.Hidden = true

	cmd.AddCommand(
		DownloadBuilder(),
		keyProvidersCmd,
		decryptCmd,
	)

	return cmd
}
