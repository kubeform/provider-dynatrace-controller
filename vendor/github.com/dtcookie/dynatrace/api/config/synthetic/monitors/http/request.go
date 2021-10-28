package http

import (
	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/http/validation"
	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/request"
	"github.com/dtcookie/hcl"
)

type Requests []*Request

func (me *Requests) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"request": {
			Type:        hcl.TypeList,
			Description: "A HTTP request to be performed by the monitor.",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Request).Schema()},
		},
	}
}

type Request struct {
	Description    *string              `json:"description,omitempty"`   // A short description of the event to appear in the web UI
	URL            string               `json:"url"`                     // The URL to check
	Method         string               `json:"method"`                  // The HTTP method of the request
	RequestBody    *string              `json:"requestBody,omitempty"`   // The body of the HTTP request—you need to escape all JSON characters. \n\n Is set to null if the request method is GET, HEAD, or OPTIONS.
	Validation     *validation.Settings `json:"validation,omitempty"`    // Validation helps you verify that your HTTP monitor loads the expected content
	Configuration  *request.Config      `json:"configuration,omitempty"` // The setup of the monitor
	PreProcessing  *string              `json:"preProcessingScript,omitempty"`
	PostProcessing *string              `json:"postProcessingScript,omitempty"`
}

func (me *Request) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the event to appear in the web UI.",
			Optional:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL to check.",
			Required:    true,
		},
		"method": {
			Type:        hcl.TypeString,
			Description: "The HTTP method of the request.",
			Required:    true,
		},
		"body": {
			Type:        hcl.TypeString,
			Description: "The body of the HTTP request.",
			Optional:    true,
		},
		"pre_processing": {
			Type:        hcl.TypeString,
			Description: "Javascript code to execute before sending the request.",
			Optional:    true,
		},
		"post_processing": {
			Type:        hcl.TypeString,
			Description: "Javascript code to execute after sending the request.",
			Optional:    true,
		},
		"validation": {
			Type:        hcl.TypeList,
			Description: "Validation helps you verify that your HTTP monitor loads the expected content",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(validation.Settings).Schema()},
		},
		"configuration": {
			Type:        hcl.TypeList,
			Description: "The setup of the monitor",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(request.Config).Schema()},
		},
	}
}

func (me *Request) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Description != nil && len(*me.Description) > 0 {
		result["description"] = *me.Description
	}
	result["url"] = me.URL
	result["method"] = me.Method
	if me.RequestBody != nil && len(*me.RequestBody) > 0 {
		result["body"] = *me.RequestBody
	}
	if me.PreProcessing != nil && len(*me.PreProcessing) > 0 {
		result["pre_processing"] = *me.PreProcessing
	}
	if me.PostProcessing != nil && len(*me.PostProcessing) > 0 {
		result["post_processing"] = *me.PostProcessing
	}
	if me.Validation != nil {
		if marshalled, err := me.Validation.MarshalHCL(); err == nil {
			result["validation"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Configuration != nil {
		if marshalled, err := me.Configuration.MarshalHCL(); err == nil {
			result["configuration"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Request) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("description", &me.Description); err != nil {
		return err
	}
	if err := decoder.Decode("url", &me.URL); err != nil {
		return err
	}
	if err := decoder.Decode("method", &me.Method); err != nil {
		return err
	}
	if err := decoder.Decode("body", &me.RequestBody); err != nil {
		return err
	}
	if err := decoder.Decode("pre_processing", &me.PreProcessing); err != nil {
		return err
	}
	if err := decoder.Decode("post_processing", &me.PostProcessing); err != nil {
		return err
	}
	if result, ok := decoder.GetOk("validation.#"); ok && result.(int) == 1 {
		me.Validation = new(validation.Settings)
		if err := me.Validation.UnmarshalHCL(hcl.NewDecoder(decoder, "validation", 0)); err != nil {
			return err
		}
	}
	if result, ok := decoder.GetOk("configuration.#"); ok && result.(int) == 1 {
		me.Configuration = new(request.Config)
		if err := me.Configuration.UnmarshalHCL(hcl.NewDecoder(decoder, "configuration", 0)); err != nil {
			return err
		}
	}
	return nil
}
