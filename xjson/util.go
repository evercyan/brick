package xjson

import (
	"encoding/json"

	"github.com/evercyan/brick/xencoding"
)

// ...
var (
	Encode = xencoding.JSONEncode
	Decode = xencoding.JSONDecode
)

// Pretty ...
func Pretty(v interface{}) string {
	if vv, ok := v.(string); ok {
		var raw json.RawMessage
		if err := json.Unmarshal([]byte(vv), &raw); err == nil {
			v = raw
		}
	}
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

// Minify ...
func Minify(v interface{}) string {
	if vv, ok := v.(string); ok {
		var raw json.RawMessage
		if err := json.Unmarshal([]byte(vv), &raw); err == nil {
			v = raw
		}
	}
	return Encode(v)
}
