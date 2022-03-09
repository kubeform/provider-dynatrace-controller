package web

import "github.com/dtcookie/hcl"

type MetaDataCaptureSettings []*MetaDataCapturing

func (me *MetaDataCaptureSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"capture": {
			Type:        hcl.TypeList,
			Description: "Java script agent meta data capture settings",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(MetaDataCapturing).Schema()},
		},
	}
}

func (me MetaDataCaptureSettings) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["capture"] = entries
	}
	return result, nil
}

func (me *MetaDataCaptureSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("capture", me); err != nil {
		return err
	}
	return nil
}

type MetaDataCapturing struct {
	Type           MetaDataCapturingType `json:"type"`               // The type of the meta data to capture. Possible values are `COOKIE`, `CSS_SELECTOR`, `JAVA_SCRIPT_FUNCTION`, `JAVA_SCRIPT_VARIABLE`, `META_TAG` and `QUERY_STRING`.
	CapturingName  string                `json:"capturingName"`      // The name of the meta data to capture
	Name           string                `json:"name"`               // Name for displaying the captured values in Dynatrace
	UniqueID       *int32                `json:"uniqueId,omitempty"` // The unique ID of the meta data to capture
	PublicMetadata bool                  `json:"publicMetadata"`     // `true` if this metadata should be captured regardless of the privacy settings, `false` otherwise
	UseLastValue   bool                  `json:"useLastValue"`       // `true` if the last captured value should be used for this metadata. By default the first value will be used.
}

func (me *MetaDataCapturing) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the meta data to capture. Possible values are `COOKIE`, `CSS_SELECTOR`, `JAVA_SCRIPT_FUNCTION`, `JAVA_SCRIPT_VARIABLE`, `META_TAG` and `QUERY_STRING`.",
			Required:    true,
		},
		"capturing_name": {
			Type:        hcl.TypeString,
			Description: "The name of the meta data to capture",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Name for displaying the captured values in Dynatrace",
			Required:    true,
		},
		"unique_id": {
			Type:        hcl.TypeInt,
			Description: "The unique ID of the meta data to capture",
			Optional:    true,
		},
		"public_metadata": {
			Type:        hcl.TypeBool,
			Description: "`true` if this metadata should be captured regardless of the privacy settings, `false` otherwise",
			Optional:    true,
		},
		"use_last_value": {
			Type:        hcl.TypeBool,
			Description: "`true` if the last captured value should be used for this metadata. By default the first value will be used.",
			Optional:    true,
		},
	}
}

func (me *MetaDataCapturing) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"type":            me.Type,
		"capturing_name":  me.CapturingName,
		"name":            me.Name,
		"unique_id":       me.UniqueID,
		"public_metadata": me.PublicMetadata,
		"use_last_value":  me.UseLastValue,
	})
}

func (me *MetaDataCapturing) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":            &me.Type,
		"capturing_name":  &me.CapturingName,
		"name":            &me.Name,
		"unique_id":       &me.UniqueID,
		"public_metadata": &me.PublicMetadata,
		"use_last_value":  &me.UseLastValue,
	})
}
