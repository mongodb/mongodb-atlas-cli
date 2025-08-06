#!/bin/bash

# --- Atlas E2E Test Environment Variables ---
export MONGODB_ATLAS_ORG_ID="5efda682a3f2ed2e7dd6cde4"
# export MONGODB_ATLAS_PROJECT_ID="68369c0d9806252ca7c427d1"
export MONGODB_ATLAS_PROJECT_ID="5efda6aea3f2ed2e7dd6ce05"
export MONGODB_ATLAS_PUBLIC_API_KEY="nmtxqlkl"
export MONGODB_ATLAS_PRIVATE_API_KEY="4a7d1c60-00e3-4284-845d-ebb4d167beea"
export MONGODB_ATLAS_OPS_MANAGER_URL="https://cloud-dev.mongodb.com"
export MONGODB_ATLAS_SERVICE="cloud"
export DO_NOT_TRACK="1"
export UPDATE_SNAPSHOTS="true"
export E2E_SKIP_CLEANUP="false"
export E2E_CLOUD_ROLE_ID="your_cloud_role_id"
export E2E_TEST_BUCKET="your_test_bucket"
export E2E_FLEX_INSTANCE_NAME="your_flex_instance_name"
export IDENTITY_PROVIDER_ID="your_identity_provider_id"
export AWS_ACCESS_KEY="your_aws_access_key"
export AWS_SECRET_ACCESS_KEY="your_aws_secret_access_key"
export AZURE_TENANT_ID="your_azure_tenant_id"
export AZURE_CLIENT_ID="your_azure_client_id"
export AZURE_CLIENT_SECRET="your_azure_client_secret"
export GCP_CREDENTIALS="your_gcp_credentials"
export E2E_TIMEOUT="3h"
export E2E_TAGS="clusters,atlas,iss"
# export E2E_TAGS="clusters,atlas,file"

# --- Clean old snapshots ---
# rm -rf test/e2e/testdata/.snapshots

# --- Run the E2E tests ---
make e2e-test

echo "E2E snapshot test run complete. Check test/e2e/testdata/.snapshots for results."