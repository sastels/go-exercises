package flatten

import (
	"fmt"
	"reflect"
)

func Flatten(input interface{}) []interface{} {

	retVal := []interface{}{}

	fmt.Printf("input:<%v>\n", input)

	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		val := s.Index(i)
		valInt, ok := val.Interface().(int)
		if ok {
			retVal = append(retVal, valInt)
		} else {
			slice := val.Interface().([]interface{})
			valFlattened := Flatten(slice)

			retVal = append(retVal, valFlattened...)
		}
	}

	return retVal
}
