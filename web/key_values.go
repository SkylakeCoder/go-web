package web

import (
	"container/list"
	"fmt"
)

// KeyValues is an utility that helps to make the key-value pairs more comfortably.
type KeyValues map[string]interface{}

// Get a new KeyValues instance.
// Example:
// keyValues := NewKeyValues(
//     "key1", 1,
//     "key2", "key2value",
//     "key3", 1.23456
// )
func NewKeyValues(params ...interface{}) (*KeyValues, error) {
	if len(params)%2 != 0 {
		return nil, fmt.Errorf("length of key-values must be even.")
	}
	vp := &KeyValues{}
	for i := 0; i < len(params); i += 2 {
		k := params[i]
		v := params[i+1]
		strKey, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("key must be string: %v", k)
		}
		switch v.(type) {
		case string:
			vp.PutString(strKey, v.(string))
		case int:
			vp.PutInt(strKey, v.(int))
		case float64:
			vp.PutFloat(strKey, v.(float64))
		case *list.List:
			vp.PutList(strKey, v.(*list.List))
		default:
			return nil, fmt.Errorf("unknown walue type: %v", v)
		}
	}
	return vp, nil
}

// Put a key-value pair and the value's type is string.
func (vp *KeyValues) PutString(k string, v string) {
	(*vp)[k] = v
}

// Put a key-value pair and the value's type is int.
func (vp *KeyValues) PutInt(k string, v int) {
	(*vp)[k] = v
}

// Put a key-value pair and the value's type is float64.
func (vp *KeyValues) PutFloat(k string, v float64) {
	(*vp)[k] = v
}

// Put a key-value pair and the value's type is *list.List.
func (vp *KeyValues) PutList(k string, v *list.List) {
	(*vp)[k] = v
}

// Get all the keys the KeyValues object.
func (vp *KeyValues) GetKeys() []string {
	result := []string{}
	for k, _ := range *vp {
		result = append(result, k)
	}
	return result
}

// Get the value by key.
func (vp *KeyValues) Get(k string) (interface{}, error) {
	v, ok := (*vp)[k]
	if ok {
		return v, nil
	} else {
		return v, fmt.Errorf("can't find key: %s", k)
	}
}

// Get the value by key and convert the value to the string type,
// it will return an error if convert failed.
func (vp *KeyValues) GetAsString(k string) (string, error) {
	v, err := vp.Get(k)
	if err != nil {
		return "", err
	}
	switch v.(type) {
	case string:
		s, _ := v.(string)
		return s, nil
	case int:
		i, _ := v.(int)
		return fmt.Sprintf("%d", i), nil
	case float64:
		f, _ := v.(float64)
		return fmt.Sprintf("%f", f), nil
	}
	return "", fmt.Errorf("GetAsString: invalid walue type. key=%s", k)
}

// Get the value by key and convert the value to the *list.List type,
// it will return an error if convert failed.
func (vp *KeyValues) GetAsList(k string) (*list.List, error) {
	v, err := vp.Get(k)
	if err != nil {
		return nil, err
	}
	l, ok := v.(*list.List)
	if !ok {
		return nil, fmt.Errorf("GetAsList: invalid value type. key=%s", k)
	}
	return l, nil
}
