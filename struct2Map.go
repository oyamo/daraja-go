package darajago

import "reflect"

func struct2Map(structure interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	var s = reflect.ValueOf(structure).Elem()
	var typeOfT = s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		result[typeOfT.Field(i).Name] = f.Interface()
	}
	return result

}
