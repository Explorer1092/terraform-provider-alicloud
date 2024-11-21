package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPrivateLinkVpcEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPrivateLinkVpcEndpointCreate,
		Read:   resourceAliCloudPrivateLinkVpcEndpointRead,
		Update: resourceAliCloudPrivateLinkVpcEndpointUpdate,
		Delete: resourceAliCloudPrivateLinkVpcEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"connection_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"endpoint_business_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"endpoint_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Interface"}, false),
			},
			"policy_document": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protected_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_group_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"vpc_endpoint_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_private_ip_address_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: IntInSlice([]int{0, 1}),
			},
		},
	}
}

func resourceAliCloudPrivateLinkVpcEndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateVpcEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewPrivatelinkClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("vpc_endpoint_name"); ok {
		request["EndpointName"] = v
	}
	if v, ok := d.GetOk("endpoint_type"); ok {
		request["EndpointType"] = v
	}
	if v, ok := d.GetOkExists("zone_private_ip_address_count"); ok {
		request["ZonePrivateIpAddressCount"] = v
	}
	if v, ok := d.GetOk("security_group_ids"); ok {
		securityGroupIdMapsArray := v.(*schema.Set).List()
		request["SecurityGroupId"] = securityGroupIdMapsArray
	}

	if v, ok := d.GetOk("policy_document"); ok {
		request["PolicyDocument"] = v
	}
	if v, ok := d.GetOkExists("protected_enabled"); ok {
		request["ProtectedEnabled"] = v
	}
	request["VpcId"] = d.Get("vpc_id")
	if v, ok := d.GetOk("service_name"); ok {
		request["ServiceName"] = v
	}
	if v, ok := d.GetOk("endpoint_description"); ok {
		request["EndpointDescription"] = v
	}
	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if v, ok := d.GetOk("service_id"); ok {
		request["ServiceId"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_privatelink_vpc_endpoint", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["EndpointId"]))

	privateLinkServiceV2 := PrivateLinkServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutCreate), 60*time.Second, privateLinkServiceV2.PrivateLinkVpcEndpointStateRefreshFunc(d.Id(), "EndpointStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudPrivateLinkVpcEndpointUpdate(d, meta)
}

func resourceAliCloudPrivateLinkVpcEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	privateLinkServiceV2 := PrivateLinkServiceV2{client}

	objectRaw, err := privateLinkServiceV2.DescribePrivateLinkVpcEndpoint(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_privatelink_vpc_endpoint DescribePrivateLinkVpcEndpoint Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("bandwidth", objectRaw["Bandwidth"])
	d.Set("connection_status", objectRaw["ConnectionStatus"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("endpoint_business_status", objectRaw["EndpointBusinessStatus"])
	d.Set("endpoint_description", objectRaw["EndpointDescription"])
	d.Set("endpoint_domain", objectRaw["EndpointDomain"])
	d.Set("endpoint_type", objectRaw["EndpointType"])
	d.Set("policy_document", objectRaw["PolicyDocument"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("service_id", objectRaw["ServiceId"])
	d.Set("service_name", objectRaw["ServiceName"])
	d.Set("status", objectRaw["EndpointStatus"])
	d.Set("vpc_endpoint_name", objectRaw["EndpointName"])
	d.Set("vpc_id", objectRaw["VpcId"])
	d.Set("zone_private_ip_address_count", objectRaw["ZonePrivateIpAddressCount"])

	objectRaw, err = privateLinkServiceV2.DescribePrivateLinkListTagResources(d.Id())
	if err != nil {
		return WrapError(err)
	}

	tagsMaps := objectRaw["TagResources"]
	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = privateLinkServiceV2.DescribeListVpcEndpointSecurityGroups(d.Id())
	if err != nil {
		return WrapError(err)
	}

	d.Set("security_group_ids", convertSecurityGroupIdToStringList(objectRaw["SecurityGroups"]))

	return nil
}

func resourceAliCloudPrivateLinkVpcEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	action := "UpdateVpcEndpointAttribute"
	conn, err := client.NewPrivatelinkClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["EndpointId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("policy_document") {
		update = true
		request["PolicyDocument"] = d.Get("policy_document")
	}

	if !d.IsNewResource() && d.HasChange("vpc_endpoint_name") {
		update = true
		request["EndpointName"] = d.Get("vpc_endpoint_name")
	}

	if !d.IsNewResource() && d.HasChange("endpoint_description") {
		update = true
		request["EndpointDescription"] = d.Get("endpoint_description")
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if d.HasChange("address_ip_version") {
		update = true
		request["AddressIpVersion"] = d.Get("address_ip_version")
	}

	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
			if err != nil {
				if IsExpectedErrors(err, []string{"EndpointLocked", "ConcurrentCallNotSupported", "EndpointConnectionOperationDenied", "EndpointOperationDenied"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	action = "ChangeResourceGroup"
	conn, err = client.NewPrivatelinkClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()

	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
	}
	request["ResourceGroupId"] = d.Get("resource_group_id")
	if !d.IsNewResource() && d.HasChange("region_id") {
		update = true
		request["ResourceRegionId"] = d.Get("region_id")
	}

	request["ResourceType"] = "VpcEndpoint"
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		privateLinkServiceV2 := PrivateLinkServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("resource_group_id"))}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, privateLinkServiceV2.PrivateLinkVpcEndpointStateRefreshFunc(d.Id(), "ResourceGroupId", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	if d.HasChange("tags") {
		privateLinkServiceV2 := PrivateLinkServiceV2{client}
		if err := privateLinkServiceV2.SetResourceTags(d, "VpcEndpoint"); err != nil {
			return WrapError(err)
		}
	}
	if !d.IsNewResource() && d.HasChange("security_group_ids") {
		oldEntry, newEntry := d.GetChange("security_group_ids")
		oldEntrySet := oldEntry.(*schema.Set)
		newEntrySet := newEntry.(*schema.Set)
		removed := oldEntrySet.Difference(newEntrySet)
		added := newEntrySet.Difference(oldEntrySet)

		if added.Len() > 0 {
			securityGroupIds := added.List()

			for _, item := range securityGroupIds {
				action := "AttachSecurityGroupToVpcEndpoint"
				conn, err := client.NewPrivatelinkClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["EndpointId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				request["DryRun"] = d.Get("dry_run")
				request["SecurityGroupId"] = item

				runtime := util.RuntimeOptions{}
				runtime.SetAutoretry(true)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				privateLinkServiceV2 := PrivateLinkServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutUpdate), 60*time.Second, privateLinkServiceV2.PrivateLinkVpcEndpointStateRefreshFunc(d.Id(), "EndpointStatus", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}

		if removed.Len() > 0 {
			securityGroupIds := removed.List()

			for _, item := range securityGroupIds {
				action := "DetachSecurityGroupFromVpcEndpoint"
				conn, err := client.NewPrivatelinkClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["EndpointId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				request["DryRun"] = d.Get("dry_run")
				request["SecurityGroupId"] = item
				runtime := util.RuntimeOptions{}
				runtime.SetAutoretry(true)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
					if err != nil {
						if IsExpectedErrors(err, []string{"ConcurrentCallNotSupported"}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				privateLinkServiceV2 := PrivateLinkServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutUpdate), 60*time.Second, privateLinkServiceV2.PrivateLinkVpcEndpointStateRefreshFunc(d.Id(), "EndpointStatus", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}

	}
	d.Partial(false)
	return resourceAliCloudPrivateLinkVpcEndpointRead(d, meta)
}

func resourceAliCloudPrivateLinkVpcEndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteVpcEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewPrivatelinkClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["EndpointId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-15"), StringPointer("AK"), query, request, &runtime)
		request["ClientToken"] = buildClientToken(action)

		if err != nil {
			if IsExpectedErrors(err, []string{"ConcurrentCallNotSupported", "EndpointOperationDenied"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"EndpointNotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	privateLinkServiceV2 := PrivateLinkServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, privateLinkServiceV2.PrivateLinkVpcEndpointStateRefreshFunc(d.Id(), "EndpointStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}

func convertSecurityGroupIdToStringList(src interface{}) (result []interface{}) {
	if src == nil {
		return
	}
	for _, v := range src.([]interface{}) {
		vv := v.(map[string]interface{})
		result = append(result, vv["SecurityGroupId"].(string))
	}
	return
}
