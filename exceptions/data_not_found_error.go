//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

var TargetDataNotFoundError = errors.New(dataNotFound)

const (
	dataNotFound = "data not found"
)

type dataNotFoundError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *dataNotFoundError) Error() string {
	return e.Message
}

func NewDataNotFoundError(err error) error {
	return &dataNotFoundError{
		Message:     dataNotFound,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e dataNotFoundError) Unwrap() error {
	log.Error(TargetDataNotFoundError, e.OriginalErr)
	return TargetDataNotFoundError
}

// Dig Returns the inner most CustomErrorWrapper
func (e dataNotFoundError) Dig() dataNotFoundError {
	var errStruct dataNotFoundError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
