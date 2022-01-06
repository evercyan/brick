package xutil

import (
	"encoding/json"
	"reflect"
	"time"
)

// is ...
func is(elem interface{}, types ...reflect.Kind) bool {
	elemType := reflect.ValueOf(elem).Kind()
	for _, t := range types {
		if t == elemType {
			return true
		}
	}
	return false
}

// IsInt ...
func IsInt(elem interface{}) bool {
	return is(
		elem,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	)
}

// IsUint ...
func IsUint(elem interface{}) bool {
	return is(
		elem,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
	)
}

// IsFloat ...
func IsFloat(elem interface{}) bool {
	return is(elem,
		reflect.Float32,
		reflect.Float64,
	)
}

// IsNumeric ...
func IsNumeric(elem interface{}) bool {
	return is(
		elem,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
	)
}

// IsBool ...
func IsBool(elem interface{}) bool {
	return is(elem, reflect.Bool)
}

// IsString ...
func IsString(elem interface{}) bool {
	return is(elem, reflect.String)
}

// IsSlice ...
func IsSlice(elem interface{}) bool {
	return is(elem, reflect.Slice)
}

// IsArray ...
func IsArray(elem interface{}) bool {
	return is(elem, reflect.Array)
}

// IsStruct ...
func IsStruct(elem interface{}) bool {
	return is(elem, reflect.Struct)
}

// IsMap ...
func IsMap(elem interface{}) bool {
	return is(elem, reflect.Map)
}

// IsFunc ...
func IsFunc(elem interface{}) bool {
	return is(elem, reflect.Func)
}

// IsChannel ...
func IsChannel(elem interface{}) bool {
	return is(elem, reflect.Chan)
}

// IsTime ...
func IsTime(elem interface{}) bool {
	if _, ok := elem.(time.Time); ok {
		return true
	}
	return false
}

// IsEmpty ...
func IsEmpty(elem interface{}) bool {
	if elem == nil {
		return true
	}
	elemValue := reflect.ValueOf(elem)
	return reflect.DeepEqual(elemValue.Interface(), reflect.Zero(elemValue.Type()).Interface())
}

// IsJSONString ...
func IsJSONString(str string) bool {
	var raw json.RawMessage
	return json.Unmarshal([]byte(str), &raw) == nil
}

// IsJSONObject ...
func IsJSONObject(elem interface{}) bool {
	b, _ := json.Marshal(elem)
	return IsJSONString(string(b))
}

// InArray ...
func InArray(elem interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == elem {
				return true
			}
		}
	case reflect.Map:
		return targetValue.MapIndex(reflect.ValueOf(elem)).IsValid()
	}
	return false
}
