package secure

import "github.com/zalando/go-keyring"

const servicePrefix = "atlascli_"

func createServiceName(profileName string) string {
	return servicePrefix + profileName
}

type SecureStore struct{}

func NewSecureStore() *SecureStore {
	return &SecureStore{}
}

func (s *SecureStore) Available() bool {
	_, err := keyring.Get(createServiceName("default"), "test")
	return err == nil
}

func (s *SecureStore) Set(profileName string, propertyName string, value string) error {
	return keyring.Set(createServiceName(profileName), propertyName, value)
}

func (s *SecureStore) Get(profileName string, propertyName string) (string, error) {
	return keyring.Get(createServiceName(profileName), propertyName)
}

func (s *SecureStore) DeleteKey(profileName string, propertyName string) error {
	return keyring.Delete(createServiceName(profileName), propertyName)
}

func (s *SecureStore) DeleteProfile(profileName string) error {
	return keyring.DeleteAll(createServiceName(profileName))
}
