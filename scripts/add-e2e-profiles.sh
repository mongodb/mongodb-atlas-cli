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

./bin/atlas config delete __e2e --force >/dev/null 2>&1 || true

# Prompt if user wants to use cloud-dev.mongodb.com 
read -p "Do you want to set ops_manager_url to cloud-dev.mongodb.com? [Y/n] " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    ops_manager_url="https://cloud-dev.mongodb.com/"
else
    ops_manager_url="https://cloud.mongodb.com/" # Default to cloud.mongodb.com
fi

./bin/atlas config set ops_manager_url $ops_manager_url -P __e2e
./bin/atlas config init -P __e2e
./bin/atlas config set output plaintext -P __e2e
./bin/atlas config set telemetry_enabled false -P __e2e

./bin/atlas config delete __e2e_snapshot --force >/dev/null 2>&1 || true

export EDITOR=echo
CONFIG_PATH=$(./bin/atlas config edit 2>/dev/null)

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

