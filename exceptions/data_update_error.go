//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type dataUpdateError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *dataUpdateError) Error() string {
	return e.Message
}

var TargetDataUpdateError = errors.New(dataUpdateFailed)

func NewDataUpdateError(err error) error {
	return &dataUpdateError{
		Message:     dataUpdateFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e dataUpdateError) Unwrap() error {
	log.Error(TargetDataUpdateError, e.OriginalErr)
	return TargetDataUpdateError
}

// Dig Returns the inner most CustomErrorWrapper
func (e dataUpdateError) Dig() dataUpdateError {
	var errStruct dataUpdateError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
