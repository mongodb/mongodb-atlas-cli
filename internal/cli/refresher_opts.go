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

package cli

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	atlasauth "go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var TokenRefreshed bool

type RefresherOpts struct {
	flow               Refresher
	refreshUserAccount func(context.Context, *config.Profile) error
}

//go:generate go tool go.uber.org/mock/mockgen -destination=../mocks/mock_refresher.go -package=mocks -source=refresher_opts.go
type Refresher interface {
	RequestCode(context.Context) (*atlasauth.DeviceCode, *atlas.Response, error)
	PollToken(context.Context, *atlasauth.DeviceCode) (*atlasauth.Token, *atlas.Response, error)
	RefreshToken(context.Context, string) (*atlasauth.Token, *atlas.Response, error)
	RegistrationConfig(ctx context.Context) (*atlasauth.RegistrationConfig, *atlas.Response, error)
}

func (opts *RefresherOpts) InitFlow(c transport.ServiceGetter) func() error {
	return func() error {
		var err error
		client := http.DefaultClient
		client.Transport = transport.Default()
		opts.flow, err = transport.FlowWithConfig(c, client, version.Version)
		return err
	}
}

// WithFlow set a flow for testing.
func (opts *RefresherOpts) WithFlow(f Refresher) {
	opts.flow = f
}

// WithUserAccountRefresher set the user_account refresher for testing.
func (opts *RefresherOpts) WithUserAccountRefresher(f func(context.Context, *config.Profile) error) {
	opts.refreshUserAccount = f
}

func (opts *RefresherOpts) RefreshAccessToken(ctx context.Context) error {
	return opts.refreshProfileAccessToken(ctx, config.Default())
}

func (opts *RefresherOpts) RefreshAccessTokenForProfile(ctx context.Context, profile *config.Profile) error {
	if err := opts.refreshProfileAccessToken(ctx, profile); err != nil {
		return err
	}
	if profile.AccessToken() == "" || profile.RefreshToken() != "" {
		return nil
	}
	current, err := profile.ServiceAccountToken()
	if err != nil {
		return err
	}
	if current != nil && !current.Valid() {
		return fmt.Errorf("access token for profile %s is expired and no refresh token is available", profile.Name())
	}
	return nil
}

func (opts *RefresherOpts) refreshProfileAccessToken(ctx context.Context, profile *config.Profile) error {
	current, err := profile.Token()
	if current == nil {
		return err
	}
	if current.Valid() {
		TokenRefreshed = true
		return nil
	}
	if profile.AuthType() == config.UserAccount {
		return opts.refreshUserAccountToken(ctx, profile)
	}
	if opts.flow == nil || profile != config.Default() {
		if err := opts.InitFlow(profile)(); err != nil {
			return err
		}
	}
	t, _, err := opts.flow.RefreshToken(ctx, profile.RefreshToken())
	if err != nil {
		return err
	}
	profile.SetAccessToken(t.AccessToken)
	profile.SetRefreshToken(t.RefreshToken)
	if err := profile.Save(); err != nil {
		return err
	}
	TokenRefreshed = true
	return nil
}

// refreshUserAccountToken refreshes through the Authorization Server; the device flow must never see these tokens.
func (opts *RefresherOpts) refreshUserAccountToken(ctx context.Context, profile *config.Profile) error {
	refresh := opts.refreshUserAccount
	if refresh == nil {
		refresh = func(ctx context.Context, p *config.Profile) error {
			return transport.RefreshUserAccountToken(ctx, p, version.Version)
		}
	}
	if err := refresh(ctx, profile); err != nil {
		return err
	}
	TokenRefreshed = true
	return nil
}

func (opts *RefresherOpts) PollToken(c context.Context, d *atlasauth.DeviceCode) (*atlasauth.Token, *atlas.Response, error) {
	return opts.flow.PollToken(c, d)
}

func (opts *RefresherOpts) RequestCode(c context.Context) (*atlasauth.DeviceCode, *atlas.Response, error) {
	return opts.flow.RequestCode(c)
}

func (opts *RefresherOpts) RegistrationConfig(c context.Context) (*atlasauth.RegistrationConfig, *atlas.Response, error) {
	return opts.flow.RegistrationConfig(c)
}
