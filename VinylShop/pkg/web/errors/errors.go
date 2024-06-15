package errors

import "errors"

var (
	ErrNotFound           = NewErr("Not Found", 404)
	ErrUnpocessableEntity = NewErr("Unprocessable Entity", 422)
	ErrBadRequest         = NewErr("Bad Request", 400)
	ErrForbidden          = NewErr("Forbidden: Not Enough Permissions", 403)
	ErrNotAuthorized      = NewErr("Forbidden: Not Authorized", 401)
	ErrTooManyRequests    = NewErr("Too Many Requests", 429)
	ErrInternal           = NewErr("Internal Server Error", 500)
	ErrEmptyField         = errors.New("fields cannot be empty(password, email)")
	ErrEmailValidation    = errors.New("wrong email format")
	ErrPasswordValidation = errors.New("password must have atleast 1 number, 1 uppercase and lowercase letters")
	ErrIncorrectPassword  = errors.New("password is incorrect")
	ErrEmailNotFound      = errors.New("email not found")
)
