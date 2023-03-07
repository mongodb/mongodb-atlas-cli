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
CACHE_URL="https://s3.amazonaws.com/mongodb-mongocli-build/dependencies/go/${SHA}.tgz"
echo "${CACHE_URL}"

pushd "${GOPATH}/pkg/mod"
ls -alfh
curl -sfL "${CACHE_URL}" | tar zx
ls -alfh