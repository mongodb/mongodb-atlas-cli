#!/bin/bash

# Copyright 2022 MongoDB Inc
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

PYKMIP_HOME="${XDG_CONFIG_HOME}/pykmip"
mkdir "${PYKMIP_HOME}"
echo "${PYKMIP_HOME} folder created"

python3 -m venv "${PYKMIP_HOME}/venv"
# shellcheck source=/dev/null
source "${PYKMIP_HOME}/venv/bin/activate"
echo "${PYKMIP_HOME}/venv created"

pip install --upgrade pip
echo "${PYKMIP_HOME}/venv created"

pip install pykmip
echo "PyKMIP package installed"

echo "${KMIP_CA}" | base64 --decode > "${PYKMIP_HOME}/tls-rootCA.pem"
echo "${KMIP_CERT}" | base64 --decode > "${PYKMIP_HOME}/tls-localhost.pem"
echo "CA and server cert copied to ${PYKMIP_HOME}"

mv kmip.db "${PYKMIP_HOME}/kmip.db"
echo "KMIP db moved to ${PYKMIP_HOME}/kmip.db"

mv start.py "${PYKMIP_HOME}/start.py"
echo "start KMIP script moved to ${PYKMIP_HOME}/start.py"

cat << EOF > "${PYKMIP_HOME}/pykmip.service"
[Unit]
Description=PyKMIP service

[Service]
User=root
Restart=always
ExecStart="${PYKMIP_HOME}/venv/bin/python" "${PYKMIP_HOME}/start.py" "${PYKMIP_HOME}"

[Install]
WantedBy=multi-user.target
EOF
sudo cp "${PYKMIP_HOME}/pykmip.service" "/etc/systemd/system/pykmip.service"
echo "PyKMIP systemd service created"

sudo systemctl daemon-reload
echo "systemctl daemon-reload done"

sudo systemctl start pykmip.service
echo "PyKMIP started"

sleep 5s
sudo systemctl status pykmip.service
sudo netstat -tunlp | grep python
