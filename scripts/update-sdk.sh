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

set -euo pipefail

LATEST_SDK_RELEASE=$(curl -sSfL -X GET  https://api.github.com/repos/mongodb/atlas-sdk-go/releases/latest | jq -r '.tag_name' | cut -d '.' -f 1)
echo  "==> Updating SDK to latest major version $LATEST_SDK_RELEASE"
gomajor get "go.mongodb.org/atlas-sdk/$LATEST_SDK_RELEASE@latest"
go mod tidy
sed -i -r "s|go.mongodb.org/atlas-sdk/v[0-9]*|go.mongodb.org/atlas-sdk/$LATEST_SDK_RELEASE|" build/ci/library_owners.json
echo "Done"
