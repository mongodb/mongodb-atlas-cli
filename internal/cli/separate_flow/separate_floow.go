package separate_flow

import (
	"context"
	"fmt"
	"github.com/mongodb/mongocli/internal/cli/auth"
)

//go:generate mockgen -destination=../../mocks/mock.go -package=mocks github.com/mongodb/mongocli/internal/cli/separate_flow RegisterFlow

type RegisterFlow interface {
	Flow(opts *auth.RegisterOpts) error
}

type Flow struct {
	Ctx context.Context
}

// Run should be used instead of run for external command dependencies.
func (f *Flow) Flow(opts *auth.RegisterOpts) error {
	_, _ = fmt.Fprintf(opts.Login.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")

	if err := opts.RegisterAndAuthenticate(f.Ctx); err != nil {
		return err
	}

	opts.Login.SetOAuthUpAccess()
	s, err := opts.Login.Config.AccessTokenSubject()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.Login.OutWriter, "Successfully logged in as %s.\n", s)
	if opts.Login.SkipConfig {
		return opts.Login.Config.Save()
	}

	return nil
}