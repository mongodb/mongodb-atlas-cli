Write-Output 'Install gotestsum...'
$GOPATH = go env GOPATH
[Environment]::SetEnvironmentVariable("PATH", "$GOPATH\bin;$env:PATH", "User")
go install gotest.tools/gotestsum@latest

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
