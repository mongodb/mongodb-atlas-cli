// Copyright 2022 MongoDB Inc
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

package setup

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
)

func (opts *Opts) promptSettings() error {
	opts.settings = defaultSettings

	p := &survey.Select{
		Message: "How do you want to set up your Atlas cluster?",
		Options: settingOptions,
		Default: opts.settings,
		Description: func(value string, _ int) string {
			return settingsDescription[value]
		},
	}

	return telemetry.TrackAskOne(p, &opts.settings, nil)
}

const loadSampleDataMsg = `
Load sample data:			Yes`

func (opts *Opts) askConfirmDefaultQuestion(values *clusterSettings) error {
	if opts.Confirm {
		return nil
	}

	loadSampleData := ""
	if !opts.SkipSampleData {
		loadSampleData = loadSampleDataMsg
	}

	clusterTier := ""
	if opts.Tier != DefaultAtlasTier {
		diskSize := defaultDiskSizeGB(values.providerName(), opts.Tier)

		clusterTier = fmt.Sprintf(`
Cluster Tier:				%s
Cluster Disk Size (GiB):		%.1f`, opts.Tier, diskSize)
	}

	clusterVersion := ""
	if values.providerName() != tenant {
		version := opts.MDBVersion
		if version == "" {
			version = values.MdbVersion
		}
		clusterVersion = `
MongoDB Version:			` + version
	}

	fmt.Printf(`
[Default Settings]
Cluster Name:				%s%s
Cloud Provider and Region:		%s - %s%s
Database User Username:			%s%s
Allow connections from (IP Address):	%s

`,
		values.ClusterName,
		clusterTier,
		values.Provider,
		values.Region,
		clusterVersion,
		values.DBUsername,
		loadSampleData,
		strings.Join(values.IPAddresses, ", "),
	)

	return opts.promptSettings()
}
