package xerror

import (
	"fmt"
)

// xError ...
type xError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	wrapped error
}

// Error ...
func (e *xError) Error() string {
	errorMsg := fmt.Sprintf("[%d] %s", e.Code, e.Message)
	if e.wrapped == nil {
		return errorMsg
	}
	return fmt.Sprintf("%s ::: %s", errorMsg, e.wrapped.Error())
}

// WithMsg ...
func (e *xError) WithMsg(message string) *xError {
	xerr := e.Clone()
	xerr.Message = message
	return xerr
}

// Wrap ...
func (e *xError) Wrap(err error) *xError {
	xerr := e.Clone()
	xerr.wrapped = err
	return xerr
}

// Unwrap ...
func (e *xError) Unwrap() error {
	return e.wrapped
}

// Clone ...
func (e *xError) Clone() *xError {
	return &xError{
		Code:    e.Code,
		Message: e.Message,
		wrapped: e.wrapped,
	}
}
