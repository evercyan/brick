package xtype

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/evercyan/brick/xencoding"
)

// ToInt ...
func ToInt(v interface{}) int {
	return int(ToInt64(v))
}

// ToInt64 ...
func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		if val == "" {
			return 0
		}
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return value
	default:
		return 0
	}
}

// ToUint ...
func ToUint(v interface{}) uint {
	return uint(ToInt64(v))
}

// ToUint ...
func ToUint64(v interface{}) uint64 {
	return uint64(ToInt64(v))
}

// ToFloat64 ...
func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		if val == "" {
			return 0
		}
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return value
	default:
		return 0
	}
}

// ToString ...
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.Itoa(int(val))
	case int16:
		return strconv.Itoa(int(val))
	case int32:
		return strconv.Itoa(int(val))
	case int64:
		return strconv.Itoa(int(val))
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case string:
		return val
	case []byte:
		return Bytes2String(val)
	default:
		if IsJSONObject(v) {
			return xencoding.JSONEncode(v)
		}
		return fmt.Sprint(v)
	}
}

// ToBool ...
func ToBool(v interface{}) bool {
	if v == nil {
		return false
	}
	if v, ok := v.(bool); ok {
		return v
	}
	b, err := strconv.ParseBool(ToString(v))
	return err == nil && b
}

// ToSlice ...
func ToSlice(v interface{}) []interface{} {
	res := make([]interface{}, 0)
	if v == nil {
		return res
	}
	vo := reflect.ValueOf(v)
	vt := vo.Kind()
	if vt == reflect.String {
		vv := v.(string)
		if vv == "" {
			return res
		}
		vvs := strings.Split(vv, ",")
		res = make([]interface{}, 0, len(vvs))
		for _, vvv := range vvs {
			vvv = strings.TrimSpace(vvv)
			if vvv != "" {
				res = append(res, vvv)
			}
		}
		return res
	} else if vt == reflect.Map {
		iter := vo.MapRange()
		for iter.Next() {
			res = append(res, iter.Value().Interface())
		}
	} else if vt == reflect.Array || vt == reflect.Slice {
		res = make([]interface{}, vo.Len())
		for i := 0; i < vo.Len(); i++ {
			res[i] = vo.Index(i).Interface()
		}
	} else {
		res = append(res, v)
	}
	return res
}

// Bytes2String ...
// unsafe(): 163837650    7.215 ns/op    0 B/op    0 allocs/op
// string(): 31958134     37.09 ns/op    0 B/op    0 allocs/op
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
