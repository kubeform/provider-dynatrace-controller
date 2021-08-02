package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// DataSource has no documentation
type DataSource struct {
	CapturingAndStorageLocation *CapturingAndStorageLocation `json:"capturingAndStorageLocation,omitempty"` // Specifies the location where the values are captured and stored.  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.   Not applicable in other cases.   If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed.
	Scope                       *ScopeConditions             `json:"scope,omitempty"`                       // Conditions for data capturing.
	ParameterName               *string                      `json:"parameterName,omitempty"`               // The name of the web request parameter to capture.  Required if the **source** is one of the following: `POST_PARAMETER`, `GET_PARAMETER`, `REQUEST_HEADER`, `RESPONSE_HEADER`, `CUSTOM_ATTRIBUTE`.  Not applicable in other cases.
	IIBMethodNodeCondition      *ValueCondition              `json:"iibMethodNodeCondition,omitempty"`      // IBM integration bus label node name condition for which the value is captured.
	Methods                     []*CapturedMethod            `json:"methods,omitempty"`                     // The method specification if the **source** value is `METHOD_PARAM`.   Not applicable in other cases.
	SessionAttributeTechnology  *SessionAttributeTechnology  `json:"sessionAttributeTechnology,omitempty"`  // The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases.
	Technology                  *Technology                  `json:"technology,omitempty"`                  // The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases.
	ValueProcessing             *ValueProcessing             `json:"valueProcessing,omitempty"`             // Process values as specified.
	CICSSDKMethodNodeCondition  *ValueCondition              `json:"cicsSDKMethodNodeCondition,omitempty"`  // IBM integration bus label node name condition for which the value is captured.
	Enabled                     *bool                        `json:"enabled"`                               // The data source is enabled (`true`) or disabled (`false`).
	Source                      Source                       `json:"source"`                                // The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**.
	IIBLabelMethodNodeCondition *ValueCondition              `json:"iibLabelMethodNodeCondition,omitempty"` // IBM integration bus label node name condition for which the value is captured.
	IIBNodeType                 *IIBNodeType                 `json:"iibNodeType,omitempty"`                 // The IBM integration bus node type for which the value is captured.  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.  Not applicable in other cases.
	Unknowns                    map[string]json.RawMessage   `json:"-"`
}

func (me *DataSource) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"capturing_and_storage_location": {
			Type:        hcl.TypeString,
			Description: "Specifies the location where the values are captured and stored.  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.   Not applicable in other cases.   If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed",
			Optional:    true,
		},
		"scope": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Conditions for data capturing",
			Elem: &hcl.Resource{
				Schema: new(ScopeConditions).Schema(),
			},
		},
		"parameter_name": {
			Type:        hcl.TypeString,
			Description: "The name of the web request parameter to capture.  Required if the **source** is one of the following: `POST_PARAMETER`, `GET_PARAMETER`, `REQUEST_HEADER`, `RESPONSE_HEADER`, `CUSTOM_ATTRIBUTE`.  Not applicable in other cases",
			Optional:    true,
		},
		"iib_method_node_condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "IBM integration bus label node name condition for which the value is captured",
			Elem: &hcl.Resource{
				Schema: new(ValueCondition).Schema(),
			},
		},
		"methods": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The method specification if the **source** value is `METHOD_PARAM`.   Not applicable in other cases",
			Elem: &hcl.Resource{
				Schema: new(CapturedMethod).Schema(),
			},
		},
		"session_attribute_technology": {
			Type:        hcl.TypeString,
			Description: "The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases",
			Optional:    true,
		},
		"technology": {
			Type:        hcl.TypeString,
			Description: "The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases",
			Optional:    true,
		},
		"value_processing": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Process values as specified",
			Elem: &hcl.Resource{
				Schema: new(ValueProcessing).Schema(),
			},
		},
		"cics_sdk_method_node_condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "IBM integration bus label node name condition for which the value is captured",
			Elem: &hcl.Resource{
				Schema: new(ValueCondition).Schema(),
			},
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The data source is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"source": {
			Type:        hcl.TypeString,
			Description: "The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**",
			Required:    true,
		},
		"iib_label_method_node_condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "IBM integration bus label node name condition for which the value is captured",
			Elem: &hcl.Resource{
				Schema: new(ValueCondition).Schema(),
			},
		},
		"iib_node_type": {
			Type:        hcl.TypeString,
			Description: "The IBM integration bus node type for which the value is captured.  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.  Not applicable in other cases",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DataSource) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.CapturingAndStorageLocation != nil {
		result["capturing_and_storage_location"] = string(*me.CapturingAndStorageLocation)
	}
	if me.Scope != nil {
		if marshalled, err := me.Scope.MarshalHCL(); err == nil {
			result["scope"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ParameterName != nil {
		result["parameter_name"] = string(*me.ParameterName)
	}
	if me.IIBMethodNodeCondition != nil {
		if marshalled, err := me.IIBMethodNodeCondition.MarshalHCL(); err == nil {
			result["iib_method_node_condition"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Methods) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Methods {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["methods"] = entries
	}
	if me.SessionAttributeTechnology != nil {
		result["session_attribute_technology"] = string(*me.SessionAttributeTechnology)
	}
	if me.Technology != nil {
		result["technology"] = string(*me.Technology)
	}
	if me.ValueProcessing != nil && !me.ValueProcessing.IsZero() {
		if marshalled, err := me.ValueProcessing.MarshalHCL(); err == nil {
			result["value_processing"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.CICSSDKMethodNodeCondition != nil {
		if marshalled, err := me.CICSSDKMethodNodeCondition.MarshalHCL(); err == nil {
			result["cics_sdk_method_node_condition"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Enabled != nil {
		result["enabled"] = opt.Bool(me.Enabled)
	}
	result["source"] = string(me.Source)
	if me.IIBLabelMethodNodeCondition != nil {
		if marshalled, err := me.IIBLabelMethodNodeCondition.MarshalHCL(); err == nil {
			result["iib_label_method_node_condition"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.IIBNodeType != nil {
		result["iib_node_type"] = string(*me.IIBNodeType)
	}
	return result, nil
}

func (me *DataSource) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "capturing_and_storage_location")
		delete(me.Unknowns, "scope")
		delete(me.Unknowns, "parameter_name")
		delete(me.Unknowns, "iib_method_node_condition")
		delete(me.Unknowns, "methods")
		delete(me.Unknowns, "session_attribute_technology")
		delete(me.Unknowns, "technology")
		delete(me.Unknowns, "value_processing")
		delete(me.Unknowns, "cics_sdk_method_node_condition")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "source")
		delete(me.Unknowns, "iib_label_method_node_condition")
		delete(me.Unknowns, "iib_node_type")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("capturing_and_storage_location"); ok {
		me.CapturingAndStorageLocation = CapturingAndStorageLocation(value.(string)).Ref()
	}
	if _, ok := decoder.GetOk("scope.#"); ok {
		me.Scope = new(ScopeConditions)
		if err := me.Scope.UnmarshalHCL(hcl.NewDecoder(decoder, "scope", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("parameter_name"); ok {
		me.ParameterName = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("iib_method_node_condition.#"); ok {
		me.IIBMethodNodeCondition = new(ValueCondition)
		if err := me.IIBMethodNodeCondition.UnmarshalHCL(hcl.NewDecoder(decoder, "iib_method_node_condition", 0)); err != nil {
			return err
		}
	}
	if result, ok := decoder.GetOk("methods.#"); ok {
		me.Methods = []*CapturedMethod{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(CapturedMethod)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "methods", idx)); err != nil {
				return err
			}
			me.Methods = append(me.Methods, entry)
		}
	}
	if value, ok := decoder.GetOk("session_attribute_technology"); ok {
		me.SessionAttributeTechnology = SessionAttributeTechnology(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("technology"); ok {
		me.Technology = Technology(value.(string)).Ref()
	}
	if _, ok := decoder.GetOk("value_processing.#"); ok {
		me.ValueProcessing = new(ValueProcessing)
		if err := me.ValueProcessing.UnmarshalHCL(hcl.NewDecoder(decoder, "value_processing", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("cics_sdk_method_node_condition.#"); ok {
		me.CICSSDKMethodNodeCondition = new(ValueCondition)
		if err := me.CICSSDKMethodNodeCondition.UnmarshalHCL(hcl.NewDecoder(decoder, "cics_sdk_method_node_condition", 0)); err != nil {
			return err
		}
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("source"); ok {
		me.Source = Source(value.(string))
	}
	if _, ok := decoder.GetOk("iib_label_method_node_condition.#"); ok {
		me.IIBLabelMethodNodeCondition = new(ValueCondition)
		if err := me.IIBLabelMethodNodeCondition.UnmarshalHCL(hcl.NewDecoder(decoder, "iib_label_method_node_condition", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("iib_node_type"); ok {
		me.IIBNodeType = IIBNodeType(value.(string)).Ref()
	}
	return nil
}

func (me *DataSource) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("capturingAndStorageLocation", me.CapturingAndStorageLocation); err != nil {
		return nil, err
	}
	if err := m.Marshal("scope", me.Scope); err != nil {
		return nil, err
	}
	if err := m.Marshal("parameterName", me.ParameterName); err != nil {
		return nil, err
	}
	if err := m.Marshal("iibMethodNodeCondition", me.IIBMethodNodeCondition); err != nil {
		return nil, err
	}
	if err := m.Marshal("methods", me.Methods); err != nil {
		return nil, err
	}
	if err := m.Marshal("sessionAttributeTechnology", me.SessionAttributeTechnology); err != nil {
		return nil, err
	}
	if err := m.Marshal("technology", me.Technology); err != nil {
		return nil, err
	}
	if err := m.Marshal("valueProcessing", me.ValueProcessing); err != nil {
		return nil, err
	}
	if err := m.Marshal("cicsSDKMethodNodeCondition", me.CICSSDKMethodNodeCondition); err != nil {
		return nil, err
	}
	if err := m.Marshal("enabled", me.Enabled); err != nil {
		return nil, err
	}
	if err := m.Marshal("source", me.Source); err != nil {
		return nil, err
	}
	if err := m.Marshal("iibLabelMethodNodeCondition", me.IIBLabelMethodNodeCondition); err != nil {
		return nil, err
	}
	if err := m.Marshal("iibNodeType", me.IIBNodeType); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *DataSource) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("capturingAndStorageLocation", &me.CapturingAndStorageLocation); err != nil {
		return err
	}
	if err := m.Unmarshal("scope", &me.Scope); err != nil {
		return err
	}
	if err := m.Unmarshal("parameterName", &me.ParameterName); err != nil {
		return err
	}
	if err := m.Unmarshal("iibMethodNodeCondition", &me.IIBMethodNodeCondition); err != nil {
		return err
	}
	if err := m.Unmarshal("methods", &me.Methods); err != nil {
		return err
	}
	if err := m.Unmarshal("sessionAttributeTechnology", &me.SessionAttributeTechnology); err != nil {
		return err
	}
	if err := m.Unmarshal("technology", &me.Technology); err != nil {
		return err
	}
	if err := m.Unmarshal("enabled", &me.Enabled); err != nil {
		return err
	}
	if err := m.Unmarshal("valueProcessing", &me.ValueProcessing); err != nil {
		return err
	}
	if err := m.Unmarshal("cicsSDKMethodNodeCondition", &me.CICSSDKMethodNodeCondition); err != nil {
		return err
	}
	if err := m.Unmarshal("source", &me.Source); err != nil {
		return err
	}
	if err := m.Unmarshal("iibLabelMethodNodeCondition", &me.IIBLabelMethodNodeCondition); err != nil {
		return err
	}
	if err := m.Unmarshal("iibNodeType", &me.IIBNodeType); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// CapturingAndStorageLocation Specifies the location where the values are captured and stored.
//  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.
//  Not applicable in other cases.
//  If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed.
type CapturingAndStorageLocation string

func (me CapturingAndStorageLocation) Ref() *CapturingAndStorageLocation {
	return &me
}

// CapturingAndStorageLocations offers the known enum values
var CapturingAndStorageLocations = struct {
	CaptureAndStoreOnBoth        CapturingAndStorageLocation
	CaptureAndStoreOnClient      CapturingAndStorageLocation
	CaptureAndStoreOnServer      CapturingAndStorageLocation
	CaptureOnClientStoreOnServer CapturingAndStorageLocation
}{
	"CAPTURE_AND_STORE_ON_BOTH",
	"CAPTURE_AND_STORE_ON_CLIENT",
	"CAPTURE_AND_STORE_ON_SERVER",
	"CAPTURE_ON_CLIENT_STORE_ON_SERVER",
}

// SessionAttributeTechnology The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases.
type SessionAttributeTechnology string

func (me SessionAttributeTechnology) Ref() *SessionAttributeTechnology {
	return &me
}

// SessionAttributeTechnologys offers the known enum values
var SessionAttributeTechnologys = struct {
	ASPNet     SessionAttributeTechnology
	ASPNetCore SessionAttributeTechnology
	Java       SessionAttributeTechnology
}{
	"ASP_NET",
	"ASP_NET_CORE",
	"JAVA",
}

// Technology The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases.
type Technology string

func (me Technology) Ref() *Technology {
	return &me
}

// Technologys offers the known enum values
var Technologys = struct {
	DotNet Technology
	Java   Technology
	PHP    Technology
}{
	"DOTNET",
	"JAVA",
	"PHP",
}

// Source The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**.
type Source string

// Sources offers the known enum values
var Sources = struct {
	CICSSdk          Source
	ClientIP         Source
	CustomAttribute  Source
	IibLabel         Source
	IibNode          Source
	MethodParam      Source
	PostParameter    Source
	QueryParameter   Source
	RequestHeader    Source
	ResponseHeader   Source
	SessionAttribute Source
	URI              Source
	URIPath          Source
}{
	"CICS_SDK",
	"CLIENT_IP",
	"CUSTOM_ATTRIBUTE",
	"IIB_LABEL",
	"IIB_NODE",
	"METHOD_PARAM",
	"POST_PARAMETER",
	"QUERY_PARAMETER",
	"REQUEST_HEADER",
	"RESPONSE_HEADER",
	"SESSION_ATTRIBUTE",
	"URI",
	"URI_PATH",
}

// IIBNodeType The IBM integration bus node type for which the value is captured.
//  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.
//  Not applicable in other cases.
type IIBNodeType string

func (me IIBNodeType) Ref() *IIBNodeType {
	return &me
}

// IIBNodeTypes offers the known enum values
var IIBNodeTypes = struct {
	AggregateControlNode       IIBNodeType
	AggregateReplyNode         IIBNodeType
	AggregateRequestNode       IIBNodeType
	CallableFlowReplyNode      IIBNodeType
	CollectorNode              IIBNodeType
	ComputeNode                IIBNodeType
	DatabaseNode               IIBNodeType
	DecisionServiceNode        IIBNodeType
	DotNetComputeNode          IIBNodeType
	FileReadNode               IIBNodeType
	FilterNode                 IIBNodeType
	FlowOrderNode              IIBNodeType
	GroupCompleteNode          IIBNodeType
	GroupGatherNode            IIBNodeType
	GroupScatterNode           IIBNodeType
	HTTPHeader                 IIBNodeType
	JavaComputeNode            IIBNodeType
	JmsClientReceive           IIBNodeType
	JmsClientReplyNode         IIBNodeType
	JmsHeader                  IIBNodeType
	MqGetNode                  IIBNodeType
	MqOutputNode               IIBNodeType
	PassthruNode               IIBNodeType
	ResetContentDescriptorNode IIBNodeType
	ReSequenceNode             IIBNodeType
	RouteNode                  IIBNodeType
	SAPReplyNode               IIBNodeType
	ScaReplyNode               IIBNodeType
	SecurityPep                IIBNodeType
	SequenceNode               IIBNodeType
	SoapExtractNode            IIBNodeType
	SoapReplyNode              IIBNodeType
	SoapWrapperNode            IIBNodeType
	SrRetrieveEntityNode       IIBNodeType
	SrRetrieveItServiceNode    IIBNodeType
	ThrowNode                  IIBNodeType
	TraceNode                  IIBNodeType
	TryCatchNode               IIBNodeType
	ValidateNode               IIBNodeType
	WsReplyNode                IIBNodeType
	XslMqsiNode                IIBNodeType
}{
	"AGGREGATE_CONTROL_NODE",
	"AGGREGATE_REPLY_NODE",
	"AGGREGATE_REQUEST_NODE",
	"CALLABLE_FLOW_REPLY_NODE",
	"COLLECTOR_NODE",
	"COMPUTE_NODE",
	"DATABASE_NODE",
	"DECISION_SERVICE_NODE",
	"DOT_NET_COMPUTE_NODE",
	"FILE_READ_NODE",
	"FILTER_NODE",
	"FLOW_ORDER_NODE",
	"GROUP_COMPLETE_NODE",
	"GROUP_GATHER_NODE",
	"GROUP_SCATTER_NODE",
	"HTTP_HEADER",
	"JAVA_COMPUTE_NODE",
	"JMS_CLIENT_RECEIVE",
	"JMS_CLIENT_REPLY_NODE",
	"JMS_HEADER",
	"MQ_GET_NODE",
	"MQ_OUTPUT_NODE",
	"PASSTHRU_NODE",
	"RESET_CONTENT_DESCRIPTOR_NODE",
	"RE_SEQUENCE_NODE",
	"ROUTE_NODE",
	"SAP_REPLY_NODE",
	"SCA_REPLY_NODE",
	"SECURITY_PEP",
	"SEQUENCE_NODE",
	"SOAP_EXTRACT_NODE",
	"SOAP_REPLY_NODE",
	"SOAP_WRAPPER_NODE",
	"SR_RETRIEVE_ENTITY_NODE",
	"SR_RETRIEVE_IT_SERVICE_NODE",
	"THROW_NODE",
	"TRACE_NODE",
	"TRY_CATCH_NODE",
	"VALIDATE_NODE",
	"WS_REPLY_NODE",
	"XSL_MQSI_NODE",
}
