package cli

import (
	"sync"

	"github.com/10gen/mcli/internal/config"
)

type globalOpts struct {
	config.Config
	profile   string
	projectID string
	once      sync.Once
}

// newGlobalOpts returns an globalOpts
func newGlobalOpts() *globalOpts {
	return new(globalOpts)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) ProjectID() string {
	_ = opts.loadConfig()
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = opts.Config.ProjectID()
	return opts.projectID
}

func (opts *globalOpts) loadConfig() error {
	var err error
	opts.once.Do(func() {
		opts.Config, err = config.New(opts.profile)
	})
	return err
}
