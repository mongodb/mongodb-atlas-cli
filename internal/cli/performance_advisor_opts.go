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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type PerformanceAdvisorOpts struct {
	ProcessName string
	HostID      string
}

// validateProcessName checks that the processName respects the format processName:port.
func (opts *PerformanceAdvisorOpts) validateProcessName() error {
	const length = 2
	process := strings.Split(opts.ProcessName, ":")
	if len(process) != length {
		return fmt.Errorf("'%v' is not valid", opts.ProcessName)
	}
	return nil
}

// Host returns the correct processName or the hostId in accordance with the service.
func (opts *PerformanceAdvisorOpts) Host() (string, error) {
	if opts.ProcessName == "" {
		return opts.HostID, nil
	}

	if err := opts.validateProcessName(); err != nil {
		return "", err
	}
	return opts.ProcessName, nil
}

func (opts *PerformanceAdvisorOpts) AddPerformanceAdvisorOptsFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opts.HostID, flag.HostID, "", usage.HostID)
	_ = cmd.Flags().MarkDeprecated(flag.HostID, "Flag is invalid for MongoDB Atlas")
	cmd.Flags().StringVar(&opts.ProcessName, flag.ProcessName, "", usage.ProcessNameAtlasCLI)
	_ = cmd.MarkFlagRequired(flag.ProcessName)
}
