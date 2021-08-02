package strategy

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

type Wrapper struct {
	Strategy MonitoringStrategy
}

func (me *Wrapper) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "An auto-adaptive baseline strategy to detect anomalies within metrics that show a regular change over time, as the baseline is also updated automatically. An example is to detect an anomaly in the number of received network packets or within the number of user actions over time",
			Elem:        &hcl.Resource{Schema: new(Auto).Schema()},
		},
		"static": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "A static threshold monitoring strategy to alert on hard limits within a given metric. An example is the violation of a critical memory limit",
			Elem:        &hcl.Resource{Schema: new(Static).Schema()},
		},
		"generic": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A generic monitoring strategy",
			Elem:        &hcl.Resource{Schema: new(BaseMonitoringStrategy).Schema()},
		},
	}
}

func (me *Wrapper) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Strategy != nil {
		switch strategy := me.Strategy.(type) {
		case *Auto:
			if marshalled, err := strategy.MarshalHCL(hcl.NewDecoder(decoder, "auto", 0)); err == nil {
				result["auto"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *Static:
			if marshalled, err := strategy.MarshalHCL(hcl.NewDecoder(decoder, "static", 0)); err == nil {
				result["static"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *BaseMonitoringStrategy:
			if marshalled, err := strategy.MarshalHCL(hcl.NewDecoder(decoder, "generic", 0)); err == nil {
				result["generic"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		default:
		}
	}
	return result, nil
}

func (me *Wrapper) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("auto.#"); ok {
		cfg := new(Auto)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "auto", 0)); err != nil {
			return err
		}
		me.Strategy = cfg
	}
	if _, ok := decoder.GetOk("static.#"); ok {
		cfg := new(Static)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "static", 0)); err != nil {
			return err
		}
		me.Strategy = cfg
	}
	if _, ok := decoder.GetOk("generic.#"); ok {
		cfg := new(BaseMonitoringStrategy)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "generic", 0)); err != nil {
			return err
		}
		me.Strategy = cfg
	}
	return nil
}

func (me *Wrapper) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if rawType, found := properties["type"]; found {
		var sType string
		if err := json.Unmarshal(rawType, &sType); err != nil {
			return err
		}
		switch sType {
		case string(Types.AutoAdaptiveBaseline):
			cfg := new(Auto)
			if err := json.Unmarshal(data, &cfg); err != nil {
				return err
			}
			me.Strategy = cfg
		case string(Types.StaticThreshold):
			cfg := new(Static)
			if err := json.Unmarshal(data, &cfg); err != nil {
				return err
			}
			me.Strategy = cfg
		default:
			cfg := new(BaseMonitoringStrategy)
			if err := json.Unmarshal(data, &cfg); err != nil {
				return err
			}
			me.Strategy = cfg
		}
	}
	return nil
}
