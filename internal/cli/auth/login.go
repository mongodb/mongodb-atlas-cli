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

package auth

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/oauth"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/mock_login.go -package=mocks github.com/mongodb/mongocli/internal/cli/auth Authenticator

type Authenticator interface {
	RequestCode(context.Context) (*auth.DeviceCode, *atlas.Response, error)
	PollToken(context.Context, *auth.DeviceCode) (*auth.Token, *atlas.Response, error)
}

type loginOpts struct {
	OutWriter      io.Writer
	Service        string
	AuthToken      string
	RefreshToken   string
	OpsManagerURL  string
	ProjectID      string
	OrgID          string
	MongoShellPath string
	Output         string
	isGov          bool
	isCloudManager bool
	noBrowser      bool
	flow           Authenticator
}

func (opts *loginOpts) initFlow() func() error {
	return func() error {
		var err error
		opts.flow, err = oauth.FlowWithConfig(config.Default())
		return err
	}
}

func (opts *loginOpts) SetUpAccess() {
	config.SetService(opts.Service)
	if opts.AuthToken != "" {
		config.SetAuthToken(opts.AuthToken)
	}
	if opts.RefreshToken != "" {
		config.SetRefreshToken(opts.RefreshToken)
	}
	if opts.OpsManagerURL != "" {
		config.SetOpsManagerURL(opts.OpsManagerURL)
	}
}

func (opts *loginOpts) Run() error {
	code, _, err := opts.flow.RequestCode(context.TODO())
	if err != nil {
		return err
	}

	codeDuration := time.Duration(code.ExpiresIn) * time.Second
	_, _ = fmt.Fprintf(opts.OutWriter, `
Here is your one-time code: %s-%s

Sign in with your browser and enter the code.

Or go to %s

Your code will expire after %.0f minutes.
`,
		code.UserCode[0:len(code.UserCode)/2],
		code.UserCode[len(code.UserCode)/2:],
		code.VerificationURI,
		codeDuration.Minutes(),
	)
	if !opts.noBrowser {
		if errBrowser := browser.OpenFile(code.VerificationURI); errBrowser != nil {
			_, _ = fmt.Fprintf(os.Stderr, "There was an issue opening your browser\n")
		}
	}

	accessToken, _, err := opts.flow.PollToken(context.TODO(), code)
	if err != nil {
		return err
	}
	opts.AuthToken = accessToken.AccessToken
	opts.RefreshToken = accessToken.RefreshToken
	opts.SetUpAccess()
	if err := config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, `Successfully logged in`)
	return nil
}

func LoginBuilder() *cobra.Command {
	opts := &loginOpts{}
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Atlas",
		Example: `  To start the interactive setup
  $ mongocli auth login
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.initFlow()()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Annotations: map[string]string{
			"toc": "true",
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.isGov, "gov", false, "Loging to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.isCloudManager, "cm", false, "Loging to Cloud Manager.")
	cmd.Flags().BoolVar(&opts.noBrowser, "noBrowser", false, "Don't try to open a browser session.")

	return cmd
}

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate me.",
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		LoginBuilder(),
	)

	return cmd
}
