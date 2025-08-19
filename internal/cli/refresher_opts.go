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
	"net/http"

	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlasauth "go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var TokenRefreshed bool

type RefresherOpts struct {
	flow Refresher
}

//go:generate go tool go.uber.org/mock/mockgen -destination=../mocks/mock_refresher.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli Refresher
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
		opts.flow, err = transport.FlowWithConfig(c, client, config.UserAgent)
		return err
	}
}

// WithFlow set a flow for testing.
func (opts *RefresherOpts) WithFlow(f Refresher) {
	opts.flow = f
}

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
