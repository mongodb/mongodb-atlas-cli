package store

import (
	"context"
	"fmt"

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type OwnerCreator interface {
	CreateOwner(*om.User, []string) (*om.CreateUserResponse, error)
}

// CreateOwner encapsulate the logic to manage different cloud providers
func (s *Store) CreateOwner(u *om.User, IPs []string) (*om.CreateUserResponse, error) {
	switch s.service {
	case config.OpsManagerService:
		var opts *om.WhitelistOpts
		if len(IPs) > 0 {
			opts = &om.WhitelistOpts{Whitelist: IPs}
		}

		result, _, err := s.client.(*om.Client).UnauthUsers.CreateFirstUser(context.Background(), u, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
