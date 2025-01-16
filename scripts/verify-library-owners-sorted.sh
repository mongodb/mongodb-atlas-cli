#!/bin/bash
# Copyright 2025 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

SCRIPT_DIR="${BASH_SOURCE%/*}"
LIBRARY_OWNERS_FILE="$SCRIPT_DIR/../build/ci/library_owners.json"


# Extract keys and check if they're sorted
SORTED_KEYS=$(jq -r 'keys | sort_by(ascii_downcase + ".")[]' "$LIBRARY_OWNERS_FILE")
ORIGINAL_KEYS=$(jq -r 'to_entries | map(.key)[]' "$LIBRARY_OWNERS_FILE")

if [ "$SORTED_KEYS" = "$ORIGINAL_KEYS" ]; then
    echo "library_owners are sorted"
    exit 0
else
    echo "library_owners are NOT sorted"
    exit 1
fi
