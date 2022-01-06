package xconvert

import (
	"bytes"
	"encoding/json"
	"reflect"
	"regexp"
	"unicode"
)

// ToCamelCase ...
func ToCamelCase(str string) string {
	chunks := regexp.MustCompile(`[\p{L}\p{N}]+`).FindAll([]byte(str), -1)
	for k, v := range chunks {
		if k == 0 {
			continue
		}
		chunks[k] = bytes.Title(v)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase ...
func ToSnakeCase(str string) string {
	str = ToCamelCase(str)
	runes := []rune(str)
	length := len(runes)
	var resp []rune
	for i := 0; i < length; i++ {
		resp = append(resp, unicode.ToLower(runes[i]))
		if i+1 < length && (unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i])) {
			resp = append(resp, '_')
		}
	}
	return string(resp)
}

// CopyStructByReflect ...
func CopyStructByReflect(src, dst interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		value := srcVal.Field(i)
		name := srcVal.Type().Field(i).Name
		dstValue := dstVal.FieldByName(name)
		if !dstValue.IsValid() || !dstValue.CanSet() {
			continue
		}
		dstValue.Set(value)
	}
}

// CopyStructByJSON ...
func CopyStructByJSON(src, dst interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
