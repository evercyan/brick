package xencoding

import (
	"encoding/json"
)

// JSONEncode ...
func JSONEncode(elem interface{}) string {
	b, _ := json.Marshal(elem)
	return string(b)
}

// JSONDecode ...
func JSONDecode(str string, dst interface{}) error {
	return json.Unmarshal([]byte(str), dst)
}
