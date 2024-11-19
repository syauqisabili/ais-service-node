package pkg

import "fmt"

// Error represents an error with a code and message.
const (
	NoError         = 0
	ErrNotFound     = 1
	ErrUnauthorized = 2
	ErrInvalid      = 3
	ErrCreateDir    = 4
	ErrDirNotExist  = 5
	ErrCreateFile   = 6
	ErrNoParam      = 7
	ErrFileNotExist = 8
	ErrReadFile     = 9
	ErrWriteFile    = 10
	ErrSaveFile     = 11
	ErrProcessFail  = 12

	ErrTargetNotFound = 1001

	// Ais Message Error
	ErrInvalidAisMessage      = 6000
	ErrNotSupportedAisMessage = 6001
)

type Error struct {
	Code    int
	Message string
}

// Error implements the error interface for Error.
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Error creates a new Error with the given code and message.
func ErrorStatus(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
