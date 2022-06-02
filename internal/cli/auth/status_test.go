package auth

import (
	"context"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_GetStatus_InvalidToken(t *testing.T) {
	t.Cleanup(cleanUpConfig)
	ctx := context.TODO()
	config.SetAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")

	status, _ := GetStatus(ctx)
	assert.Equal(t, LoggedInWithInvalidToken, status)
}

func Test_GetStatus_APIKeys(t *testing.T) {
	t.Cleanup(cleanUpConfig)
	ctx := context.TODO()

	config.SetPublicAPIKey("publicKey")
	config.SetPrivateAPIKey("privateKey")

	status, _ := GetStatus(ctx)
	assert.Equal(t, LoggedInWithAPIKeys, status)
}

func Test_GetStatus_NotLoggedIn(t *testing.T) {
	t.Cleanup(cleanUpConfig)
	ctx := context.TODO()

	status, _ := GetStatus(ctx)
	assert.Equal(t, NotLoggedIn, status)
}

func cleanUpConfig() {
	config.SetAccessToken("")
	config.SetPublicAPIKey("")
	config.SetPrivateAPIKey("")
}
