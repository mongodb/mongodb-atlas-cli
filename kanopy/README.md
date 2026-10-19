# Atlas CLI OAuth callback on Kanopy

This directory packages and deploys the static callback page in `../oauth-callback/index.html` as the `atlas-cli-oauth-callback` web app in the `apix` Kanopy namespace.

## Deployment

The root `.drone.yml` builds a narrow image context so Kaniko never receives the rest of the Atlas CLI repository.

- Pull request updates: run the callback-page tests only; no Drone secrets are used.
- Push callback/deployment changes to `master`: build once, deploy to staging, then deploy the same image to production.
- Manual pre-merge staging: trigger the `deploy-staging-manual` custom build for `CLOUDP-459029` in Drone.
- Atlas CLI `v*` release tags do not deploy the callback page.
- Helm chart: `mongodb/web-app` 4.37.1.
- Container endpoint: `/atlas-cli/oauth-callback.html`.
- Liveness/readiness endpoint: `/healthz`.
- Kanopy service health: [staging dashboard](https://grafana.staging.corp.mongodb.com/d/DgX5qJmWz/kanopy-home-dashboard?orgId=1) and [production dashboard](https://grafana.corp.mongodb.com/d/DgX5qJmWz/kanopy-home-dashboard?orgId=1).

The standard Kanopy origins are CorpSecure-protected:

- Staging: `https://atlas-cli-oauth-callback.apix.staging.corp.mongodb.com/atlas-cli/oauth-callback.html`
- Production: `https://atlas-cli-oauth-callback.apix.prod.corp.mongodb.com/atlas-cli/oauth-callback.html`

They are internal origins, not the public OAuth redirect URI. The public URI remains `https://tools.mongodb.com/atlas-cli/oauth-callback.html` and requires the CloudFront/vanity route and approval work tracked by CLOUDP-459030.

## One-time Kanopy/Drone setup

1. Activate `mongodb/mongodb-atlas-cli` in [Drone](https://drone.corp.mongodb.com). If the `mongodb`-organization repository is not available there, ask the Drone/Kanopy administrators to onboard this repository; keep the callback source in this repository.
2. Add these repository secrets in Drone → repository Settings → Secrets, with event access as listed. Keep pull-request access disabled:
   - `ecr_access_key` and `ecr_secret_key` from the `ecr` Secret in the staging `apix` namespace.
     - Events: `push`, `custom`.
   - `staging_kubernetes_token` from `kanopy-cicd-token` in staging `apix`.
     - Events: `push`, `custom`.
   - `prod_kubernetes_token` from `kanopy-cicd-token` in production `apix`.
     - Event: `push`.
3. Verify that you can access the `apix` namespace in both clusters. Kanopy’s namespace admin group is `10gen-kanopy-apix-admin`; request membership through the Kanopy/MANA access process if you lack it.
4. For a pre-merge smoke test, push the reviewed feature branch, then select `CLOUDP-459029` → **New Build** in Drone. The manual custom pipeline deploys to the shared staging release, so it replaces the currently deployed callback version there.
5. After merge, a callback/deployment path change on `master` deploys staging and then production from the same commit. Atlas CLI release tags do not trigger this pipeline.

The secret values must be entered directly in Drone. Do not commit them or paste them into chat, tickets, or this repository.

## Local preview

From the repository root:

```sh
sh kanopy/prepare-context.sh
docker build -f kanopy/.context/Dockerfile -t atlas-cli-oauth-callback:local kanopy/.context
docker run --rm -p 8080:8080 atlas-cli-oauth-callback:local
curl -i http://localhost:8080/healthz
curl -i 'http://localhost:8080/atlas-cli/oauth-callback.html?code=synthetic&state=synthetic&iss=https%3A%2F%2Fauthorize.mongodb.com'
```

Use synthetic callback parameters only.

The shared encoder/decoder vectors are in `oauth-callback/testdata/paste-v1.json`. Validate the browser page and Go decoder with:

```sh
node --test oauth-callback/paste-v1.test.mjs
go test ./internal/cli/auth
```

## Callback logging and public exposure

The NGINX access log records `$uri`, not `$request_uri` or `$args`, so it omits query values. This does not control Kanopy mesh/edge or CloudFront logging; verify that `code`, `state`, and `iss` are redacted from every platform log sink before sending real OAuth responses through the endpoint. Error-log behavior also needs verification by Kanopy owners.

The `apix` namespace is mesh-only and its normal hostname enforces CorpSecure. Do not use either internal hostname as the OAuth redirect. Public traffic requires the approved CloudFront/vanity route, public-access review, and log redaction tracked separately in CLOUDP-459030.

## References

- [Kanopy static HTML example](https://github.com/10gen/kanopy-docs/blob/master/docs/examples/static-html/README.md)
- [Kanopy Drone pipeline guide](https://kanopy.corp.mongodb.com/docs/getting_started/drone_pipeline/)
- [Kanopy mesh ingress documentation](https://github.com/10gen/kanopy-docs/blob/master/docs/ingress/README.md)
- [Kanopy public workload isolation policy](https://github.com/10gen/kanopy-docs/blob/master/docs/policies/public_workload_isolation.md)
