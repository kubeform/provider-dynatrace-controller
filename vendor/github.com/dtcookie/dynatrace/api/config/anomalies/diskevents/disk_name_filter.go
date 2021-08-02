package diskevents

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// DiskNameFilter Narrows the rule usage down to disks, matching the specified criteria.
type DiskNameFilter struct {
	Operator Operator `json:"operator"` // Comparison operator.
	Value    string   `json:"value"`    // Value to compare to.
}

func (me *DiskNameFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operator": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Possible values are: `CONTAINS`, `DOES_NOT_CONTAIN`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `EQUALS` and `STARTS_WITH`",
		},
		"value": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Value to compare to",
		},
	}
}

func (me *DiskNameFilter) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	return decoder.MarshalAll(map[string]interface{}{
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *DiskNameFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	reader := hcl.NewReader(decoder, nil)
	me.Operator = Operator(opt.String(reader.String("operator")))
	me.Value = opt.String(reader.String("value"))
	return nil
}

// Operator Comparison operator.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	Contains         Operator
	DoesNotContain   Operator
	DoesNotEqual     Operator
	DoesNotStartWith Operator
	Equals           Operator
	StartsWith       Operator
}{
	"CONTAINS",
	"DOES_NOT_CONTAIN",
	"DOES_NOT_EQUAL",
	"DOES_NOT_START_WITH",
	"EQUALS",
	"STARTS_WITH",
}
