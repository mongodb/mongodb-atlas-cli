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

# Notarize generated binaries with Apple and replace the original binary with the notarized one
# This depends on binaries being generated in a goreleaser manner and gon being set up.
# goreleaser should already take care of calling this script as a hook.

# this script could run in parallel for both x86_64 and arm64
# we need to make sure to call the right one at the right time

curl "${NOTARY_SERVICE_URL}" --output macos-notary.zip
unzip -u macos-notary.zip
chmod 755 ./darwin_amd64/macnotary

if [[ -f "./dist/macos_darwin_amd64/bin/mongocli" && ! -f "./dist/mongocli_macos_signed_x86_64.zip" ]]; then
  echo "notarizing x86_64"
  zip -rj ./dist/macos_darwin_amd64/mongocli_amd64_bin.zip ./dist/macos_darwin_amd64/bin/mongocli # The Notarization Service takes an archive as input
  ./darwin_amd64/macnotary \
    -f ./dist/macos_darwin_amd64/mongocli_amd64_bin.zip \
    -m notarizeAndSign -u https://dev.macos-notary.build.10gen.cc/api \
    -s "${NOTARY_SERVICE_SECRET:?}" \
    -k "${NOTARY_SERVICE_KEY_ID:?}"  \
    -b com.mongodb.mongocli \
    -o ./dist/mongocli_macos_signed_x86_64.zip

  echo "replacing original file"
  unzip -od ./dist/macos_darwin_amd64/bin/ ./dist/mongocli_macos_signed_x86_64.zip
fi

if [[ -f "./dist/macos_darwin_arm64/bin/mongocli" && ! -f "./dist/mongocli_macos_signed_arm64.zip" ]]; then
  echo "notarizing arm64"
  zip -rj ./dist/macos_darwin_arm64/mongocli_arm64_bin.zip ./dist/macos_darwin_arm64/bin/mongocli # The Notarization Service takes an archive as input
  ./darwin_amd64/macnotary \
    -f ./dist/macos_darwin_arm64/mongocli_arm64_bin.zip \
    -m notarizeAndSign -u https://dev.macos-notary.build.10gen.cc/api \
    -s "${NOTARY_SERVICE_SECRET:?}" \
    -k "${NOTARY_SERVICE_KEY_ID:?}"  \
    -b com.mongodb.mongocli \
    -o ./dist/mongocli_macos_signed_arm64.zip

  echo "replacing original file"
  unzip -od ./dist/macos_darwin_arm64/bin/ ./dist/mongocli_macos_signed_arm64.zip
fi
