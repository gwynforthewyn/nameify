data "azurerm_resource_group" "{{ .ResourceGroup }}" {
  name     = "{{ .ResourceGroup }}"
  location = "{{ .Location }}"
}

data "azurerm_private_dns_zone" "" {
  name                = "mydomain.com"
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_private_dns_a_record" "example" {
  name                = "test"
  zone_name           = azurerm_private_dns_zone.example.name
  resource_group_name = azurerm_resource_group.example.name
  ttl                 = 300
  records             = ["10.0.180.17"]
}
