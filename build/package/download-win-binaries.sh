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

VERSION_GIT="$(git tag --list "mongocli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"
VERSION_NAME="$VERSION_GIT"
if [[ "${unstable-}" == "-unstable" ]]; then
	VERSION_NAME="$VERSION_GIT-next"
fi

PACKAGE_NAME="mongocli_${VERSION_NAME}_windows_x86_64.msi"
BINARY_NAME="mongocli.exe"

pushd bin

echo "downloading https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/${PACKAGE_NAME} into $PWD"
curl "https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/${PACKAGE_NAME}" --output "${PACKAGE_NAME}"

echo "downloading https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/${BINARY_NAME} into $PWD"
curl "https://${BUCKET}.s3.amazonaws.com/${project}/dist/${revision}_${created_at}/${BINARY_NAME}" --output "${BINARY_NAME}"
