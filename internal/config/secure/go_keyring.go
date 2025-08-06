package secure

import "github.com/zalando/go-keyring"

const servicePrefix = "atlascli_"

func createServiceName(profileName string) string {
	return servicePrefix + profileName
}

type KeyringStore struct{}

func NewSecureStore() *KeyringStore {
	return &KeyringStore{}
}

func (*KeyringStore) Available() bool {
	_, err := keyring.Get(createServiceName("default"), "test")
	return err == nil
}

func (*KeyringStore) Set(profileName string, propertyName string, value string) error {
	return keyring.Set(createServiceName(profileName), propertyName, value)
}

func (*KeyringStore) Get(profileName string, propertyName string) (string, error) {
	return keyring.Get(createServiceName(profileName), propertyName)
}

func (*KeyringStore) DeleteKey(profileName string, propertyName string) error {
	return keyring.Delete(createServiceName(profileName), propertyName)
}

func (*KeyringStore) DeleteProfile(profileName string) error {
	return keyring.DeleteAll(createServiceName(profileName))
}
