package comparisoninfo

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// ComparisonInfo Type-specific comparison for attributes. The actual set of fields depends on the `type` of the comparison.
// See the [Service metrics API - JSON models](https://dt-url.net/9803svb) help topic for example models of every notification type.
type ComparisonInfo interface {
	GetType() Type
	SetNegate(bool)
	IsNegate() bool
}

// BaseComparisonInfo Type-specific comparison for attributes. The actual set of fields depends on the `type` of the comparison.
// See the [Service metrics API - JSON models](https://dt-url.net/9803svb) help topic for example models of every notification type.
type BaseComparisonInfo struct {
	Negate   bool                       `json:"negate"` // Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**.
	Type     Type                       `json:"type"`   // Defines the actual set of fields depending on the value. See one of the following objects:  * `STRING` -> StringComparisonInfo  * `NUMBER` -> NumberComparisonInfo  * `BOOLEAN` -> BooleanComparisonInfo  * `HTTP_METHOD` -> HttpMethodComparisonInfo  * `STRING_REQUEST_ATTRIBUTE` -> StringRequestAttributeComparisonInfo  * `NUMBER_REQUEST_ATTRIBUTE` -> NumberRequestAttributeComparisonInfo  * `ZOS_CALL_TYPE` -> ZosComparisonInfo  * `IIB_INPUT_NODE_TYPE` -> IIBInputNodeTypeComparisonInfo  * `ESB_INPUT_NODE_TYPE` -> ESBInputNodeTypeComparisonInfo  * `FAILED_STATE` -> FailedStateComparisonInfo  * `FLAW_STATE` -> FlawStateComparisonInfo  * `FAILURE_REASON` -> FailureReasonComparisonInfo  * `HTTP_STATUS_CLASS` -> HttpStatusClassComparisonInfo  * `TAG` -> TagComparisonInfo  * `FAST_STRING` -> FastStringComparisonInfo  * `SERVICE_TYPE` -> ServiceTypeComparisonInfo
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *BaseComparisonInfo) SetNegate(negate bool) {
	me.Negate = negate
}

func (me *BaseComparisonInfo) IsNegate() bool {
	return me.Negate
}

func (me *BaseComparisonInfo) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Defines the actual set of fields depending on the value",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *BaseComparisonInfo) GetType() Type {
	return me.Type
}

func (me *BaseComparisonInfo) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	fmt.Printf("BaseComparisonInfo.Unknowns: %v", me.Unknowns)
	return properties.EncodeAll(map[string]interface{}{
		"type":     me.Type,
		"unknowns": me.Unknowns,
	})
}

func (me *BaseComparisonInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":     &me.Type,
		"unknowns": &me.Unknowns,
	})
}

func (me *BaseComparisonInfo) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"negate": me.Negate,
		"type":   me.Type,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *BaseComparisonInfo) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"negate": &me.Negate,
		"type":   &me.Type,
	}); err != nil {
		return err
	}
	me.Unknowns = properties
	return nil
}
