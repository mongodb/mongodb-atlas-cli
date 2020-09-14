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

package cloudmanager

import (
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/atlas/performanceadvisor"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/agents"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/automation"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/backup"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/clusters"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/dbusers"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/logs"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/maintenance"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/metrics"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/processes"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/security"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/servers"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const cloudManager = "MongoDB Cloud Manager operations."

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-manager",
		Aliases: []string{"cm"},
		Short:   cloudManager,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config.SetService(config.CloudManagerService)
			return validate.Credentials()
		},
	}

	cmd.AddCommand(
		clusters.Builder(),
		alerts.Builder(),
		backup.Builder(),
		servers.Builder(),
		automation.Builder(),
		security.Builder(),
		dbusers.Builder(),
		events.Builder(),
		processes.Builder(),
		metrics.Builder(),
		logs.Builder(),
		agents.Builder(),
		maintenance.Builder(),
		performanceadvisor.Builder(),
	)

	return cmd
}
