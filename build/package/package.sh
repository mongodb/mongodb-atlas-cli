#!/bin/bash

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

set -vx

export GOROOT="${go_root:?}"
export PATH="./bin:${go_bin:?}:$PATH"
export GITHUB_TOKEN="${github_token:?}"
export NOTARY_SERVICE_URL="${notary_service_url:?}"
export MACOS_NOTARY_KEY="${notary_service_key_id:?}"
export MACOS_NOTARY_SECRET="${notary_service_secret:?}"
export GORELEASER_KEY="${goreleaser_key:?}"
export VERSION_GIT?="$(git tag --list "${tool_name:?}/v*" --sort=committerdate | tail -1 | cut -d "v" -f 2)"

if [[ -z "${VERSION_GIT:?}" ]]; then
    VERSION_GIT="$(git describe | cut -d "v" -f 2)"
fi

# avoid race conditions on the notarization step by using `-p 1`
${goreleaser_cmd:?|goreleaser --rm-dist --snapshot -p 1}





