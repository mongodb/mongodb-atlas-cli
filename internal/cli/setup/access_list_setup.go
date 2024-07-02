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

package setup

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func (opts *Opts) createAccessList() error {
	entries := opts.newProjectIPAccessList()
	if _, err := opts.store.CreateProjectIPAccessList(entries); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) askAccessListOptions() error {
	if !opts.shouldAskForValue(flag.AccessListIP) {
		return nil
	}
	message := ""

	if len(opts.IPAddresses) == 0 {
		publicIP := store.IPAddress()
		if publicIP != "" {
			message = fmt.Sprintf(" [Press Enter to use your public IP address '%s']", publicIP)
		}
		opts.IPAddresses = append(opts.IPAddresses, publicIP)
	}
	fmt.Print(`
[Set up your database network access details]
`)
	err := telemetry.TrackAskOne(
		newAccessListQuestion(strings.Join(opts.IPAddresses, ", "), message),
		&opts.IPAddressesResponse,
		survey.WithValidator(survey.Required),
	)

	if err == nil && opts.IPAddressesResponse != "" {
		ips := strings.Split(opts.IPAddressesResponse, ",")
		opts.IPAddresses = ips
	}
	return err
}

func (opts *Opts) newProjectIPAccessList() []*atlasv2.NetworkPermissionEntry {
	var accessListComment = "IP added with atlas quickstart"

	accessListArray := make([]*atlasv2.NetworkPermissionEntry, len(opts.IPAddresses))
	for i, addr := range opts.IPAddresses {
		accessList := &atlasv2.NetworkPermissionEntry{
			GroupId:   pointer.Get(opts.ConfigProjectID()),
			Comment:   &accessListComment,
			IpAddress: pointer.Get(addr),
		}

		accessListArray[i] = accessList
	}
	return accessListArray
}
