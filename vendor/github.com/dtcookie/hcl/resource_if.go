package hcl

// Resource has no documentation
type ResourceIF interface {
	GetOk(key string) (interface{}, bool)
	Get(key string) interface{}
	Set(key string, value interface{}) error
	Append(key string) ResourceIF
}
