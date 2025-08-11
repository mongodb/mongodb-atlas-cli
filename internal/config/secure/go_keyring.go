package secure

import (
	"errors"
	"slices"

	"github.com/zalando/go-keyring"
)

//go:generate go tool go.uber.org/mock/mockgen -destination=./mocks.go -package=secure github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure KeyringClient

const servicePrefix = "atlascli_"

func createServiceName(profileName string) string {
	return servicePrefix + profileName
}

// KeyringClient abstracts keyring operations for easier testing
type KeyringClient interface {
	Set(service, user, password string) error
	Get(service, user string) (string, error)
	Delete(service, user string) error
	DeleteAll(service string) error
}

// DefaultKeyringClient implements KeyringClient using the zalando/go-keyring library
type DefaultKeyringClient struct{}

func NewDefaultKeyringClient() *DefaultKeyringClient {
	return &DefaultKeyringClient{}
}

func (*DefaultKeyringClient) Set(service, user, password string) error {
	return keyring.Set(service, user, password)
}

func (*DefaultKeyringClient) Get(service, user string) (string, error) {
	value, err := keyring.Get(service, user)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return "", err
	}

	return value, nil
}

func (*DefaultKeyringClient) Delete(service, user string) error {
	return keyring.Delete(service, user)
}

func (*DefaultKeyringClient) DeleteAll(service string) error {
	return keyring.DeleteAll(service)
}

// Operation types for tracking changes
type operationType int

const (
	opSet operationType = iota
	opDelete
	opDeleteProfile
)

// pendingOperation represents a change that needs to be persisted
type pendingOperation struct {
	opType       operationType
	profileName  string
	propertyName string
	value        string
}

type KeyringStore struct {
	// Available indicates if the keyring is available.
	available bool
	// In-memory cache: map[profileName]map[propertyName]value
	cache map[string]map[string]string
	// List of operations to perform when Save() is called
	pendingOps []pendingOperation
	// Properties that are considered secure
	secureProperties []string
	// KeyringClient for keyring operations
	keyringClient KeyringClient
}

func NewSecureStore(profileNames []string, secureProperties []string) *KeyringStore {
	return NewSecureStoreWithClient(profileNames, secureProperties, NewDefaultKeyringClient())
}

func NewSecureStoreWithClient(profileNames []string, secureProperties []string, keyringClient KeyringClient) *KeyringStore {
	store := &KeyringStore{
		cache:            make(map[string]map[string]string),
		pendingOps:       make([]pendingOperation, 0),
		secureProperties: secureProperties,
		keyringClient:    keyringClient,
	}

	// Check if the keyring is available.
	// We do this my marking the store as available if we can get a value from the keyring.
	available := false
	attemptedToRead := false

	// Load all existing secure properties for all profiles into memory
outer:
	for _, profileName := range profileNames {
		store.cache[profileName] = make(map[string]string)
		for _, propertyName := range secureProperties {
			attemptedToRead = true

			// Attempt to read the value from the keyring.
			value, err := keyringClient.Get(createServiceName(profileName), propertyName)

			// If the store returns an error, break the loop.
			if err != nil {
				break outer
			}

			store.cache[profileName][propertyName] = value
			available = true
		}
	}

	// If we didn't attempt to read, try to read a value from the default service.
	if !attemptedToRead {
		_, err := keyringClient.Get(createServiceName("default"), "test")
		available = err == nil
	}

	// Set the available flag.
	store.available = available

	return store
}

func (k *KeyringStore) Available() bool {
	return k.available
}

func (k *KeyringStore) Save() error {
	// Process all pending operations
	for _, op := range k.pendingOps {
		switch op.opType {
		case opSet:
			if err := k.keyringClient.Set(createServiceName(op.profileName), op.propertyName, op.value); err != nil {
				return err
			}
		case opDelete:
			if err := k.keyringClient.Delete(createServiceName(op.profileName), op.propertyName); err != nil {
				return err
			}
		case opDeleteProfile:
			if err := k.keyringClient.DeleteAll(createServiceName(op.profileName)); err != nil {
				return err
			}
		}
	}

	// Clear pending operations after successful save
	k.pendingOps = make([]pendingOperation, 0)
	return nil
}

func (k *KeyringStore) Set(profileName string, propertyName string, value string) {
	// Ignore properties that are not in SecureProperties
	if !slices.Contains(k.secureProperties, propertyName) {
		return
	}

	// Initialize profile map if it doesn't exist
	if k.cache[profileName] == nil {
		k.cache[profileName] = make(map[string]string)
	}

	// Update in-memory cache
	k.cache[profileName][propertyName] = value

	// Add to pending operations
	k.pendingOps = append(k.pendingOps, pendingOperation{
		opType:       opSet,
		profileName:  profileName,
		propertyName: propertyName,
		value:        value,
	})
}

func (k *KeyringStore) Get(profileName string, propertyName string) string {
	// Check if profile exists in cache
	if profileCache, exists := k.cache[profileName]; exists {
		if value, exists := profileCache[propertyName]; exists {
			return value
		}
	}
	return ""
}

func (k *KeyringStore) DeleteKey(profileName string, propertyName string) {
	// Remove from in-memory cache if it exists
	if profileCache, exists := k.cache[profileName]; exists {
		delete(profileCache, propertyName)
	}

	// Add to pending operations
	k.pendingOps = append(k.pendingOps, pendingOperation{
		opType:       opDelete,
		profileName:  profileName,
		propertyName: propertyName,
	})
}

func (k *KeyringStore) DeleteProfile(profileName string) {
	// Remove from in-memory cache
	delete(k.cache, profileName)

	// Add to pending operations
	k.pendingOps = append(k.pendingOps, pendingOperation{
		opType:      opDeleteProfile,
		profileName: profileName,
	})
}
