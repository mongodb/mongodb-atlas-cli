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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/oauth"
	atlasauth "go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var TokenRefreshed bool

type RefresherOpts struct {
	flow Refresher
}

//go:generate mockgen -destination=../mocks/mock_refresher.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli Refresher
type Refresher interface {
	RequestCode(context.Context) (*atlasauth.DeviceCode, *atlas.Response, error)
	PollToken(context.Context, *atlasauth.DeviceCode) (*atlasauth.Token, *atlas.Response, error)
	RefreshToken(context.Context, string) (*atlasauth.Token, *atlas.Response, error)
	RegistrationConfig(ctx context.Context) (*atlasauth.RegistrationConfig, *atlas.Response, error)
}

func (opts *RefresherOpts) InitFlow(c oauth.ServiceGetter) func() error {
	return func() error {
		var err error
		opts.flow, err = oauth.FlowWithConfig(c)
		return err
	}
}

// WithFlow set a flow for testing.
func (opts *RefresherOpts) WithFlow(f Refresher) {
	opts.flow = f
}

var ErrInvalidRefreshToken = errors.New("session expired")

func (opts *RefresherOpts) RefreshAccessToken(ctx context.Context) error {
	current, err := config.Token()
	if current == nil {
		return err
	}
	if current.Valid() {
		TokenRefreshed = true
		return nil
	}
	t, _, err := opts.flow.RefreshToken(ctx, config.RefreshToken())
	if err != nil {
		var target *atlas.ErrorResponse
		if errors.As(err, &target) && target.ErrorCode == "INVALID_REFRESH_TOKEN" {
			return fmt.Errorf(
				`%w

Please note that your session expires periodically. 
If you use Atlas CLI for automation, see https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-automate/ for best practices.
To login, run: atlas auth login`,
				ErrInvalidRefreshToken)
		}
		return err
	}
	config.SetAccessToken(t.AccessToken)
	config.SetRefreshToken(t.RefreshToken)
	if err := config.Save(); err != nil {
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
