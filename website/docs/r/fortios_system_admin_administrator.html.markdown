---
layout: "fortios"
page_title: "FortiOS: fortios_system_admin_administrator"
sidebar_current: "docs-fortios-resource-system-admin-administrator"
description: |-
  Provides a resource to configure administrator accounts of FortiOS.
---

# fortios_system_admin_administrator
Provides a resource to configure administrator accounts of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_system_admin_administrator" "admintest" {
	name = "testadminacc"
	password = "cc37331AC1"
	trusthost1 = "1.1.1.0 255.255.255.0"
	trusthost2 = "2.2.2.0 255.255.255.0"
	accprofile = "3d3"
	vdom = ["root"]
	comments = "comments"
}
```

## Argument Reference
The following arguments are supported:
* `name` - (Required) User name.
* `password` - (Required) Admin user password.
* `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
* `vdom` - Virtual domain(s) that the administrator can access.
* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the administrator account item.
* `name` - User name.
* `password` - Admin user password.
* `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
* `vdom` - Virtual domain(s) that the administrator can access.
* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `comments` - Comment.

