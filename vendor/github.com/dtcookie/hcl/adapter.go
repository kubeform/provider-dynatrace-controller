package hcl

import "github.com/dtcookie/opt"

type Adapter interface {
	GetString(key string) *string
	GetBool(key string) *bool
}

func Adapt(d Decoder) Adapter {
	return &adapter{Decoder: d}
}

type adapter struct {
	Decoder Decoder
}

func (a *adapter) GetString(key string) *string {
	if value, _ := a.Decoder.GetOk(key); value != nil {
		return opt.NewString(value.(string))
	}
	return nil
}

func (a *adapter) GetBool(key string) *bool {
	if value, ok := a.Decoder.GetOkExists(key); ok {
		return opt.NewBool(value.(bool))
	}
	return nil
}
