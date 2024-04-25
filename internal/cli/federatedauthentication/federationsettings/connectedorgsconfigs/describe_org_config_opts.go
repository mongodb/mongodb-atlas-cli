package connectedorgsconfigs

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115012/admin"
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

func (opts *DescribeOrgConfigsOpts) GetConnectedOrgConfig(federationSettings string, orgId string) (*atlasv2.ConnectedOrgConfig, error) {
	params := &atlasv2.GetConnectedOrgConfigApiParams{
		FederationSettingsId: federationSettings,
		OrgId:                orgId,
	}

	return opts.describeStore.GetConnectedOrgConfig(params)
}
