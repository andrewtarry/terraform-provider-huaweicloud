// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product SecMaster
// ---------------------------------------------------------------

package secmaster

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API SecMaster DELETE /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules
// @API SecMaster POST /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules
// @API SecMaster GET /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/{id}
// @API SecMaster PUT /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/{id}
// @API SecMaster POST /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/enable
// @API SecMaster POST /v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/disable
func ResourceAlertRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAlertRuleCreate,
		UpdateContext: resourceAlertRuleUpdate,
		ReadContext:   resourceAlertRuleRead,
		DeleteContext: resourceAlertRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceAlertRuleImportState,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the workspace to which the alert rule belongs.`,
			},
			"pipeline_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the pipeline ID of the alert rule.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the alert rule name.`,
			},
			"severity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the severity of the alert rule.`,
			},
			"type": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: `Specifies the type of the alert rule.`,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the description of the alert rule.`,
			},
			"status": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the status of the alert rule.`,
			},
			"query_rule": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the query rule of the alert rule.`,
			},
			"query_plan": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        alertRuleScheduleSchema(),
				Required:    true,
				Description: `Specifies the query plan of the alert rule.`,
			},
			"triggers": {
				Type:        schema.TypeList,
				MaxItems:    5,
				MinItems:    1,
				Elem:        alertRuleAlertRuleTriggerSchema(),
				Required:    true,
				Description: `Specifies the triggers of the alert rule.`,
			},
			"query_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the query type of the alert rule.`,
			},
			"custom_information": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Specifies the custom information of the alert rule.`,
			},
			"event_grouping": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Specifies whether to put events in a group.`,
			},
			"debugging_alarm": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Specifies whether to generate debugging alarms.`,
			},
			"suppression": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Specifies whether to stop the query when an alarm is generated.`,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The created time.`,
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The updated time.`,
			},
		},
	}
}

func alertRuleScheduleSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"query_interval": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the query interval.`,
			},
			"query_interval_unit": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the query interval unit.`,
			},
			"time_window": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the time window.`,
			},
			"time_window_unit": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the time window unit.`,
			},
			"execution_delay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the execution delay in minutes.`,
			},
			"overtime_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the overtime interval in minutes.`,
			},
		},
	}
	return &sc
}

func alertRuleAlertRuleTriggerSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the expression.`,
			},
			"operator": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the operator.`,
			},
			"accumulated_times": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the accumulated times.`,
			},
			"mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the trigger mode.`,
			},
			"severity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the severity of the trigger.`,
			},
		},
	}
	return &sc
}

func resourceAlertRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createAlertRule: Create a SecMaster alert rule.
	var (
		createAlertRuleHttpUrl = "v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules"
		createAlertRuleProduct = "secmaster"
	)
	createAlertRuleClient, err := cfg.NewServiceClient(createAlertRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating SecMaster client: %s", err)
	}

	createAlertRulePath := createAlertRuleClient.Endpoint + createAlertRuleHttpUrl
	createAlertRulePath = strings.ReplaceAll(createAlertRulePath, "{project_id}", createAlertRuleClient.ProjectID)
	createAlertRulePath = strings.ReplaceAll(createAlertRulePath, "{workspace_id}", d.Get("workspace_id").(string))

	createAlertRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	createAlertRuleOpt.JSONBody = utils.RemoveNil(buildCreateAlertRuleBodyParams(d))
	createAlertRuleResp, err := createAlertRuleClient.Request("POST", createAlertRulePath, &createAlertRuleOpt)
	if err != nil {
		return diag.Errorf("error creating AlertRule: %s", err)
	}

	createAlertRuleRespBody, err := utils.FlattenResponse(createAlertRuleResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("rule_id", createAlertRuleRespBody)
	if err != nil {
		return diag.Errorf("error creating AlertRule: ID is not found in API response")
	}
	d.SetId(id.(string))

	return resourceAlertRuleRead(ctx, d, meta)
}

func buildCreateAlertRuleBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"pipe_id":           d.Get("pipeline_id"),
		"rule_name":         d.Get("name"),
		"severity":          d.Get("severity"),
		"alert_type":        d.Get("type"),
		"description":       d.Get("description"),
		"status":            d.Get("status"),
		"query":             d.Get("query_rule"),
		"query_type":        d.Get("query_type"),
		"schedule":          buildCreateAlertRuleRequestBodySchedule(d.Get("query_plan")),
		"custom_properties": utils.ValueIngoreEmpty(d.Get("custom_information")),
		"event_grouping":    utils.ValueIngoreEmpty(d.Get("event_grouping")),
		"simulation":        utils.ValueIngoreEmpty(d.Get("debugging_alarm")),
		"triggers":          buildCreateAlertRuleRequestBodyAlertRuleTrigger(d.Get("triggers")),
		"suppression":       utils.ValueIngoreEmpty(d.Get("suppression")),
	}
	return bodyParams
}

func buildCreateAlertRuleRequestBodySchedule(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"frequency_interval": utils.ValueIngoreEmpty(raw["query_interval"]),
			"frequency_unit":     utils.ValueIngoreEmpty(raw["query_interval_unit"]),
			"period_interval":    utils.ValueIngoreEmpty(raw["time_window"]),
			"period_unit":        utils.ValueIngoreEmpty(raw["time_window_unit"]),
			"delay_interval":     utils.ValueIngoreEmpty(raw["execution_delay"]),
			"overtime_interval":  utils.ValueIngoreEmpty(raw["overtime_interval"]),
		}
		return params
	}
	return nil
}

func buildCreateAlertRuleRequestBodyAlertRuleTrigger(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"expression":        utils.ValueIngoreEmpty(raw["expression"]),
				"operator":          utils.ValueIngoreEmpty(raw["operator"]),
				"accumulated_times": utils.ValueIngoreEmpty(raw["accumulated_times"]),
				"mode":              utils.ValueIngoreEmpty(raw["mode"]),
				"severity":          utils.ValueIngoreEmpty(raw["severity"]),
			}
		}
		return rst
	}
	return nil
}

func resourceAlertRuleRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getAlertRule: Query the SecMaster alert rule detail
	var (
		getAlertRuleHttpUrl = "v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/{id}"
		getAlertRuleProduct = "secmaster"
	)
	getAlertRuleClient, err := cfg.NewServiceClient(getAlertRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating SecMaster client: %s", err)
	}

	getAlertRulePath := getAlertRuleClient.Endpoint + getAlertRuleHttpUrl
	getAlertRulePath = strings.ReplaceAll(getAlertRulePath, "{project_id}", getAlertRuleClient.ProjectID)
	getAlertRulePath = strings.ReplaceAll(getAlertRulePath, "{workspace_id}", d.Get("workspace_id").(string))
	getAlertRulePath = strings.ReplaceAll(getAlertRulePath, "{id}", d.Id())

	getAlertRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	getAlertRuleResp, err := getAlertRuleClient.Request("GET", getAlertRulePath, &getAlertRuleOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving AlertRule")
	}

	getAlertRuleRespBody, err := utils.FlattenResponse(getAlertRuleResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("pipeline_id", utils.PathSearch("pipe_id", getAlertRuleRespBody, nil)),
		d.Set("name", utils.PathSearch("rule_name", getAlertRuleRespBody, nil)),
		d.Set("severity", utils.PathSearch("severity", getAlertRuleRespBody, nil)),
		d.Set("description", utils.PathSearch("description", getAlertRuleRespBody, nil)),
		d.Set("status", utils.PathSearch("status", getAlertRuleRespBody, nil)),
		d.Set("query_rule", utils.PathSearch("query", getAlertRuleRespBody, nil)),
		d.Set("query_type", utils.PathSearch("query_type", getAlertRuleRespBody, nil)),
		d.Set("query_plan", flattenGetAlertRuleResponseBodySchedule(getAlertRuleRespBody)),
		d.Set("custom_information", utils.PathSearch("custom_properties", getAlertRuleRespBody, nil)),
		d.Set("event_grouping", utils.PathSearch("event_grouping", getAlertRuleRespBody, nil)),
		d.Set("debugging_alarm", utils.PathSearch("simulation", getAlertRuleRespBody, nil)),
		d.Set("triggers", flattenGetAlertRuleResponseBodyAlertRuleTrigger(getAlertRuleRespBody)),
		d.Set("suppression", utils.PathSearch("suppression", getAlertRuleRespBody, nil)),
		d.Set("created_at", utils.FormatTimeStampRFC3339(
			int64(utils.PathSearch("create_time", getAlertRuleRespBody, 0).(float64))/1000, false)),
		d.Set("updated_at", utils.FormatTimeStampRFC3339(
			int64(utils.PathSearch("update_time", getAlertRuleRespBody, 0).(float64))/1000, false)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenGetAlertRuleResponseBodySchedule(resp interface{}) []interface{} {
	var rst []interface{}
	curJson, err := jmespath.Search("schedule", resp)
	if err != nil {
		log.Printf("[ERROR] error parsing schedule from response= %#v", resp)
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"query_interval":      utils.PathSearch("frequency_interval", curJson, nil),
			"query_interval_unit": utils.PathSearch("frequency_unit", curJson, nil),
			"time_window":         utils.PathSearch("period_interval", curJson, nil),
			"time_window_unit":    utils.PathSearch("period_unit", curJson, nil),
			"execution_delay":     utils.PathSearch("delay_interval", curJson, nil),
			"overtime_interval":   utils.PathSearch("overtime_interval", curJson, nil),
		},
	}
	return rst
}

func flattenGetAlertRuleResponseBodyAlertRuleTrigger(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("triggers", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"expression":        utils.PathSearch("expression", v, nil),
			"operator":          utils.PathSearch("operator", v, nil),
			"accumulated_times": utils.PathSearch("accumulated_times", v, nil),
			"mode":              utils.PathSearch("mode", v, nil),
			"severity":          utils.PathSearch("severity", v, nil),
		})
	}
	return rst
}

func resourceAlertRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var (
		updateAlertRuleProduct = "secmaster"
	)
	updateAlertRuleClient, err := cfg.NewServiceClient(updateAlertRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating SecMaster client: %s", err)
	}

	updateAlertRuleChanges := []string{
		"name",
		"severity",
		"type",
		"description",
		"query_rule",
		"query_type",
		"query_plan",
		"custom_information",
		"event_grouping",
		"debugging_alarm",
		"triggers",
		"suppression",
	}

	if d.HasChanges(updateAlertRuleChanges...) {
		// updateAlertRule: Update the configuration of SecMaster alert rule
		var (
			updateAlertRuleHttpUrl = "v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/{id}"
		)

		updateAlertRulePath := updateAlertRuleClient.Endpoint + updateAlertRuleHttpUrl
		updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{project_id}", updateAlertRuleClient.ProjectID)
		updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{workspace_id}", d.Get("workspace_id").(string))
		updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{id}", d.Id())

		updateAlertRuleOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
			MoreHeaders: map[string]string{"Content-Type": "application/json"},
		}

		updateAlertRuleOpt.JSONBody = utils.RemoveNil(buildUpdateAlertRuleBodyParams(d))
		_, err = updateAlertRuleClient.Request("PUT", updateAlertRulePath, &updateAlertRuleOpt)
		if err != nil {
			return diag.Errorf("error updating AlertRule: %s", err)
		}
	}

	if d.HasChange("status") {
		// Update the status of SecMaster alert rule
		if d.Get("status").(string) == "ENABLED" {
			err := updateAlertStatus(updateAlertRuleClient, d, "enable")
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := updateAlertStatus(updateAlertRuleClient, d, "disable")
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return resourceAlertRuleRead(ctx, d, meta)
}

func updateAlertStatus(updateAlertRuleClient *golangsdk.ServiceClient, d *schema.ResourceData, action string) error {
	var (
		updateAlertRuleStatusHttpUrl = "v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules/{action}"
	)

	updateAlertRulePath := updateAlertRuleClient.Endpoint + updateAlertRuleStatusHttpUrl
	updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{project_id}", updateAlertRuleClient.ProjectID)
	updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{workspace_id}", d.Get("workspace_id").(string))
	updateAlertRulePath = strings.ReplaceAll(updateAlertRulePath, "{action}", action)

	updateAlertRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	updateAlertRuleOpt.JSONBody = []string{d.Id()}
	_, err := updateAlertRuleClient.Request("POST", updateAlertRulePath, &updateAlertRuleOpt)
	if err != nil {
		return fmt.Errorf("error updating alert rule status: %s", err)
	}

	return nil
}

func buildUpdateAlertRuleBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"rule_name":         utils.ValueIngoreEmpty(d.Get("name")),
		"severity":          utils.ValueIngoreEmpty(d.Get("severity")),
		"alert_type":        utils.ValueIngoreEmpty(d.Get("type")),
		"description":       utils.ValueIngoreEmpty(d.Get("description")),
		"status":            utils.ValueIngoreEmpty(d.Get("status")),
		"query":             utils.ValueIngoreEmpty(d.Get("query_rule")),
		"query_type":        utils.ValueIngoreEmpty(d.Get("query_type")),
		"schedule":          buildUpdateAlertRuleRequestBodySchedule(d.Get("query_plan")),
		"custom_properties": utils.ValueIngoreEmpty(d.Get("custom_information")),
		"event_grouping":    utils.ValueIngoreEmpty(d.Get("event_grouping")),
		"simulation":        utils.ValueIngoreEmpty(d.Get("debugging_alarm")),
		"triggers":          buildUpdateAlertRuleRequestBodyAlertRuleTrigger(d.Get("triggers")),
		"suppression":       utils.ValueIngoreEmpty(d.Get("suppression")),
	}
	return bodyParams
}

func buildUpdateAlertRuleRequestBodySchedule(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"frequency_interval": utils.ValueIngoreEmpty(raw["query_interval"]),
			"frequency_unit":     utils.ValueIngoreEmpty(raw["query_interval_unit"]),
			"period_interval":    utils.ValueIngoreEmpty(raw["time_window"]),
			"period_unit":        utils.ValueIngoreEmpty(raw["time_window_unit"]),
			"delay_interval":     utils.ValueIngoreEmpty(raw["execution_delay"]),
			"overtime_interval":  utils.ValueIngoreEmpty(raw["overtime_interval"]),
		}
		return params
	}
	return nil
}

func buildUpdateAlertRuleRequestBodyAlertRuleTrigger(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"expression":        utils.ValueIngoreEmpty(raw["expression"]),
				"operator":          utils.ValueIngoreEmpty(raw["operator"]),
				"accumulated_times": utils.ValueIngoreEmpty(raw["accumulated_times"]),
				"mode":              utils.ValueIngoreEmpty(raw["mode"]),
				"severity":          utils.ValueIngoreEmpty(raw["severity"]),
			}
		}
		return rst
	}
	return nil
}

func resourceAlertRuleDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// deleteAlertRule: Delete an existing SecMaster alert rule
	var (
		deleteAlertRuleHttpUrl = "v1/{project_id}/workspaces/{workspace_id}/siem/alert-rules"
		deleteAlertRuleProduct = "secmaster"
	)
	deleteAlertRuleClient, err := cfg.NewServiceClient(deleteAlertRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating SecMaster client: %s", err)
	}

	deleteAlertRulePath := deleteAlertRuleClient.Endpoint + deleteAlertRuleHttpUrl
	deleteAlertRulePath = strings.ReplaceAll(deleteAlertRulePath, "{project_id}", deleteAlertRuleClient.ProjectID)
	deleteAlertRulePath = strings.ReplaceAll(deleteAlertRulePath, "{workspace_id}", d.Get("workspace_id").(string))

	deleteAlertRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	deleteAlertRuleOpt.JSONBody = []string{d.Id()}
	_, err = deleteAlertRuleClient.Request("DELETE", deleteAlertRulePath, &deleteAlertRuleOpt)
	if err != nil {
		return diag.Errorf("error deleting AlertRule: %s", err)
	}

	return nil
}

func resourceAlertRuleImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format specified for import id, must be <workspace_id>/<rule_id>")
	}

	d.SetId(parts[1])

	err := d.Set("workspace_id", parts[0])
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
