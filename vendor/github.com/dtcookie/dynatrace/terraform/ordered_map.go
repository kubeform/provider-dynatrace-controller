package terraform

import (
	"encoding/json"
	"sort"
)

// OrderedMapKey has no documentation
type OrderedMapKey struct {
	Value string
	Order int
}

// String ensures that the string representation of a key just contains its name
func (omk OrderedMapKey) String() string {
	return omk.Value
}

// MarshalText ensures that OrdredeMapKeys implement the interface `encoding.TextMarshaler`
func (omk OrderedMapKey) MarshalText() ([]byte, error) {
	return []byte(omk.Value), nil
}

// OrderedMapKeyCollection is an ordinary slice of OrderedMapKeys - it just provides a method for convenient sorting conveniently
type OrderedMapKeyCollection []OrderedMapKey

// Less represents the sort function for this collection
func (omkc OrderedMapKeyCollection) Less(i, j int) bool { return omkc[i].Order < omkc[j].Order }

// OrderedMap implements a map that allows the user to specify the output order of the entries during JSON Marshalling
type OrderedMap map[OrderedMapKey]interface{}

// Delete removes an element from the map
func (om OrderedMap) Delete(key string) {
	// var omk *OrderedMapKey
	for k := range om {
		if k.Value == key {
			delete(om, k)
		}
	}
}

// NextKey produces a key with an order higher than any key currently used within this map and the given value
func (om OrderedMap) NextKey(value string) OrderedMapKey {
	if len(om) == 0 {
		return OrderedMapKey{Value: value, Order: 0}
	}
	max := 0
	for k := range om {
		if k.Order > max {
			max = k.Order
		}
	}
	return OrderedMapKey{Value: value, Order: max + 1}
}

// PrefixKey produces a key with an order lower than any key currently used within this map and the given value
func (om OrderedMap) PrefixKey(value string) OrderedMapKey {
	if len(om) == 0 {
		return OrderedMapKey{Value: value, Order: 0}
	}
	min := 0
	for k := range om {
		if k.Order < min {
			min = k.Order
		}
	}
	return OrderedMapKey{Value: value, Order: min - 1}
}

// OrderedKeys returns the ordered keys - ordered by the `order` field
func (om OrderedMap) OrderedKeys() []OrderedMapKey {
	keys := OrderedMapKeyCollection{}
	for k := range om {
		keys = append(keys, k)
	}
	sort.Slice(keys, keys.Less)
	return keys
}

// MarshalJSON ensures that the values of this map is gettign marshalled in the order as defined by the keys
func (om OrderedMap) MarshalJSON() ([]byte, error) {
	s := "{"
	keys := om.OrderedKeys()
	for _, omk := range keys {
		if bytes, err := json.Marshal(omk.Value); err == nil {
			s = s + string(bytes) + ":"
		} else {
			return nil, err
		}
		// add value
		v := om[omk]
		if bytes, err := json.Marshal(v); err == nil {
			s = s + string(bytes) + ","
		} else {
			return nil, err
		}
	}
	if len(keys) > 0 {
		s = s[0 : len(s)-1]
	}
	s = s + "}"
	return []byte(s), nil
}
