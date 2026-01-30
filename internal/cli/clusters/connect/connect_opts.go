// Copyright 2023 MongoDB Inc
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

// Package connect holds shared constants and validation for atlas clusters connect and atlas deployments connect.
package connect

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
)

const (
	ConnectWithConnectionString = "connectionString"
	ConnectWithMongosh          = "mongosh"
	ConnectWithCompass          = "compass"
	ConnectWithVsCode           = "vscode"
)

var ConnectWithOptions = []string{ConnectWithMongosh, ConnectWithCompass, ConnectWithVsCode, ConnectWithConnectionString}

const (
	PausedState  = "PAUSED"
	StoppedState = "STOPPED"
)

const MaxItemsPerPage = 500

const AtlasCluster = "atlas"

var ErrInvalidConnectWith = errors.New("invalid --connectWith option")

const (
	ConnectionStringTypeStandard = "standard"
	ConnectionStringTypePrivate  = "private"
)

var ConnectionStringTypeOptions = []string{ConnectionStringTypeStandard, ConnectionStringTypePrivate}

var ErrConnectionStringTypeNotImplemented = errors.New("connection string type not implemented")

const PromptConnectionStringType = "What type of connection string type would you like to use?"

// ValidateConnectWith returns an error if s is not one of ConnectWithOptions (case-insensitive).
func ValidateConnectWith(s string) error {
	if !search.StringInSliceFold(ConnectWithOptions, s) {
		return fmt.Errorf("%w: %s", ErrInvalidConnectWith, s)
	}
	return nil
}

// ValidateConnectionStringType returns an error if s is not one of ConnectionStringTypeOptions (case-insensitive).
func ValidateConnectionStringType(s string) error {
	if !search.StringInSliceFold(ConnectionStringTypeOptions, s) {
		return fmt.Errorf("%w: %s", ErrConnectionStringTypeNotImplemented, s)
	}
	return nil
}
