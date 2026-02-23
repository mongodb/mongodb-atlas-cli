Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope Process -Force
# ssh
Write-Output 'Configure SSH...'
(Get-Content 'C:/ProgramData/ssh/sshd_config') -replace '#PubkeyAuthentication yes', 'PubkeyAuthentication yes' -replace '#PasswordAuthentication yes', 'PasswordAuthentication no' | Set-Content 'C:/ProgramData/ssh/sshd_config'
# choco
Write-Output 'Install Chocolatey packages...'
choco install golang git.install make mongodb-database-tools -y
choco install docker-desktop -y --install-arguments="--accept-license --backend=wsl-2 --always-run-service"
