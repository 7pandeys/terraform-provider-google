// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package gkehub

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func suppressGkeHubEndpointSelfLinkDiff(_, old, new string, _ *schema.ResourceData) bool {
	// The custom expander injects //container.googleapis.com/ if a selflink is supplied.
	selfLink := strings.TrimPrefix(old, "//container.googleapis.com/")
	if selfLink == new {
		return true
	}

	return false
}

func ResourceGKEHubMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHubMembershipCreate,
		Read:   resourceGKEHubMembershipRead,
		Update: resourceGKEHubMembershipUpdate,
		Delete: resourceGKEHubMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHubMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"membership_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The client-provided identifier of the membership.`,
			},
			"authority": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Authority encodes how Google will recognize identities from this Membership.
See the workload identity documentation for more details:
https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://' and // be a valid
with length <2000 characters. For example: 'https://container.googleapis.com/v1/projects/my-project/locations/us-west1/clusters/my-cluster' (must be 'locations' rather than 'zones'). If the cluster is provisioned with Terraform, this is '"https://container.googleapis.com/v1/${google_container_cluster.my-cluster.id}"'.`,
						},
					},
				},
			},
			"endpoint": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gke_cluster": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_link": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: suppressGkeHubEndpointSelfLinkDiff,
										Description: `Self-link of the GCP resource for the GKE cluster.
For example: '//container.googleapis.com/projects/my-project/zones/us-west1-a/clusters/my-cluster'.
It can be at the most 1000 characters in length. If the cluster is provisioned with Terraform,
this can be '"//container.googleapis.com/${google_container_cluster.my-cluster.id}"' or
'google_container_cluster.my-cluster.id'.`,
									},
								},
							},
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this membership.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of the membership.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceGKEHubMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	endpointProp, err := expandGKEHubMembershipEndpoint(d.Get("endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(endpointProp)) && (ok || !reflect.DeepEqual(v, endpointProp)) {
		obj["endpoint"] = endpointProp
	}
	authorityProp, err := expandGKEHubMembershipAuthority(d.Get("authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authority"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorityProp)) && (ok || !reflect.DeepEqual(v, authorityProp)) {
		obj["authority"] = authorityProp
	}
	labelsProp, err := expandGKEHubMembershipEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships?membershipId={{membership_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Membership: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Membership: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHubOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Membership", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Membership: %s", err)
	}

	if err := d.Set("name", flattenGKEHubMembershipName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Membership %q: %#v", d.Id(), res)

	return resourceGKEHubMembershipRead(d, meta)
}

func resourceGKEHubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHubMembership %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}

	if err := d.Set("name", flattenGKEHubMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("labels", flattenGKEHubMembershipLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("endpoint", flattenGKEHubMembershipEndpoint(res["endpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("authority", flattenGKEHubMembershipAuthority(res["authority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("terraform_labels", flattenGKEHubMembershipTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("effective_labels", flattenGKEHubMembershipEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}

	return nil
}

func resourceGKEHubMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	authorityProp, err := expandGKEHubMembershipAuthority(d.Get("authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authority"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorityProp)) {
		obj["authority"] = authorityProp
	}
	labelsProp, err := expandGKEHubMembershipEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Membership %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("authority") {
		updateMask = append(updateMask, "authority")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Membership %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Membership %q: %#v", d.Id(), res)
	}

	err = GKEHubOperationWaitTime(
		config, res, project, "Updating Membership", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHubMembershipRead(d, meta)
}

func resourceGKEHubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Membership %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Membership")
	}

	err = GKEHubOperationWaitTime(
		config, res, project, "Deleting Membership", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Membership %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHubMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/memberships/(?P<membership_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<membership_id>[^/]+)",
		"(?P<membership_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHubMembershipName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGKEHubMembershipEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gke_cluster"] =
		flattenGKEHubMembershipEndpointGkeCluster(original["gkeCluster"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_link"] =
		flattenGKEHubMembershipEndpointGkeClusterResourceLink(original["resourceLink"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeClusterResourceLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipAuthority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["issuer"] =
		flattenGKEHubMembershipAuthorityIssuer(original["issuer"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipAuthorityIssuer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGKEHubMembershipEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHubMembershipEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGkeCluster, err := expandGKEHubMembershipEndpointGkeCluster(original["gke_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGkeCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gkeCluster"] = transformedGkeCluster
	}

	return transformed, nil
}

func expandGKEHubMembershipEndpointGkeCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceLink, err := expandGKEHubMembershipEndpointGkeClusterResourceLink(original["resource_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceLink"] = transformedResourceLink
	}

	return transformed, nil
}

func expandGKEHubMembershipEndpointGkeClusterResourceLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if strings.HasPrefix(v.(string), "//") {
		return v, nil
	} else {
		v = "//container.googleapis.com/" + v.(string)
		return v, nil
	}
}

func expandGKEHubMembershipAuthority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIssuer, err := expandGKEHubMembershipAuthorityIssuer(original["issuer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["issuer"] = transformedIssuer
	}

	return transformed, nil
}

func expandGKEHubMembershipAuthorityIssuer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHubMembershipEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
