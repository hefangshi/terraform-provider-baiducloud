package baiducloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	testAccPeerConnAcceptorResourceType     = "baiducloud_peer_conn_acceptor"
	testAccPeerConnAcceptorResourceName     = testAccPeerConnAcceptorResourceType + "." + BaiduCloudTestResourceName
	testAccPeerConnAcceptorResourceAttrName = BaiduCloudTestResourceAttrNamePrefix + "PeerConnAcceptor"
)

func TestAccBaiduCloudPeerConnAcceptor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,

		Steps: []resource.TestStep{
			{
				Config: testAccPeerConnAcceptorConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccPeerConnAcceptorResourceName),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "bandwidth_in_mbps", "20"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "description", "test peer conn"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "local_if_name", "local-interface"),
					resource.TestCheckResourceAttrSet(testAccPeerConnResourceName, "local_if_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "local_vpc_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_vpc_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_region"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_account_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnResourceName, "status"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "created_time"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "dns_status", "close"),
				),
			},
			{
				ResourceName:            testAccPeerConnAcceptorResourceName,
				ImportState:             true,
				ImportStateVerifyIgnore: []string{"dns_sync"},
			},
			{
				Config: testAccPeerConnAcceptorConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccPeerConnAcceptorResourceName),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "bandwidth_in_mbps", "20"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "description", "test peer conn"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "local_if_name", "local-interface"),
					resource.TestCheckResourceAttrSet(testAccPeerConnResourceName, "local_if_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "local_vpc_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_vpc_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_region"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "peer_account_id"),
					resource.TestCheckResourceAttrSet(testAccPeerConnAcceptorResourceName, "created_time"),
					resource.TestCheckResourceAttr(testAccPeerConnAcceptorResourceName, "dns_status", "open"),
					resource.TestCheckResourceAttrSet(testAccPeerConnResourceName, "status"),
					resource.TestCheckResourceAttrSet(testAccPeerConnResourceName, "created_time"),
				),
			},
		},
	})
}

func testAccPeerConnAcceptorConfig() string {
	region := os.Getenv("BAIDUCLOUD_REGION")
	return fmt.Sprintf(`
provider "baiducloud" {
  alias = "local"
  // credential
}

provider "baiducloud" {
  alias = "peer"
  // credential
}

resource "baiducloud_vpc" "local-vpc" {
  provider = "baiducloud.local"
  name = "%s"
  cidr = "172.17.0.0/16"
}

resource "baiducloud_vpc" "peer-vpc" {
  provider = "baiducloud.peer"
  name = "%s"
  cidr = "172.18.0.0/16"
}

resource "baiducloud_peer_conn" "default" {
  provider = "baiducloud.local"
  bandwidth_in_mbps = 20
  local_vpc_id = "${baiducloud_vpc.local-vpc.id}"
  peer_vpc_id = "${baiducloud_vpc.peer-vpc.id}"
  peer_region = "%s"
  description = "test peer conn"
  local_if_name = "local-interface"
  billing = {
    payment_timing = "Postpaid"
  }
}

resource "baiducloud_peer_conn_acceptor" "default" {
  provider = "baiducloud.peer"
  peer_conn_id = "${baiducloud_peer_conn.default.id}"
  auto_accept = true
  dns_sync = false
}
`, BaiduCloudTestResourceAttrNamePrefix+"VPC-local",
		BaiduCloudTestResourceAttrNamePrefix+"VPC-peer", region)
}

func testAccPeerConnAcceptorConfigUpdate() string {
	region := os.Getenv("BAIDUCLOUD_REGION")
	return fmt.Sprintf(`
provider "baiducloud" {
  alias = "local"
  // credential
}

provider "baiducloud" {
  alias = "peer"
  // credential
}

resource "baiducloud_vpc" "local-vpc" {
  provider = "baiducloud.local"
  name = "%s"
  cidr = "172.17.0.0/16"
}

resource "baiducloud_vpc" "peer-vpc" {
  provider = "baiducloud.peer"
  name = "%s"
  cidr = "172.18.0.0/16"
}

resource "baiducloud_peer_conn" "default" {
  provider = "baiducloud.local"
  bandwidth_in_mbps = 20
  local_vpc_id = "${baiducloud_vpc.local-vpc.id}"
  peer_vpc_id = "${baiducloud_vpc.peer-vpc.id}"
  peer_region = "%s"
  description = "test peer conn"
  local_if_name = "local-interface"
  billing = {
    payment_timing = "Postpaid"
  }
}

resource "baiducloud_peer_conn_acceptor" "default" {
  provider = "baiducloud.peer"
  peer_conn_id = "${baiducloud_peer_conn.default.id}"
  auto_accept = true
  dns_sync = true
}
`, BaiduCloudTestResourceAttrNamePrefix+"VPC-local",
		BaiduCloudTestResourceAttrNamePrefix+"VPC-peer", region)
}
