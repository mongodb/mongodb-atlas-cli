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
	"fmt"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerAutomationWatchOpts struct {
	globalOpts
	store store.AutomationStatusGetter
}

func (opts *opsManagerAutomationWatchOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *opsManagerAutomationWatchOpts) Run() error {
	for {
		result, err := opts.store.GetAutomationStatus(opts.ProjectID())
		if err != nil {
			return err
		}
		reachedGoal := true
		for _, p := range result.Processes {
			if p.LastGoalVersionAchieved != result.GoalVersion {
				reachedGoal = false
				break
			}
		}
		if reachedGoal {
			break
		}
		fmt.Print(".")
		time.Sleep(4 * time.Second)
	}
	fmt.Printf("\nChanges deployed successfully\n")
	return nil
}

// mongocli ops-manager automation watch [--projectId projectId]
func OpsManagerAutomationWatchBuilder() *cobra.Command {
	opts := &opsManagerAutomationWatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch",
		Short: description.WatchAutomationStatus,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
