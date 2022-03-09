package web

import "github.com/dtcookie/hcl"

type UserActionNamingPlaceholders []*UserActionNamingPlaceholder

func (me *UserActionNamingPlaceholders) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"placeholder": {
			Type:        hcl.TypeList,
			Description: "User action placeholders",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingPlaceholder).Schema()},
		},
	}
}

func (me UserActionNamingPlaceholders) MarshalHCL() (map[string]interface{}, error) {
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
		result["placeholder"] = entries
	}
	return result, nil
}

func (me *UserActionNamingPlaceholders) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("placeholder", me); err != nil {
		return err
	}
	return nil
}

// UserActionNamingPlaceholder The placeholder settings
type UserActionNamingPlaceholder struct {
	Name                        string          `json:"name"`                        // Placeholder name. Valid length needs to be between 1 and 50 characters.
	Input                       Input           `json:"input"`                       // The input for the place holder. Possible values are `ELEMENT_IDENTIFIER`, `INPUT_TYPE`, `METADATA`, `PAGE_TITLE`, `PAGE_URL`, `SOURCE_URL`, `TOP_XHR_URL` and `XHR_URL`.
	ProcessingPart              ProcessingPart  `json:"processingPart"`              // The part to process. Possible values are `ALL`, `ANCHOR` and `PATH`.
	ProcessingSteps             ProcessingSteps `json:"processingSteps,omitempty"`   // The processing step settings
	MetaDataID                  *int32          `json:"metadataId,omitempty"`        // The ID of the metadata
	UseGuessedElementIdentifier bool            `json:"useGuessedElementIdentifier"` // Use the element identifier that was selected by Dynatrace
}

func (me *UserActionNamingPlaceholder) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "Placeholder name. Valid length needs to be between 1 and 50 characters",
			Required:    true,
		},
		"input": {
			Type:        hcl.TypeString,
			Description: "The input for the place holder. Possible values are `ELEMENT_IDENTIFIER`, `INPUT_TYPE`, `METADATA`, `PAGE_TITLE`, `PAGE_URL`, `SOURCE_URL`, `TOP_XHR_URL` and `XHR_URL`",
			Required:    true,
		},
		"processing_part": {
			Type:        hcl.TypeString,
			Description: "The part to process. Possible values are `ALL`, `ANCHOR` and `PATH`",
			Required:    true,
		},
		"processing_steps": {
			Type:        hcl.TypeList,
			Description: "The processing step settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ProcessingSteps).Schema()},
		},
		"metadata_id": {
			Type:        hcl.TypeInt,
			Description: "The ID of the metadata",
			Optional:    true,
		},
		"use_guessed_element_identifier": {
			Type:        hcl.TypeBool,
			Description: "Use the element identifier that was selected by Dynatrace",
			Optional:    true,
		},
	}
}

func (me *UserActionNamingPlaceholder) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":                           me.Name,
		"input":                          me.Input,
		"processing_part":                me.ProcessingPart,
		"processing_steps":               me.ProcessingSteps,
		"metadata_id":                    me.MetaDataID,
		"use_guessed_element_identifier": me.UseGuessedElementIdentifier,
	})
}

func (me *UserActionNamingPlaceholder) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":                           &me.Name,
		"input":                          &me.Input,
		"processing_part":                &me.ProcessingPart,
		"processing_steps":               &me.ProcessingSteps,
		"metadata_id":                    &me.MetaDataID,
		"use_guessed_element_identifier": &me.UseGuessedElementIdentifier,
	})
}
