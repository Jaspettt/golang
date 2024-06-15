package errors

type Error struct {
	err  string
	code int
}

func NewErr(err string, code int) *Error {
	return &Error{err: err, code: code}
}

func (e *Error) Error() string {
	return e.err
}

func (e *Error) Code() int {
	return e.code
}
