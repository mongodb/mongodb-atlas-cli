package store

import (
	"context"
	"fmt"

	cm "github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/mongodb/mongocli/internal/config"
)

type OwnerCreator interface {
	CreateOwner(*cm.User, []string) (*cm.CreateUserResponse, error)
}

// CreateOwner encapsulate the logic to manage different cloud providers
func (s *Store) CreateOwner(u *cm.User, IPs []string) (*cm.CreateUserResponse, error) {
	switch s.service {
	case config.OpsManagerService:
		var opts *cm.WhitelistOpts
		if len(IPs) > 0 {
			opts = &cm.WhitelistOpts{Whitelist: IPs}
		}

		result, _, err := s.client.(*cm.Client).UnauthUsers.CreateFirstUser(context.Background(), u, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
