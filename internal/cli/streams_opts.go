// Copyright 2025 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type StreamsOpts struct {
	Instance string
}

// ValidateInstance validates instance.
func (opts *StreamsOpts) ValidateInstance() error {
	if opts.Instance == "" {
		return errMissingInstance
	}
	return nil
}

func (opts *StreamsOpts) AddStreamsOptsFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.Instance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
}
