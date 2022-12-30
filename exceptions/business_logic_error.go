//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

const (
	businessLogicFailed = "business logic error"
)

type BusinessLogicError struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	Err       error  `json:"-"`
}

func (e *BusinessLogicError) Error() string {
	return e.Message
}

var TargetBusinessLogicError = errors.New(businessLogicFailed)

func NewBusinessLogicError(ErrorCode int, err error) error {
	return &BusinessLogicError{
		Message:   businessLogicFailed,
		ErrorCode: ErrorCode,
		Err:       err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e BusinessLogicError) Unwrap() error {
	log.Error(TargetBusinessLogicError, e.ErrorCode, e.Err)
	return TargetBusinessLogicError
}

// Dig Returns the inner most CustomErrorWrapper
func (e BusinessLogicError) Dig() BusinessLogicError {
	var errStruct BusinessLogicError
	if errors.As(e.Err, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
