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
	"github.com/mongodb/mongocli/internal/cli/backup"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atlas",
		Short: description.Atlas,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return validate.Credentials()
		},
	}
	// hide old backup until we have a better name for it
	b := backup.Builder()
	b.Hidden = true

	cmd.AddCommand(ClustersBuilder())
	cmd.AddCommand(DBUsersBuilder())
	cmd.AddCommand(WhitelistBuilder())
	cmd.AddCommand(alerts.Builder())
	cmd.AddCommand(b)
	cmd.AddCommand(events.Builder())
	cmd.AddCommand(MetricsBuilder())
	cmd.AddCommand(LogsBuilder())
	cmd.AddCommand(ProcessesBuilder())

	return cmd
}
