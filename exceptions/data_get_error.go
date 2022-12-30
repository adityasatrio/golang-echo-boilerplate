//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type DataGetError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *DataGetError) Error() string {
	return e.Message
}

var TargetDataGetError = errors.New(DataGetFailed)

func NewDataGetError(err error) error {
	return &DataGetError{
		Message:     DataGetFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e DataGetError) Unwrap() error {
	log.Error(TargetDataGetError, e.OriginalErr)
	return TargetDataGetError
}

// Dig Returns the inner most CustomErrorWrapper
func (e DataGetError) Dig() DataGetError {
	var errStruct DataGetError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
