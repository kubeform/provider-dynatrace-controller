package terraform

import (
	"fmt"
	"strconv"
	"strings"
)

// NGResourceCollector has no documentation
type NGResourceCollector interface {
	Set(string, interface{}) error
}

type ngResource struct {
	values map[string]interface{}
}

type path string

func (p path) asInt() int {
	if res, err := strconv.Atoi(string(p)); err == nil {
		return res
	}
	return -1
}

func (p path) split() (path, path) {
	pos := -1
	if pos = strings.Index(string(p), "."); pos == -1 {
		return p, ""
	}
	return path(p[:pos]), path(p[pos+1:])
}

func (res *ngResource) Set(key string, v interface{}) error {
	logger.Println(fmt.Sprintf("[SET][%s] %v", key, v))
	var err error
	var result interface{}
	result, err = res.set(res.values, path(key), v)
	res.values = result.(map[string]interface{})
	return err
}

func (res *ngResource) set(values interface{}, key path, v interface{}) (interface{}, error) {
	first, second := key.split()
	iFirst := first.asInt()
	if iFirst != -1 {
		var param []interface{}
		if values == nil {
			param = []interface{}{}
		} else {
			param = values.([]interface{})
		}
		return res.setIdx(param, iFirst, second, v)
	}
	var param map[string]interface{}
	if values == nil {
		param = map[string]interface{}{}
	} else {
		param = values.(map[string]interface{})
	}
	return res.setMap(param, string(first), second, v)
}

func (res *ngResource) setMap(values map[string]interface{}, mkey string, key path, v interface{}) (map[string]interface{}, error) {
	var err error
	if values == nil {
		values = map[string]interface{}{}
	}
	var cur interface{}
	var found bool
	if cur, found = values[mkey]; !found {
		cur = nil
	}
	if len(key) > 0 {
		if cur, err = res.set(cur, key, v); err != nil {
			return nil, err
		}
		values[mkey] = cur
	} else {
		values[mkey] = v
	}
	return values, nil
}

func (res *ngResource) setIdx(values []interface{}, idx int, key path, v interface{}) ([]interface{}, error) {
	var err error
	if values == nil {
		values = []interface{}{}
	}
	if idx >= len(values) {
		if idx > len(values) {
			return nil, fmt.Errorf("len: %d, idx: %d", len(values), idx)
		}
		values = append(values, nil)
	}
	cur := values[idx]
	if len(key) > 0 {
		if cur, err = res.set(cur, key, v); err != nil {
			return nil, err
		}
		values[idx] = cur
	} else {
		values[idx] = v
	}
	return values, nil
}
