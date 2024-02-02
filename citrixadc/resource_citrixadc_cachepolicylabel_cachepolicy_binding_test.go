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

const testAccCachepolicylabel_cachepolicy_binding_basic = `

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}

	resource "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_cachepolicylabel_cachepolicy_binding" {
		labelname  = citrixadc_cachepolicylabel.tf_policylabel.labelname
		priority   = 100
		policyname = citrixadc_cachepolicy.tf_cachepolicy.policyname
	}
  
  
`

const testAccCachepolicylabel_cachepolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
`

func TestAccCachepolicylabel_cachepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCachepolicylabel_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicylabel_cachepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", nil),
				),
			},
			{
				Config: testAccCachepolicylabel_cachepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingNotExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "my_cachepolicylabel,my_cachepolicy"),
				),
			},
		},
	})
}

func testAccCheckCachepolicylabel_cachepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachepolicylabel_cachepolicy_binding id is set")
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
			ResourceType:             "cachepolicylabel_cachepolicy_binding",
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
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachepolicylabel_cachepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "cachepolicylabel_cachepolicy_binding",
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
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCachepolicylabel_cachepolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachepolicylabel_cachepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Cachepolicylabel_cachepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
