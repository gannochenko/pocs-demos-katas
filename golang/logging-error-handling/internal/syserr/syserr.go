package syserr

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
