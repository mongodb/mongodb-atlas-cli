#!/usr/bin/env bash

# Copyright 2026 MongoDB Inc
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

# Runs from the apix-action Dependabot workflow after a dependency bump.
# Only edit files here: the workflow stages the working tree and creates the
# signed commit itself.

set -Eeou pipefail

# The reusable workflow does not set up Go, so fall back to installing the
# version pinned in go.mod when the runner image lacks a toolchain.
if ! command -v go >/dev/null 2>&1; then
    GO_VERSION=$(sed -n 's/^go //p' go.mod)
    curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | sudo tar -C /usr/local -xzf -
    export PATH="/usr/local/go/bin:$PATH"
fi

make gen-purls
