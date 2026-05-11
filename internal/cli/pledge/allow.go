// Copyright 2026 MongoDB Inc
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

package pledge

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/spf13/cobra"
)

// AllowBuilder returns the cobra command for `atlas pledge --allow <token>`.
// It is run in a second terminal to approve a blocked operation.
func AllowBuilder() *cobra.Command {
	return &cobra.Command{
		Use:   "allow <token>",
		Short: "Approve a blocked operation using an out-of-band token.",
		Long: `Approves a specific blocked operation for the session that requested it.

Run this command in a second terminal when atlas reports that an operation is
blocked by the active pledge. The token is printed by the blocked command.

The approver's own pledge must permit the requested operation tier. The grant
is valid for 5 minutes.`,
		Args:    cobra.ExactArgs(1),
		Example: "  atlas pledge allow a1b2c3d4e5f6...",
		RunE: func(cmd *cobra.Command, args []string) error {
			token := args[0]

			approverSID, err := pledge.Session()
			if err != nil {
				return fmt.Errorf("pledge is not supported on this platform: %w", err)
			}

			// Load approver's own pledge (nil = no pledge = can approve anything).
			approverPledge, err := pledge.Load(approverSID)
			if errors.Is(err, pledge.ErrNoPledge) {
				approverPledge = nil
			} else if err != nil {
				return err
			}

			if err := pledge.Approve(token, approverPledge, approverSID); err != nil {
				switch {
				case errors.Is(err, pledge.ErrApprovalExpired):
					return fmt.Errorf("token has expired (tokens are valid for 5 minutes)")
				case errors.Is(err, pledge.ErrApprovalConsumed):
					return fmt.Errorf("token has already been consumed")
				case errors.Is(err, pledge.ErrApprovalNotFound):
					return fmt.Errorf("token not found — check the token and ensure the requesting terminal is still waiting")
				case errors.Is(err, pledge.ErrApproverUnderPledged):
					return fmt.Errorf("your pledge does not permit the requested operation tier; use a less-restrictive pledge to approve")
				default:
					return err
				}
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Approved. The waiting command will now proceed.")
			return nil
		},
	}
}
