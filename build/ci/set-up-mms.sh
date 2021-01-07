#!/bin/bash

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

set -euo pipefail

while getopts 'i:h:g:u:a:b:' opt; do
  case ${opt} in
  h) hostsFile="${OPTARG}" ;; # Output of Evergreen host.list
  *) exit 1 ;;
  esac
done

hosts=$(
  cat <<EOF | python - "${hostsFile}"
import sys
import json
with open(sys.argv[1]) as hostsfile:
    hosts = json.load(hostsfile)
    for host in hosts:
        print(host["dns_name"])
EOF

)

cd ..
cd ..

for host in ${hosts}; do
  echo "set base_urs"
  ./bin/mongocli config set base_url "http://ec2-63-35-176-208.eu-west-1.compute.amazonaws.com:9080/"
done

echo "create first user"
./bin/mongocli om owner create --firstName evergreen --lastName evergreen --email evergreenTest@gmail.com --password "evergreen1234_" -o json > apikeys.json

cat apikeys.json
export PUBLIC_KEY=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["publicKey"])
EOF

)

export PRIVATE_KEY=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["privateKey"])
EOF

)

echo "set public_api_key"
./bin/mongocli config set public_api_key "${PUBLIC_KEY}"

echo "set private_api_key"
./bin/mongocli config set private_api_key "${PRIVATE_KEY}"

echo "set service"
./bin/mongocli config set service ops-manager

echo "create organization"
./bin/mongocli iam organizations create myOrg -o json > organization.json

cat organization.json
export ORGANIZATION_ID=$(
  cat <<EOF | python - organization.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    org = json.load(jsonfile)
    print(org["id"])
EOF

)

echo "create project"
./bin/mongocli iam projects create myProj -o json > project.json

cat project.json
export PROJECT_ID=$(
  cat <<EOF | python - project.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    proj = json.load(jsonfile)
    print(proj["id"])
EOF

)
