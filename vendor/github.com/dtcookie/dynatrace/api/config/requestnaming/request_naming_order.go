package requestnaming

import (
	"github.com/dtcookie/hcl"
)

type Order struct {
	Values []Ref `json:"values"`
}

func (me *Order) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ids": {
			Type:        hcl.TypeList,
			Required:    true,
			Description: "The IDs of the request namings in the order they should be taken into consideration",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Order) MarshalHCL() (map[string]interface{}, error) {
	refs := []interface{}{}
	for _, ref := range me.Values {
		refs = append(refs, ref.ID)
	}
	return map[string]interface{}{
		"ids": refs,
	}, nil
}

func (me *Order) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Values = []Ref{}
	values, ok := decoder.GetOk("ids")
	if ok {
		vals := values.([]interface{})
		for _, val := range vals {
			me.Values = append(me.Values, Ref{ID: val.(string)})
		}
	}
	return nil
}

type Ref struct {
	ID string `json:"id"`
}
