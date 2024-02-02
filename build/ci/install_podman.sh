#!/bin/bash

# Copyright 2024 MongoDB Inc
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

set -Eeou pipefail
if [ -n "$(which yum 2>/dev/null)" ]; then
  sudo yum update
  sudo yum install -y podman
elif [ -n "$(which apt-get 2>/dev/null)" ]; then
  sudo apt-get update
  sudo apt-get install -y podman
elif [ -n "$(which zypper 2>/dev/null)" ]; then
  sudo zypper install -y podman
elif [ -n "$(which brew 2>/dev/null)" ]; then
  sudo brew install podman
fi

podman --version
