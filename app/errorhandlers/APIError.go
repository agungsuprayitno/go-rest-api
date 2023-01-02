package errorhandlers

import (
	"net/http"
)

type APIError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	APIError() (int, string, string)
}

type sentinelAPIError struct {
	status int
	msg    string
}

func (e sentinelAPIError) Error() string {
	return e.msg
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.msg
}

type sentinelWrappedError struct {
	error
	sentinel *sentinelAPIError
}

func (e sentinelWrappedError) Is(err error) bool {
	return e.sentinel == err
}

func (e sentinelWrappedError) APIError() (int, string) {
	return e.sentinel.APIError()
}

var (
	ErrAuth      = &sentinelAPIError{status: http.StatusUnauthorized, msg: "invalid token"}
	ErrNotFound  = &sentinelAPIError{status: http.StatusNotFound, msg: "not found"}
	ErrDuplicate = &sentinelAPIError{status: http.StatusBadRequest, msg: "duplicate"}
)

func WrapError(err error, sentinel *sentinelAPIError) error {
	return sentinelWrappedError{error: err, sentinel: sentinel}
}

// func wrapError(err error) error {
// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		return app.WrapError(err, app.ErrNotFound)
// 	case isMySQLError(err, codeDuplicate):
// 		return app.WrapError(err, app.ErrDuplicate)
// 	default:
// 		return err
// 	}
// }
