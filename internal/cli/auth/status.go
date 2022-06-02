package auth

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	LoggedInWithValidToken
	LoggedInWithInvalidToken
	LoggedInWithAPIKeys
	NotLoggedIn
)

// GetStatus get user authentication status.
func GetStatus(ctx context.Context) (int, error) {
	var err error

	if config.PublicAPIKey() != "" && config.PrivateAPIKey() != "" {
		return LoggedInWithAPIKeys, nil
	}
	if _, err = AccountWithAccessToken(); err == nil {
		// token exists but it is not refreshed
		if err = cli.RefreshToken(ctx); err != nil || validate.Token() != nil {
			return LoggedInWithInvalidToken, nil
		}
		return LoggedInWithValidToken, nil
	}

	return NotLoggedIn, err
}
