//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type DataUpdateError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *DataUpdateError) Error() string {
	return e.Message
}

var TargetDataUpdateError = errors.New(DataUpdateFailed)

func NewDataUpdateError(err error) error {
	return &DataUpdateError{
		Message:     DataUpdateFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e DataUpdateError) Unwrap() error {
	log.Error(TargetDataUpdateError, e.OriginalErr)
	return TargetDataUpdateError
}

// Dig Returns the inner most CustomErrorWrapper
func (e DataUpdateError) Dig() DataUpdateError {
	var errStruct DataUpdateError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
