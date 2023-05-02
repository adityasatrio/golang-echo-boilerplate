//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
)

var TargetBusinessLogicError = errors.New(businessLogicFailed)

const (
	businessLogicFailed = "business logic error"
)

type (
	BusinessLogicError struct {
		Message   string `json:"message"`
		ErrorCode int    `json:"errorCode"`
		Err       error  `json:"-"`
	}
)

func (e *BusinessLogicError) Error() string {
	return e.Message
}

func NewBusinessLogicError(ErrorCode int, err error) error {
	return &BusinessLogicError{
		Message:   businessLogicFailed,
		ErrorCode: ErrorCode,
		Err:       err,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e *BusinessLogicError) Unwrap() error {
	errorLogic := BusinessLogicReason(e.ErrorCode)
	jsonByte, err := json.Marshal(errorLogic)
	if err != nil {
		log.Error(errorLogic, e.Err)
	} else {
		log.Error(string(jsonByte), e.Err)
	}

	return TargetBusinessLogicError
}

// Dig Returns the innermost CustomErrorWrapper
func (e *BusinessLogicError) Dig() *BusinessLogicError {
	var errStruct *BusinessLogicError
	if errors.As(e.Err, &errStruct) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return errStruct.Dig()
	}
	return e
}
