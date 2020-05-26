#!/usr/bin/env bash

set -o errexit

OUTPUT_PATH="./dist/mongocli_macos_signed_x86_64.zip"

# gon settings
cat <<EOF_GON_JSON > gon.json
{
    "source" : ["./dist/macos_darwin_amd64/mongocli"],
    "bundle_id" : "com.mongodb.mongocli",
    "apple_id": {
      "username": "${ac_username}",
      "password": "${ac_password}"
    },
    "sign" :{
      "application_identity" : "Developer ID Application: MongoDB, Inc. (4XWMY46275)"
    },
    "zip" :{
      "output_path": "$OUTPUT_PATH"
    }
}
EOF_GON_JSON
