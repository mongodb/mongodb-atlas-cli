Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
$GOPATH = go env GOPATH
[Environment]::SetEnvironmentVariable("PATH", "$GOPATH\bin;$env:PATH", "User")
go install gotest.tools/gotestsum@latest
cd $env:HOME
git clone --depth 1 --branch master "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
gotestsum --junitfile e2e-tests.xml --format standard-verbose -- -v -p 1 -timeout 1h -tags="atlas,deployments,local" ./test/e2e...
