package syserr

import (
	"errors"
	"fmt"
	"strings"

	pkgError "github.com/pkg/errors"
)

type Code string

const (
	Internal Code = "internal"
	BadInput Code = "bad_input"
	NotFound Code = "not_found"
)

func NewInternalError(message string) *Error {
	return &Error{
		message: message,
		code:    Internal,
	}
}

func NewBadInputError(message string) *Error {
	return &Error{
		message: message,
		code:    BadInput,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		message: message,
		code:    NotFound,
	}
}

type Error struct {
	message string
	code    Code
	stack   []ErrorStackItem
}

type ErrorStackItem struct {
	file     string
	line     int
	function string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) GetCode() Code {
	return e.code
}

func (e Error) GetStack() []ErrorStackItem {
	return e.stack
}

type stackTracer interface {
	StackTrace() pkgError.StackTrace
}

func GetStack(err error) pkgError.StackTrace {
	var traceableError stackTracer
	ok := errors.As(err, &traceableError)
	if ok {
		return (traceableError).StackTrace()
	}

	return pkgError.StackTrace{}
}

func ConvertStackToStrings(stackTrace pkgError.StackTrace) []string {
	result := make([]string, 0)

	for _, frame := range stackTrace {
		result = append(result, fmt.Sprintf("%s:%d\t%n", getFrameFilePath(frame), frame, frame))
	}

	return result
}

func getFrameFilePath(frame pkgError.Frame) string {
	frameString := strings.Split(fmt.Sprintf("%+s", frame), "\n\t")
	return frameString[1]
}
