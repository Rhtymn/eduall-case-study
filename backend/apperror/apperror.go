package apperror

import (
	"fmt"
	"runtime/debug"
)

type AppError struct {
	code    int
	message string
	err     error
	stack   []byte
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		code:    code,
		message: message,
		err:     err,
		stack:   debug.Stack(),
	}
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) Code() int {
	return e.code
}

func (e *AppError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("(%d) %s: %s", e.code, e.message, e.err)
	}
	return fmt.Sprintf("(%d) %s", e.code, e.message)
}

func (e *AppError) StackTrace() []byte {
	return e.stack
}

func IsErrorCode(err error, code int) bool {
	aerr, ok := err.(*AppError)
	return ok && aerr.code == code
}

func Wrap(err error) error {
	if _, ok := err.(*AppError); ok {
		return err
	}
	return NewInternal(err)
}

func NewInternal(err error) error {
	return NewAppError(CodeInternal, "internal error", err)
}

func NewInvalidValidation(err error) error {
	return NewAppError(CodeInvalidValidation, "invalid validation", err)
}
