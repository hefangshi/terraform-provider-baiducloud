package baiducloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	testAccEipsDataSourceName          = "data.baiducloud_eips.default"
	testAccEipsDataSourceAttrKeyPrefix = "eips.0."
)

func TestAccBaiduCloudEipsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEipsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccEipsDataSourceName),
					resource.TestCheckResourceAttrSet(testAccEipsDataSourceName, testAccEipsDataSourceAttrKeyPrefix+"eip"),
					resource.TestCheckResourceAttr(testAccEipsDataSourceName, testAccEipsDataSourceAttrKeyPrefix+"bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(testAccEipsDataSourceName, testAccEipsDataSourceAttrKeyPrefix+"tags.#", "1"),
					resource.TestCheckResourceAttr(testAccEipsDataSourceName, testAccEipsDataSourceAttrKeyPrefix+"tags.0.tag_key", "testKey"),
					resource.TestCheckResourceAttr(testAccEipsDataSourceName, testAccEipsDataSourceAttrKeyPrefix+"tags.0.tag_value", "testValue"),
				),
			},
		},
	})
}

const testAccEipsDataSourceConfig = `
resource "baiducloud_eip" "my-eip" {
  name              = "test-BaiduAccEip"
  bandwidth_in_mbps = 100
  payment_timing    = "Postpaid"
  billing_method    = "ByTraffic"

  tags {
    tag_key   = "testKey"
    tag_value = "testValue"
  }
}

data "baiducloud_eips" "default" {
  eip = "${baiducloud_eip.my-eip.id}"
}
`
