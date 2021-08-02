package hcl

type Set interface {
	List() []interface{}
	Len() int
}
