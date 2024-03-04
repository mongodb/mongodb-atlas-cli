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
export GITHUB_TOKEN=${github_token:?}
export NOTARY_SERVICE_URL=${notary_service_url:?}
export MACOS_NOTARY_KEY=${notary_service_key_id:?}
export MACOS_NOTARY_SECRET=${notary_service_secret:?}
export GORELEASER_KEY=${goreleaser_key:?}
export VERSION_GIT

VERSION_GIT="$(git tag --list "${tool_name:?}/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"

if [[ "${unstable-}" == "-unstable" ]]; then
	# avoid race conditions on the notarization step by using `-p 1`
	./bin/goreleaser --config "build/package/.goreleaser.yml" --rm-dist --release-notes "CHANGELOG.md" --snapshot -p 1
else
	# avoid race conditions on the notarization step by using `-p 1`
	./bin/goreleaser --config "build/package/.goreleaser.yml" --rm-dist --release-notes "CHANGELOG.md" -p 1
fi

# check that the notarization service signed the mac binaries
SIGNED_FILE_NAME=mongodb-atlas-cli_macos_signed.zip
if [[ -f "dist/$SIGNED_FILE_NAME" ]]; then
	echo "$SIGNED_FILE_NAME exists. The Mac notarization service has run."
else
	echo "ERROR: $SIGNED_FILE_NAME does not exist. The Mac notarization service has not run."
	exit 1 # ERROR
fi
