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

package atlas

import (
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/atlas/backup"
	"github.com/mongodb/mongocli/internal/cli/atlas/clusters"
	"github.com/mongodb/mongocli/internal/cli/atlas/datalake"
	"github.com/mongodb/mongocli/internal/cli/atlas/dbusers"
	"github.com/mongodb/mongocli/internal/cli/atlas/logs"
	"github.com/mongodb/mongocli/internal/cli/atlas/metrics"
	"github.com/mongodb/mongocli/internal/cli/atlas/networking"
	"github.com/mongodb/mongocli/internal/cli/atlas/privateendpoints"
	"github.com/mongodb/mongocli/internal/cli/atlas/processes"
	"github.com/mongodb/mongocli/internal/cli/atlas/security"
	"github.com/mongodb/mongocli/internal/cli/atlas/whitelist"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const atlasShort = "Atlas operations."

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atlas",
		Short: atlasShort,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config.SetService(config.CloudService)
			return validate.Credentials()
		},
	}
	cmd.AddCommand(datalake.Builder())
	cmd.AddCommand(clusters.Builder())
	cmd.AddCommand(dbusers.Builder())
	cmd.AddCommand(whitelist.Builder())
	cmd.AddCommand(alerts.Builder())
	cmd.AddCommand(backup.Builder())
	cmd.AddCommand(events.Builder())
	cmd.AddCommand(metrics.Builder())
	cmd.AddCommand(logs.Builder())
	cmd.AddCommand(processes.Builder())
	cmd.AddCommand(privateendpoints.Builder())
	cmd.AddCommand(networking.Builder())
	cmd.AddCommand(security.Builder())

	return cmd
}
