package entity

import "github.com/pkg/errors"

// bad request error: status code 400
type BadRequestError struct {
	Message string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
	}
}

func (e *BadRequestError) Error() string    { return e.Message }
func (e *BadRequestError) BadRequest() bool { return true }

type badRequestError interface {
	BadRequest() bool
}

func IsBadRequestError(err error) bool {
	e, ok := errors.Cause(err).(badRequestError)
	return ok && e.BadRequest()
}

// unauthorized error: status code 401
type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
	}
}

func (e *UnauthorizedError) Error() string      { return e.Message }
func (e *UnauthorizedError) Unauthorized() bool { return true }

type unauthorizedError interface {
	Unauthorized() bool
}

func IsUnauthorizedError(err error) bool {
	e, ok := errors.Cause(err).(unauthorizedError)
	return ok && e.Unauthorized()
}

// precondition failed: status code 412
type PreconditionFailedError struct {
	Message string
}

func NewPreconditionFailedError(message string) *PreconditionFailedError {
	return &PreconditionFailedError{
		Message: message,
	}
}

func (e *PreconditionFailedError) Error() string            { return e.Message }
func (e *PreconditionFailedError) PreconditionFailed() bool { return true }

type preconditionFailed interface {
	Precondition() bool
}

func IsPreconditionFailedError(err error) bool {
	e, ok := errors.Cause(err).(preconditionFailed)
	return ok && e.Precondition()
}

// system error: status code 500
type SystemError struct {
	Message string
}

func NewSystemError(message string) *SystemError {
	return &SystemError{
		Message: message,
	}
}

func (e *SystemError) Error() string { return e.Message }
func (e *SystemError) System() bool  { return true }

type systemError interface {
	system() bool
}

func IsSyatemError(err error) bool {
	e, ok := errors.Cause(err).(systemError)
	return ok && e.system()
}
