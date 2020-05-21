#!/bin/bash
set -ex

hostsfile=$1
connection_attempts=$2

while getopts 'i:h:t:u:' opt; do
    case $opt in
        i) keyfile="$OPTARG";;      # SSH identity file
        u) user="$OPTARG";;         # Username on the remote host
        h) hostsfile="$OPTARG";;    # Output of Evergreen host.list
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

echo "$hosts"
for host in $hosts; do
    set +e

    while ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "root@$host" != 0;
      do
       [ "$attempts" -ge "$connection_attempts" ] && exit 1
          ((attempts++))
          printf "SSH connection attempt %d/%d failed. Retrying...\n" "$attempts" "$connection_attempts"
          # sleep for Permission denied (publickey) errors
          sleep 10
      done
done