#!/bin/bash
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




