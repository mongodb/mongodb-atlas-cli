#!/bin/bash
set -ex

declare -i attempts
declare -i connection_attempts
declare -ri timeout=10

while getopts 'i:h:t:u:' opt; do
    case $opt in
        i) keyfile="$OPTARG";;              # SSH identity file
        u) user="$OPTARG";;                 # Username on the remote host
        h) hostsfile="$OPTARG";;            # Output of Evergreen host.list
        t) connection_attempts="$OPTARG";;  # How many times to attempt to connect via SSH
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

attempts=0
connection_attempts=${connection_attempts:-60} # Total timeout = timeout * timeout_attempts

for host in $hosts; do
    set +e
    echo "Waiting for $host to become available..."
    while ! ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "$user@$host" exit 2> /dev/null; do
        if [ "$attempts" -ge "$connection_attempts" ]; then
            echo 'Connect to spawn host failed'
            exit 1
        fi
        ((attempts++))

        echo "SSH connection attempt $attempts/$connection_attempts failed. Retrying ($host)..."
        # sleep for Permission denied (publickey) errors
        sleep "$timeout"
    done
    set -e
done
