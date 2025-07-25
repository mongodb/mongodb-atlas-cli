package config

import (
	"slices"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure"
	"github.com/spf13/afero"
)

var secureProperties = []string{
	publicAPIKey,
	privateAPIKey,
	AccessTokenField,
	RefreshTokenField,
}

type ProxyStore struct {
	insecure Store
	secure   secure.Store
}

func NewDefaultStore() (Store, error) {
	insecure, err := NewViperStore(afero.NewOsFs())

	if err != nil {
		return nil, err
	}
	secure := secure.NewSecureStore()

	return NewStore(insecure, secure), nil
}

func NewStore(insecure Store, secure secure.Store) Store {
	if !secure.Available() {
		return insecure
	}

	return &ProxyStore{
		insecure: insecure,
		secure:   secure,
	}
}

func isSecureProperty(propertyName string) bool {
	return slices.Contains(secureProperties, propertyName)
}

// Store interface implementation for ProxyStore

func (p *ProxyStore) IsSecure() bool {
	return true
}

func (p *ProxyStore) Save() error {
	return p.insecure.Save()
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
		if val, err := p.secure.Get(profileName, propertyName); err == nil {
			return val
		}
		return ""
	}
	return p.insecure.GetHierarchicalValue(profileName, propertyName)
}

func (p *ProxyStore) SetProfileValue(profileName string, propertyName string, value any) {
	if isSecureProperty(propertyName) {
		if v, ok := value.(string); ok {
			_ = p.secure.Set(profileName, propertyName, v)
		}
		return
	}
	p.insecure.SetProfileValue(profileName, propertyName, value)
}

func (p *ProxyStore) GetProfileValue(profileName string, propertyName string) any {
	if isSecureProperty(propertyName) {
		if val, err := p.secure.Get(profileName, propertyName); err == nil {
			return val
		}
		return ""
	}
	return p.insecure.GetProfileValue(profileName, propertyName)
}

func (p *ProxyStore) GetProfileStringMap(profileName string) map[string]string {
	return p.insecure.GetProfileStringMap(profileName)
}

func (p *ProxyStore) SetGlobalValue(propertyName string, value any) {
	if isSecureProperty(propertyName) {
		if v, ok := value.(string); ok {
			_ = p.secure.Set(DefaultProfile, propertyName, v)
		}
		return
	}
	p.insecure.SetGlobalValue(propertyName, value)
}

func (p *ProxyStore) GetGlobalValue(propertyName string) any {
	if isSecureProperty(propertyName) {
		if val, err := p.secure.Get(DefaultProfile, propertyName); err == nil {
			return val
		}
		return ""
	}
	return p.insecure.GetGlobalValue(propertyName)
}

func (p *ProxyStore) IsSetGlobal(propertyName string) bool {
	return p.insecure.IsSetGlobal(propertyName)
}
