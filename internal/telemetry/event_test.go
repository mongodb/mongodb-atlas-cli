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

package telemetry

import (
	"errors"
	"runtime"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestTelemetry_withCommandPath(t *testing.T) {
	config.ToolName = config.AtlasCLI
	testCmd := &cobra.Command{
		Use: "test",
	}
	rootCmd := &cobra.Command{
		Use: "root",
	}
	rootCmd.AddCommand(testCmd)

	e := newEvent(withCommandPath(testCmd))

	a := assert.New(t)
	a.Equal("root-test", e.Properties["command"])
}

func TestTelemetry_withProfileDefault(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withProfile())

	a := assert.New(t)
	a.Equal(config.DefaultProfile, e.Properties["profile"])
}

func TestTelemetry_withProfileCustom(t *testing.T) {
	config.ToolName = config.AtlasCLI

	const profile = "test"
	config.SetName(profile)

	e := newEvent(withProfile())

	a := assert.New(t)
	a.NotEqual(e.Properties["profile"], config.DefaultProfile)
	a.NotEqual(e.Properties["profile"], profile) // should be a base64
}

func TestTelemetry_withDuration(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withDuration(cmd))

	a := assert.New(t)
	a.GreaterOrEqual(e.Properties["duration"], int64(10))
}

func TestTelemetry_withFlags(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}
	_ = cmd.Flags().Bool("test", false, "")
	_ = cmd.Flags().Bool("test2", false, "")
	_ = cmd.ParseFlags([]string{"--test"})
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withFlags(cmd))

	a := assert.New(t)
	a.Equal(e.Properties["flags"], []string{"test"})
}

func TestTelemetry_withVersion(t *testing.T) {
	config.ToolName = config.AtlasCLI

	version.Version = "vTest"
	version.GitCommit = "sha-test"

	e := newEvent(withVersion())

	a := assert.New(t)
	a.Equal(e.Properties["version"], "vTest")
	a.Equal(e.Properties["git_commit"], "sha-test")
}

func TestTelemetry_withOS(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withOS())

	a := assert.New(t)
	a.Equal(e.Properties["os"], runtime.GOOS)
	a.Equal(e.Properties["arch"], runtime.GOARCH)
}

func TestTelemetry_withAuthMethod_apiKey(t *testing.T) {
	config.ToolName = config.AtlasCLI

	config.SetPublicAPIKey("test-public")
	config.SetPrivateAPIKey("test-private")

	e := newEvent(withAuthMethod())

	a := assert.New(t)
	a.Equal(e.Properties["auth_method"], "api_key")
}

func TestTelemetry_withAuthMethod_oauth(t *testing.T) {
	config.ToolName = config.AtlasCLI

	config.SetPublicAPIKey("")
	config.SetPrivateAPIKey("")

	e := newEvent(withAuthMethod())

	a := assert.New(t)
	a.Equal(e.Properties["auth_method"], "oauth")
}

func TestTelemetry_withService(t *testing.T) {
	config.ToolName = config.AtlasCLI

	const url = "http://host.test"
	config.SetService(config.CloudService)
	config.SetOpsManagerURL(url)

	e := newEvent(withService())

	a := assert.New(t)
	a.Equal(config.CloudService, e.Properties["service"])
	a.Equal(url, e.Properties["ops_manager_url"])
}

func TestTelemetry_withProjectID_Flag(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	const projectID = "test"
	var p string
	cmd.Flags().StringVarP(&p, flag.ProjectID, "", "", "")
	_ = cmd.ParseFlags([]string{"--" + flag.ProjectID, projectID})
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withProjectID(cmd))

	a := assert.New(t)
	a.Equal(projectID, e.Properties["project_id"])
}

func TestTelemetry_withProjectID_Config(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	const projectID = "test"
	config.SetProjectID(projectID)
	var p string
	cmd.Flags().StringVarP(&p, flag.ProjectID, "", "", "")
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withProjectID(cmd))

	a := assert.New(t)
	a.Equal(projectID, e.Properties["project_id"])
}

func TestTelemetry_withProjectID_NoFlagOrConfig(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	config.SetProjectID("")
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withProjectID(cmd))

	a := assert.New(t)
	_, ok := e.Properties["project_id"]
	a.Equal(false, ok)
}

func TestTelemetry_withOrgID_Flag(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	const orgID = "test"
	var p string
	cmd.Flags().StringVarP(&p, flag.OrgID, "", "", "")
	_ = cmd.ParseFlags([]string{"--" + flag.OrgID, orgID})
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withOrgID(cmd))

	a := assert.New(t)
	a.Equal(orgID, e.Properties["org_id"])
}

func TestTelemetry_withOrgID_Config(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	const orgID = "test"
	config.SetOrgID(orgID)
	var p string
	cmd.Flags().StringVarP(&p, flag.OrgID, "", "", "")
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withOrgID(cmd))

	a := assert.New(t)
	a.Equal(orgID, e.Properties["org_id"])
}

func TestTelemetry_withOrgID_NoFlagOrConfig(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}

	config.SetOrgID("")
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withOrgID(cmd))

	a := assert.New(t)
	_, ok := e.Properties["org_id"]
	a.Equal(false, ok)
}

func TestTelemetry_withError(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withError(errors.New("test")))

	a := assert.New(t)
	a.Equal("ERROR", e.Properties["result"])
	a.Equal("test", e.Properties["error"])
}

func TestTelemetry_withExtraProps(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withExtraProps(map[string]interface{}{"a": true}))

	a := assert.New(t)
	a.Equal(true, e.Properties["a"])
}
