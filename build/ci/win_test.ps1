<#
    .SYNOPSIS
    Runs the e2e test in atlas cli.

    .PARAMETER GOPROXY
    Specifies the value of GOPROXY environment variable.

    .PARAMETER REVISION
    Specifies the value of REVISION to be fetched by git.

    .PARAMETER DOCKER_USERNAME
    Specifies the value of DOCKER_USERNAME for Docker Hub authentication.

    .PARAMETER DOCKER_PAT
    Specifies the value of DOCKER_PAT for Docker Hub authentication.
#>
param(
  [String]$GOPROXY = "https://proxy.golang.org,direct",
  [String]$REVISION = "HEAD",
  [String]$DOCKER_USERNAME,
  [String]$DOCKER_PAT
)

Write-Output 'Disable Windows Defender...'
Set-MpPreference -DisableRealtimeMonitoring $true -Force

Write-Output 'Setting up docker credentials...'
$credentialString = "${DOCKER_USERNAME}:${DOCKER_PAT}"
$credentialBytes = [System.Text.Encoding]::UTF8.GetBytes($credentialString)
$base64Credentials = [System.Convert]::ToBase64String($credentialBytes)
New-Item -Path "$env:HOME\.docker" -ItemType "Directory" -Force
$dockerConfig = @"
{
  "auths": {
    "https://index.docker.io/v1/": {
      "auth": "$base64Credentials"
    }
  }
}
"@
$dockerConfig | Set-Content -Path "$env:HOME\.docker\config.json" -Encoding ASCII
Write-Output "Docker credentials saved successfully to $env:HOME\.docker\config.json"

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

Write-Output "Waiting for Docker daemon to start..."
$maxAttempts = 120  # Wait up to 120 seconds
$attempt = 0

while ($attempt -lt $maxAttempts) {
    try {
        docker ps | Out-Null
        if ($LASTEXITCODE -eq 0) {
            Write-Output "Docker daemon is ready!"
            break
        }
    }
    catch {
        # Ignore errors while waiting
    }
    
    Start-Sleep -Seconds 1
    $attempt++
    Write-Output "Waiting for Docker daemon to start... $attempt/$maxAttempts"
}

if ($attempt -eq $maxAttempts) {
    Write-Output "`nTimeout: Docker daemon did not start within $maxAttempts seconds"
    exit 1
}
Write-Output "Docker pull"
docker pull "mongodb/mongodb-atlas-local:latest"
Write-Output "Run tests"
$env:TEST_CMD="gotestsum --junitfile e2e-tests.xml --format standard-verbose --"
$env:E2E_TEST_PACKAGES="./test/e2e/atlas/deployments/local/..."
$env:BUILD_FLAGS="-mod vendor -x"
make e2e-test
