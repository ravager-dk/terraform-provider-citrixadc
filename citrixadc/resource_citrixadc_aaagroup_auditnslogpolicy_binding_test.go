/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package citrixadc

import (
	"fmt"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccAaagroup_auditnslogpolicy_binding_basic = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority  = 150
	}
`

const testAccAaagroup_auditnslogpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

func TestAccAaagroup_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAaagroup_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingExist("citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingNotExist("citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", "my_group,tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckAaagroup_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_auditnslogpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditnslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditnslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Aaagroup_auditnslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
