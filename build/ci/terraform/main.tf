provider "azurerm" {
  features {}
}

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.19.0"
    }
  }
  backend "azurerm" {
    resource_group_name  = "atlascli-image-resources"
    storage_account_name = "atlascliterraform"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

variable "image_id" {
  type = string
  default = "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/atlascli-image-resources/providers/Microsoft.Compute/images/atlascli-win11-image"
}

variable "certificate_path" {
  type = string
  default = "~/.ssh/id_rsa.pub"
}

resource "azurerm_resource_group" "atlascli_vm_rg" {
  name     = "atlascli-resources"
  location = "East US"
}

resource "azurerm_virtual_network" "atlascli_vm_vnet" {
  name                = "atlascli-network"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.atlascli_vm_rg.location
  resource_group_name = azurerm_resource_group.atlascli_vm_rg.name
}

resource "azurerm_subnet" "atlascli_vm_subnet" {
  name                 = "atlascli-subnet"
  resource_group_name  = azurerm_resource_group.atlascli_vm_rg.name
  virtual_network_name = azurerm_virtual_network.atlascli_vm_vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_network_interface" "atlascli_vm_nic" {
  name                = "atlascli-nic"
  location            = azurerm_resource_group.atlascli_vm_rg.location
  resource_group_name = azurerm_resource_group.atlascli_vm_rg.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.atlascli_vm_subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.atlascli_vm_pip.id
  }
}

resource "azurerm_network_security_group" "atlascli_vm_nsg" {
  name                = "atlascli-nsg"
  location            = azurerm_resource_group.atlascli_vm_rg.location
  resource_group_name = azurerm_resource_group.atlascli_vm_rg.name

  security_rule {
    name                       = "RDP"
    priority                   = 1001
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "3389"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "SSH"
    priority                   = 1002
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_network_interface_security_group_association" "atlascli_vm_nic_nsg_assoc" {
  network_interface_id      = azurerm_network_interface.atlascli_vm_nic.id
  network_security_group_id = azurerm_network_security_group.atlascli_vm_nsg.id
}

resource "azurerm_public_ip" "atlascli_vm_pip" {
  name                = "atlascli-pip"
  location            = azurerm_resource_group.atlascli_vm_rg.location
  resource_group_name = azurerm_resource_group.atlascli_vm_rg.name
  allocation_method   = "Static"
}

resource "azurerm_windows_virtual_machine" "atlascli_vm" {
  name                  = "atlascli-vm"
  location              = azurerm_resource_group.atlascli_vm_rg.location
  resource_group_name   = azurerm_resource_group.atlascli_vm_rg.name
  size                  = "Standard_D2s_v3"
  admin_username        = "atlascli"
  admin_password        = "P@ssw0rd1234!"
  network_interface_ids = [azurerm_network_interface.atlascli_vm_nic.id]
  computer_name         = "atlasclivm"
  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }
  source_image_id = var.image_id
}

resource "azurerm_virtual_machine_extension" "atlascli_vm_extension" {
  name                        = "atlascli_vm_extension"
  virtual_machine_id          = azurerm_windows_virtual_machine.atlascli_vm.id
  publisher                   = "Microsoft.Compute"
  type                        = "CustomScriptExtension"
  type_handler_version        = "1.10"
  auto_upgrade_minor_version  = true

  settings = <<SETTINGS
 {
  "commandToExecute": "powershell.exe -Command \"$keyPath = $env:ProgramData + '\\ssh\\administrators_authorized_keys'; Add-Content -Force -Path $keyPath -Value '${local.ssh_pub_key}'; icacls.exe $keyPath /inheritance:r /grant 'Administrators:F' /grant 'SYSTEM:F'\""
 }
SETTINGS
}

locals {
  ssh_pub_key = trimspace(file(var.certificate_path))
}

data "azurerm_public_ip" "ip" {
  name = azurerm_public_ip.atlascli_vm_pip.name
  resource_group_name = azurerm_resource_group.atlascli_vm_rg.name
  depends_on = [azurerm_windows_virtual_machine.atlascli_vm]
}

output "public_ip_address" {
  value      = data.azurerm_public_ip.ip.ip_address
}

output "key" {
  value      = local.ssh_pub_key
}
