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
	"github.com/andreangiolillo/mongocli-test/internal/cli/mongocli/alerts"
	"github.com/andreangiolillo/mongocli-test/internal/cli/mongocli/events"
	"github.com/andreangiolillo/mongocli-test/internal/cli/mongocli/performanceadvisor"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/admin"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/agents"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/automation"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/backup"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/clusters"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/dbusers"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/diagnosearchive"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/featurepolicies"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/livemigrations"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/logs"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/maintenance"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/metrics"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/monitoring"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/owner"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/processes"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/security"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/servers"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/serverusage"
	softwarecompotents "github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/softwarecomponents"
	"github.com/andreangiolillo/mongocli-test/internal/cli/opsmanager/versionmanifest"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/log"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/andreangiolillo/mongocli-test/internal/validate"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	var debugLevel bool
	cmd := &cobra.Command{
		Use:     "ops-manager",
		Aliases: []string{"om"},
		Short:   "MongoDB Ops Manager operations.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetWriter(cmd.ErrOrStderr())
			if debugLevel {
				log.SetLevel(log.DebugLevel)
			}

			config.SetService(config.OpsManagerService)
			// do not validate to create an owner
			if cmd.CommandPath() != "mongocli ops-manager owner create" {
				return validate.Credentials()
			}
			return nil
		},
		Annotations: map[string]string{
			"toc": "true",
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
		owner.Builder(),
		events.Builder(),
		monitoring.Builder(),
		processes.Builder(),
		metrics.Builder(),
		logs.Builder(),
		agents.Builder(),
		diagnosearchive.Builder(),
		maintenance.Builder(),
		performanceadvisor.Builder(),
		versionmanifest.Builder(),
		admin.Builder(),
		softwarecompotents.Builder(),
		featurepolicies.Builder(),
		serverusage.Builder(),
		livemigrations.Builder())

	cmd.PersistentFlags().BoolVarP(&debugLevel, flag.Debug, flag.DebugShort, false, usage.Debug)
	_ = cmd.PersistentFlags().MarkHidden(flag.Debug)

	return cmd
}
