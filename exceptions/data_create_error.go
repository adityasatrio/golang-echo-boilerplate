//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

var TargetDataCreateError = errors.New(dataCreateFailed)

const (
	dataCreateFailed = "create data error"
)

type dataCreateError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *dataCreateError) Error() string {
	return e.Message
}

func NewDataCreateError(err error) error {
	return &dataCreateError{
		Message:     dataCreateFailed,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e dataCreateError) Unwrap() error {
	log.Error(TargetDataCreateError, e.OriginalErr)
	return TargetDataCreateError
}

// Dig Returns the inner most CustomErrorWrapper
func (e dataCreateError) Dig() dataCreateError {
	var errStruct dataCreateError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
