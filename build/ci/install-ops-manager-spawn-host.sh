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

while getopts 'i:h:g:u:a:b:' opt; do
  case ${opt} in
  i) keyfile="${OPTARG}" ;; # SSH identity file
  u) user="${OPTARG}" ;; # Username on the remote host
  h) hostsFile="${OPTARG}" ;; # Output of Evergreen host.list
  *) exit 1 ;;
  esac
done

# Install ego
curl -sL https://raw.githubusercontent.com/mongodb-labs/ego/master/install.sh | bash
export EGO_DEBUG=1

export SSH_OPTS="-i ${keyfile} -o SendEnv=LC_GROUP_ID -o SendEnv=LC_AGENT_KEY"

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
for host in ${hosts}; do
ssh -i "$keyfile" -o ConnectTimeout=10  -o StrictHostKeyChecking=no -tt "${user}@${host}" 'bash -s' <<'ENDSSH'
  # commands to run on remote host

  #install ego
  curl -sL https://raw.githubusercontent.com/mongodb-labs/ego/master/install.sh | bash

  source ~/.bashrc

  #install mms
  ego ops_manager_install_version --version 4.4.6 --mongodb-version 4.2.8

  exit

ENDSSH
done
