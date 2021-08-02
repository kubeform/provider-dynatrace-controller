/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package aws

import (
	"context"

	"github.com/dtcookie/dynatrace/api/config/credentials/aws"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	res := &schema.Resource{
		Schema:        hcl2sdk.Convert(new(aws.AWSCredentialsConfig).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
	return res
}

func NewService(m interface{}) *aws.ServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := aws.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(aws.AWSCredentialsConfig)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = nil
	credentials.Metadata = nil
	objStub, err := NewService(m).Create(credentials)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(aws.AWSCredentialsConfig)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = opt.NewString(d.Id())
	credentials.Metadata = nil
	if err := NewService(m).Update(credentials); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials, err := NewService(m).Get(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if credentials.AuthenticationData != nil {
		if credentials.AuthenticationData.KeyBasedAuthentication != nil {
			if credentials.AuthenticationData.KeyBasedAuthentication.SecretKey == nil {
				if secretKey, ok := d.GetOk("authentication_data.0.secret_key"); ok && len(secretKey.(string)) > 0 {
					credentials.AuthenticationData.KeyBasedAuthentication.SecretKey = opt.NewString(secretKey.(string))
				}
			}
		}
	}
	marshalled, err := credentials.MarshalHCL()
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range marshalled {
		d.Set(k, v)
	}

	return diag.Diagnostics{}
}

// Delete the configuration
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := NewService(m).Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
