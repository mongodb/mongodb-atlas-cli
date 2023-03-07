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

# shellcheck disable=SC2154
if [[ "${go_dep_cache_exists}" == "yes" ]]; then
  exit 0
fi
# shellcheck disable=SC2154
FILE="gomod-${go_dep_sha}.tar.xz"

pushd "${GOPATH}"

export XZ_OPT=-9
tar -Jcf "${FILE}" "pkg/mod"
popd

mv "${GOPATH}/${FILE}" .
