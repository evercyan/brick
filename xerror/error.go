package xerror

import (
	"errors"
	"fmt"
)

// New ...
func New(code int, msg string) *xError {
	return &xError{
		Code:    code,
		Message: msg,
	}
}

// IsXErr ...
func IsXErr(err error) bool {
	var xerr *xError
	return errors.As(err, &xerr)
}

// ToXErr ...
func ToXErr(err error) *xError {
	var xerr *xError
	if errors.As(err, &xerr) {
		return xerr
	}
	return nil
}

// Wrap ...
func Wrap(err1, err2 error) error {
	if xerr := ToXErr(err1); xerr != nil {
		return xerr.Wrap(err2)
	}
	return fmt.Errorf("%s ::: %w", err1.Error(), err2)
}

// Unwrap ...
func Unwrap(err error) error {
	if xerr := ToXErr(err); xerr != nil {
		return xerr.Unwrap()
	}
	return errors.Unwrap(err)
}
