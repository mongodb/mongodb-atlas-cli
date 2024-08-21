Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
$goPath = go env GOPATH
$env:Path += ";$goPath\bin" 
go install gotest.tools/gotestsum@latest
cd $env:HOME
git clone --depth 1 --branch master "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
gotestsum --junitfile e2e-tests.xml --format standard-verbose -- -timeout 1h -tags e2e -v -run ^TestDeploymentsLocal*$
