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

export GOCACHE="$(cygpath --mixed "${workdir}\.gocache")"
export CGO_ENABLED=0

go-msi check-env

VERSION=$(git describe | cut -d "v" -f 2)

env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X github.com/mongodb/mongocli/internal/version.Version=${VERSION}" -o ./bin/mongocli.exe

go-msi make --msi "dist/mongocli_${VERSION}_windows_x86_64.msi" --version ${VERSION}
