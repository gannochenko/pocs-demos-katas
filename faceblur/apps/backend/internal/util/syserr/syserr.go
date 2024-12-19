package syserr

import (
	"errors"
	"fmt"
	"strings"

	pkgError "github.com/pkg/errors"
	"github.com/samber/lo"
)

type Code string

const (
	InternalCode       Code = "internal"
	BadInputCode       Code = "bad_input"
	UnauthorisedCode   Code = "unauthorised"
	NotFoundCode       Code = "not_found"
	NotImplementedCode Code = "not_implemented"
)

type Error struct {
	Message      string
	code         *Code
	Stack        []*ErrorStackItem
	fields       []*Field
	WrappedError error
}

type ErrorStackItem struct {
	File     string
	Line     string
	Function string
}

type Field struct {
	Key   string
	Value any
}

func F(key string, value any) *Field {
	return &Field{
		Key:   key,
		Value: value,
	}
}

func NewInternal(message string, fields ...*Field) *Error {
	return New(lo.ToPtr(InternalCode), message, fields...)
}

func NewBadInput(message string, fields ...*Field) *Error {
	return New(lo.ToPtr(BadInputCode), message, fields...)
}

func NewUnauthorized(message string, fields ...*Field) *Error {
	return New(lo.ToPtr(UnauthorisedCode), message, fields...)
}

func NewNotFound(message string, fields ...*Field) *Error {
	return New(lo.ToPtr(NotFoundCode), message, fields...)
}

func NewNotImplemented(message string, fields ...*Field) *Error {
	return New(lo.ToPtr(NotImplementedCode), message, fields...)
}

func New(code *Code, message string, fields ...*Field) *Error {
	stack := extractStackFromGenericError(pkgError.New(""))

	return &Error{
		Message: message,
		code:    code,
		fields:  fields,
		Stack:   stack,
	}
}

func Wrap(err error, message string, fields ...*Field) *Error {
	newError := New(nil, message, fields...)
	newError.WrappedError = err

	return newError
}

func WrapAs(err error, code Code, message string, fields ...*Field) *Error {
	newError := New(&code, message, fields...)
	newError.WrappedError = err

	return newError
}

func (e *Error) Unwrap() error {
	return e.WrappedError
}

func (e *Error) Error() string {
	if e.WrappedError != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.WrappedError.Error())
	}

	return e.Message
}

func (e *Error) GetCode() *Code {
	return e.code
}

func (e *Error) GetStack() []*ErrorStackItem {
	return e.Stack
}

func (e *Error) GetFields() []*Field {
	return e.fields
}

type stackTracer interface {
	StackTrace() pkgError.StackTrace
}

func (e *Error) GetStackFormatted() []string {
	return formatStack(e.Stack)
}

func GetStackFormatted(err error) []string {
	var sysErr *Error
	ok := errors.As(err, &sysErr)
	if ok {
		return sysErr.GetStackFormatted()
	}

	return formatStack(extractStackFromGenericError(err))
}

func GetCode(err error) Code {
	if err == nil {
		return InternalCode
	}

	for {
		if sErr, ok := err.(*Error); ok {
			code := sErr.GetCode()
			if code == nil {
				return InternalCode
			}

			return *code
		}

		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return InternalCode
			}
		default:
			return InternalCode
		}
	}
}

func GetFields(err error) []*Field {
	var result []*Field

	for {
		if err == nil {
			return result
		}

		if sErr, ok := err.(*Error); ok {
			result = append(result, sErr.GetFields()...)
		}

		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
		default:
			return result
		}
	}
}

func formatStack(stack []*ErrorStackItem) []string {
	result := make([]string, len(stack))
	for index, stackItem := range stack {
		result[index] = fmt.Sprintf("%s:%s %s", stackItem.File, stackItem.Line, stackItem.Function)
	}

	return result
}

func extractStackFromGenericError(err error) []*ErrorStackItem {
	stackTrace := extractStackTraceFromGenericError(err)

	result := make([]*ErrorStackItem, len(stackTrace))

	for index, frame := range stackTrace {
		result[index] = &ErrorStackItem{
			File:     getFrameFilePath(frame),
			Line:     fmt.Sprintf("%d", frame),
			Function: fmt.Sprintf("%s", frame),
		}
	}

	return result
}

func extractStackTraceFromGenericError(err error) pkgError.StackTrace {
	var result pkgError.StackTrace

	var traceableError stackTracer
	ok := errors.As(err, &traceableError)
	if ok {
		result = traceableError.StackTrace()
	}

	return result
}

func getFrameFilePath(frame pkgError.Frame) string {
	frameString := strings.Split(fmt.Sprintf("%+s", frame), "\n\t")
	return frameString[1]
}
