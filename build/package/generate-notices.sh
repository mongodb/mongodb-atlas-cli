#!/usr/bin/env bash

set -Eeou pipefail

export GOPATH="${workdir:?}"
export PATH="$GOPATH/bin:$PATH"

pushd "${GOPATH}"

go get github.com/google/go-licenses

popd

go-licenses save "github.com/mongodb/mongocli" --save_path=third_party_notices
