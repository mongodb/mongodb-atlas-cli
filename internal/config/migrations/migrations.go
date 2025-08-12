package migrations

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure"
	"github.com/spf13/afero"
)

var (
	ErrConfigVersionUnsupported = errors.New("existing config version is not supported, please update the cli")

	VersionKey = "version"
)

type Migrator struct {
	dependencies MigrationDependencies
	migrations   []MigrationFunc
}

func NewMigrator(dependencies MigrationDependencies, migrations []MigrationFunc) *Migrator {
	return &Migrator{dependencies: dependencies, migrations: migrations}
}

func NewDefaultMigrator() *Migrator {
	dependencies := MigrationDependencies{
		GetInsecureStore: func() (config.Store, error) {
			return config.NewViperStore(afero.NewOsFs(), false)
		},
		GetSecureStore: func() (config.SecureStore, error) {
			// For migrations, we need to create a temporary insecure store to get profile names
			insecureStore, err := config.NewViperStore(afero.NewOsFs(), false)
			if err != nil {
				return nil, fmt.Errorf("failed to create insecure store for migrations: %w", err)
			}
			profileNames := insecureStore.GetProfileNames()
			return secure.NewSecureStore(profileNames, config.SecureProperties), nil
		},
	}

	migrations := []MigrationFunc{
		NewMigrateToVersion2(),
	}

	return NewMigrator(dependencies, migrations)
}

func (m *Migrator) currentVersion() (int, error) {
	store, err := m.dependencies.GetInsecureStore()
	if err != nil {
		return 0, fmt.Errorf("failed to get store: %w", err)
	}

	// Get the current version from the store.
	// It is possible that the version is not set, in which case it could return nil or an empty string.
	rawVersion := store.GetGlobalValue(VersionKey)
	if rawVersion == nil || rawVersion == "" {
		return 1, nil
	}

	// Convert the version to an int.
	version, ok := rawVersion.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid version type: %T", rawVersion)
	}

	return int(version), nil
}

func (m *Migrator) persistCurrentVersion(version int) error {
	store, err := m.dependencies.GetInsecureStore()
	if err != nil {
		return fmt.Errorf("failed to get store: %w", err)
	}
	store.SetGlobalValue(VersionKey, version)
	return store.Save()
}

func (m *Migrator) Migrate() error {
	currentVersion, err := m.currentVersion()
	if err != nil {
		return fmt.Errorf("failed to get current config version: %w", err)
	}

	maxSupportedVersion := len(m.migrations) + 1

	// If the config version is the same as the max supported version, we don't need to migrate.
	if currentVersion == maxSupportedVersion {
		return nil
	}

	// If the config version is greater than the max supported version then the cli is outdated.
	if currentVersion > maxSupportedVersion {
		return ErrConfigVersionUnsupported
	}

	// In all other cases, we need to migrate the config.
	//
	// Notes:
	// - Migrations are in order and 0-indexed
	// -We skip the migrations that have already been applied.
	for _, migration := range m.migrations[currentVersion-1:] {
		if err := migration(m.dependencies); err != nil {
			return err
		}

		// In case the migration succeeds, update the config version.
		if err := m.persistCurrentVersion(currentVersion + 1); err != nil {
			return fmt.Errorf("failed to set current config version: %w", err)
		}
	}
	return nil
}

type MigrationDependencies struct {
	GetInsecureStore func() (config.Store, error)
	GetSecureStore   func() (config.SecureStore, error)
}

// MigrationFunc is a function that migrates the config.
// It is passed the dependencies and should return an error if the migration fails.
//
// Migrations should be idempotent.
// They should only be run once, but in theory they could run multiple times if updating the version number in the config fails.
type MigrationFunc func(dependencies MigrationDependencies) error
