package errors

import "errors"

var (
	Unknown     = errors.New("unknown error")
	NotFound    = errors.New("not found")
	InvalidData = errors.New("invalid data")
	NoAffected  = errors.New("0 row affected")
)

type ErrorResponse struct {
	Msg string `json:"message,omitempty"`
}
