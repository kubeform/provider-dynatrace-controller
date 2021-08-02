package terraform

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceData has no documentation
type ResourceData interface {
	GetOk(key string) (interface{}, bool)
	Keys() []string
	Split() ([]ResourceData, error)
}

// NewResourceData has no documentation
func NewResourceData(source interface{}) (ResourceData, error) {
	if source == nil {
		panic(errors.New("ResourceData for nil source unsupported"))
	}
	switch typedSource := source.(type) {
	case *schema.ResourceData:
		return &terraformResourceData{d: typedSource}, nil
	case *schema.Set:
		return &setResourceData{s: typedSource}, nil
	case []interface{}:
		if len(typedSource) > 0 {
			log.Printf("creating sliceResourceData (elemType: %T)\n", typedSource[0])
		} else {
			log.Printf("creating sliceResourceData (0 size array)\n")
		}
		return &sliceResourceData{s: typedSource}, nil
	case map[string]interface{}:
		return &mapResourceData{m: typedSource}, nil
	case string, int, bool:
		return &singleResourceData{v: source}, nil
	default:
		panic(fmt.Errorf("ResourceData for source %T unsupported", source))
		// return nil, fmt.Errorf("ResourceData for source %T unsupported", source)
	}
}

type terraformResourceData struct {
	d *schema.ResourceData
}

func (rd *terraformResourceData) GetOk(key string) (interface{}, bool) {
	return rd.d.GetOk(key)
}

func (rd *terraformResourceData) Keys() []string {
	return []string{}
}

func (rd *terraformResourceData) Split() ([]ResourceData, error) {
	panic(fmt.Errorf("ResourceData based on %T cannot split", rd.d))
}

type mapResourceData struct {
	m map[string]interface{}
}

func (mrd *mapResourceData) GetOk(key string) (interface{}, bool) {
	if stored, ok := mrd.m[key]; ok {
		return stored, true
	}
	return nil, false
}

func (mrd *mapResourceData) Keys() []string {
	keys := []string{}
	for k := range mrd.m {
		keys = append(keys, k)
	}
	return keys
}

func (mrd *mapResourceData) Split() ([]ResourceData, error) {
	panic(fmt.Errorf("ResourceData based on %T cannot split", mrd.m))
}

type singleResourceData struct {
	v interface{}
}

func (srd *singleResourceData) GetOk(key string) (interface{}, bool) {
	return srd.v, true
}

func (srd *singleResourceData) Keys() []string {
	return []string{}
}

func (srd *singleResourceData) Split() ([]ResourceData, error) {
	panic(fmt.Errorf("ResourceData based on %T cannot split", srd.v))
}

type setResourceData struct {
	s *schema.Set
}

func (srd *setResourceData) GetOk(key string) (interface{}, bool) {
	if srd.s.Len() > 1 {
		panic(fmt.Errorf("ResourceData based on %T (len: %v) doesn't provide single values", srd.s, srd.s.Len()))
	}
	if srd.s.Len() == 0 {
		// panic(fmt.Errorf("ResourceData based on %T doesn't contain any entries", srd.s))
		return nil, false
	}
	var child ResourceData
	var err error
	if child, err = NewResourceData(srd.s.List()[0]); err != nil {
		panic(err)
	}
	return child.GetOk(key)
}

func (srd *setResourceData) Keys() []string {
	if srd.s.Len() > 1 {
		panic(fmt.Errorf("ResourceData based on %T (len: %v) doesn't provide keys", srd.s, srd.s.Len()))
	}
	if srd.s.Len() == 0 {
		panic(fmt.Errorf("ResourceData based on %T doesn't contain any entries", srd.s))
	}
	var child ResourceData
	var err error
	if child, err = NewResourceData(srd.s.List()[0]); err != nil {
		panic(err)
	}
	return child.Keys()
}

func (srd *setResourceData) Split() ([]ResourceData, error) {
	log.Println("setResourceData.split")
	result := []ResourceData{}
	for _, elem := range srd.s.List() {
		log.Printf("  elem: %v, type: %T\n", elem, elem)
		var resourceData ResourceData
		var err error
		if resourceData, err = NewResourceData(elem); err != nil {
			return nil, err
		}
		result = append(result, resourceData)
	}
	return result, nil
}

type sliceResourceData struct {
	s []interface{}
}

func (srd *sliceResourceData) GetOk(key string) (interface{}, bool) {
	if len(srd.s) > 1 {
		panic(fmt.Errorf("ResourceData based on %T (len: %v) doesn't provide single values", srd.s, len(srd.s)))
	}
	if len(srd.s) == 0 {
		panic(fmt.Errorf("ResourceData based on %T doesn't contain any entries", srd.s))
	}
	var child ResourceData
	var err error
	if child, err = NewResourceData(srd.s[0]); err != nil {
		panic(err)
	}
	return child.GetOk(key)
}

func (srd *sliceResourceData) Keys() []string {
	if len(srd.s) > 1 {
		panic(fmt.Errorf("ResourceData based on %T (len: %v) doesn't provide keys", srd.s, len(srd.s)))
	}
	if len(srd.s) == 0 {
		panic(fmt.Errorf("ResourceData based on %T doesn't contain any entries", srd.s))
	}
	var child ResourceData
	var err error
	if child, err = NewResourceData(srd.s[0]); err != nil {
		panic(err)
	}
	return child.Keys()
}

func (srd *sliceResourceData) Split() ([]ResourceData, error) {
	log.Println("sliceResourceData.split")
	result := []ResourceData{}
	for _, elem := range srd.s {
		log.Printf("  elem: %v, type: %T\n", elem, elem)
		var resourceData ResourceData
		var err error
		if resourceData, err = NewResourceData(elem); err != nil {
			return nil, err
		}
		result = append(result, resourceData)
	}
	return result, nil
}
