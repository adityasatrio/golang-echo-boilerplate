//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type DataCreateError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *DataCreateError) Error() string {
	return e.Message
}

var TargetDataCreateError = errors.New(DataCreateFailed)

func NewDataCreateError(err error) error {
	return &DataCreateError{
		Message:     DataCreateFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e DataCreateError) Unwrap() error {
	log.Error(TargetDataCreateError, e.OriginalErr)
	return TargetDataCreateError
}

// Dig Returns the inner most CustomErrorWrapper
func (e DataCreateError) Dig() DataCreateError {
	var errStruct DataCreateError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
