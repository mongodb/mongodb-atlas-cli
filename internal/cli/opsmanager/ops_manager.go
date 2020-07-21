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

package opsmanager

import (
	"github.com/mongodb/mongocli/internal/cli/agents"
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/automation"
	"github.com/mongodb/mongocli/internal/cli/backup"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/clusters"
	"github.com/mongodb/mongocli/internal/cli/owner"
	"github.com/mongodb/mongocli/internal/cli/security"
	"github.com/mongodb/mongocli/internal/cli/servers"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ops-manager",
		Aliases: []string{"om"},
		Short:   description.OpsManager,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config.SetService(config.OpsManagerService)
			return validate.Credentials()
		},
	}

	cmd.AddCommand(clusters.Builder())
	cmd.AddCommand(alerts.Builder())
	cmd.AddCommand(backup.Builder())
	cmd.AddCommand(servers.Builder())
	cmd.AddCommand(automation.Builder())
	cmd.AddCommand(security.Builder())
	cmd.AddCommand(DBUsersBuilder())
	cmd.AddCommand(owner.Builder())
	cmd.AddCommand(events.Builder())
	cmd.AddCommand(ProcessesBuilder())
	cmd.AddCommand(MetricsBuilder())
	cmd.AddCommand(LogsBuilder())
	cmd.AddCommand(agents.Builder())
	cmd.AddCommand(DiagnoseArchive())

	return cmd
}
