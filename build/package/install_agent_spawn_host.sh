#!/bin/bash

set -ex

while getopts 'i:h:t:u:' opt; do
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
    ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "$user@$host" \  < src/github.com/mongodb/mongocli/build/package/automation_agent.sh "$groupid" "$apiKey"
    set -e
done