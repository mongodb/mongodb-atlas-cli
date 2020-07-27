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
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/agents"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/automation"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/backup"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/clusters"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/dbusers"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/logs"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/metrics"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/owner"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/security"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/servers"
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
			// do not validate to create an owner
			if cmd.CommandPath() != "mongocli ops-manager owner create" {
				return validate.Credentials()
			}
			return nil
		},
	}

	cmd.AddCommand(clusters.Builder())
	cmd.AddCommand(alerts.Builder())
	cmd.AddCommand(backup.Builder())
	cmd.AddCommand(servers.Builder())
	cmd.AddCommand(automation.Builder())
	cmd.AddCommand(security.Builder())
	cmd.AddCommand(dbusers.Builder())
	cmd.AddCommand(owner.Builder())
	cmd.AddCommand(events.Builder())
	cmd.AddCommand(ProcessesBuilder())
	cmd.AddCommand(metrics.Builder())
	cmd.AddCommand(logs.Builder())
	cmd.AddCommand(agents.Builder())
	cmd.AddCommand(DiagnoseArchive())

	return cmd
}
