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
	"strings"

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerLogsCollectOpts struct {
	globalOpts
	resourceType              string
	resourceName              string
	logTypes                  []string
	sizeRequestedPerFileBytes int64
	redacted                  bool
	store                     store.LogCollector
}

func (opts *opsManagerLogsCollectOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerLogsCollectOpts) Run() error {
	result, err := opts.store.Collect(opts.ProjectID(), opts.newLog())
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

func (opts *opsManagerLogsCollectOpts) newLog() *om.LogCollectionJob {
	return &om.LogCollectionJob{
		ResourceType:              opts.resourceType,
		ResourceName:              opts.resourceName,
		Redacted:                  &opts.redacted,
		SizeRequestedPerFileBytes: opts.sizeRequestedPerFileBytes,
		LogTypes:                  opts.logTypes,
	}
}

// mongocli om logs collect resourceType resourceName --sizeRequestedPerFileBytes size --logTypes type --redacted redacted [--projectId projectId]
func OpsManagerLogsCollectOptsBuilder() *cobra.Command {
	opts := &opsManagerLogsCollectOpts{}
	cmd := &cobra.Command{
		Use:   "collect",
		Short: description.StartLogCollectionJob,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("accepts %d arg(s), received %d", 2, len(args))
			}

			args[0] = strings.ToLower(args[0])
			if !search.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid Resource Type. The Resource Type must be cluster, process or replicaset but was %q", args[0])
			}
			return nil
		},
		ValidArgs: []string{"cluster", "process", "replicaset"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.resourceType = args[0]
			opts.resourceName = args[1]
			return opts.Run()
		},
	}

	cmd.Flags().StringArrayVar(&opts.logTypes, flags.LogTypes, nil, usage.LogTypes)
	cmd.Flags().Int64Var(&opts.sizeRequestedPerFileBytes, flags.SizeRequestedPerFileBytes, 0, usage.SizeRequestedPerFileBytes)
	cmd.Flags().BoolVar(&opts.redacted, flags.LogRedacted, false, usage.LogRedacted)

	_ = cmd.MarkFlagRequired(flags.SizeRequestedPerFileBytes)
	_ = cmd.MarkFlagRequired(flags.LogTypes)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
