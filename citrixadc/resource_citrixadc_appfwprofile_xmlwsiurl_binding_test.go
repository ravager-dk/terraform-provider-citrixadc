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

const testAccAppfwprofile_xmlwsiurl_binding_basic = `
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
	name                     = "tf_appfwprofile"
	bufferoverflowaction     = ["none"]
	contenttypeaction        = ["none"]
	cookieconsistencyaction  = ["none"]
	creditcard               = ["none"]
	creditcardaction         = ["none"]
	crosssitescriptingaction = ["none"]
	csrftagaction            = ["none"]
	denyurlaction            = ["none"]
	dynamiclearning          = ["none"]
	fieldconsistencyaction   = ["none"]
	fieldformataction        = ["none"]
	fileuploadtypesaction    = ["none"]
	inspectcontenttypes      = ["none"]
	jsondosaction            = ["none"]
	jsonsqlinjectionaction   = ["none"]
	jsonxssaction            = ["none"]
	multipleheaderaction     = ["none"]
	sqlinjectionaction       = ["none"]
	starturlaction           = ["none"]
	type                     = ["HTML"]
	xmlattachmentaction      = ["none"]
	xmldosaction             = ["none"]
	xmlformataction          = ["none"]
	xmlsoapfaultaction       = ["none"]
	xmlsqlinjectionaction    = ["none"]
	xmlvalidationaction      = ["none"]
	xmlwsiaction             = ["none"]
	xmlxssaction             = ["none"]
  }
  resource "citrixadc_appfwprofile_xmlwsiurl_binding" "tf_binding" {
	name           = citrixadc_appfwprofile.tf_appfwprofile.name
	xmlwsiurl      = ".*"
	state          = "DISABLED"
	xmlwsichecks   = "R1140"
	isautodeployed = "AUTODEPLOYED"
	comment        = "Testing"
	alertonly      = "ON"
  }
`

const testAccAppfwprofile_xmlwsiurl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		bufferoverflowaction     = ["none"]
		contenttypeaction        = ["none"]
		cookieconsistencyaction  = ["none"]
		creditcard               = ["none"]
		creditcardaction         = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction            = ["none"]
		denyurlaction            = ["none"]
		dynamiclearning          = ["none"]
		fieldconsistencyaction   = ["none"]
		fieldformataction        = ["none"]
		fileuploadtypesaction    = ["none"]
		inspectcontenttypes      = ["none"]
		jsondosaction            = ["none"]
		jsonsqlinjectionaction   = ["none"]
		jsonxssaction            = ["none"]
		multipleheaderaction     = ["none"]
		sqlinjectionaction       = ["none"]
		starturlaction           = ["none"]
		type                     = ["HTML"]
		xmlattachmentaction      = ["none"]
		xmldosaction             = ["none"]
		xmlformataction          = ["none"]
		xmlsoapfaultaction       = ["none"]
		xmlsqlinjectionaction    = ["none"]
		xmlvalidationaction      = ["none"]
		xmlwsiaction             = ["none"]
		xmlxssaction             = ["none"]
	  }
`

func TestAccAppfwprofile_xmlwsiurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_xmlwsiurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlwsiurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlwsiurl_bindingExist("citrixadc_appfwprofile_xmlwsiurl_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccAppfwprofile_xmlwsiurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlwsiurl_bindingNotExist("citrixadc_appfwprofile_xmlwsiurl_binding.tf_binding", "tf_appfwprofile,.*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlwsiurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlwsiurl_binding id is set")
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

		name := idSlice[0]
		xmlwsiurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlwsiurl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["xmlwsiurl"].(string) == xmlwsiurl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlwsiurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlwsiurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		xmlwsiurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlwsiurl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["xmlwsiurl"].(string) == xmlwsiurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlwsiurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlwsiurl_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlwsiurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwprofile_xmlwsiurl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlwsiurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
