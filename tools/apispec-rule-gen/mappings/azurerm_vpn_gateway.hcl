mapping "azurerm_vpn_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/virtualWan.json"

  scale_unit = VpnGatewayProperties.vpnGatewayScaleUnit
}
