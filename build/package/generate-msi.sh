#!/usr/bin/env bash
# Copyright 2020 MongoDB Inc
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

GOCACHE="$(cygpath --mixed "${workdir:?}\.gocache")"
CGO_ENABLED=0
export GOCACHE
export CGO_ENABLED

go-msi check-env

VERSION_GIT="$(git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"
VERSION_NAME="$VERSION_GIT"
if [[ "${unstable-}" == "-unstable" ]]; then
	VERSION_NAME="$VERSION_GIT-next"
fi

COMMIT=$(git log -n1 --format=format:"%H")

SOURCE_FILES=./cmd/atlas
PACKAGE_NAME=mongodb-atlas-cli_${VERSION_NAME}_windows_x86_64.msi
OUTPUT=./bin/atlas.exe
LINKER_FLAGS="-s -w -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.Version=${VERSION_GIT} -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.GitCommit=${COMMIT}"
WIX_MANIFEST_FILE="./build/package/wix/atlascli.json"

env GOOS=windows GOARCH=amd64 go build \
	-ldflags "${LINKER_FLAGS}" -o ${OUTPUT} "${SOURCE_FILES}"

go-msi make --path "${WIX_MANIFEST_FILE}" --msi "dist/${PACKAGE_NAME}" --version "${VERSION_GIT}"
