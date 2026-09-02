// Copyright 2026 MongoDB Inc
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

// Package useragent builds the Atlas CLI User-Agent, including the identifier of
// the AI agent invoking the CLI when one is detected.
package useragent

import (
	"github.com/mongodb/atlas-cli-core/config"
	"go.mongodb.org/atlas-sdk/v20250312024/detectaiagent"
)

// UserAgent returns the User-Agent to send to the Atlas API. When the CLI is
// invoked by a known AI agent, an ai-agent/<id> identifier is appended.
func UserAgent(version string) string {
	userAgent := config.UserAgent(version)

	if agent, ok := detectaiagent.Detect(); ok {
		return userAgent + " " + agent.UserAgentIdentifier
	}

	return userAgent
}

// AgentID returns the identifier of the detected AI agent, or an empty string
// when the CLI is not being invoked by a known agent.
func AgentID() string {
	if agent, ok := detectaiagent.Detect(); ok {
		return agent.ID
	}

	return ""
}
