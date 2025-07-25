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

package config

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
)

var (
	CLIUserType = newCLIUserTypeFromEnvs()
	HostName    = getConfigHostnameFromEnvs()
	UserAgent   = fmt.Sprintf("%s/%s (%s;%s;%s)", AtlasCLI, version.Version, runtime.GOOS, runtime.GOARCH, HostName)
)

// newCLIUserTypeFromEnvs patches the user type information based on set env vars.
func newCLIUserTypeFromEnvs() string {
	if value, ok := os.LookupEnv(CLIUserTypeEnv); ok {
		return value
	}

	return DefaultUser
}

// getConfigHostnameFromEnvs patches the agent hostname based on set env vars.
func getConfigHostnameFromEnvs() string {
	var builder strings.Builder

	envVars := []struct {
		envName  string
		hostName string
	}{
		{AtlasActionHostNameEnv, AtlasActionHostName},
		{GitHubActionsHostNameEnv, GitHubActionsHostName},
		{ContainerizedHostNameEnv, DockerContainerHostName},
	}

	for _, envVar := range envVars {
		if envIsTrue(envVar.envName) {
			appendToHostName(&builder, envVar.hostName)
		} else {
			appendToHostName(&builder, "-")
		}
	}
	configHostName := builder.String()

	if isDefaultHostName(configHostName) {
		return NativeHostName
	}
	return configHostName
}

func appendToHostName(builder *strings.Builder, configVal string) {
	if builder.Len() > 0 {
		builder.WriteString("|")
	}
	builder.WriteString(configVal)
}

// isDefaultHostName checks if the hostname is the default placeholder.
func isDefaultHostName(hostname string) bool {
	// Using strings.Count for a more dynamic approach.
	return strings.Count(hostname, "-") == strings.Count(hostname, "|")+1
}

func envIsTrue(env string) bool {
	return IsTrue(os.Getenv(env))
}
