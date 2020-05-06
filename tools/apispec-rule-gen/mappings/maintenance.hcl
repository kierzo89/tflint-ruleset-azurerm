mapping "azurerm_maintenance_configuration" {
  import_path = "azure-rest-api-specs/specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2020-04-01/Maintenance.json"

  location = MaintenanceConfiguration.location
  scope    = MaintenanceConfigurationProperties.maintenanceScope
}
