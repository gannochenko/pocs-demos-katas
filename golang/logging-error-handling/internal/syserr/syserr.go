package syserr

import (
	"errors"
	"fmt"
	"strings"

	pkgError "github.com/pkg/errors"
)

type Code string

const (
	InternalCode Code = "internal"
	BadInputCode Code = "bad_input"
	NotFoundCode Code = "not_found"
)

type Field struct {
	Key   string
	Value any
}

func NewFiled(key string, value any) *Field {
	return &Field{
		Key:   key,
		Value: value,
	}
}

func Internal(message string, fields ...*Field) *Error {
	return NewError(InternalCode, message, fields...)
}

func BadInput(message string, fields ...*Field) *Error {
	return NewError(BadInputCode, message, fields...)
}

func NotFound(message string, fields ...*Field) *Error {
	return NewError(NotFoundCode, message, fields...)
}

func NewError(code Code, message string, fields ...*Field) *Error {
	stack := GetStack(pkgError.New(""))

	return &Error{
		message: message,
		code:    code,
		fields:  fields,
		stack:   stack,
	}
}

type Error struct {
	message string
	code    Code
	stack   []*ErrorStackItem
	fields  []*Field
}

type ErrorStackItem struct {
	File     string
	Line     string
	Function string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) SetError(message string) {
	e.message = message
}

func (e Error) GetCode() Code {
	return e.code
}

func (e Error) GetStack() []*ErrorStackItem {
	return e.stack
}

func (e Error) GetFields() []*Field {
	return e.fields
}

func (e Error) AppendField(field *Field) {
	e.fields = append(e.fields, field)
}

func WrapError(err error, message string, fields ...*Field) *Error {
	wrappedMessage := err.Error()
	if message != "" {
		wrappedMessage = fmt.Sprintf("%s: %s", message, wrappedMessage)
	}

	var systemError *Error
	ok := errors.As(err, &systemError)
	if ok {
		systemError.SetError(wrappedMessage)

		for _, field := range fields {
			systemError.AppendField(field)
		}

		return systemError
	} else {
		return NewError(InternalCode, wrappedMessage, fields...)
	}
}

type stackTracer interface {
	StackTrace() pkgError.StackTrace
}

func GetStack(err error) []*ErrorStackItem {
	var traceableError stackTracer
	ok := errors.As(err, &traceableError)
	if ok {
		stackTrace := (traceableError).StackTrace()

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

	return make([]*ErrorStackItem, 0)
}

func getFrameFilePath(frame pkgError.Frame) string {
	frameString := strings.Split(fmt.Sprintf("%+s", frame), "\n\t")
	return frameString[1]
}
