//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

var TargetDataDeleteError = errors.New(dataDeleteFailed)

const (
	dataDeleteFailed = "delete data error"
)

type dataDeleteError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *dataDeleteError) Error() string {
	return e.Message
}

func NewDeleteError(err error) error {
	return &dataDeleteError{
		Message:     dataDeleteFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e dataDeleteError) Unwrap() error {
	log.Error(TargetDataDeleteError, e.OriginalErr)
	return TargetDataDeleteError
}

// Dig Returns the inner most CustomErrorWrapper
func (e dataDeleteError) Dig() dataDeleteError {
	var errStruct dataDeleteError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
