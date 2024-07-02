// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectedorgsconfigs

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type DescribeOrgConfigsOpts struct {
	describeStore store.ConnectedOrgConfigsDescriber
}

func (opts *DescribeOrgConfigsOpts) InitDescribeStore(ctx context.Context) func() error {
	return func() error {
		if opts.describeStore != nil {
			return nil
		}

		var err error
		opts.describeStore, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOrgConfigsOpts) GetConnectedOrgConfig(federationSettings string, orgID string) (*atlasv2.ConnectedOrgConfig, error) {
	params := &atlasv2.GetConnectedOrgConfigApiParams{
		FederationSettingsId: federationSettings,
		OrgId:                orgID,
	}

	return opts.describeStore.GetConnectedOrgConfig(params)
}
