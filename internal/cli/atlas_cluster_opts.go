package cli

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/store"
)

var defaultMongoDBMajorVersion string

func DefaultMongoDBMajorVersion() (string, error) {
	if defaultMongoDBMajorVersion != "" {
		return defaultMongoDBMajorVersion, nil
	}
	s, err := store.New(store.PrivateUnauthenticatedPreset(config.Default()))
	if err != nil {
		return "", err
	}
	defaultMongoDBMajorVersion, _ = s.DefaultMongoDBVersion()

	return defaultMongoDBMajorVersion, nil
}
