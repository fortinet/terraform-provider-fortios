---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_address"
sidebar_current: "docs-fortios-resource-firewall-object-address"
description: |-
  Provides a resource to configure firewall addresses used in firewall policies of FortiOS.
---

# fortios_firewall_object_address
Provides a resource to configure firewall addresses used in firewall policies of FortiOS.

## Example Usage for Iprange Address
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_address" "s1" {
	name = "s1"
	type = "iprange"
	start_ip = "1.0.0.0"
	end_ip = "2.0.0.0"
	comment = "dd"
}
```

## Example Usage for Geography Address
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_address" "s2" {
	name = "s2"
	type = "geography"
	country = "AO"
	comment = "dd"
}
```

## Example Usage for Fqdn Address
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_address" "s3" {
	name = "s3"
	type = "fqdn"
	fqdn = "baid.com"
	comment = "dd"
}
```

## Example Usage for Ipmask Address
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_address" "s4" {
	name = "s4"
	type = "ipmask"
	subnet = "0.0.0.0 0.0.0.0"
	comment = "dd"
}
```

## Argument Reference
The following arguments are supported:
* `name` - (Required) Address name.
* `type` - (Required) Type of address(Support ipmask, iprange, fqdn and geography).
* `subnet` - IP address and subnet mask of address.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `fqdn` - Fully Qualified Domain Name address.
* `country` - IP addresses associated to a specific country.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the address item.
* `name` - Address name.
* `type` - Type of address(Support ipmask, iprange, fqdn and geography).
* `subnet` - IP address and subnet mask of address.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `fqdn` - Fully Qualified Domain Name address.
* `country` - IP addresses associated to a specific country.
* `comment` - Comment.
