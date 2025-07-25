package secure

//go:generate go tool go.uber.org/mock/mockgen -destination=./mocks.go -package=secure github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure Store

type Store interface {
	Available() bool
	Set(profileName string, propertyName string, value string) error
	Get(profileName string, propertyName string) (string, error)
	DeleteKey(profileName string, propertyName string) error
	DeleteProfile(profileName string) error
}
