packer {
  required_plugins {
    azure = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/azure"
    }
    windows-update = {
      version = ">= 0.16.0"
      source  = "github.com/rgl/windows-update"
    }
  }
}

variable "client_id" {
  type    = string
  default = env("AZURE_APP_ID")
}

variable "client_secret" {
  type    = string
  default = env("AZURE_PASSWORD")
}

variable "subscription_id" {
  type    = string
  default = env("AZURE_SUBSCRIPTION_ID")
}

variable "tenant_id" {
  type    = string
  default = env("AZURE_TENANT")
}

source "azure-arm" "windows-11" {
  client_id                         = var.client_id
  client_secret                     = var.client_secret
  subscription_id                   = var.subscription_id
  tenant_id                         = var.tenant_id
  managed_image_resource_group_name = "atlascli-image-resources"
  managed_image_name                = "atlascli-win11-image"
  os_type                           = "Windows"
  image_publisher                   = "MicrosoftWindowsDesktop"
  image_offer                       = "windows-11"
  image_sku                         = "win11-23h2-pro"
  location                          = "East US"
  vm_size                           = "Standard_D2s_v3"
  communicator                      = "winrm"
  winrm_use_ssl                     = true
  winrm_insecure                    = true
  winrm_timeout                     = "5m"
  winrm_username                    = "packer"
  public_ip_sku                     = "Standard"
}

build {
  sources = ["source.azure-arm.windows-11"]

  provisioner "windows-update" {}

  provisioner "powershell" {
    script = "setup0.ps1"
    elevated_user = "SYSTEM"
    elevated_password = ""
  }

  provisioner "windows-restart" {}

  provisioner "powershell" {
    script = "setup1-win11.ps1"
  }

  provisioner "windows-restart" {}

  provisioner "powershell" {
    script = "setup2.ps1"
  }
}
