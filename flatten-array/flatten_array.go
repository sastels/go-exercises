package flatten

import (
	"reflect"
)

func Flatten(input interface{}) []interface{} {
	retVal := []interface{}{}
	s := reflect.ValueOf(input)

	for i := 0; i < s.Len(); i++ {
		val := s.Index(i)
		if val.IsNil() {
			continue
		}
		slice, isSlice := val.Interface().([]interface{})
		if isSlice {
			retVal = append(retVal, Flatten(slice)...)
		} else {
			retVal = append(retVal, val.Interface().(int))
		}
	}

	return retVal
}
