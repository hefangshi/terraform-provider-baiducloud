package baiducloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	testAccVPCsDataSourceName          = "data.baiducloud_vpcs.default"
	testAccVPCsDataSourceAttrKeyPrefix = "vpcs.0."
)

func TestAccBaiduCloudVPCsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccVPCsDataSourceName),
					resource.TestCheckResourceAttrSet(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"vpc_id"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"cidr", "192.168.0.0/24"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"name", "test-BaiduAccVPC"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"description", "created by terraform"),
					resource.TestCheckResourceAttrSet(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"route_table_id"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"secondary_cidrs.#", "0"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"tags.#", "1"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"tags.0.tag_key", "tagK"),
					resource.TestCheckResourceAttr(testAccVPCsDataSourceName, testAccVPCsDataSourceAttrKeyPrefix+"tags.0.tag_value", "tagV"),
				),
			},
		},
	})
}

const testAccVPCsDataSourceConfig = `
resource "baiducloud_vpc" "default" {
  name = "test-BaiduAccVPC"
  description = "created by terraform"
  cidr = "192.168.0.0/24"
  tags {
	tag_key = "tagK"
    tag_value = "tagV"
  }
}

data "baiducloud_vpcs" "default" {
  vpc_id = "${baiducloud_vpc.default.id}"
}
`
