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

package vault

import (
	"context"
	"reflect"

	"github.com/dtcookie/dynatrace/api/config/credentials/vault"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/terraform"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"
	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceRead,
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func DataSourceRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(uuid.New().String())
	conf := m.(*config.ProviderConfiguration)
	apiService := vault.NewService(conf.DTenvURL, conf.APIToken)
	stubList, err := apiService.ListAll()
	if err != nil {
		return err
	}
	if len(stubList.Credentials) > 0 {
		profiles := map[string]interface{}{}
		for _, stub := range stubList.Credentials {
			profiles[*stub.ID] = stub.Name
		}
		d.Set("credentials", profiles)
	}
	return nil
}

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        hcl2sdk.Convert(new(vault.Credentials).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func Resolve(d *schema.ResourceData) (*vault.Credentials, error) {
	resolver, err := terraform.NewResolver(d)
	if err != nil {
		return nil, err
	}
	untypedConfig, err := resolver.Resolve(reflect.TypeOf(vault.Credentials{}))
	if err != nil {
		return nil, err
	}
	typedConfig := untypedConfig.(vault.Credentials)
	return &typedConfig, nil
}

func NewService(m interface{}) *vault.ServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := vault.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(vault.Credentials)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = nil
	objStub, err := NewService(m).Create(credentials)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(vault.Credentials)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = opt.NewString(d.Id())
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
