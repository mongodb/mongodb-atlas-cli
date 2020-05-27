#!/bin/bash

set -ex

mmsGroupId=$1
mmsApiKey=$2
replace_property_in_file() {
    # Parameter check
    if [[ "$#" -lt 3 ]]; then
        echo "Invalid call: 'replace_property_in_file $*'"
        echo "Usage: replace_property_in_file FILENAME PROPERTY VALUE"
        echo
        exit 1
    fi

    # Set the new property
    temp_file=$(mktemp)
    grep -vE "^\\s*${2}\\s*=" "${1}" > "${temp_file}" # Export contents minus any lines containing the specified property
    echo "${2}=${3}" >> "${temp_file}"                # Set the new property value
    cat "${temp_file}" > "${1}"                       # Replace the contents of the original file, while preserving any permissions
    rm "${temp_file}"
}

clear

echo "Download and extract the automation agent"
curl -OL https://cloud-dev.mongodb.com/download/agent/automation/mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb
sudo dpkg -i mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb

echo "Replacing mmsGroupId and mmsApiKey properties"
replace_property_in_file "/etc/mongodb-mms/automation-agent.config" "mmsGroupId" "$mmsGroupId"
replace_property_in_file "/etc/mongodb-mms/automation-agent.config" "mmsApiKey" "$mmsApiKey"

echo "Preparing the /data directory to store your MongoDB data. This directory must be owned by the mongod user"
sudo mkdir -p /data
sudo chown mongod:mongod /data

echo echo "$(</etc/mongodb-mms/automation-agent.config )"
echo "Starting the agent"
sudo systemctl start mongodb-mms-automation-agent.service

