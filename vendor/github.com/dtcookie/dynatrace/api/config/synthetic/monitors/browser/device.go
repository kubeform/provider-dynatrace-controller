package browser

import (
	"github.com/dtcookie/hcl"
)

type Device struct {
	Name         *string      `json:"deviceName,omitempty"`   // The name of the preconfigured device—when editing in the browser, press `Crtl+Spacebar` to see the list of available devices
	Orientation  *Orientation `json:"orientation,omitempty"`  // The orientation of the device. Possible values are `portrait` or `landscape`. Desktop and laptop devices are not allowed to use the `portrait` orientation
	Mobile       *bool        `json:"mobile,omitempty"`       // The flag of the mobile device.\nSet to `true` for mobile devices or `false` for a desktop or laptop. Required if `touchEnabled` is specified.
	TouchEnabled *bool        `json:"touchEnabled,omitempty"` // The flag of the touchscreen.\nSet to `true` if the device uses touchscreen. In that case, use can set interaction event as `tap`. Required if `mobile` is specified.
	Width        *int         `json:"width,omitempty"`        // The width of the screen in pixels.\nThe maximum allowed width is `1920`
	Height       *int         `json:"height,omitempty"`       // The height of the screen in pixels.\nThe maximum allowed width is `1080`
	ScaleFactor  *int         `json:"scaleFactor,omitempty"`  // The pixel ratio of the device
}

func (me *Device) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the preconfigured device—when editing in the browser, press `Crtl+Spacebar` to see the list of available devices",
			Optional:    true,
		},
		"orientation": {
			Type:        hcl.TypeString,
			Description: "The orientation of the device. Possible values are `portrait` or `landscape`. Desktop and laptop devices are not allowed to use the `portrait` orientation",
			Optional:    true,
		},
		"mobile": {
			Type:        hcl.TypeBool,
			Description: "The flag of the mobile device.\nSet to `true` for mobile devices or `false` for a desktop or laptop.",
			Optional:    true,
		},
		"touch_enabled": {
			Type:        hcl.TypeBool,
			Description: "The flag of the touchscreen.\nSet to `true` if the device uses touchscreen. In that case, use can set interaction event as `tap`.",
			Optional:    true,
		},
		"width": {
			Type:        hcl.TypeInt,
			Description: "The width of the screen in pixels.\nThe maximum allowed width is `1920`.",
			Optional:    true,
		},
		"height": {
			Type:        hcl.TypeInt,
			Description: "The height of the screen in pixels.\nThe maximum allowed width is `1080`.",
			Optional:    true,
		},
		"scale_factor": {
			Type:        hcl.TypeInt,
			Description: "The pixel ratio of the device.",
			Optional:    true,
		},
	}
}

func (me *Device) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Name != nil && len(*me.Name) > 0 {
		result["name"] = *me.Name
	}
	if me.Orientation != nil {
		result["orientation"] = string(*me.Orientation)
	}
	if me.Mobile != nil && *me.Mobile {
		result["mobile"] = me.Mobile
	} else {
		result["mobile"] = false
	}
	if me.TouchEnabled != nil && *me.TouchEnabled {
		result["touch_enabled"] = me.TouchEnabled
	} else {
		result["touch_enabled"] = false
	}
	if me.Width != nil {
		result["width"] = *me.Width
	}
	if me.Height != nil {
		result["height"] = *me.Height
	}
	if me.ScaleFactor != nil {
		result["scale_factor"] = *me.ScaleFactor
	}
	return result, nil
}

func (me *Device) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("name", &me.Name); err != nil {
		return err
	}
	if err := decoder.Decode("orientation", &me.Orientation); err != nil {
		return err
	}
	if err := decoder.Decode("mobile", &me.Mobile); err != nil {
		return err
	}
	if err := decoder.Decode("touch_enabled", &me.TouchEnabled); err != nil {
		return err
	}
	if err := decoder.Decode("width", &me.Width); err != nil {
		return err
	}
	if err := decoder.Decode("height", &me.Height); err != nil {
		return err
	}
	if err := decoder.Decode("scale_factor", &me.ScaleFactor); err != nil {
		return err
	}
	return nil
}
