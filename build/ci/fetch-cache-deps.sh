#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
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

GO_SUM_SHA=$(shasum -a256 go.sum | awk '{print $1}')
FILE="gomod-${GO_SUM_SHA}.tar.xz"
CACHE_URL="https://s3.amazonaws.com/mongodb-mongocli-build/dependencies/go/${FILE}"
mkdir -p "${GOPATH}"
pushd "${GOPATH}"
GO_CACHE_EXISTS="no"

set +e # ignore download errors, assume no cache available
curl -sLf "${CACHE_URL}" -o "${FILE}"
set -e
if [[ -f "${FILE}" ]]; then
  GO_CACHE_EXISTS="yes"
  tar -xf "${FILE}"
  rm "${FILE}"
fi

# used so other tasks can figure out the sha without bash
popd
cat <<EOF >"sha_expansion.yaml"
go_dep_sha: "${GO_SUM_SHA}"
go_dep_cache_exists: "${GO_CACHE_EXISTS}"
EOF
