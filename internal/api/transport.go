package api

import (
	"net/http"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
)

func authenticatedTransport(authenticatedConfig store.AuthenticatedConfig, httpTransport http.RoundTripper) http.RoundTripper {
	username := authenticatedConfig.PublicAPIKey()
	password := authenticatedConfig.PrivateAPIKey()

	if username != "" && password != "" {
		return &digest.Transport{
			Username:  username,
			Password:  password,
			Transport: httpTransport,
		}
	}

	return &transport{
		authenticatedConfig: authenticatedConfig,
		base:                httpTransport,
	}
}

type transport struct {
	authenticatedConfig store.AuthenticatedConfig
	base                http.RoundTripper
}

func (tr *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := tr.authenticatedConfig.Token()
	if err == nil {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	return tr.base.RoundTrip(req)
}
