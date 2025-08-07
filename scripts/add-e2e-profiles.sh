#!/bin/bash

set -euo pipefail

./bin/atlas config delete __e2e --force >/dev/null 2>&1 || true
./bin/atlas config init -P __e2e
./bin/atlas config set output plaintext -P __e2e
./bin/atlas config set telemetry_enabled false -P __e2e

./bin/atlas config delete __e2e_snapshot --force >/dev/null 2>&1 || true

CONFIG_PATH=$(EDITOR=echo atlas config edit 2>/dev/null)

cat <<EOF >> "$CONFIG_PATH"

[__e2e_snapshot]
org_id = 'a0123456789abcdef012345a'
project_id = 'b0123456789abcdef012345b'
public_api_key = 'ABCDEF01'
private_api_key = '12345678-abcd-ef01-2345-6789abcdef01'
ops_manager_url = 'http://localhost:8080/'
service = 'cloud'
telemetry_enabled = false
output = 'plaintext'
EOF

echo "Added e2e profiles to $CONFIG_PATH"

