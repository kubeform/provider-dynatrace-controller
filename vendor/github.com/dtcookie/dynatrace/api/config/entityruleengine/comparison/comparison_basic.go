package comparison

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// Comparison Defines how the matching is actually performed: what and how are we comparing.
// The actual set of fields and possible values of the **operator** field depend on the **type** of the comparison. \n\nFind the list of actual models in the description of the **type** field and check the description of the model you need.
type Comparison interface {
	MarshalHCL() (map[string]interface{}, error)
	UnmarshalHCL(decoder hcl.Decoder) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	Schema() map[string]*hcl.Schema
	IsNegated() bool
	GetType() ComparisonBasicType
}

// BaseComparison Defines how the matching is actually performed: what and how are we comparing.
// The actual set of fields and possible values of the **operator** field depend on the **type** of the comparison. \n\nFind the list of actual models in the description of the **type** field and check the description of the model you need.
type BaseComparison struct {
	Type     ComparisonBasicType        `json:"type"`   // Defines the actual set of fields depending on the value. See one of the following objects:  * `STRING` -> StringComparison  * `INDEXED_NAME` -> IndexedNameComparison  * `INDEXED_STRING` -> IndexedStringComparison  * `INTEGER` -> IntegerComparison  * `SERVICE_TYPE` -> ServiceTypeComparison  * `PAAS_TYPE` -> PaasTypeComparison  * `CLOUD_TYPE` -> CloudTypeComparison  * `AZURE_SKU` -> AzureSkuComparision  * `AZURE_COMPUTE_MODE` -> AzureComputeModeComparison  * `ENTITY_ID` -> EntityIdComparison  * `SIMPLE_TECH` -> SimpleTechComparison  * `SIMPLE_HOST_TECH` -> SimpleHostTechComparison  * `SERVICE_TOPOLOGY` -> ServiceTopologyComparison  * `DATABASE_TOPOLOGY` -> DatabaseTopologyComparison  * `OS_TYPE` -> OsTypeComparison  * `HYPERVISOR_TYPE` -> HypervisorTypeComparision  * `IP_ADDRESS` -> IpAddressComparison  * `OS_ARCHITECTURE` -> OsArchitectureComparison  * `BITNESS` -> BitnessComparision  * `APPLICATION_TYPE` -> ApplicationTypeComparison  * `MOBILE_PLATFORM` -> MobilePlatformComparison  * `CUSTOM_APPLICATION_TYPE` -> CustomApplicationTypeComparison  * `DCRUM_DECODER_TYPE` -> DcrumDecoderComparison  * `SYNTHETIC_ENGINE_TYPE` -> SyntheticEngineTypeComparison  * `TAG` -> TagComparison  * `INDEXED_TAG` -> IndexedTagComparison
	Negate   bool                       `json:"negate"` // Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (bcb *BaseComparison) IsNegated() bool {
	return bcb.Negate
}

func (bcb *BaseComparison) GetType() ComparisonBasicType {
	return bcb.Type
}

func (bcb *BaseComparison) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of comparison",
			Required:    true,
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns EQUALS into DOES NOT EQUAL",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (bcb *BaseComparison) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(bcb.Unknowns) > 0 {
		data, err := json.Marshal(bcb.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = bcb.Negate
	result["type"] = string(bcb.Type)
	return result, nil
}

func (bcb *BaseComparison) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), bcb); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &bcb.Unknowns); err != nil {
			return err
		}
		delete(bcb.Unknowns, "type")
		delete(bcb.Unknowns, "negate")
		if len(bcb.Unknowns) == 0 {
			bcb.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		bcb.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		bcb.Negate = value.(bool)
	}
	return nil
}

func (bcb *BaseComparison) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(bcb.Unknowns) > 0 {
		for k, v := range bcb.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(bcb.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(bcb.Type)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	return json.Marshal(m)
}

func (bcb *BaseComparison) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &bcb.Negate); err != nil {
			return err
		}
	}
	if v, found := m["type"]; found {
		if err := json.Unmarshal(v, &bcb.Type); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "type")
	if len(m) > 0 {
		bcb.Unknowns = m
	}
	return nil
}

// ComparisonBasicType Defines the actual set of fields depending on the value. See one of the following objects:
// * `STRING` -> StringComparison
// * `INDEXED_NAME` -> IndexedNameComparison
// * `INDEXED_STRING` -> IndexedStringComparison
// * `INTEGER` -> IntegerComparison
// * `SERVICE_TYPE` -> ServiceTypeComparison
// * `PAAS_TYPE` -> PaasTypeComparison
// * `CLOUD_TYPE` -> CloudTypeComparison
// * `AZURE_SKU` -> AzureSkuComparision
// * `AZURE_COMPUTE_MODE` -> AzureComputeModeComparison
// * `ENTITY_ID` -> EntityIdComparison
// * `SIMPLE_TECH` -> SimpleTechComparison
// * `SIMPLE_HOST_TECH` -> SimpleHostTechComparison
// * `SERVICE_TOPOLOGY` -> ServiceTopologyComparison
// * `DATABASE_TOPOLOGY` -> DatabaseTopologyComparison
// * `OS_TYPE` -> OsTypeComparison
// * `HYPERVISOR_TYPE` -> HypervisorTypeComparision
// * `IP_ADDRESS` -> IpAddressComparison
// * `OS_ARCHITECTURE` -> OsArchitectureComparison
// * `BITNESS` -> BitnessComparision
// * `APPLICATION_TYPE` -> ApplicationTypeComparison
// * `MOBILE_PLATFORM` -> MobilePlatformComparison
// * `CUSTOM_APPLICATION_TYPE` -> CustomApplicationTypeComparison
// * `DCRUM_DECODER_TYPE` -> DcrumDecoderComparison
// * `SYNTHETIC_ENGINE_TYPE` -> SyntheticEngineTypeComparison
// * `TAG` -> TagComparison
// * `INDEXED_TAG` -> IndexedTagComparison
type ComparisonBasicType string

// ComparisonBasicTypes offers the known enum values
var ComparisonBasicTypes = struct {
	ApplicationType       ComparisonBasicType
	AzureComputeMode      ComparisonBasicType
	AzureSku              ComparisonBasicType
	Bitness               ComparisonBasicType
	CloudType             ComparisonBasicType
	CustomApplicationType ComparisonBasicType
	DatabaseTopology      ComparisonBasicType
	DCRumDecoderType      ComparisonBasicType
	EntityID              ComparisonBasicType
	HypervisorType        ComparisonBasicType
	IndexedName           ComparisonBasicType
	IndexedString         ComparisonBasicType
	IndexedTag            ComparisonBasicType
	Integer               ComparisonBasicType
	IPAddress             ComparisonBasicType
	MobilePlatform        ComparisonBasicType
	OSArchitecture        ComparisonBasicType
	OSType                ComparisonBasicType
	PaasType              ComparisonBasicType
	ServiceTopology       ComparisonBasicType
	ServiceType           ComparisonBasicType
	SimpleHostTech        ComparisonBasicType
	SimpleTech            ComparisonBasicType
	String                ComparisonBasicType
	SyntheticEngineType   ComparisonBasicType
	Tag                   ComparisonBasicType
}{
	"APPLICATION_TYPE",
	"AZURE_COMPUTE_MODE",
	"AZURE_SKU",
	"BITNESS",
	"CLOUD_TYPE",
	"CUSTOM_APPLICATION_TYPE",
	"DATABASE_TOPOLOGY",
	"DCRUM_DECODER_TYPE",
	"ENTITY_ID",
	"HYPERVISOR_TYPE",
	"INDEXED_NAME",
	"INDEXED_STRING",
	"INDEXED_TAG",
	"INTEGER",
	"IP_ADDRESS",
	"MOBILE_PLATFORM",
	"OS_ARCHITECTURE",
	"OS_TYPE",
	"PAAS_TYPE",
	"SERVICE_TOPOLOGY",
	"SERVICE_TYPE",
	"SIMPLE_HOST_TECH",
	"SIMPLE_TECH",
	"STRING",
	"SYNTHETIC_ENGINE_TYPE",
	"TAG",
}
