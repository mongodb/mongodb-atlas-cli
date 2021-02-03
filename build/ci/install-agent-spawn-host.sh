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

_print_usage() {
    echo
    echo '  -i <keyfile>                SSH identity file'
    echo '  -u <user>                   Username on the remote host'
    echo '  -h <hostsFile>              Output of Evergreen host.list'
}

while getopts 'i:h:g:u:a:b:' opt; do
  case ${opt} in
  i) keyfile="${OPTARG}" ;; # SSH identity file
  u) user="${OPTARG}" ;; # Username on the remote host
  h) hostsFile="${OPTARG}" ;; # Output of Evergreen host.list
  *) echo "invalid option for install-agent-spawn-host $1" ; _print_usage "$@" ; exit 1 ;;
  esac
done


flags=()
# get the information that you need from the toml file if exists
if [[ -f "${XDG_CONFIG_HOME}/mongocli.toml" ]]; then
  IFS="="
  while read -r name value
  do
    name=${name//[[:blank:]]/}
    if [[ "${name}" == "project_id" ]]; then
      LC_GROUP_ID=$(echo "${value}" | tr -d '"')
      export LC_GROUP_ID
    fi

    if [[ "${name}" == "agent_api_key" ]]; then
      LC_AGENT_KEY=$(echo "${value}" | tr -d '"')
      export LC_AGENT_KEY
    fi

    if [[ "${name}" == "ops_manager_url" ]]; then
      BASE_URL=$(echo "${value}" | tr -d '"')
      export BASE_URL
    fi
  done < "${XDG_CONFIG_HOME}/mongocli.toml"

  flags+=("--ops-manager")
fi

flags+=("--baseUrl ${BASE_URL}")

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
    echo "Seeding ${host}"
    ./ego seed "${user}@${host}"

    echo "bin/ego scenario_install_agent"
    ./ego run "${user}@${host}" bin/ego scenario_install_agent "${flags[@]}"

done
