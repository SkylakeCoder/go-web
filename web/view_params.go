package web

import (
	"container/list"
	"fmt"
)

type ViewParams map[string]interface{}

func NewViewParams(params ...interface{}) (*ViewParams, error) {
	if len(params)%2 != 0 {
		return nil, fmt.Errorf("length of view params must be even.")
	}
	vp := &ViewParams{}
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

func (vp *ViewParams) PutString(k string, v string) {
	(*vp)[k] = v
}

func (vp *ViewParams) PutInt(k string, v int) {
	(*vp)[k] = v
}

func (vp *ViewParams) PutFloat(k string, v float64) {
	(*vp)[k] = v
}

func (vp *ViewParams) PutList(k string, v *list.List) {
	(*vp)[k] = v
}

func (vp *ViewParams) GetKeys() []string {
	result := []string{}
	for k, _ := range *vp {
		result = append(result, k)
	}
	return result
}

func (vp *ViewParams) Get(k string) (interface{}, error) {
	v, ok := (*vp)[k]
	if ok {
		return v, nil
	} else {
		return v, fmt.Errorf("can't find key: %s", k)
	}
}

func (vp *ViewParams) GetAsString(k string) (string, error) {
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

func (vp *ViewParams) GetAsList(k string) (*list.List, error) {
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
