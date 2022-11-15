package utils

import (
	"fmt"
	"reflect"
)

// 参考项目 https://github.com/dablelv/go-huge-util

// UniqueSlice 删除切片中的重复值
func UniqueSlice(slice interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// unique the slice
	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len())
	m := make(map[interface{}]struct{})
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			dst = reflect.Append(dst, v.Index(i))
			m[v.Index(i).Interface()] = struct{}{}
		}
	}
	return dst.Interface(), nil
}
