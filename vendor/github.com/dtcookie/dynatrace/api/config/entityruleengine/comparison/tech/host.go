package tech

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// Host The value to compare to.
type Host struct {
	Type         *SimpleHostTechType        `json:"type,omitempty"`         // Predefined technology, if technology is not predefined, then the verbatim type must be set
	VerbatimType *string                    `json:"verbatimType,omitempty"` // Non-predefined technology, use for custom technologies.
	Unknowns     map[string]json.RawMessage `json:"-"`
}

func (sht *Host) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "Predefined technology, if technology is not predefined, then the verbatim type must be set. Possible values are APPARMOR, BOSH, BOSHBPM, CLOUDFOUNDRY, CONTAINERD, CRIO, DIEGO_CELL, DOCKER, GARDEN, GRSECURITY, KUBERNETES, OPENSHIFT, OPENSTACK_COMPUTE, OPENSTACK_CONTROLLER and SELINUX",
			Optional:    true,
		},
		"verbatim_type": {
			Type:        hcl.TypeString,
			Description: "Non-predefined technology, use for custom technologies",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (sht *Host) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(sht.Unknowns) > 0 {
		data, err := json.Marshal(sht.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if sht.Type != nil {
		result["type"] = sht.Type.String()
	}
	if sht.VerbatimType != nil {
		result["verbatim_type"] = *sht.VerbatimType
	}
	return result, nil
}

func (sht *Host) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), sht); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &sht.Unknowns); err != nil {
			return err
		}
		delete(sht.Unknowns, "type")
		delete(sht.Unknowns, "verbatim_type")
		if len(sht.Unknowns) == 0 {
			sht.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		sht.Type = SimpleHostTechType(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("verbatim_type"); ok {
		sht.VerbatimType = opt.NewString(value.(string))
	}
	return nil
}

func (sht *Host) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(sht.Unknowns) > 0 {
		for k, v := range sht.Unknowns {
			m[k] = v
		}
	}
	if sht.Type != nil {
		rawMessage, err := json.Marshal(sht.Type)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	if sht.VerbatimType != nil {
		rawMessage, err := json.Marshal(sht.VerbatimType)
		if err != nil {
			return nil, err
		}
		m["verbatimType"] = rawMessage
	}
	return json.Marshal(m)
}

func (sht *Host) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["type"]; found {
		if err := json.Unmarshal(v, &sht.Type); err != nil {
			return err
		}
	}
	if v, found := m["verbatimType"]; found {
		if err := json.Unmarshal(v, &sht.VerbatimType); err != nil {
			return err
		}
	}
	delete(m, "verbatimType")
	delete(m, "type")
	if len(m) > 0 {
		sht.Unknowns = m
	}
	return nil
}

// SimpleHostTechType Predefined technology, if technology is not predefined, then the verbatim type must be set
type SimpleHostTechType string

func (v SimpleHostTechType) Ref() *SimpleHostTechType {
	return &v
}

func (v *SimpleHostTechType) String() string {
	return string(*v)
}

// SimpleHostTechTypes offers the known enum values
var SimpleHostTechTypes = struct {
	Apparmor            SimpleHostTechType
	Bosh                SimpleHostTechType
	Boshbpm             SimpleHostTechType
	CloudFoundry        SimpleHostTechType
	Containerd          SimpleHostTechType
	Crio                SimpleHostTechType
	DiegoCell           SimpleHostTechType
	Docker              SimpleHostTechType
	Garden              SimpleHostTechType
	Grsecurity          SimpleHostTechType
	Kubernetes          SimpleHostTechType
	Openshift           SimpleHostTechType
	OpenStackCompute    SimpleHostTechType
	OpenStackController SimpleHostTechType
	Selinux             SimpleHostTechType
}{
	"APPARMOR",
	"BOSH",
	"BOSHBPM",
	"CLOUDFOUNDRY",
	"CONTAINERD",
	"CRIO",
	"DIEGO_CELL",
	"DOCKER",
	"GARDEN",
	"GRSECURITY",
	"KUBERNETES",
	"OPENSHIFT",
	"OPENSTACK_COMPUTE",
	"OPENSTACK_CONTROLLER",
	"SELINUX",
}
