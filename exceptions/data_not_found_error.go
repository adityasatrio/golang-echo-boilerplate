//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type DataNotFoundError struct {
	Message     string `json:"message"`
	OriginalErr error  `json:"-"`
}

func (e *DataNotFoundError) Error() string {
	return e.Message
}

var TargetDataNotFoundError = errors.New(DataNotFound)

func NewDataNotFoundError(err error) error {
	return &DataNotFoundError{
		Message:     DataNotFound,
		OriginalErr: err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e DataNotFoundError) Unwrap() error {
	log.Error(TargetDataNotFoundError, e.OriginalErr)
	return TargetDataNotFoundError
}

// Dig Returns the inner most CustomErrorWrapper
func (e DataNotFoundError) Dig() DataNotFoundError {
	var errStruct DataNotFoundError
	if errors.As(e.OriginalErr, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
