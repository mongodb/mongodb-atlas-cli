Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
cd $env:HOME
git clone "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
git checkout CLOUDP-265394
$env:TEST_CMD="go run gotest.tools/gotestsum@latest --junitfile e2e-tests.xml --format standard-verbose --"
$env:E2E_TAGS="atlas,deployments,local"
make e2e-test
