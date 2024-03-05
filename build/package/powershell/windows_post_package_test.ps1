# Install current version of MongoDB Atlas CLI
$NewestVersion = git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2
$PackageName = "mongodb-atlas-cli_${NewestVersion}_windows_x86_64.msi"
if("${Env:UNSTABLE}" -eq "-unstable") {
    $PackageName = "mongodb-atlas-cli_${NewestVersion}-next_windows_x86_64.msi"
}
$Source = "https://mongodb-mongocli-build.s3.amazonaws.com/${Env:PROJECT}/dist/${Env:REVISION}_${Env:CREATED_AT}/${PackageName}"
Invoke-WebRequest -Uri $Source -OutFile $PackageName | Out-Null
Start-Process -Wait -FilePath msiexec -ArgumentList /i, $PackageName, /quiet, /norestart -Verb RunAs

# Set missing %APPDATA% environment variable
cd "C:\Program Files (x86)\MongoDB Atlas CLI"
$Env:APPDATA = "C:\Users\Administrator\AppData\"

# Get Atlas CLI version
$InstalledVersion = .\atlas --version
$InstalledVersion = ($InstalledVersion.Split([Environment]::NewLine) | Select -First 1)

# Uninstall MongoDB Atlas CLI and remove environment variable
$App = Get-WmiObject -Class Win32_Product | Where-Object{$_.Name -eq "MongoDB Atlas CLI"}
$App.Uninstall()
$Env:APPDATA = $null

# Check MongoDB Atlas CLI version
if (-Not ($InstalledVersion -Match $NewestVersion)) {
    throw "MSI test unsuccessful: version doesn't match the newest one."
}
Write-Host "Test successful: MSI installed and uninstalled properly."
