package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectAddressGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectAddressGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectAddressGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectAddressGroupExists("fortios_firewall_object_addressgroup.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_addressgroup.test1", "name", "s1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_addressgroup.test1", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectAddressGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object AddressGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object AddressGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectAddressGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object AddressGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object AddressGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectAddressGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_addressgroup" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectAddressGroup(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object AddressGroup %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSFirewallObjectAddressGroupConfig = `
resource "fortios_firewall_object_addressgroup" "test1" {
	name = "s1"
	member = ["google-play","swscan.apple.com"]
	comment = "Terraform Test"
}
`
