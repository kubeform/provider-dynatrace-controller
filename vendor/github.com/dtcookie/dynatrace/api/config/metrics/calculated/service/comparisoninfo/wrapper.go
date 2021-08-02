package comparisoninfo

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

type Wrapper struct {
	Negate     bool
	Comparison ComparisonInfo
}

func (me *Wrapper) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"negate": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**",
		},
		"boolean": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Boolean Comparison for `BOOLEAN` attributes",
			Elem:        &hcl.Resource{Schema: new(Boolean).Schema()},
		},
		"esb_input_node_type": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Type-specific comparison information for attributes of type 'ESB_INPUT_NODE_TYPE'",
			Elem:        &hcl.Resource{Schema: new(ESBInputNodeType).Schema()},
		},
		"failed_state": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `FAILED_STATE` attributes",
			Elem:        &hcl.Resource{Schema: new(FailedState).Schema()},
		},
		"failure_reason": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `FAILURE_REASON` attributes",
			Elem:        &hcl.Resource{Schema: new(FailureReason).Schema()},
		},
		"fast_string": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `FAST_STRING` attributes. Use it for all service property attributes",
			Elem:        &hcl.Resource{Schema: new(FastString).Schema()},
		},
		"flaw_state": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `FLAW_STATE` attributes",
			Elem:        &hcl.Resource{Schema: new(FlawState).Schema()},
		},
		"http_method": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `HTTP_METHOD` attributes",
			Elem:        &hcl.Resource{Schema: new(HTTPMethod).Schema()},
		},
		"http_status_class": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `HTTP_STATUS_CLASS` attributes",
			Elem:        &hcl.Resource{Schema: new(HTTPStatusClass).Schema()},
		},
		"iib_input_node_type": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `IIB_INPUT_NODE_TYPE` attributes",
			Elem:        &hcl.Resource{Schema: new(IIBInputNodeType).Schema()},
		},
		"number": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `NUMBER` attributes",
			Elem:        &hcl.Resource{Schema: new(Number).Schema()},
		},
		"number_request_attribute": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `NUMBER_REQUEST_ATTRIBUTE` attributes",
			Elem:        &hcl.Resource{Schema: new(NumberRequestAttribute).Schema()},
		},
		"service_type": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `SERVICE_TYPE` attributes",
			Elem:        &hcl.Resource{Schema: new(ServiceType).Schema()},
		},
		"string": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `STRING` attributes",
			Elem:        &hcl.Resource{Schema: new(String).Schema()},
		},
		"string_request_attribute": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `STRING_REQUEST_ATTRIBUTE` attributes",
			Elem:        &hcl.Resource{Schema: new(StringRequestAttribute).Schema()},
		},
		"tag": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `TAG` attributes",
			Elem:        &hcl.Resource{Schema: new(Tag).Schema()},
		},
		"zos_call_type": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `ZOS_CALL_TYPE` attributes",
			Elem:        &hcl.Resource{Schema: new(ZOSCallType).Schema()},
		},
		"generic": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Comparison for `NUMBER` attributes",
			Elem:        &hcl.Resource{Schema: new(BaseComparisonInfo).Schema()},
		},
	}
}

func (me *Wrapper) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	properties.Encode("negate", me.Comparison.IsNegate())
	switch cmp := me.Comparison.(type) {
	case *Boolean:
		if err := properties.Encode("boolean", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ESBInputNodeType:
		if err := properties.Encode("esb_input_node_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FailedState:
		if err := properties.Encode("failed_state", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FailureReason:
		if err := properties.Encode("failure_reason", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FastString:
		if err := properties.Encode("fast_string", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *FlawState:
		if err := properties.Encode("flaw_state", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *HTTPMethod:
		if err := properties.Encode("http_method", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *HTTPStatusClass:
		if err := properties.Encode("http_status_class", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *IIBInputNodeType:
		if err := properties.Encode("iib_input_node_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *NumberRequestAttribute:
		if err := properties.Encode("number_request_attribute", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *Number:
		if err := properties.Encode("number", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ServiceType:
		if err := properties.Encode("service_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *StringRequestAttribute:
		if err := properties.Encode("string_request_attribute", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *String:
		if err := properties.Encode("string", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *Tag:
		if err := properties.Encode("tag", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *ZOSCallType:
		if err := properties.Encode("zos_call_type", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	case *BaseComparisonInfo:
		if err := properties.Encode("generic", cmp); err != nil {
			return nil, err
		}
		return properties, nil
	default:
		return nil, fmt.Errorf("cannot HCL marshal objects (xxx) of type %T", cmp)
	}
}

func (me *Wrapper) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("negate", &me.Negate); err != nil {
		return err
	}
	var err error
	var cmp interface{}
	if cmp, err = decoder.DecodeAny(map[string]interface{}{
		"boolean":                  new(Boolean),
		"esb_input_node_type":      new(ESBInputNodeType),
		"failed_state":             new(FailedState),
		"failure_reason":           new(FailureReason),
		"fast_string":              new(FastString),
		"flaw_state":               new(FlawState),
		"http_method":              new(HTTPMethod),
		"http_status_class":        new(HTTPStatusClass),
		"iib_input_node_type":      new(IIBInputNodeType),
		"number":                   new(Number),
		"number_request_attribute": new(NumberRequestAttribute),
		"service_type":             new(ServiceType),
		"string":                   new(String),
		"string_request_attribute": new(StringRequestAttribute),
		"tag":                      new(Tag),
		"zos_call_type":            new(ZOSCallType),
		"generic":                  new(BaseComparisonInfo)}); err != nil {
		return err
	}
	if cmp != nil {
		me.Comparison = cmp.(ComparisonInfo)
		me.Comparison.SetNegate(me.Negate)
	}
	return nil
}

func (me *Wrapper) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	var compType string
	if err := properties.UnmarshalAll(map[string]interface{}{
		"negate": &me.Negate,
		"type":   &compType,
	}); err != nil {
		return err
	}
	switch compType {
	case "STRING":
		cfg := new(String)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "NUMBER":
		cfg := new(Number)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "BOOLEAN":
		cfg := new(Boolean)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "HTTP_METHOD":
		cfg := new(HTTPMethod)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "STRING_REQUEST_ATTRIBUTE":
		cfg := new(StringRequestAttribute)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "NUMBER_REQUEST_ATTRIBUTE":
		cfg := new(NumberRequestAttribute)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "ZOS_CALL_TYPE":
		cfg := new(ZOSCallType)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "IIB_INPUT_NODE_TYPE":
		cfg := new(IIBInputNodeType)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "ESB_INPUT_NODE_TYPE":
		cfg := new(ESBInputNodeType)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "FAILED_STATE":
		cfg := new(FailedState)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "FLAW_STATE":
		cfg := new(FlawState)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "FAILURE_REASON":
		cfg := new(FailureReason)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "HTTP_STATUS_CLASS":
		cfg := new(HTTPStatusClass)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "TAG":
		cfg := new(Tag)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "FAST_STRING":
		cfg := new(FastString)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	case "SERVICE_TYPE":
		cfg := new(ServiceType)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	default:
		cfg := new(BaseComparisonInfo)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.Negate = me.Negate
		me.Comparison = cfg
	}
	return nil
}
