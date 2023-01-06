package darajago

import "reflect"

func struct2Map(structure interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(structure)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		m[v.Type().Field(i).Name] = v.Field(i).Interface()
	}
	return m
}
