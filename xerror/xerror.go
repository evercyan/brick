package xerror

import (
	"sort"
	"sync"
)

var (
	errors = make(map[int]*Error)
	mu     sync.Mutex
)

// New ...
func New(code int, msg string) *Error {
	err := &Error{
		Code: code,
		Msg:  msg,
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := errors[code]; !ok {
		errors[code] = err
	}
	return err
}

// Errors 返回所有 error 列表
// 按 code 升序
// 通过 WithMsg 声明的错误码不重复显示
func Errors() []*Error {
	res := make([]*Error, 0)
	if len(errors) == 0 {
		return res
	}
	codes := make([]int, 0)
	for code := range errors {
		codes = append(codes, code)
	}
	sort.Ints(codes)
	for _, code := range codes {
		res = append(res, errors[code])
	}
	return res
}
