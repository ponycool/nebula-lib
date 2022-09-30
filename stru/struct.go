package stru

import "reflect"

type Struct struct {
}

// Assign 结构体赋值
func (s Struct) Assign(ref interface{}, value interface{}) {
	refVal := reflect.ValueOf(ref).Elem()
	vVal := reflect.ValueOf(value).Elem()
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		name := vTypeOfT.Field(i).Name
		if ok := refVal.FieldByName(name).IsValid() && refVal.FieldByName(name).Type() == vVal.Field(i).Type(); ok {
			refVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
