# Atlas CLI OAuth callback

This directory contains the static callback page and Kanopy app configuration for Atlas CLI authorization. The page packages the authorization code for the waiting CLI prompt and uses the response issuer as HKDF info when present.

The callback route is `/atlas-cli/oauth-callback.html`; the health check is `/healthz`.

## Local preview

From the repository root:

```sh
docker build -t atlas-cli-oauth-callback:local build/kanopy
docker run --rm -p 8080:8080 atlas-cli-oauth-callback:local
curl -i http://localhost:8080/healthz
curl -i 'http://localhost:8080/atlas-cli/oauth-callback.html?code=synthetic&state=synthetic&iss=https%3A%2F%2Fauthorize.mongodb.com'
```

Use synthetic callback values for local preview.

The shared vectors are in `build/kanopy/testdata/paste-v1.json`. Run the page and decoder tests with:

```sh
node --test build/kanopy/tests/paste-v1.test.mjs
go test ./internal/cli/auth
```
