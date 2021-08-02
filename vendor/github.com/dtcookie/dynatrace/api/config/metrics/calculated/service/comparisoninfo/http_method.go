package comparisoninfo

import (
	"encoding/json"
	"log"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// HTTPMethod Comparison for `HTTP_METHOD` attributes.
type HTTPMethod struct {
	BaseComparisonInfo
	Comparison HTTPMethodComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *HTTPMethodValue     `json:"value,omitempty"`  // The value to compare to.
	Values     []HTTPMethodValue    `json:"values,omitempty"` // The values to compare to.
}

func (me *HTTPMethod) GetType() Type {
	return Types.HTTPMethod
}

func (me *HTTPMethod) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to. Possible values are `CONNECT`, `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to. Possible values are `CONNECT`, `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`",
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

func (me *HTTPMethod) MarshalHCL() (map[string]interface{}, error) {
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

func (me *HTTPMethod) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
	log.Printf("values: %v", me.Values)
	return err
}

func (me *HTTPMethod) MarshalJSON() ([]byte, error) {
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

func (me *HTTPMethod) UnmarshalJSON(data []byte) error {
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

// HTTPMethodComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type HTTPMethodComparison string

// HTTPMethodComparisons offers the known enum values
var HTTPMethodComparisons = struct {
	Equals      HTTPMethodComparison
	EqualsAnyOf HTTPMethodComparison
	Exists      HTTPMethodComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// HTTPMethodValue The values to compare to.
type HTTPMethodValue string

// HTTPMethodValues offers the known enum values
var HTTPMethodValues = struct {
	Connect HTTPMethodValue
	Delete  HTTPMethodValue
	Get     HTTPMethodValue
	Head    HTTPMethodValue
	Options HTTPMethodValue
	Patch   HTTPMethodValue
	Post    HTTPMethodValue
	Put     HTTPMethodValue
	Trace   HTTPMethodValue
}{
	"CONNECT",
	"DELETE",
	"GET",
	"HEAD",
	"OPTIONS",
	"PATCH",
	"POST",
	"PUT",
	"TRACE",
}
