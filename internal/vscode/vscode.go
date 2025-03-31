// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vscode

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
)

var ErrVsCodeCliNotInstalled = errors.New("did not find vscode cli, install vscode and vscode cli: https://code.visualstudio.com/download")

func Detect() bool {
	return binPath() != ""
}

func binPath() string {
	p, err := exec.LookPath(vsCodeCliBin)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err == nil {
		return p
	}

	return ""
}

func SaveConnection(connStr, deploymentName, deploymentType string) error {
	params := url.Values{}

	connectionName := connectionNameBuilder(deploymentName, deploymentType)

	params.Add("connectionString", connStr)
	params.Add("name", connectionName)
	params.Add("reuseExisting", "true")
	if config.TelemetryEnabled() {
		// telemetry event will be seen within the extension's telemetry with "AtlasCLI" as the identifying name that the event origin is AtlasCLI
		params.Add("utm_source", "AtlasCLI")
	}
	doubleEncodedParams := url.PathEscape(params.Encode())
	deeplink := "vscode://mongodb.mongodb-vscode/connectWithURI?" + doubleEncodedParams

	return execCommand("--open-url", deeplink)
}

func connectionNameBuilder(deploymentName, deploymentType string) string {
	typeSuffix := "Atlas"
	if strings.EqualFold(deploymentType, "local") {
		typeSuffix = "Local"
	}
	return fmt.Sprintf("%s (%s)", deploymentName, typeSuffix)
}

func execCommand(args ...string) error {
	cmd := exec.Command(vsCodeCliBin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}
