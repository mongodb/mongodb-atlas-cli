#!/bin/bash

set -ex

while getopts 'i:h:t:u:v:' opt; do
    case $opt in
        i) keyfile="$OPTARG";;      # SSH identity file
        u) user="$OPTARG";;         # Username on the remote host
        h) hostsfile="$OPTARG";;    # Output of Evergreen host.list
        t) groupid="$OPTARG";;      # APIGroupId
        v) apiKey="$OPTARG";;       # APIKey
        *) exit 1;;
    esac
done

hosts=$(cat << EOF | python - "$hostsfile"
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
    echo "installing the automation agent on $host"
    ssh -i "$keyfile" -o ConnectTimeout=50 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "$user@$host" 'bash -s' <<'ENDSSH'
        free -m
        echo "Installing dependeces"
        sudo apt-get install -y --no-install-recommends ca-certificates curl logrotate openssl snmp && exit
ENDSSH
    ssh -i "$keyfile" -o ConnectTimeout=50 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "$user@$host" ARG1="$groupid" ARG2="$apiKey" 'bash -s' <<'ENDSSH'
        free -m
        echo "Downloadind and extracting the automation agent"
        curl -OL https://cloud-dev.mongodb.com/download/agent/automation/mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb
        sudo dpkg -i mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb

        echo "Replacing mmsGroupId and mmsApiKey properties"
        sudo sed -i "s/\(mmsGroupId *= *\).*/\1$ARG1/" /etc/mongodb-mms/automation-agent.config
        sudo sed -i "s/\(mmsApiKey *= *\).*/\1$ARG2/" /etc/mongodb-mms/automation-agent.config

        echo "Preparing the /data directory to store your MongoDB data"
        sudo mkdir -p /data
        sudo chown mongod:mongod /data

        echo "Starting the agent"
        sudo systemctl start mongodb-mms-automation-agent.service
        exit
ENDSSH
    echo "Storing $host in src/github.com/mongodb/mongocli/e2e/cloud_manager/e2e.properties"
    sudo sed -i "s/\(hostname *= *\).*/\1$host/" src/github.com/mongodb/mongocli/e2e/cloud_manager/e2e.env
done