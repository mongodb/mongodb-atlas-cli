package migrations

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
)

func NewMigrateToVersion2() MigrationFunc {
	return func(dependencies MigrationDependencies) error {
		// First, we upgrade the auth type for each profile.
		insecureStore, err := dependencies.GetInsecureStore()
		if err != nil {
			return fmt.Errorf("failed to get store: %w", err)
		}

		setAuthTypes(insecureStore, getAuthType)
		if err := insecureStore.Save(); err != nil {
			return fmt.Errorf("failed to save store: %w", err)
		}

		// Once the auth type is set, we can migrate secrets from insecure store to the secure store.
		secureStore, err := dependencies.GetSecureStore()
		if err != nil {
			return fmt.Errorf("failed to get secure store: %w", err)
		}

		// Migrate secrets from insecure store to the secure store if the secure store is available.
		if secureStore.Available() {
			migrateSecrets(insecureStore, secureStore)

			if err := secureStore.Save(); err != nil {
				return fmt.Errorf("failed to save secure store: %w", err)
			}
		}

		return nil
	}
}

// setAuthTypes sets the auth type for each profile based on the credentials available.
// Nothing is set if no credentials are found.
func setAuthTypes(store config.Store, getAuthType func(*config.Profile) config.AuthMechanism) {
	profileNames := store.GetProfileNames()

	for _, name := range profileNames {
		profile := config.NewProfile(name, store)
		authType := getAuthType(profile)
		if authType != "" {
			profile.SetAuthType(authType)
		}
	}
}

func getAuthType(profile *config.Profile) config.AuthMechanism {
	switch {
	case profile.PublicAPIKey() != "" && profile.PrivateAPIKey() != "":
		return config.APIKeys
	case profile.AccessToken() != "" && profile.RefreshToken() != "":
		return config.UserAccount
	case profile.ClientID() != "" && profile.ClientSecret() != "":
		return config.ServiceAccount
	}
	return config.AuthMechanism("") // This should not happen unless profile is not properly initialized.
}

// migrateSecrets migrates secrets from insecure store to the secure store.
// It also deletes the secrets from the insecure store.
func migrateSecrets(insecureStore config.Store, secureStore config.SecureStore) {
	profileNames := insecureStore.GetProfileNames()

	for _, name := range profileNames {
		for _, property := range config.SecureProperties {
			if value, ok := insecureStore.GetProfileValue(name, property).(string); ok && value != "" {
				secureStore.Set(name, property, value)
				insecureStore.SetProfileValue(name, property, nil)
			}
		}
	}
}
