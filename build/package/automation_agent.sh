#!/bin/bash

set -ex

mmsGroupId=$1
mmsApiKey=$2

echo "Installing dependeces"
sudo apt-get -y install libcurl3 libgssapi-krb5-2 \
     libkrb5-dbg libldap-2.4-2 libpci3 libsasl2-2 snmp \
     liblzma5 openssl

clear

echo "Downloadind and extracting the automation agent"
curl -OL https://cloud-dev.mongodb.com/download/agent/automation/mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb
sudo dpkg -i mongodb-mms-automation-agent-manager_10.15.0.6409-1_amd64.ubuntu1604.deb

echo "Replacing mmsGroupId and mmsApiKey properties"
sudo sed -i "s/\(mmsGroupId *= *\).*/\1$mmsGroupId/" /etc/mongodb-mms/automation-agent.config
sudo sed -i "s/\(mmsApiKey *= *\).*/\1$mmsApiKey/" /etc/mongodb-mms/automation-agent.config

echo "Preparing the /data directory to store your MongoDB data"
sudo mkdir -p /data
sudo chown mongod:mongod /data


echo "Starting the agent"
sudo systemctl start mongodb-mms-automation-agent.service
