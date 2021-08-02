package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/cloud_type"
	"github.com/dtcookie/hcl"
)

// CloudType Comparison for `CLOUD_TYPE` attributes.
type CloudType struct {
	BaseComparison
	Operator cloud_type.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *cloud_type.Value   `json:"value,omitempty"` // The value to compare to.
}

func (ctc *CloudType) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.CloudType
}

func (ctc *CloudType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be CLOUD_TYPE",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns EQUALS into DOES NOT EQUAL",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Either EQUALS or EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to. Possible values are AZURE, EC2, GOOGLE_CLOUD_PLATFORM, OPENSTACK, ORACLE and UNRECOGNIZED.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (ctc *CloudType) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ctc.Unknowns) > 0 {
		data, err := json.Marshal(ctc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = ctc.Negate
	result["operator"] = string(ctc.Operator)
	if ctc.Value != nil {
		result["value"] = ctc.Value.String()
	}
	return result, nil
}

func (ctc *CloudType) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ctc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ctc.Unknowns); err != nil {
			return err
		}
		delete(ctc.Unknowns, "type")
		delete(ctc.Unknowns, "negate")
		delete(ctc.Unknowns, "operator")
		delete(ctc.Unknowns, "value")
		if len(ctc.Unknowns) == 0 {
			ctc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		ctc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		ctc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		ctc.Operator = cloud_type.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		ctc.Value = cloud_type.Value(value.(string)).Ref()
	}
	return nil
}

func (ctc *CloudType) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ctc.Unknowns) > 0 {
		for k, v := range ctc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(ctc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ctc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&ctc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if ctc.Value != nil {
		rawMessage, err := json.Marshal(ctc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (ctc *CloudType) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	ctc.Type = ctc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &ctc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &ctc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &ctc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		ctc.Unknowns = m
	}
	return nil
}
