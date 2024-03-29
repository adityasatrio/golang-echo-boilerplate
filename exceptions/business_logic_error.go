//reference : https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7

package exceptions

import (
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
	"runtime"
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
		File      string `json:"file"`
		Line      int    `json:"line"`
	}
)

func (e *BusinessLogicError) Error() string {
	return e.Message
}

func NewBusinessLogicError(ErrorCode int, err error) error {
	_, file, line, ok := runtime.Caller(1) // Capture caller info
	if !ok {
		file = "???"
		line = 0
	}

	return &BusinessLogicError{
		Message:   businessLogicFailed,
		ErrorCode: ErrorCode,
		Err:       err,
		File:      file,
		Line:      line,
	}
}

// Unwrap Implements the errors.Unwrap interface
func (e *BusinessLogicError) Unwrap() error {
	errorLogic := BusinessLogicReason(e.ErrorCode)
	jsonByte, err := json.Marshal(errorLogic)
	if err != nil {
		log.Errorf("BusinessException: %s, FileCaller: %s, LineCaller: %d, error: %v", errorLogic, e.File, e.Line, e.Err)
	} else {
		log.Errorf("BusinessException: %s, FileCaller: %s, LineCaller: %d, error: %v", string(jsonByte), e.File, e.Line, e.Err)
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
