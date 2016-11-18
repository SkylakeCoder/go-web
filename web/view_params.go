package web

import (
	"container/list"
	"errors"
	"fmt"
)

type ViewParams map[string]interface{}

func NewViewParams(params ...interface{}) (*ViewParams, error) {
	if len(params)%2 != 0 {
		return nil, errors.New("invalid view params.")
	}
	vp := &ViewParams{}
	for i := 0; i < len(params); i += 2 {
		k := params[i]
		v := params[i+1]
		strKey, ok := k.(string)
		if !ok {
			return nil, errors.New("key must be string!")
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
			return nil, errors.New("unknown walue type.")
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
		return v, errors.New("can't find key:" + k)
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
	return "", errors.New("GetAsString: invalid walue type. key=" + k)
}

func (vp *ViewParams) GetAsList(k string) (*list.List, error) {
	v, err := vp.Get(k)
	if err != nil {
		return nil, err
	}
	l, ok := v.(*list.List)
	if !ok {
		return nil, errors.New("GetAsList: invalid value type. key=" + k)
	}
	return l, nil
}
