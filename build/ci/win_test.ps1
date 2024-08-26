Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
cd $env:HOME
git clone "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
git checkout CLOUDP-265394
go install gotest.tools/gotestsum@latest
$env:TEST_CMD="gotestsum --junitfile e2e-tests.xml --format standard-verbose --"
$env:E2E_TAGS="atlas,deployments,local,auth,noauth,nocli"
echo '' | make e2e-test
