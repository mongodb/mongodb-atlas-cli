#!/usr/bin/env bash
# Copyright 2020 MongoDB Inc
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

GOROOT="$(cygpath --mixed "c:\\golang\\go1.24")"
PATH="$(cygpath --mixed "c:\\golang\\go1.24\\bin"):${PATH}"
GOCACHE="off"
CGO_ENABLED=0

export GOROOT PATH GOCACHE CGO_ENABLED

choco install -y "go-msi" --force

go-msi make --path "wix.json" --msi "out.msi" --version "$(cat version.txt)"

choco uninstall -y "go-msi" --force
