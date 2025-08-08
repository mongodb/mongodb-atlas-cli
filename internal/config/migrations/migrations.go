package migrations

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure"
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
		GetInsecureStore: config.NewDefaultStore,
		GetSecureStore: func() (config.SecureStore, error) {
			return secure.NewSecureStore(), nil
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
	version, ok := store.GetGlobalValue(VersionKey).(int)
	if !ok {
		return 0, errors.New("invalid version type")
	}
	return version, nil
}

func (m *Migrator) setCurrentVersion(version int) error {
	store, err := m.dependencies.GetInsecureStore()
	if err != nil {
		return fmt.Errorf("failed to get store: %w", err)
	}
	store.SetGlobalValue(VersionKey, version)
	return nil
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
		if err := m.setCurrentVersion(currentVersion + 1); err != nil {
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
