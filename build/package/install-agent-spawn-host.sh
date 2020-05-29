#!/bin/bash

set -euo pipefail

while getopts 'i:h:g:u:a:b:' opt; do
  case $opt in
  i) keyfile="$OPTARG" ;; # SSH identity file
  u) user="$OPTARG" ;; # Username on the remote host
  h) hostsfile="$OPTARG" ;; # Output of Evergreen host.list
  g) groupid="$OPTARG" ;; # APIGroupId
  a) apiKey="$OPTARG" ;; # APIKey
  b) baseUrl="$OPTARG";; #Cloud Manager URL
  *) exit 1 ;;
  esac
done

declare -a ssh_opts
ssh_opts[0]="-o"
ssh_opts[1]="UserKnownHostsFile=/dev/null"
ssh_opts[2]="-o"
ssh_opts[3]="StrictHostKeyChecking=no"
ssh_opts[4]="-q"
ssh_opts[5]="-o"
ssh_opts[6]="ConnectTimeout=30"
ssh_opts[7]="-o"
ssh_opts[8]="TCPKeepAlive=yes"
ssh_opts[9]="-o"
ssh_opts[10]="ServerAliveInterval=15"
ssh_opts[11]="-o"
ssh_opts[12]="ServerAliveCountMax=20"
ssh_opts[13]="-i"
ssh_opts[14]="$keyfile"

hosts=$(
  cat <<EOF | python - "$hostsfile"
import sys
import json
with open(sys.argv[1]) as hostsfile:
    hosts = json.load(hostsfile)
    for host in hosts:
        print(host["dns_name"])
EOF
)
for host in $hosts; do
  set +e
  echo "Installing dependeces on $host"
  ssh "${ssh_opts[@]}" -tt "$user@$host" 'bash -s' <<'ENDSSH'
        sudo apt-get install -y --no-install-recommends ca-certificates curl logrotate openssl snmp && exit
ENDSSH
  echo "Installing the automation agent on $host"
  ssh "${ssh_opts[@]}" -tt "$user@$host" ARG1="$groupid" ARG2="$apiKey" ARG3="$baseUrl" 'bash -s' <<'ENDSSH'
        echo "Downloadind and extracting the automation agent"
        curl -OL ${ARG3}download/agent/automation/mongodb-mms-automation-agent-manager_latest_amd64.ubuntu1604.deb
        sudo dpkg -i mongodb-mms-automation-agent-manager_latest_amd64.ubuntu1604.deb

        echo "Replacing mmsGroupId and mmsApiKey"
        sudo sed -i "s/\(mmsGroupId *= *\).*/\1$ARG1/" /etc/mongodb-mms/automation-agent.config
        sudo sed -i "s/\(mmsApiKey *= *\).*/\1$ARG2/" /etc/mongodb-mms/automation-agent.config

        echo "Preparing the /data directory to store your MongoDB data"
        sudo mkdir -p /data
        sudo chown mongod:mongod /data

        echo "Starting the agent"
        sudo systemctl start mongodb-mms-automation-agent.service
        exit
ENDSSH
  echo "Storing $host in src/github.com/mongodb/mongocli/e2e/cloud_manager/e2e.env"
  sudo sed -i "s/\(hostname *= *\).*/\1$host/" src/github.com/mongodb/mongocli/e2e/cloud_manager/e2e.env
done
