//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type BusinessLogicError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *BusinessLogicError) Error() string {
	return e.Message
}

var TargetBusinessLogicError = errors.New(BusinessLogicFailed)

func NewBusinessLogicError(err error) error {
	return &BusinessLogicError{
		Message:     BusinessLogicFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e BusinessLogicError) Unwrap() error {
	log.Error(TargetBusinessLogicError, e.OriginalErr)
	return TargetBusinessLogicError
}

// Dig Returns the inner most CustomErrorWrapper
func (e BusinessLogicError) Dig() BusinessLogicError {
	var errStruct BusinessLogicError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
