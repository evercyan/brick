package xerror

import (
	"encoding/json"
)

// Error ...
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error ...
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// WithMsg 自定义错误文案
func (e *Error) WithMsg(msg string) *Error {
	return New(e.Code, msg)
}
