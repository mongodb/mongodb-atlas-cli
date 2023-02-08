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

export PATH="$ADD_PATH:$PATH"
mkdir golangci-lint-cache
# don't use the user home cache since running on evergreen
GOLANGCI_LINT_CACHE="$(pwd)/golangci-lint-cache"
export GOLANGCI_LINT_CACHE

golangci-lint run --out-format junit-xml >lint-tests.xml
