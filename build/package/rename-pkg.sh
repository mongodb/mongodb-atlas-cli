#!/usr/bin/env bash

set -Eeou pipefail

VERSION=$(git describe --abbrev=0 | cut -d "v" -f 2)
FILENAME=mongocli_"${VERSION}"_linux_x86_64
if [[ "${unstable}" == "-unstable" ]]; then
    FILENAME="mongocli_next_linux_x86_64"
fi

cd dist

mkdir yum apt

# we could generate a similar name with goreleaser but we want to keep the vars evg compatible to use later
cp "$FILENAME.deb" apt/
mv "apt/$FILENAME.deb" "apt/mongodb-cli${unstable}_${VERSION}${latest_deb}_amd64.deb"
cp "$FILENAME.rpm" yum/
mv "yum/$FILENAME.rpm" "yum/mongodb-cli${unstable}-${VERSION}${latest_rpm}.x86_64.rpm"
