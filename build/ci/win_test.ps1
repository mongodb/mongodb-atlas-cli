Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
go install gotest.tools/gotestsum@latest
cd $env:HOME
git clone --depth 1 --branch master "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
$GOPATH = go env GOPATH
"$GOPATH\bin\gotestsum.exe" --junitfile e2e-tests.xml --format standard-verbose -- -v -p 1 -timeout 1h -tags="atlas,deployments,local" ./test/e2e...
