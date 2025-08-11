package config

import (
	"errors"
	"slices"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure"
	"github.com/spf13/afero"
)

var SecureProperties = []string{
	publicAPIKey,
	privateAPIKey,
	AccessTokenField,
	RefreshTokenField,
	ClientIDField,
	ClientSecretField,
}

type ProxyStore struct {
	insecure Store
	secure   SecureStore
}

func NewDefaultStore() (Store, error) {
	insecure, err := NewViperStore(afero.NewOsFs(), true)

	if err != nil {
		return nil, err
	}

	profileNames := insecure.GetProfileNames()
	secureStore := secure.NewSecureStore(profileNames, SecureProperties)

	return NewStore(insecure, secureStore), nil
}

func NewStore(insecureStore Store, secureStore SecureStore) Store {
	if !secureStore.Available() {
		return insecureStore
	}

	return &ProxyStore{
		insecure: insecureStore,
		secure:   secureStore,
	}
}

func isSecureProperty(propertyName string) bool {
	return slices.Contains(SecureProperties, propertyName)
}

// Store interface implementation for ProxyStore

func (*ProxyStore) IsSecure() bool {
	return true
}

func (p *ProxyStore) Save() error {
	errs := []error{}

	if err := p.insecure.Save(); err != nil {
		errs = append(errs, err)
	}

	if err := p.secure.Save(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (p *ProxyStore) GetProfileNames() []string {
	return p.insecure.GetProfileNames()
}

func (p *ProxyStore) RenameProfile(oldProfileName string, newProfileName string) error {
	return p.insecure.RenameProfile(oldProfileName, newProfileName)
}

func (p *ProxyStore) DeleteProfile(profileName string) error {
	return p.insecure.DeleteProfile(profileName)
}

func (p *ProxyStore) GetHierarchicalValue(profileName string, propertyName string) any {
	if isSecureProperty(propertyName) {
		return p.secure.Get(profileName, propertyName)
	}
	return p.insecure.GetHierarchicalValue(profileName, propertyName)
}

func (p *ProxyStore) SetProfileValue(profileName string, propertyName string, value any) {
	if isSecureProperty(propertyName) {
		if v, ok := value.(string); ok {
			p.secure.Set(profileName, propertyName, v)
		}
		return
	}
	p.insecure.SetProfileValue(profileName, propertyName, value)
}

func (p *ProxyStore) GetProfileValue(profileName string, propertyName string) any {
	if isSecureProperty(propertyName) {
		return p.secure.Get(profileName, propertyName)
	}
	return p.insecure.GetProfileValue(profileName, propertyName)
}

func (p *ProxyStore) GetProfileStringMap(profileName string) map[string]string {
	return p.insecure.GetProfileStringMap(profileName)
}

func (p *ProxyStore) SetGlobalValue(propertyName string, value any) {
	if isSecureProperty(propertyName) {
		if v, ok := value.(string); ok {
			p.secure.Set(DefaultProfile, propertyName, v)
		}
		return
	}
	p.insecure.SetGlobalValue(propertyName, value)
}

func (p *ProxyStore) GetGlobalValue(propertyName string) any {
	if isSecureProperty(propertyName) {
		return p.secure.Get(DefaultProfile, propertyName)
	}
	return p.insecure.GetGlobalValue(propertyName)
}

func (p *ProxyStore) IsSetGlobal(propertyName string) bool {
	return p.insecure.IsSetGlobal(propertyName)
}
