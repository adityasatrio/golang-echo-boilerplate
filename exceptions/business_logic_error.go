//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type businessLogicError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *businessLogicError) Error() string {
	return e.Message
}

var TargetBusinessLogicError = errors.New(businessLogicFailed)

func NewBusinessLogicError(err error) error {
	return &businessLogicError{
		Message:     businessLogicFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e businessLogicError) Unwrap() error {
	log.Error(TargetBusinessLogicError, e.OriginalErr)
	return TargetBusinessLogicError
}

// Dig Returns the inner most CustomErrorWrapper
func (e businessLogicError) Dig() businessLogicError {
	var errStruct businessLogicError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
