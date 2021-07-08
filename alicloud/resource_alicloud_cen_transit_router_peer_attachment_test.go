package alicloud

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudCenTransitRouterPeerAttachment_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_peer_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPeerAttachmentMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CbnService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPeerAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000000, 9999999)
	name := fmt.Sprintf("tf-testAccCenTransitRouterPeerAttachment%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPeerAttachmentBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"cen_id":                                "cen-f6rslz7pzbnj8sshxc",
					"transit_router_id":                     "tr-bp1p0oqyc5iv22yjpymgu",
					"peer_transit_router_region_id":         "us-east-1",
					"peer_transit_router_id":                "${alicloud_cen_transit_router.default.transit_router_id}",
					"cen_bandwidth_package_id":              "cenbwp-buw65zk0606xh0ukvd",
					"bandwidth":                             "2",
					"transit_router_attachment_description": name,
					"transit_router_attachment_name":        name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cen_id":                                CHECKSET,
						"peer_transit_router_id":                CHECKSET,
						"transit_router_id":                     CHECKSET,
						"peer_transit_router_region_id":         "us-east-1",
						"cen_bandwidth_package_id":              "cenbwp-buw65zk0606xh0ukvd",
						"bandwidth":                             "2",
						"transit_router_attachment_description": name,
						"transit_router_attachment_name":        name,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cen_id", "dry_run", "route_table_association_enabled", "route_table_propagation_enabled"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_publish_route_enabled": `false`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_publish_route_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": `2`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cen_bandwidth_package_id": "cenbwp-buw65zk0606xh0ukvd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cen_bandwidth_package_id": "cenbwp-buw65zk0606xh0ukvd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"transit_router_attachment_description": "desp1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transit_router_attachment_description": "desp1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"transit_router_attachment_name": "name1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transit_router_attachment_name": "name1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_publish_route_enabled":            `true`,
					"bandwidth":                             `2`,
					"cen_bandwidth_package_id":              "cenbwp-buw65zk0606xh0ukvd",
					"transit_router_attachment_description": "desp",
					"transit_router_attachment_name":        "name",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_publish_route_enabled":            "true",
						"bandwidth":                             "2",
						"cen_bandwidth_package_id":              "cenbwp-buw65zk0606xh0ukvd",
						"transit_router_attachment_description": "desp",
						"transit_router_attachment_name":        "name",
					}),
				),
			},
		},
	})
}

var AlicloudCenTransitRouterPeerAttachmentMap = map[string]string{
	"auto_publish_route_enabled":            CHECKSET,
	"bandwidth":                             CHECKSET,
	"cen_bandwidth_package_id":              CHECKSET,
	"cen_id":                                CHECKSET,
	"dry_run":                               NOSET,
	"peer_transit_router_id":                CHECKSET,
	"peer_transit_router_region_id":         CHECKSET,
	"resource_type":                         "TR",
	"route_table_association_enabled":       NOSET,
	"route_table_propagation_enabled":       NOSET,
	"status":                                CHECKSET,
	"transit_router_attachment_description": CHECKSET,
	"transit_router_attachment_name":        CHECKSET,
	"transit_router_id":                     CHECKSET,
}

func AlicloudCenTransitRouterPeerAttachmentBasicDependence(name string) string {
	return fmt.Sprintf(`

variable "name" {	
	default = "%s"
}

resource "alicloud_cen_transit_router" "default" {
  cen_id = "cen-f6rslz7pzbnj8sshxc"
}

resource "alicloud_cen_bandwidth_package_attachment" "default" {
  instance_id        = "cen-f6rslz7pzbnj8sshxc"
  bandwidth_package_id = "cenbwp-buw65zk0606xh0ukvd"
  depends_on = [alicloud_cen_transit_router.default]
}

`, name)
}
