#!/bin/sh
set -eu

context=kanopy/.context
rm -rf "$context"
mkdir -p "$context/content/atlas-cli/oauth-callback"
cp kanopy/Dockerfile "$context/Dockerfile"
cp kanopy/nginx.conf "$context/nginx.conf"
cp oauth-callback/index.html "$context/content/atlas-cli/oauth-callback/index.html"
