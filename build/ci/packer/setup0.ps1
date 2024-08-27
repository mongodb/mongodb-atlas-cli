Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope Process -Force
# defender
Write-Output 'Disable Windows Defender...'
Set-MpPreference -DisableRealtimeMonitoring $true -Force
# ssh
Write-Output 'Install SSH...'
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
Set-Service -Name sshd -StartupType 'Automatic'
Set-Service -Name ssh-agent -StartupType 'Automatic'
New-NetFirewallRule -Name sshd -DisplayName 'OpenSSH Server (sshd)' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
# hyper-v
Write-Output 'Install Hyper-V...'
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
dism.exe /online /enable-feature /featurename:Microsoft-Hyper-V  /all /norestart
# choco
Write-Output 'Install Chocolatey...'
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
