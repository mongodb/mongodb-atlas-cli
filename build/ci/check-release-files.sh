#!/usr/bin/env bash

# Copyright 2025 MongoDB Inc
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

if [[ -z "${version}" ]]; then
    echo "version environment variable is not set"
    exit 1
fi

# shellcheck disable=SC2154 # unstable is set by evergreen
if [[ "${unstable}" == "-unstable" ]]; then
    version="${version}-next"
fi

REQUIRED_FILES=(
    "mongodb-atlas-cli_${version}_linux_arm64.deb"
    "mongodb-atlas-cli_${version}_linux_arm64.deb.sig"
    "mongodb-atlas-cli_${version}_linux_arm64.rpm"
    "mongodb-atlas-cli_${version}_linux_arm64.rpm.sig"
    "mongodb-atlas-cli_${version}_linux_arm64.tar.gz"
    "mongodb-atlas-cli_${version}_linux_arm64.tar.gz.sig"
    "mongodb-atlas-cli_${version}_linux_x86_64.deb"
    "mongodb-atlas-cli_${version}_linux_x86_64.deb.sig"
    "mongodb-atlas-cli_${version}_linux_x86_64.rpm"
    "mongodb-atlas-cli_${version}_linux_x86_64.rpm.sig"
    "mongodb-atlas-cli_${version}_linux_x86_64.tar.gz"
    "mongodb-atlas-cli_${version}_linux_x86_64.tar.gz.sig"
    "mongodb-atlas-cli_${version}_macos_arm64.zip"
    "mongodb-atlas-cli_${version}_macos_x86_64.zip"
    "mongodb-atlas-cli_${version}_windows_x86_64.msi"
    "mongodb-atlas-cli_${version}_windows_x86_64.zip"
    "sbom.json"
)

for file in "${REQUIRED_FILES[@]}"; do
    if [[ ! -f "dist/${file}" ]]; then
        echo "dist/${file} is missing"
        exit 1
    fi
done
