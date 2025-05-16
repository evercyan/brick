package xlodash

import (
	"reflect"
)

// IsZero 判断变量是否其对应类型的零值
func IsZero[V comparable](v V) bool {
	vv := reflect.ValueOf(v)
	return vv.Interface() == reflect.Zero(vv.Type()).Interface()
}
