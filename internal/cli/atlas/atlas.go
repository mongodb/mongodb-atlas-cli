// Copyright 2021 MongoDB Inc
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
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/alerts"
	"github.com/mongodb/mongocli/internal/cli/atlas/accesslists"
	"github.com/mongodb/mongocli/internal/cli/atlas/accesslogs"
	"github.com/mongodb/mongocli/internal/cli/atlas/backup"
	"github.com/mongodb/mongocli/internal/cli/atlas/cloudproviders"
	"github.com/mongodb/mongocli/internal/cli/atlas/clusters"
	"github.com/mongodb/mongocli/internal/cli/atlas/customdbroles"
	"github.com/mongodb/mongocli/internal/cli/atlas/customdns"
	"github.com/mongodb/mongocli/internal/cli/atlas/datalake"
	"github.com/mongodb/mongocli/internal/cli/atlas/dbusers"
	"github.com/mongodb/mongocli/internal/cli/atlas/integrations"
	"github.com/mongodb/mongocli/internal/cli/atlas/livemigrations"
	"github.com/mongodb/mongocli/internal/cli/atlas/logs"
	"github.com/mongodb/mongocli/internal/cli/atlas/maintenance"
	"github.com/mongodb/mongocli/internal/cli/atlas/metrics"
	"github.com/mongodb/mongocli/internal/cli/atlas/networking"
	"github.com/mongodb/mongocli/internal/cli/atlas/privateendpoints"
	"github.com/mongodb/mongocli/internal/cli/atlas/processes"
	"github.com/mongodb/mongocli/internal/cli/atlas/quickstart"
	"github.com/mongodb/mongocli/internal/cli/atlas/security"
	"github.com/mongodb/mongocli/internal/cli/atlas/serverless"
	"github.com/mongodb/mongocli/internal/cli/events"
	"github.com/mongodb/mongocli/internal/cli/performanceadvisor"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const (
	Use = "atlas"
)

func Builder() *cobra.Command {
	deprecatedMessage := deprecatedMessage()
	opts := &cli.RefresherOpts{}
	cmd := &cobra.Command{
		Use:   Use,
		Short: "MongoDB Atlas operations.",
		Long:  deprecatedMessage,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintf(os.Stderr, deprecatedMessage)
			if err := opts.InitFlow(); err != nil {
				return err
			}
			if err := opts.RefreshAccessToken(cmd.Context()); err != nil {
				return err
			}
			if config.Service() == "" {
				config.SetService(config.CloudService)
			}

			if cmd.Name() == "quickstart" { // quickstart has its own check
				return nil
			}

			return validate.Credentials()
		},
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		quickstart.Builder(),
		clusters.Builder(),
		dbusers.Builder(),
		customdbroles.Builder(),
		accesslists.Builder(),
		datalake.Builder(),
		alerts.Builder(),
		backup.Builder(),
		events.Builder(),
		metrics.Builder(),
		performanceadvisor.Builder(),
		logs.Builder(),
		processes.Builder(),
		privateendpoints.Builder(),
		networking.Builder(),
		security.Builder(),
		integrations.Builder(),
		maintenance.Builder(),
		customdns.Builder(),
		cloudproviders.Builder(),
		serverless.Builder(),
		livemigrations.Builder(),
		accesslogs.Builder(),
	)

	updateHelpString(cmd, deprecatedMessage)
	return cmd
}

// updateHelpString updates the help template to include the deprecated message.
func updateHelpString(cmd *cobra.Command, deprecatedMessage string) {
	for _, c := range cmd.Commands() {
		c.SetHelpTemplate(fmt.Sprintf("%s \n%s", deprecatedMessage, c.HelpTemplate()))
	}
}

func deprecatedMessage() string {
	if strings.Contains(runtime.GOOS, "darwin") {
		deprecatedMessage := `There’s a new, dedicated Atlas CLI available for Atlas users.  Install the Atlas CLI to enjoy the same capabilities and keep getting new features. Run brew install mongodb-atlas or visit https://dochub.mongodb.org/core/migrate-to-atlas-cli. Atlas commands for MongoCLI are now deprecated, but they will keep receiving support for 12 months (until April 30, 2023).

`

		return fmt.Sprintf("%s %s", "\u26A0", deprecatedMessage)
	}

	return "\u26A0 There’s a new, dedicated Atlas CLI available for Atlas users. Install the Atlas CLI to enjoy the same capabilities and keep getting new features: https://dochub.mongodb.org/core/migrate-to-atlas-cli. Atlas commands for MongoCLI are now deprecated, but they will keep receiving support for 12 months (until April 30, 2023).\n\n"
}
