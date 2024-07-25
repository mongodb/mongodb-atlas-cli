#!/usr/bin/env bash

# Copyright 2022 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -Eeou pipefail

export GOROOT="${GOROOT:?}"
export GITHUB_APP_ID=${github_app_id:?}
export GITHUB_APP_PEM=${github_app_pem:?}
export NOTARY_SERVICE_URL=${notary_service_url:?}
export MACOS_NOTARY_KEY=${notary_service_key_id:?}
export MACOS_NOTARY_SECRET=${notary_service_secret:?}
export GORELEASER_KEY=${goreleaser_key:?}
export VERSION_GIT

echo "$GITHUB_APP_PEM" > app.pem
GITHUB_INSTALLATION_ID=$(gh-token installations --app-id "$GITHUB_APP_ID" --key ./app.pem | jq '.[] | select(.account.login == "mongodb") | .id' | head -1)
GITHUB_TOKEN=$(gh-token generate --app-id "$GITHUB_APP_ID" --key ./app.pem --installation-id "$GITHUB_INSTALLATION_ID" -t)
rm -rf app.pem
export GITHUB_TOKEN

VERSION_GIT="$(git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"

if [[ "${unstable-}" == "-unstable" ]]; then
	# avoid race conditions on the notarization step by using `-p 1`
	./bin/goreleaser --config "build/package/.goreleaser.yml" --clean --release-notes "CHANGELOG.md" -p 1 --snapshot
else
	# avoid race conditions on the notarization step by using `-p 1`
	./bin/goreleaser --config "build/package/.goreleaser.yml" --clean --release-notes "CHANGELOG.md" -p 1
fi

gh-token revoke -t "$GITHUB_TOKEN"

# check that the notarization service signed the mac binaries
SIGNED_FILE_NAME=mongodb-atlas-cli_macos_signed.zip
if [[ -f "dist/$SIGNED_FILE_NAME" ]]; then
	echo "$SIGNED_FILE_NAME exists. The Mac notarization service has run."
else
	echo "ERROR: $SIGNED_FILE_NAME does not exist. The Mac notarization service has not run."
	exit 1 # ERROR
fi
