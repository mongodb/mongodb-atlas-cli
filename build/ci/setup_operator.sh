#!/bin/bash

# Copyright 2021 MongoDB Inc
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

set -o errexit

workdir=$(pwd)
os=$(uname | tr '[:upper:]' '[:lower:]')
arch=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n "$(uname -m)" ;; esac)

if ! [ -x "$(command -v "$workdir"/bin/kind)" ]; then
    echo "Kind is not installed. Installing Kind"

    kind_version="v0.17.0"

    mkdir -p "${workdir:?}/bin/"

    echo "Downloading Kind binary"
    curl --retry 3 --silent -L "https://github.com/kubernetes-sigs/kind/releases/download/${kind_version}/kind-${os}-${arch}" -o kind

    echo "Installing Kind in ${workdir}/bin"
    chmod +x kind
    mv kind "${workdir}/bin"
    "${workdir}/bin/kind" version
fi

if ! [ "$("$workdir"/bin/kind get clusters | grep -x kind)" = "kind" ]; then
  echo "Starting Kind cluster"
  "${workdir}/bin/kind" create cluster
fi

if ! [ -x "$(command -v "$workdir"/bin/kubectl)" ]; then
  echo "kubectl is not installed. Installing kubectl"

  echo "Downloading latest kubectl"
  curl -s --retry 3 -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/${os}/${arch}/kubectl"

  echo "Installing kubectl in ${workdir}/bin"
  chmod +x kubectl
  mv kubectl "${workdir}/bin"
  "${workdir}/bin/kubectl" version --client
fi

if [ "$("$workdir"/bin/kubectl get pods --namespace mongodb-atlas-system --selector app.kubernetes.io/instance=mongodb-atlas-kubernetes-operator --output name | wc -l)" = 0 ]; then
  echo "AKO is not installed. Installing AKO"

  echo "Installing AKO in local cluster"
  curl --retry 3 --silent -L https://raw.githubusercontent.com/mongodb/mongodb-atlas-kubernetes/main/deploy/all-in-one.yaml -o all-in-one.yaml
  awk -v url="$MCLI_OPS_MANAGER_URL" '{gsub(/https:\/\/cloud.mongodb.com\//, url, $0); print}' all-in-one.yaml > all.yaml
  "${workdir}/bin/kubectl" apply -f all.yaml
  "${workdir}/bin/kubectl" delete secrets mongodb-atlas-operator-api-key --ignore-not-found --namespace mongodb-atlas-system
  "${workdir}/bin/kubectl" create secret generic mongodb-atlas-operator-api-key --from-literal="orgId=${MCLI_ORG_ID}" --from-literal="publicApiKey=${MCLI_PUBLIC_API_KEY}" --from-literal="privateApiKey=${MCLI_PRIVATE_API_KEY}" --namespace mongodb-atlas-system
  "${workdir}/bin/kubectl" label secret mongodb-atlas-operator-api-key atlas.mongodb.com/type=credentials --namespace mongodb-atlas-system
  rm all-in-one.yaml # all.yaml

  echo "waiting until operator is ready"
  checkCmd="while ! ${workdir}/bin/kubectl --namespace mongodb-atlas-system get pods --selector app.kubernetes.io/instance=mongodb-atlas-kubernetes-operator --output jsonpath={.items[0].status.phase} 2>/dev/null | grep -q Running ; do printf .; sleep 1; done"
  timeout --foreground 1m bash -c "${checkCmd}" || true
  if [ "$("$workdir"/bin/kubectl --namespace mongodb-atlas-system get pods --selector app.kubernetes.io/instance=mongodb-atlas-kubernetes-operator --output jsonpath='{.items[0].status.phase}' | grep -q Running)" = 0 ]; then
      echo "Operator hasn't reached RUNNING state after 1 minute. The full yaml configuration for the pod is:"
      kubectl --namespace mongodb-atlas-system get pods --selector app.kubernetes.io/instance=mongodb-atlas-kubernetes-operator --output yaml#

      echo "Operator failed to start, exiting"
      exit 1
  fi
fi