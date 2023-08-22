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

package encryptionatrest

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/spf13/cobra"
)

const (
	active = "ACTIVE"
)

type combinedStore interface {
	store.CompliancePolicyEncryptionAtRestUpdater
	store.CompliancePolicyDescriber
}

func baseCommand() *cobra.Command {
	const use = "encryptionatrest"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage Encryption at Rest of the backup compliance policy for your project. Encryption at rest enforces all clusters with a Backup Compliance Policy to use Customer Key Management.",
	}

	return cmd
}

func Builder() *cobra.Command {
	cmd := baseCommand()

	cmd.AddCommand(
		EnableBuilder(),
	)

	return cmd
}
