package xconvert

import (
	"fmt"
	"strconv"
)

// ToUint ...
func ToUint(elem interface{}) uint64 {
	if elem == nil {
		return 0
	}
	switch val := elem.(type) {
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return val
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		if val == "" {
			return 0
		}
		value, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return 0
		}
		return value
	default:
		return 0
	}
}

// ToFloat64 ...
func ToFloat64(elem interface{}) float64 {
	if elem == nil {
		return 0
	}
	switch val := elem.(type) {
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
func ToString(elem interface{}) string {
	return fmt.Sprintf("%v", elem)
}

// ToBool ...
func ToBool(elem interface{}) bool {
	if elem == nil {
		return false
	}
	if v, ok := elem.(bool); ok {
		return v
	}
	v, err := strconv.ParseBool(ToString(elem))
	return err == nil && v
}
