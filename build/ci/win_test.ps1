<#
    .SYNOPSIS
    Runs the e2e test in atlas cli.

    .PARAMETER GOPROXY
    Specifies the value of GOPROXY environment variable.

    .PARAMETER REVISION
    Specifies the value of REVISION to be fetched by git.
#>
param(
  [String]$GOPROXY = "https://proxy.golang.org,direct",
  [String]$REVISION = "HEAD"
)

Write-Output 'Disable Windows Defender...'
Set-MpPreference -DisableRealtimeMonitoring $true -Force
Write-Output "Start Docker"
Start "C:\Program Files\Docker\Docker\Docker Desktop.exe"
Write-Output "Clone"
cd $env:HOME
git clone "http://github.com/mongodb/mongodb-atlas-cli.git"
cd "mongodb-atlas-cli"
git checkout $REVISION
Write-Output "Vendor dependencies"
$env:GOPROXY=$GOPROXY
tar -xzf ../vendor.tar.gz
Write-Output "Run tests"
$env:TEST_CMD="gotestsum --junitfile e2e-tests.xml --format standard-verbose --"
$env:E2E_TAGS="atlas,deployments,local,auth,noauth,nocli"
$env:BUILD_FLAGS="-mod vendor"
make e2e-test
