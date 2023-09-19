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

//go:build unit

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
	"github.com/stretchr/testify/require"
)

func TestWithCommandPath(t *testing.T) {
	config.ToolName = config.AtlasCLI
	testCmd := &cobra.Command{
		Use: "test",
	}
	rootCmd := &cobra.Command{
		Use: "root",
	}
	rootCmd.AddCommand(testCmd)

	e := newEvent(withCommandPath(testCmd))
	assert.Equal(t, "root-test", e.Properties["command"])
}

func TestWithCommandPathAndAlias(t *testing.T) {
	config.ToolName = config.AtlasCLI
	rootCmd := &cobra.Command{
		Use: "root",
	}
	rootCmd.AddCommand(&cobra.Command{
		Use:     "test",
		Aliases: []string{"t"},
	})
	rootCmd.SetArgs([]string{"t"})
	calledCmd, _ := rootCmd.ExecuteContextC(NewContext())

	e := newEvent(withCommandPath(calledCmd))

	a := assert.New(t)
	a.Equal("root-test", e.Properties["command"])
	a.Equal("t", e.Properties["alias"])
}

func TestWithProfile(t *testing.T) {
	config.ToolName = config.AtlasCLI
	t.Run("default", func(t *testing.T) {
		e := newEvent(withProfile(&configMock{name: config.DefaultProfile}))
		assert.Equal(t, config.DefaultProfile, e.Properties["profile"])
	})
	t.Run("named", func(t *testing.T) {
		const profile = "test"

		e := newEvent(withProfile(&configMock{name: profile}))

		a := assert.New(t)
		a.NotEqual(e.Properties["profile"], config.DefaultProfile)
		a.NotEqual(e.Properties["profile"], profile) // should be a base64
	})
}

func TestWithDuration(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withDuration(cmd))
	assert.GreaterOrEqual(t, e.Properties["duration"], int64(10))
}

func TestWithFlags(t *testing.T) {
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
	assert.Equal(t, e.Properties["flags"], []string{"test"})
}

func TestWithVersion(t *testing.T) {
	config.ToolName = config.AtlasCLI

	version.Version = "vTest"
	version.GitCommit = "sha-test"

	e := newEvent(withVersion())

	a := assert.New(t)
	a.Equal(e.Properties["version"], "vTest")
	a.Equal(e.Properties["git_commit"], "sha-test")
}

func TestWithOS(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withOS())

	a := assert.New(t)
	a.Equal(e.Properties["os"], runtime.GOOS)
	a.Equal(e.Properties["arch"], runtime.GOARCH)
}

func TestWithUserAgent(t *testing.T) {
	e := newEvent(withUserAgent())

	a := assert.New(t)
	a.Equal(e.Properties["UserAgent"], config.UserAgent)
	a.Equal(e.Properties["HostName"], config.HostName)
}

func TestWithAuthMethod(t *testing.T) {
	config.ToolName = config.AtlasCLI
	t.Run("api key", func(t *testing.T) {
		c := &configMock{
			publicKey:  "test-public",
			privateKey: "test-private",
		}
		e := newEvent(withAuthMethod(c))
		assert.Equal(t, e.Properties["auth_method"], "api_key")
	})
	t.Run("Oauth", func(t *testing.T) {
		e := newEvent(withAuthMethod(&configMock{}))
		assert.Equal(t, e.Properties["auth_method"], "oauth")
	})
}

func TestWithService(t *testing.T) {
	config.ToolName = config.AtlasCLI
	const url = "http://host.test"
	c := &configMock{
		service: config.CloudService,
		url:     url,
	}
	e := newEvent(withService(c))

	a := assert.New(t)
	a.Equal(config.CloudService, e.Properties["service"])
	a.Equal(url, e.Properties["ops_manager_url"])
}

func TestWithProjectID(t *testing.T) {
	config.ToolName = config.AtlasCLI
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	var p string
	cmd.Flags().StringVarP(&p, flag.ProjectID, "", "", "")
	const projectID = "test"
	t.Run("From Flag", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.ProjectID, projectID))
		require.NoError(t, cmd.ExecuteContext(NewContext()))
		e := newEvent(withProjectID(cmd, &configMock{}))

		assert.Equal(t, projectID, e.Properties["project_id"])
	})
	t.Run("From Config", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.ProjectID, ""))
		require.NoError(t, cmd.ExecuteContext(NewContext()))
		c := &configMock{project: projectID}
		e := newEvent(withProjectID(cmd, c))
		assert.Equal(t, projectID, e.Properties["project_id"])
	})
	t.Run("no value", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.ProjectID, ""))
		e := newEvent(withProjectID(cmd, &configMock{}))
		require.NoError(t, cmd.ExecuteContext(NewContext()))
		_, ok := e.Properties["project_id"]
		assert.False(t, ok)
	})
}

func TestWithOrgID(t *testing.T) {
	config.ToolName = config.AtlasCLI

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	const orgID = "test"
	var p string
	cmd.Flags().StringVarP(&p, flag.OrgID, "", "", "")
	t.Run("From Flag", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.OrgID, orgID))
		require.NoError(t, cmd.ExecuteContext(NewContext()))

		e := newEvent(withOrgID(cmd, &configMock{}))
		assert.Equal(t, orgID, e.Properties["org_id"])
	})
	t.Run("From Config", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.OrgID, ""))
		require.NoError(t, cmd.ExecuteContext(NewContext()))
		c := &configMock{org: orgID}
		e := newEvent(withOrgID(cmd, c))
		assert.Equal(t, orgID, e.Properties["org_id"])
	})
	t.Run("no value", func(t *testing.T) {
		require.NoError(t, cmd.Flags().Set(flag.OrgID, ""))
		e := newEvent(withOrgID(cmd, &configMock{}))
		require.NoError(t, cmd.ExecuteContext(NewContext()))
		_, ok := e.Properties["org_id"]
		assert.False(t, ok)
	})
}

func TestWithError(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withError(errors.New("test")))

	a := assert.New(t)
	a.Equal("ERROR", e.Properties["result"])
	a.Equal("test", e.Properties["error"])
}

func TestSanitizePrompt(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Test",
			expected: "Test",
		},
		{
			input:    "Test [test1]",
			expected: "Test []",
		},
		{
			input:    "Test [test1] test2 [test3]",
			expected: "Test [] test2 []",
		},
	}

	for _, testCase := range testCases {
		got := sanitizePrompt(testCase.input)
		if got != testCase.expected {
			t.Errorf("expected: %v, got %v", testCase.expected, got)
		}
	}
}

func TestSanitizeSelectOption(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Test",
			expected: "Test",
		},
		{
			input:    "Test (test1)", // org id or projet id
			expected: "test1",
		},
	}

	for _, testCase := range testCases {
		got := sanitizeSelectOption(testCase.input)
		if got != testCase.expected {
			t.Errorf("expected: %v, got %v", testCase.expected, got)
		}
	}
}

func TestWithPrompt(t *testing.T) {
	config.ToolName = config.AtlasCLI

	q := "random question"
	k := "select"

	e := newEvent(withPrompt(q, k))

	a := assert.New(t)
	a.Equal(q, e.Properties["prompt"])
	a.Equal(k, e.Properties["prompt_type"])
}

func TestWithChoice(t *testing.T) {
	config.ToolName = config.AtlasCLI

	c := "test choice"

	e := newEvent(withChoice(c))
	assert.Equal(t, c, e.Properties["choice"])
}

func TestWithDefault(t *testing.T) {
	config.ToolName = config.AtlasCLI
	e := newEvent(withDefault(true))
	assert.Equal(t, true, e.Properties["default"])
}

func TestWithEmpty(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withEmpty(true))
	assert.Equal(t, true, e.Properties["empty"])
}

func TestWithAnonymousID(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(withAnonymousID())
	assert.NotEmpty(t, e.Properties["anonymous_id"])
}

func TestWithDeploymentType(t *testing.T) {
	config.ToolName = config.AtlasCLI

	e := newEvent(WithDeploymentType("test"))
	assert.Equal(t, e.Properties["deployment_type"], "test")
}

func TestWithSignal(t *testing.T) {
	config.ToolName = config.AtlasCLI

	q := "interrupt"
	e := newEvent(withSignal(q))
	assert.Equal(t, q, e.Properties["signal"])
}

func TestWithHelpCommand(t *testing.T) {
	config.ToolName = config.AtlasCLI
	testCmd := &cobra.Command{
		Use: "test",
	}
	rootCmd := &cobra.Command{
		Use: "root",
	}
	rootCmd.AddCommand(testCmd)
	rootCmd.InitDefaultHelpCmd()
	helpCmd := rootCmd.Commands()[0]

	args := []string{"test"}

	e := newEvent(withHelpCommand(helpCmd, args))

	assert.Equal(t, "root-test", e.Properties["help_command"])
}

func TestWithHelpCommand_NotFound(t *testing.T) {
	config.ToolName = config.AtlasCLI
	testCmd := &cobra.Command{
		Use: "test",
	}
	rootCmd := &cobra.Command{
		Use: "root",
	}
	rootCmd.AddCommand(testCmd)
	rootCmd.InitDefaultHelpCmd()
	helpCmd := rootCmd.Commands()[0]

	args := []string{"test2"}

	e := newEvent(withHelpCommand(helpCmd, args))

	_, ok := e.Properties["help_command"]
	assert.False(t, ok)
}

type configMock struct {
	name       string
	publicKey  string
	privateKey string
	service    string
	url        string
	project    string
	org        string
}

func (c configMock) Name() string {
	return c.name
}

func (c configMock) OrgID() string {
	return c.org
}

func (c configMock) ProjectID() string {
	return c.project
}

func (c configMock) Service() string {
	return c.service
}

func (c configMock) OpsManagerURL() string {
	return c.url
}

func (c configMock) PublicAPIKey() string {
	return c.publicKey
}

func (c configMock) PrivateAPIKey() string {
	return c.privateKey
}
