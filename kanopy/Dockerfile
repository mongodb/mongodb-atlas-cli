# Multi-arch NGINX image, pinned to the reviewed OCI index digest.
FROM nginxinc/nginx-unprivileged:1.28-alpine@sha256:7377697a821c131a924a7105fafbe7414db4e9fcc77a6f08f776f33f141ec3f8

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY content/atlas-cli/oauth-callback/index.html /usr/share/nginx/html/atlas-cli/oauth-callback/index.html
