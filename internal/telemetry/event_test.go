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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithCommandPath(t *testing.T) {
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
	t.Run("default", func(t *testing.T) {
		e := newEvent(withProfile(&configMock{name: config.DefaultProfile}))
		assert.Equal(t, config.DefaultProfile, e.Properties["profile"])
	})
	t.Run("named", func(t *testing.T) {
		const profile = "test"

		e := newEvent(withProfile(&configMock{name: profile}))

		a := assert.New(t)
		a.NotEqual(config.DefaultProfile, e.Properties["profile"])
		a.NotEqual(profile, e.Properties["profile"]) // should be a base64
	})
}

func TestWithDuration(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withDuration(cmd))
	assert.GreaterOrEqual(t, e.Properties["duration"], int64(10))
}

func TestWithFlags(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
			time.Sleep(10 * time.Millisecond)
		},
	}
	_ = cmd.Flags().Bool("test", false, "")
	_ = cmd.Flags().Bool("test2", false, "")
	_ = cmd.ParseFlags([]string{"--test"})
	_ = cmd.ExecuteContext(NewContext())

	e := newEvent(withFlags(cmd))
	assert.Equal(t, []string{"test"}, e.Properties["flags"])
}

func TestWithVersion(t *testing.T) {
	version.Version = "vTest"
	version.GitCommit = "sha-test"

	e := newEvent(withVersion())

	a := assert.New(t)
	a.Equal("vTest", e.Properties["version"])
	a.Equal("sha-test", e.Properties["git_commit"])
}

func TestWithOS(t *testing.T) {
	e := newEvent(withOS())

	a := assert.New(t)
	a.Equal(runtime.GOOS, e.Properties["os"])
	a.Equal(runtime.GOARCH, e.Properties["arch"])
}

func TestWithUserAgent(t *testing.T) {
	e := newEvent(withUserAgent())

	a := assert.New(t)
	a.Equal(e.Properties["UserAgent"], config.UserAgent)
	a.Equal(e.Properties["HostName"], config.HostName)
}

func TestWithAuthMethod(t *testing.T) {
	t.Run("api key", func(t *testing.T) {
		c := &configMock{
			publicKey:  "test-public",
			privateKey: "test-private",
		}
		e := newEvent(withAuthMethod(c))
		assert.Equal(t, "api_key", e.Properties["auth_method"])
	})
	t.Run("Oauth", func(t *testing.T) {
		e := newEvent(withAuthMethod(&configMock{
			accessToken: "test",
		}))
		assert.Equal(t, "oauth", e.Properties["auth_method"])
	})
}

func TestWithService(t *testing.T) {
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

func TestWithCLIUserType(t *testing.T) {
	config.CLIUserType = config.DefaultUser
	a := assert.New(t)
	e := newEvent(withCLIUserType())
	a.Equal(config.DefaultUser, e.Properties["cli_user_type"])

	config.CLIUserType = config.UniversityUser
	e = newEvent(withCLIUserType())
	a.Equal(config.UniversityUser, e.Properties["cli_user_type"])
}

func TestWithProjectID(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
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
		assert.NotContains(t, e.Properties, "project_id")
	})
}

func TestWithOrgID(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
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
		assert.NotContains(t, e.Properties, "org_id")
	})
}

func TestWithError(t *testing.T) {
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
		assert.Equal(t, testCase.expected, got)
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
		assert.Equal(t, testCase.expected, got)
	}
}

func TestWithPrompt(t *testing.T) {
	q := "random question"
	k := "select"

	e := newEvent(withPrompt(q, k))

	a := assert.New(t)
	a.Equal(q, e.Properties["prompt"])
	a.Equal(k, e.Properties["prompt_type"])
}

func TestWithChoice(t *testing.T) {
	c := "test choice"

	e := newEvent(withChoice(c))
	assert.Equal(t, c, e.Properties["choice"])
}

func TestWithDefault(t *testing.T) {
	e := newEvent(withDefault(true))
	assert.Contains(t, e.Properties, "default")
}

func TestWithEmpty(t *testing.T) {
	e := newEvent(withEmpty(true))
	assert.Contains(t, e.Properties, "empty")
}

func TestWithAnonymousID(t *testing.T) {
	e := newEvent(withAnonymousID())
	assert.Contains(t, e.Properties, "device_id")
}

func TestWithDeploymentType(t *testing.T) {
	e := newEvent(WithDeploymentType("test"))
	assert.Equal(t, "test", e.Properties["deployment_type"])
}

func TestWithSignal(t *testing.T) {
	q := "interrupt"
	e := newEvent(withSignal(q))
	assert.Equal(t, q, e.Properties["signal"])
}

func TestWithHelpCommand(t *testing.T) {
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
	assert.NotContains(t, e.Properties, "help_command")
}

type configMock struct {
	name        string
	publicKey   string
	privateKey  string
	accessToken string
	service     string
	url         string
	project     string
	org         string
}

var _ Authenticator = configMock{}

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

func (c configMock) AccessToken() string {
	return c.accessToken
}
