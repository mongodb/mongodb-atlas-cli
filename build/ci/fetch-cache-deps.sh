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

SHA=$(shasum -a256 go.sum | awk '{print $1}')
FILE="gomod-${SHA}.tgz"
CACHE_URL="https://s3.amazonaws.com/mongodb-mongocli-build/dependencies/go/${FILE}"
echo "${CACHE_URL}"
mkdir -p "${GOMODCACHE}"
pushd "${GOMODCACHE}"

echo "start"
ls -alfh
set +e # ignore get errors
curl -sLf "${CACHE_URL}" -o "${FILE}"
set -e
if [[ -f "${FILE}" ]]; then
  echo "file exists start"
  tar zx "${FILE}"
  ls -alfh
  rm "${FILE}"
fi

popd
cat <<EOF >"sha_expansion.yaml"
dep_sha: "${SHA}"
EOF

echo "update expansions"
pwd
ls -alfh
cat sha_expansion.yaml