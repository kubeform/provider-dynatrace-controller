package validation

import "github.com/dtcookie/hcl"

// Rules is a list of validation rules
type Rules []*Rule

type Rule struct {
	Type        Type   `json:"type"`        // The type of the rule. Possible values are `patternConstraint`, `regexConstraint`, `httpStatusesList` and `certificateExpiryDateConstraint`
	PassIfFound bool   `json:"passIfFound"` // The validation condition. `true` means validation succeeds if the specified content/element is found. `false` means validation fails if the specified content/element is found. Always specify `false` for `certificateExpiryDateConstraint` to fail the monitor if SSL cedrtificate expiry is within the specified number of days
	Value       string `json:"value"`       // The content to look for
}

func (me *Rule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the rule. Possible values are `patternConstraint`, `regexConstraint`, `httpStatusesList` and `certificateExpiryDateConstraint`",
			Required:    true,
		},
		"pass_if_found": {
			Type:        hcl.TypeBool,
			Description: " The validation condition. `true` means validation succeeds if the specified content/element is found. `false` means validation fails if the specified content/element is found. Always specify `false` for `certificateExpiryDateConstraint` to fail the monitor if SSL certificate expiry is within the specified number of days",
			Optional:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The content to look for",
			Required:    true,
		},
	}
}

func (me *Rule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	result["pass_if_found"] = me.PassIfFound
	result["value"] = me.Value
	return result, nil
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("pass_if_found", &me.PassIfFound); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	return nil
}

// Type The type of the rule. Possible values are `patternConstraint`, `regexConstraint`, `httpStatusesList` and `certificateExpiryDateConstraint`
type Type string

// ValidationRuleTypes offers the known enum values
var Types = struct {
	PatternConstraint               Type
	RegexConstraint                 Type
	HTTPStatusesList                Type
	CertificateExpiryDateConstraint Type
}{
	`patternConstraint`,
	`regexConstraint`,
	`httpStatusesList`,
	`certificateExpiryDateConstraint`,
}
