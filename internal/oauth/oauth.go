package oauth

import (
	"net"
	"net/http"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"go.mongodb.org/atlas/auth"
)

const (
	timeout               = 5 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 5
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
	cloudGovServiceURL    = "https://cloud.mongodbgov.com/"
)

var defaultTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepAlive,
	}).DialContext,
	MaxIdleConns:          maxIdleConns,
	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	Proxy:                 http.ProxyFromEnvironment,
	IdleConnTimeout:       idleConnTimeout,
	ExpectContinueTimeout: expectContinueTimeout,
}

type ServiceGetter interface {
	Service() string
	OpsManagerURL() string
}

const clientID = "0oadn4hoajpzxeSEy357"

func FlowWithConfig(c ServiceGetter) (*auth.Config, error) {
	client := http.DefaultClient
	client.Transport = defaultTransport
	authOpts := []auth.ConfigOpt{
		auth.SetUserAgent(config.UserAgent),
		auth.SetClientID(clientID),
		auth.SetScopes([]string{"openid", "profile", "offline_access"}),
	}
	if configURL := c.OpsManagerURL(); configURL != "" {
		authOpts = append(authOpts, auth.SetAuthURL(c.OpsManagerURL()))
	} else if c.Service() == config.CloudGovService {
		authOpts = append(authOpts, auth.SetAuthURL(cloudGovServiceURL))
	}
	return auth.NewConfigWithOptions(client, authOpts...)
}
