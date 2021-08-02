package hcl

import (
	"encoding/json"

	"github.com/dtcookie/opt"
)

type Reader interface {
	String(string) *string
	Float64(string) *float64
	Int32(string) *int32
	Bool(string) *bool
	Count(string) int
	Decode(v interface{}) error
}

func NewReader(decoder Decoder, unknowns map[string]json.RawMessage) Reader {
	return &reader{decoder: decoder, unknowns: unknowns}
}

type reader struct {
	decoder  Decoder
	unknowns map[string]json.RawMessage
}

func (r *reader) rmk(key string) {
	if r.unknowns != nil {
		delete(r.unknowns, key)
	}
}

func (r *reader) String(key string) *string {
	r.rmk(key)
	if value, _ := r.decoder.GetOk(key); value != nil {
		return opt.NewString(value.(string))
	}
	return nil
}

func (r *reader) Int32(key string) *int32 {
	r.rmk(key)
	if value, _ := r.decoder.GetOk(key); value != nil {
		return opt.NewInt32(int32(value.(int)))
	}
	return nil
}

func (r *reader) Float64(key string) *float64 {
	r.rmk(key)
	if value, _ := r.decoder.GetOk(key); value != nil {
		return opt.NewFloat64(value.(float64))
	}
	return nil
}

func (r *reader) Bool(key string) *bool {
	r.rmk(key)
	if value, ok := r.decoder.GetOkExists(key); ok {
		return opt.NewBool(value.(bool))
	}
	return nil
}

func (r *reader) Count(key string) int {
	r.rmk(key)
	if result, ok := r.decoder.GetOk(key + ".#"); ok {
		return result.(int)
	}
	return 0
}

func (r *reader) Decode(v interface{}) error {
	return nil
}
