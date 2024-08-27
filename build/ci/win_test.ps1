Write-Output "Start Docker"
Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
Write-Output "Clone"
cd $env:HOME
git clone "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
git checkout CLOUDP-265394
Write-Output "Debug"
$env:GOPROXY=$args[0]
go env
Write-Output "Download dependencies"
go mod download
Write-Output "Run tests"
$env:TEST_CMD="gotestsum --junitfile e2e-tests.xml --format standard-verbose --"
$env:E2E_TAGS="atlas,deployments,local,auth,noauth,nocli"
make e2e-test
