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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/alerts"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/events"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/mongocli/performanceadvisor"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/agents"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/automation"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/backup"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/clusters"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/featurepolicies"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/livemigrations"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/logs"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/maintenance"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/metrics"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/monitoring"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/processes"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/security"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/opsmanager/servers"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	opts := &cli.RefresherOpts{}
	var debugLevel bool
	cmd := &cobra.Command{
		Use:     "cloud-manager",
		Aliases: []string{"cm"},
		Short:   "MongoDB Cloud Manager operations.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetOutput(cmd.ErrOrStderr())
			if debugLevel {
				log.SetLevel(log.DebugLevel)
			}
			if err := opts.InitFlow(config.Default())(); err != nil {
				return err
			}
			if err := opts.RefreshAccessToken(cmd.Context()); err != nil {
				return err
			}
			config.SetService(config.CloudManagerService)
			return validate.Credentials()
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
		events.Builder(),
		monitoring.Builder(),
		processes.Builder(),
		metrics.Builder(),
		logs.Builder(),
		agents.Builder(),
		maintenance.Builder(),
		performanceadvisor.Builder(),
		featurepolicies.Builder(),
		livemigrations.Builder())

	cmd.PersistentFlags().BoolVarP(&debugLevel, flag.Debug, flag.DebugShort, false, usage.Debug)
	_ = cmd.PersistentFlags().MarkHidden(flag.Debug)

	return cmd
}
