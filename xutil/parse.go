package xutil

import (
	"reflect"
	"strconv"
	"strings"
)

// Parse ...
func Parse(obj interface{}, path string) interface{} {
	if path == "" {
		return obj
	}
	segments := strings.Split(path, ".")
	for _, key := range segments {
		switch v := obj.(type) {
		case map[string]interface{}:
			if _, ok := v[key]; !ok {
				return nil
			}
			obj = v[key]
		default:
			ov := reflect.ValueOf(obj)
			if ov.Kind() == reflect.Slice || ov.Kind() == reflect.Array {
				index, err := strconv.Atoi(key)
				if err != nil {
					return nil
				}
				newObj := make([]interface{}, ov.Len())
				for i := 0; i < ov.Len(); i++ {
					newObj[i] = ov.Index(i).Interface()
				}
				if index < 0 || index >= len(newObj) {
					return nil
				}
				obj = newObj[index]
			} else {
				return nil
			}
		}
	}
	return obj
}
