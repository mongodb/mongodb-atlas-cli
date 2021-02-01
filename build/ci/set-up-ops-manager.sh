#!/bin/bash

# Copyright 2021 MongoDB Inc
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

while getopts 'h:' opt; do
  case ${opt} in
  h) hostsFile="${OPTARG}" ;; # Output of Evergreen host.list
  *) exit 1 ;;
  esac
done

host=$(
  cat <<EOF | python - "${hostsFile}"
import sys
import json
with open(sys.argv[1]) as hostsfile:
    hosts = json.load(hostsfile)
    print(hosts[0]["dns_name"])
EOF

)

cd ../..

export MCLI_OPS_MANAGER_URL="http://${host}:9080/"

echo "set service"
  ./bin/mongocli config set service "${ops_manager_service:?}"

echo "set ops_manager_url"
  ./bin/mongocli config set ops_manager_url "${MCLI_OPS_MANAGER_URL}"

echo "generate password for owner"
password=$(date +%s | sha256sum | base64 | head -c 8)0_

echo "generate email for owner"
email=$(date +%s | sha256sum | base64 | head -c 8)@ops-manager-team.com

echo "create first user"
MCLI_PRIVATE_API_KEY=$(./bin/mongocli om owner create --firstName evergreen --lastName evergreen --email "${email}" --password "${password}" -o="go-template={{.APIKey}}")

echo "MCLI_PRIVATE_API_KEY= ${MCLI_PRIVATE_API_KEY}"

echo "set public_api_key"
./bin/mongocli config set public_api_key "${email}"

echo "set private_api_key"
./bin/mongocli config set private_api_key "${MCLI_PRIVATE_API_KEY}"

echo "create organization"
MCLI_ORG_ID=$(./bin/mongocli iam organizations create myOrg -o="go-template={{.ID}}")

cat "${XDG_CONFIG_HOME}/mongocli.toml"

echo "create project"
MCLI_PROJECT_ID=$(./bin/mongocli iam projects create myProj --orgId "${MCLI_ORG_ID}" -o="go-template={{.ID}}")


cat <<EOF > "${XDG_CONFIG_HOME}/mongocli.toml"
[default]
  ops_manager_url = "${MCLI_OPS_MANAGER_URL}"
  service = "${ops_manager_service:?}"
  public_api_key = "${email}"
  private_api_key = "${MCLI_PRIVATE_API_KEY}"
  org_id = "${MCLI_ORG_ID}"
  project_id = "${MCLI_PROJECT_ID}"

EOF

echo "print mongocli.toml"
cat "${XDG_CONFIG_HOME}/mongocli.toml"
