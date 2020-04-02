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

package cli

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

func OpsManagerBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ops-manager",
		Aliases: []string{"om"},
		Short:   description.OpsManager,
	}

	cmd.AddCommand(CloudManagerClustersBuilder())
	cmd.AddCommand(AtlasAlertsBuilder())
	cmd.AddCommand(AtlasBackupsBuilder())
	cmd.AddCommand(OpsManagerServersBuilder())
	cmd.AddCommand(OpsManagerAutomationBuilder())
	cmd.AddCommand(OpsManagerSecurityBuilder())
	cmd.AddCommand(OpsManagerDBUsersBuilder())
	cmd.AddCommand(OpsManagerOwnerBuilder())
	cmd.AddCommand(AtlasEventsBuilder())

	return cmd
}
