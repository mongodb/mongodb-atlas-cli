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

# Notarize generated binaries with Apple and replace the original binary with the notarized one
# This depends on binaries being generated in a goreleaser manner and gon being set up.
# goreleaser should already take care of calling this script as a hook.

FILE="./dist/windows_windows_amd64_v1/bin/mongocli.exe"

if [[ -f "$FILE" ]]; then
  echo "notarizing windows binaries"
  notary-client.py \
    --key-name "$NOTARY_SIGNING_KEY_MONGOCLI" \
    --comment "$NOTARY_SIGNING_COMMENT" \
    --auth-token "$NOTARY_AUTH_TOKEN" \
    --notary-url "$NOTARY_URL" \
    $FILE
  ls -la
fi
