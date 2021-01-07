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

# shellcheck disable=SC2034
while getopts 's:c:t:x:h:' opt; do
  case ${opt} in
  s) MCLI_SERVICE="${OPTARG}";;
  c) TEST_CMD="${OPTARG}";;
  t) E2E_TAGS="${OPTARG}";;
  x) XDG_CONFIG_HOME="${OPTARG}";;
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

export MCLI_OPS_MANAGER_URL="http://${host}:9080/"

cd ..
cd ..

cat apikeys.json
MCLI_PUBLIC_API_KEY=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["publicKey"])
EOF

)
export MCLI_PUBLIC_API_KEY

MCLI_PRIVATE_API_KEY=$(
  cat <<EOF | python - apikeys.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    user = json.load(jsonfile)
    print(user["programmaticApiKey"]["privateKey"])
EOF

)
export MCLI_PRIVATE_API_KEY


cat organization.json
MCLI_ORG_ID=$(
  cat <<EOF | python - organization.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    org = json.load(jsonfile)
    print(org["id"])
EOF

)
export MCLI_ORG_ID

cat project.json
MCLI_PROJECT_ID=$(
  cat <<EOF | python - project.json
import sys
import json
with open(sys.argv[1]) as jsonfile:
    proj = json.load(jsonfile)
    print(proj["results"][0]["id"])
EOF

)
export MCLI_PROJECT_ID

echo "run e2e test"

make e2e-test