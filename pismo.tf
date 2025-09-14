provider "azurerm" {
  features {}
  subscription_id = "3dc39de8-30fa-4d79-816f-de0b0301f19c"
}

resource "azurerm_resource_group" "rg" {
  name     = "rg-postgres-pismo"
  location = "South India"
}

resource "azurerm_cosmosdb_postgresql_cluster" "pg_cluster" {
  name                = "postgrescluster"   # must be unique
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location

  administrator_login_password = "H@Sh1CoR3!"
  coordinator_storage_quota_in_mb = 131072        # 128 GB
  coordinator_vcore_count         = 2

  node_count            = 0 # For making it burstable since small use case

  coordinator_public_ip_access_enabled = true
}