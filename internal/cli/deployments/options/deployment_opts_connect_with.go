// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
)

const (
	ConnectWithConnectionString = "connectionString"
	ConnectWithMongosh          = "mongosh"
	ConnectWithCompass          = "compass"
)

var (
	ErrInvalidConnectWith  = errors.New("invalid --connectWith option")
	ConnectWithOptions     = []string{ConnectWithMongosh, ConnectWithCompass, ConnectWithConnectionString}
	connectWithDescription = map[string]string{
		ConnectWithConnectionString: "Connection String",
		ConnectWithMongosh:          "MongoDB Shell",
		ConnectWithCompass:          "MongoDB Compass",
	}
)

func ValidateConnectWith(s string) error {
	if !search.StringInSliceFold(ConnectWithOptions, s) {
		return fmt.Errorf("%w: %s", ErrInvalidConnectWith, s)
	}
	return nil
}

func (opts *DeploymentOpts) PromptConnectWith() (string, error) {
	p := &survey.Select{
		Message: fmt.Sprintf("How would you like to connect to %s?", opts.DeploymentName),
		Options: ConnectWithOptions,
		Description: func(value string, _ int) string {
			return connectWithDescription[value]
		},
	}

	var response string
	err := telemetry.TrackAskOne(p, &response, nil)
	return response, err
}
