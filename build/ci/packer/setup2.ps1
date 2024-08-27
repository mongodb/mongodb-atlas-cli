Write-Output 'Env...'
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v GOROOT /t REG_EXPAND_SZ /d $(go env GOROOT) /f
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v GOBIN /t REG_EXPAND_SZ /d "%GOROOT%\bin" /f
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v GOPATH /t REG_EXPAND_SZ /d "%USERPROFILE%\go" /f
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v CGO_ENABLED /t REG_EXPAND_SZ /d "0" /f

$reg_value = reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v PATH
$reg_value_lines = $reg_value -split "\n"
$reg_value_lines[2] -match "PATH\s+REG_EXPAND_SZ\s+(.+)"
$value = $matches[1]
$value += ";%GOBIN%;%GOPATH%\bin"

reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v PATH /t REG_EXPAND_SZ /d $value /f

Write-Output 'Install gotestsum...'
$env:GOROOT = go env GOROOT
$env:GOBIN = "$env:GOROOT/bin"
$GOTESTSUM_VER = '1.12.0'
New-Item -Path "gotestsum" -ItemType "Directory"
Set-Location -Path "gotestsum"
Invoke-WebRequest "https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VER}/gotestsum_${GOTESTSUM_VER}_windows_amd64.tar.gz" -OutFile gotestsum.tar.gz
tar -xzf gotestsum.tar.gz
Move-Item -Path "gotestsum.exe" -Destination "$env:GOBIN"
Set-Location -Path ".."
Remove-Item -Recurse -Force -Path "gotestsum"

Write-Output 'Sysprep...'
if( Test-Path $Env:SystemRoot\windows\system32\Sysprep\unattend.xml ){
    rm $Env:SystemRoot\windows\system32\Sysprep\unattend.xml -Force
}
& $Env:SystemRoot\System32\Sysprep\Sysprep.exe /oobe /generalize /quit /quiet /mode:vm
while($true) {
    $imageState = Get-ItemProperty HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Setup\State | Select ImageState
    Write-Output $imageState.ImageState
    if($imageState.ImageState -ne 'IMAGE_STATE_GENERALIZE_RESEAL_TO_OOBE') {
        Start-Sleep -s 10 
    } else { 
        break
    }
}
Write-Output 'Done'
