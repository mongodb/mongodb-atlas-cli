#!/usr/bin/env bash
# Copyright 2021 MongoDB Inc
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

export project
export revision
export created_at

VERSION_GIT="$(git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"
VERSION_NAME="$VERSION_GIT"
if [[ "${unstable-}" == "-unstable" ]]; then
	VERSION_NAME="$VERSION_GIT-next"
fi

PACKAGE_NAME="mongodb-atlas-cli_${VERSION_NAME}_windows_x86_64.msi"
BINARY_NAME="atlas.exe"

PACKAGE_URL=https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/unsigned_${PACKAGE_NAME}
BINARY_URL=https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/unsigned_${BINARY_NAME}

pushd bin

echo "downloading $PACKAGE_URL into $PWD/$PACKAGE_NAME"
curl "$PACKAGE_URL" --output "${PACKAGE_NAME}"

echo "downloading $BINARY_URL into $PWD/$BINARY_NAME"
curl "$BINARY_URL" --output "${BINARY_NAME}"
