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

package compliancepolicy

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/copyprotection"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/encryptionatrest"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/pointintimerestore"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/backup/compliancepolicy/policies"
	"github.com/spf13/cobra"
)

const (
	active = "ACTIVE"
)

func Builder() *cobra.Command {
	const use = "compliancePolicy"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   `Manage cloud backup compliance policy for your project. Use "atlas backups compliancepolicy setup" to enable backup compliance policy with a full configuration. Use "atlas backups compliancepolicy enable" to enable backup compliance policy without any configuration.`,
	}

	cmd.AddCommand(
		SetupBuilder(),
		EnableBuilder(),
		DescribeBuilder(),
		policies.Builder(),
		copyprotection.Builder(),
		pointintimerestore.Builder(),
		encryptionatrest.Builder(),
	)

	return cmd
}
