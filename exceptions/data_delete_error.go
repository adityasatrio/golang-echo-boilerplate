//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type DataDeleteError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *DataDeleteError) Error() string {
	return e.Message
}

var TargetDataDeleteError = errors.New(DataDeleteFailed)

func NewDeleteError(err error) error {
	return &DataDeleteError{
		Message:     DataDeleteFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e DataDeleteError) Unwrap() error {
	log.Error(TargetDataDeleteError, e.OriginalErr)
	return TargetDataDeleteError
}

// Dig Returns the inner most CustomErrorWrapper
func (e DataDeleteError) Dig() DataDeleteError {
	var errStruct DataDeleteError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
