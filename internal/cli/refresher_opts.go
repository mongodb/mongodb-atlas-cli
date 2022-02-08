package cli

import (
	"context"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/oauth"
	"go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type RefresherOpts struct {
	flow Refresher
}

type Refresher interface {
	RefreshToken(context.Context, string) (*auth.Token, *atlas.Response, error)
}

func (opts *RefresherOpts) InitFlow() error {
	var err error
	opts.flow, err = oauth.FlowWithConfig(config.Default())
	return err
}

func (opts *RefresherOpts) RefreshAccessToken(ctx context.Context) error {
	current, err := config.Token()
	if current == nil {
		return err
	}
	if current.Valid() {
		return nil
	}
	t, _, err := opts.flow.RefreshToken(ctx, config.RefreshToken())
	if err != nil {
		return err
	}
	config.SetAccessToken(t.AccessToken)
	config.SetRefreshToken(t.RefreshToken)
	return config.Save()
}
