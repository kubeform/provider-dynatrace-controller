package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// HTTPStatusClass Comparison for `HTTP_STATUS_CLASS` attributes.
type HTTPStatusClass struct {
	BaseComparisonInfo
	Values     []HTTPStatusClassValue    `json:"values,omitempty"` // The values to compare to.
	Comparison HTTPStatusClassComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *HTTPStatusClassValue     `json:"value,omitempty"`  // The value to compare to.
}

func (me *HTTPStatusClass) GetType() Type {
	return Types.HTTPStatusClass
}

func (me *HTTPStatusClass) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to. Possible values are `C_1XX`, `C_2XX`, `C_3XX`, `C_4XX`, `C_5XX` and `NO_RESPONSE`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to. Possible values are `C_1XX`, `C_2XX`, `C_3XX`, `C_4XX`, `C_5XX` and `NO_RESPONSE`",
		},
		"operator": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `EQUALS`, `EQUALS_ANY_OF` and `EXISTS`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *HTTPStatusClass) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":   me.Values,
		"value":    me.Value,
		"operator": me.Comparison,
		"unknowns": me.Unknowns,
	})
}

func (me *HTTPStatusClass) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *HTTPStatusClass) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":       me.GetType(),
		"negate":     me.Negate,
		"values":     me.Values,
		"value":      me.Value,
		"comparison": me.Comparison,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *HTTPStatusClass) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":     &me.Negate,
		"values":     &me.Values,
		"value":      &me.Value,
		"comparison": &me.Comparison,
	})
}

// HTTPStatusClassComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type HTTPStatusClassComparison string

// HTTPStatusClassComparisons offers the known enum values
var HTTPStatusClassComparisons = struct {
	Equals      HTTPStatusClassComparison
	EqualsAnyOf HTTPStatusClassComparison
	Exists      HTTPStatusClassComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// HTTPStatusClassValue The value to compare to.
type HTTPStatusClassValue string

// HTTPStatusClassValues offers the known enum values
var HTTPStatusClassValues = struct {
	C1xx       HTTPStatusClassValue
	C2xx       HTTPStatusClassValue
	C3xx       HTTPStatusClassValue
	C4xx       HTTPStatusClassValue
	C5xx       HTTPStatusClassValue
	NoResponse HTTPStatusClassValue
}{
	"C_1XX",
	"C_2XX",
	"C_3XX",
	"C_4XX",
	"C_5XX",
	"NO_RESPONSE",
}
