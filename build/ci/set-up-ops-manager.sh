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

cd ..
cd ..


echo "set service"
  ./bin/mongocli config set service ops-manager

echo "set ops_manager_url"
  ./bin/mongocli config set ops_manager_url "http://${host}:9080/"

echo "create first user"
./bin/mongocli om owner create --firstName evergreen --lastName evergreen --email evergreenTest@gmail.com --password "evergreen1234_" -o json > apikeys.json

cat apikeys.json
publicKey=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["publicKey"])
EOF

)

privateKey=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["privateKey"])
EOF

)

echo "set public_api_key"
./bin/mongocli config set public_api_key "${publicKey}"

echo "set private_api_key"
./bin/mongocli config set private_api_key "${privateKey}"

echo "create organization"
./bin/mongocli iam organizations create myOrg -o json > organization.json

cat organization.json
organizationID=$(
  cat <<EOF | python - organization.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    org = json.load(jsonfile)
    print(org["id"])
EOF

)

# This mongocli command returns an error when the user has been created with "mongocli om owner create" but the project is created anyway
# More info: https://jira.mongodb.org/browse/CLOUDP-76824
set +e
echo "create project"
./bin/mongocli iam projects create myProj --orgId "${organizationID}" -o json

set -e
echo "get project id"
./bin/mongocli iam project list --orgId "${organizationID}" -o json > project.json
cat project.json