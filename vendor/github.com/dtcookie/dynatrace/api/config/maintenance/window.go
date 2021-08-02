package maintenance

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// Window Configuration of a maintenance window.
type Window struct {
	ID                                 *string                    `json:"id,omitempty"`                                 // The ID of the maintenance window.
	Name                               string                     `json:"name"`                                         // The name of the maintenance window, displayed in the UI.
	Description                        string                     `json:"description"`                                  // A short description of the maintenance purpose.
	Schedule                           *Schedule                  `json:"schedule"`                                     // The schedule of the maintenance window.
	Scope                              *Scope                     `json:"scope,omitempty"`                              // The scope of the maintenance window.   The scope restricts the alert/problem detection suppression to certain Dynatrace entities. It can contain a list of entities and/or matching rules for dynamic formation of the scope.   If no scope is specified, the alert/problem detection suppression applies to the entire environment.
	Suppression                        Suppression                `json:"suppression"`                                  // The type of suppression of alerting and problem detection during the maintenance.
	Type                               WindowType                 `json:"type"`                                         // The type of the maintenance: planned or unplanned
	SuppressSyntheticMonitorsExecution *bool                      `json:"suppressSyntheticMonitorsExecution,omitempty"` // Suppress execution of synthetic monitors during the maintenance
	Metadata                           *api.ConfigMetadata        `json:"metadata,omitempty"`                           // Metadata useful for debugging
	Unknowns                           map[string]json.RawMessage `json:"-"`
}

func (me *Window) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the maintenance window, displayed in the UI",
			Required:    true,
		},
		"schedule": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The schedule of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(Schedule).Schema(),
			},
		},
		"scope": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "the tiles this Dashboard consist of",
			Elem: &hcl.Resource{
				Schema: new(Scope).Schema(),
			},
		},
		"suppression": {
			Type:        hcl.TypeString,
			Description: "The type of suppression of alerting and problem detection during the maintenance",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the maintenance: planned or unplanned",
			Required:    true,
		},
		"suppress_synth_mon_exec": {
			Type:        hcl.TypeBool,
			Description: "Suppress execution of synthetic monitors during the maintenance",
			Optional:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the maintenance purpose",
			Optional:    true,
		},
		"metadata": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Deprecated:  "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(api.ConfigMetadata).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Window) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	if len(me.Description) > 0 {
		result["description"] = me.Description
	}
	if me.Schedule != nil {
		if marshalled, err := me.Schedule.MarshalHCL(); err == nil {
			result["schedule"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.SuppressSyntheticMonitorsExecution != nil {
		result["suppress_synth_mon_exec"] = opt.Bool(me.SuppressSyntheticMonitorsExecution)
	}
	if me.Scope != nil {
		if marshalled, err := me.Scope.MarshalHCL(); err == nil {
			result["scope"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	result["suppression"] = string(me.Suppression)
	result["type"] = string(me.Type)
	return result, nil
}

func (me *Window) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "description")
		delete(me.Unknowns, "schedule")
		delete(me.Unknowns, "scope")
		delete(me.Unknowns, "suppression")
		delete(me.Unknowns, "type")
		delete(me.Unknowns, "metadata")
		delete(me.Unknowns, "suppress_synth_mon_exec")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("suppress_synth_mon_exec"); value != nil {
		me.SuppressSyntheticMonitorsExecution = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = value.(string)
	}

	if _, ok := decoder.GetOk("schedule.#"); ok {
		me.Schedule = new(Schedule)
		if err := me.Schedule.UnmarshalHCL(hcl.NewDecoder(decoder, "schedule", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("scope.#"); ok {
		me.Scope = new(Scope)
		if err := me.Scope.UnmarshalHCL(hcl.NewDecoder(decoder, "scope", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("suppression"); ok {
		me.Suppression = Suppression(value.(string))
	}
	if value, ok := decoder.GetOk("type"); ok {
		me.Type = WindowType(value.(string))
	}
	return nil
}

func (me *Window) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("suppressSyntheticMonitorsExecution", me.SuppressSyntheticMonitorsExecution); err != nil {
		return nil, err
	}
	if err := m.Marshal("name", me.Name); err != nil {
		return nil, err
	}
	if err := m.Marshal("description", me.Description); err != nil {
		return nil, err
	}
	if err := m.Marshal("schedule", me.Schedule); err != nil {
		return nil, err
	}
	if err := m.Marshal("scope", me.Scope); err != nil {
		return nil, err
	}
	if err := m.Marshal("suppression", me.Suppression); err != nil {
		return nil, err
	}
	if err := m.Marshal("type", me.Type); err != nil {
		return nil, err
	}
	if err := m.Marshal("metadata", me.Metadata); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *Window) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("id", &me.ID); err != nil {
		return err
	}
	if err := m.Unmarshal("suppressSyntheticMonitorsExecution", &me.SuppressSyntheticMonitorsExecution); err != nil {
		return err
	}
	if err := m.Unmarshal("name", &me.Name); err != nil {
		return err
	}
	if err := m.Unmarshal("description", &me.Description); err != nil {
		return err
	}
	if err := m.Unmarshal("schedule", &me.Schedule); err != nil {
		return err
	}
	if err := m.Unmarshal("scope", &me.Scope); err != nil {
		return err
	}
	if err := m.Unmarshal("suppression", &me.Suppression); err != nil {
		return err
	}
	if err := m.Unmarshal("type", &me.Type); err != nil {
		return err
	}
	if err := m.Unmarshal("metadata", &me.Metadata); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// WindowType The type of the maintenance: planned or unplanned.
type WindowType string

// MaintenanceWindowTypes offers the known enum values
var MaintenanceWindowTypes = struct {
	Planned   WindowType
	Unplanned WindowType
}{
	"PLANNED",
	"UNPLANNED",
}

// Suppression The type of suppression of alerting and problem detection during the maintenance.
type Suppression string

// Suppressions offers the known enum values
var Suppressions = struct {
	DetectProblemsAndAlert  Suppression
	DetectProblemsDontAlert Suppression
	DontDetectProblems      Suppression
}{
	"DETECT_PROBLEMS_AND_ALERT",
	"DETECT_PROBLEMS_DONT_ALERT",
	"DONT_DETECT_PROBLEMS",
}
