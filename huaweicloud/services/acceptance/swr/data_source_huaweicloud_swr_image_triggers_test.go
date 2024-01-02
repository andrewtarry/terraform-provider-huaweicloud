package swr

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceImageTriggers_basic(t *testing.T) {
	dataSourceName := "data.huaweicloud_swr_image_triggers.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckWorkloadName(t)
			acceptance.TestAccPreCheckCceClusterId(t)
			acceptance.TestAccPreCheckWorkloadNameSpace(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceImageTriggers_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSourceName, "triggers.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "triggers.0.enabled"),
					resource.TestCheckResourceAttrSet(dataSourceName, "triggers.0.condition_type"),
					resource.TestCheckResourceAttrSet(dataSourceName, "triggers.0.cluster_name"),
					resource.TestCheckOutput("name_filter_is_useful", "true"),
					resource.TestCheckOutput("name_filter_not_found_validation_pass", "true"),
					resource.TestCheckOutput("enabled_filter_is_useful", "true"),
					resource.TestCheckOutput("condition_type_filter_is_useful", "true"),
					resource.TestCheckOutput("cluster_name_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testAccDataSourceImageTriggers_basic(rName string) string {
	return fmt.Sprintf(`
%[1]s

locals {
  organization = huaweicloud_swr_organization.test.name
  repository   = huaweicloud_swr_repository.test.name
}

resource "huaweicloud_swr_image_trigger" "test_tag" {
  organization    = local.organization
  repository      = local.repository
  workload_type   = "deployments"
  workload_name   = "%[2]s"
  cluster_id      = "%[3]s"
  namespace       = "%[4]s"
  condition_value = "v1.2"
  enabled         = "false"
  name            = "%[5]s_tag"
  type            = "cce"
  condition_type  = "tag"
}

resource "huaweicloud_swr_image_trigger" "test_all" {
  organization    = local.organization
  repository      = local.repository
  workload_type   = "deployments"
  workload_name   = "%[2]s"
  cluster_id      = "%[3]s"
  namespace       = "%[4]s"
  condition_value = ".*"
  enabled         = "true"
  name            = "%[5]s_all"
  type            = "cce"
  condition_type  = "all"
}

resource "huaweicloud_swr_image_trigger" "test_reg" {
  organization    = local.organization
  repository      = local.repository
  workload_type   = "deployments"
  workload_name   = "%[2]s"
  cluster_id      = "%[3]s"
  namespace       = "%[4]s"
  condition_value = ".*"
  enabled         = "false"
  name            = "%[5]s_reg"
  type            = "cce"
  condition_type  = "regular"
}

data "huaweicloud_swr_image_triggers" "test" {
  depends_on = [
	huaweicloud_swr_image_trigger.test_tag,
	huaweicloud_swr_image_trigger.test_all,
	huaweicloud_swr_image_trigger.test_reg,
  ]

  organization = local.organization
  repository   = local.repository
}

locals {
  name = data.huaweicloud_swr_image_triggers.test.triggers[0].name
}
data "huaweicloud_swr_image_triggers" "filter_by_name" {
  organization = local.organization
  repository   = local.repository
  name         = local.name
}
output "name_filter_is_useful" {
  value = length(data.huaweicloud_swr_image_triggers.filter_by_name.triggers) > 0 && alltrue(
	[for v in data.huaweicloud_swr_image_triggers.filter_by_name.triggers[*].name : v == local.name]
  )
}

data "huaweicloud_swr_image_triggers" "filter_by_name_not_found" {
  organization = local.organization
  repository   = local.repository
  name         = "%[5]s_not_found"
  }
output "name_filter_not_found_validation_pass" {
  value = length(data.huaweicloud_swr_image_triggers.filter_by_name_not_found.triggers) == 0
}

locals {
  enabled = data.huaweicloud_swr_image_triggers.test.triggers[0].enabled
}
data "huaweicloud_swr_image_triggers" "filter_by_enabled" {
  organization = local.organization
  repository   = local.repository
  enabled      = local.enabled
}
output "enabled_filter_is_useful" {
  value = length(data.huaweicloud_swr_image_triggers.filter_by_enabled.triggers) > 0 && alltrue(
	[for v in data.huaweicloud_swr_image_triggers.filter_by_enabled.triggers[*].enabled : v == local.enabled]
  )
}

locals {
  condition_type = data.huaweicloud_swr_image_triggers.test.triggers[0].condition_type
}
data "huaweicloud_swr_image_triggers" "filter_by_condition_type" {
  organization   = local.organization
  repository     = local.repository
  condition_type = local.condition_type
}
output "condition_type_filter_is_useful" {
  value = length(data.huaweicloud_swr_image_triggers.filter_by_condition_type.triggers) > 0 && alltrue(
	[for v in data.huaweicloud_swr_image_triggers.filter_by_condition_type.triggers[*].condition_type : v == local.condition_type]
  )
}

locals {
  cluster_name = data.huaweicloud_swr_image_triggers.test.triggers[0].cluster_name
}
data "huaweicloud_swr_image_triggers" "filter_by_cluster_name" {
  organization   = local.organization
  repository     = local.repository
  cluster_name   = local.cluster_name
  }
output "cluster_name_filter_is_useful" {
  value = length(data.huaweicloud_swr_image_triggers.filter_by_cluster_name.triggers) > 0 && alltrue(
	[for v in data.huaweicloud_swr_image_triggers.filter_by_cluster_name.triggers[*].cluster_name : v == local.cluster_name]
  )
}
`, testAccSWRRepository_basic(rName), acceptance.HW_WORKLOAD_NAME,
		acceptance.HW_CCE_CLUSTER_ID, acceptance.HW_WORKLOAD_NAMESPACE, rName)
}
