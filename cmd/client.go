package cmd

import (
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/viper"
)

func newAuthenticatedClient() opsmanager.Client {
	baseURL := viper.GetString("base_url")
	publicKey := viper.GetString("public_key")
	privateKey := viper.GetString("private_key")
	resolver := httpclient.NewURLResolverWithPrefix(baseURL, opsmanager.PublicAPIPrefix)

	return opsmanager.NewClientWithDigestAuth(resolver, publicKey, privateKey)
}

func newDefaultClient() opsmanager.Client {
	baseURL := viper.GetString("base_url")
	resolver := httpclient.NewURLResolverWithPrefix(baseURL, opsmanager.PublicAPIPrefix)

	return opsmanager.NewDefaultClient(resolver)
}
