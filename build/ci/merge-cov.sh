#!/usr/bin/env bash

# Copyright 2024 MongoDB Inc
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

dirs=""
for f in cov/*.tgz; do
    dir="${f/.coverage.tgz/}"
    dirs="$dirs,$dir"
    mkdir -p "$dir"
    tar -xzvf "$f" -C "$dir"
done
dirs=${dirs:1}
mkdir -p cov/merged
go tool covdata merge -i "$dirs" -o cov/merged
go tool covdata textfmt -i cov/merged -o coverage.out
