# Atlas CLI OAuth callback

This directory contains the static callback page and the Kanopy app configuration that serves it in the `apix` namespace.

The page packages an authorization code for the waiting Atlas CLI prompt. When the response includes `iss`, the page uses it as HKDF info. Callback parameters are removed from the browser address bar after processing.

The callback endpoint is `/atlas-cli/oauth-callback.html`. The health check is `/healthz`.

## Local preview

From the repository root:

```sh
sh kanopy/prepare-context.sh
docker build -f kanopy/.context/Dockerfile -t atlas-cli-oauth-callback:local kanopy/.context
docker run --rm -p 8080:8080 atlas-cli-oauth-callback:local
curl -i http://localhost:8080/healthz
curl -i 'http://localhost:8080/atlas-cli/oauth-callback.html?code=synthetic&state=synthetic&iss=https%3A%2F%2Fauthorize.mongodb.com'
```

Use synthetic callback parameters for local preview.

The shared encoder and decoder vectors are in `oauth-callback/testdata/paste-v1.json`. Run the tests with:

```sh
node --test oauth-callback/paste-v1.test.mjs
go test ./internal/cli/auth
```
