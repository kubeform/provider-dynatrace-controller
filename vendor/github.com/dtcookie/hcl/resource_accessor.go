package hcl

// ResourceAccessor has no documentation
type ResourceAccessor interface {
	Set(key string, value interface{}) error
}
