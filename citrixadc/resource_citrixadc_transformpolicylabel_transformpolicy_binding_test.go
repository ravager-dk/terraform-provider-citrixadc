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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccTransformpolicylabel_transformpolicy_binding_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
  }
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
  }
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
  }
  resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding"{
	 policyname = citrixadc_transformpolicy.tf_trans_policy.name
	  labelname = citrixadc_transformpolicylabel.transformpolicylabel.labelname
	  priority = 2
  }
`

const testAccTransformpolicylabel_transformpolicy_binding_basic_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
  }
  
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
  }
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
  }
`

func TestAccTransformpolicylabel_transformpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", nil),
				),
			},
			{
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingNotExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "label_1,tf_trans_policy"),
				),
			},
		},
	})
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformpolicylabel_transformpolicy_binding id is set")
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

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "transformpolicylabel_transformpolicy_binding",
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "transformpolicylabel_transformpolicy_binding",
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformpolicylabel_transformpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Transformpolicylabel_transformpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
