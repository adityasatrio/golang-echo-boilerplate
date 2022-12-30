//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type dataGetError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *dataGetError) Error() string {
	return e.Message
}

var TargetDataGetError = errors.New(dataGetFailed)

func NewDataGetError(err error) error {
	return &dataGetError{
		Message:     dataGetFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e dataGetError) Unwrap() error {
	log.Error(TargetDataGetError, e.OriginalErr)
	return TargetDataGetError
}

// Dig Returns the inner most CustomErrorWrapper
func (e dataGetError) Dig() dataGetError {
	var errStruct dataGetError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
