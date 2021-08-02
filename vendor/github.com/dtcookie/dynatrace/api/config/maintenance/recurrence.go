package maintenance

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// Recurrence The recurrence of the maintenance window.
type Recurrence struct {
	DayOfMonth      *int32                     `json:"dayOfMonth,omitempty"` // The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February.
	DayOfWeek       *DayOfWeek                 `json:"dayOfWeek,omitempty"`  // The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`.
	DurationMinutes int32                      `json:"durationMinutes"`      // The duration of the maintenance window in minutes.
	StartTime       string                     `json:"startTime"`            // The start time of the maintenance window in HH:mm format.
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *Recurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"day_of_month": {
			Type:        hcl.TypeInt,
			Description: "The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February",
			Optional:    true,
		},
		"day_of_week": {
			Type:        hcl.TypeString,
			Description: "The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`",
			Optional:    true,
		},
		"duration_minutes": {
			Type:        hcl.TypeInt,
			Description: "The duration of the maintenance window in minutes",
			Required:    true,
		},
		"start_time": {
			Type:        hcl.TypeString,
			Description: "The start time of the maintenance window in HH:mm format",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Recurrence) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.DayOfMonth != nil {
		result["day_of_month"] = int(opt.Int32(me.DayOfMonth))
	}
	if me.DayOfWeek != nil {
		result["day_of_week"] = string(*me.DayOfWeek)
	}
	result["duration_minutes"] = int(me.DurationMinutes)
	result["start_time"] = me.StartTime
	return result, nil
}

func (me *Recurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "day_of_month")
		delete(me.Unknowns, "day_of_week")
		delete(me.Unknowns, "duration_minutes")
		delete(me.Unknowns, "start_time")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("day_of_month"); ok {
		me.DayOfMonth = opt.NewInt32(int32(value.(int)))
	}
	if value, ok := decoder.GetOk("day_of_week"); ok {
		me.DayOfWeek = DayOfWeek(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("duration_minutes"); ok {
		me.DurationMinutes = int32(value.(int))
	}
	if value, ok := decoder.GetOk("start_time"); ok {
		me.StartTime = value.(string)
	}
	return nil
}

func (me *Recurrence) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("dayOfMonth", me.DayOfMonth); err != nil {
		return nil, err
	}
	if err := m.Marshal("dayOfWeek", me.DayOfWeek); err != nil {
		return nil, err
	}
	if err := m.Marshal("durationMinutes", me.DurationMinutes); err != nil {
		return nil, err
	}
	if err := m.Marshal("startTime", me.StartTime); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *Recurrence) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("dayOfMonth", &me.DayOfMonth); err != nil {
		return err
	}
	if err := m.Unmarshal("dayOfWeek", &me.DayOfWeek); err != nil {
		return err
	}
	if err := m.Unmarshal("durationMinutes", &me.DurationMinutes); err != nil {
		return err
	}
	if err := m.Unmarshal("startTime", &me.StartTime); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// DayOfWeek The day of the week for weekly maintenance.
// The format is the full name of the day in upper case, for example `THURSDAY`.
type DayOfWeek string

func (me DayOfWeek) Ref() *DayOfWeek {
	return &me
}

// DayOfWeeks offers the known enum values
var DayOfWeeks = struct {
	Friday    DayOfWeek
	Monday    DayOfWeek
	Saturday  DayOfWeek
	Sunday    DayOfWeek
	Thursday  DayOfWeek
	Tuesday   DayOfWeek
	Wednesday DayOfWeek
}{
	"FRIDAY",
	"MONDAY",
	"SATURDAY",
	"SUNDAY",
	"THURSDAY",
	"TUESDAY",
	"WEDNESDAY",
}
