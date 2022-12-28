package xconvert

import (
	"bytes"
	"encoding/json"
	"reflect"
	"regexp"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Title ...
func Title(s string) string {
	return cases.Title(language.English, cases.NoLower).String(s)
}

// ToCamelCase ...
func ToCamelCase(s string) string {
	chunks := regexp.MustCompile(`[\p{L}\p{N}]+`).FindAll([]byte(s), -1)
	for k, v := range chunks {
		if k == 0 {
			continue
		}
		chunks[k] = cases.Title(language.English, cases.NoLower).Bytes(v)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase ...
func ToSnakeCase(s string) string {
	s = ToCamelCase(s)
	runes := []rune(s)
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
