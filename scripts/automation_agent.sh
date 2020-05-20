#!/bin/bash

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
echo "Installing dependeces"
sudo yum install cyrus-sasl cyrus-sasl-gssapi \
     cyrus-sasl-plain krb5-libs libcurl \
     lm_sensors-libs net-snmp net-snmp-agent-libs \
     openldap openssl tcp_wrappers-libs xz-libs

echo "Download and extract the automation agent"
curl -OL https://cloud-dev.mongodb.com/download/agent/automation/mongodb-mms-automation-agent-manager-10.15.0.6396-1.x86_64.rhel7.rpm
sudo rpm -U mongodb-mms-automation-agent-manager-10.15.0.6396-1.x86_64.rhel7.rpm

echo "Replacing mmsGroupId and mmsApiKey properties"
replace_property_in_file "/etc/mongodb-mms/automation-agent.config" "mmsGroupId" ""
replace_property_in_file "/etc/mongodb-mms/automation-agent.config" "mmsApiKey" ""

echo "Preparing the /data directory to store your MongoDB data. This directory must be owned by the mongod user"
sudo mkdir -p /data
sudo chown mongod:mongod /data

echo "Starting the agent"
sudo systemctl start mongodb-mms-automation-agent.service


