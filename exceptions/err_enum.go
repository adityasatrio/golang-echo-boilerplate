package exceptions

import "net/http"

type ErrorLogic struct {
	ErrCode  int
	HttpCode int
	Message  string
}

const (
	EBL10001 = 10001
	EBL10002 = 10002
	EBL10003 = 10003
	EBL10004 = 10004
	EBL10005 = 10005
	EBL10006 = 10006
	EBL10007 = 10007
)

var businessLogicReason = map[int]ErrorLogic{
	EBL10001: {ErrCode: EBL10001, HttpCode: http.StatusUnprocessableEntity, Message: "data is already exist"},
	EBL10002: {ErrCode: EBL10002, HttpCode: http.StatusNotFound, Message: "data not found"},
	EBL10003: {ErrCode: EBL10003, HttpCode: http.StatusUnprocessableEntity, Message: "create data failed"},
	EBL10004: {ErrCode: EBL10004, HttpCode: http.StatusUnprocessableEntity, Message: "update data failed"},
	EBL10005: {ErrCode: EBL10005, HttpCode: http.StatusUnprocessableEntity, Message: "delete data failed"},
	EBL10006: {ErrCode: EBL10006, HttpCode: http.StatusUnprocessableEntity, Message: "get data failed"},
	EBL10007: {ErrCode: EBL10007, HttpCode: http.StatusInternalServerError, Message: "your explanation of error EBL = error business logic"},
}

func BusinessLogicReason(code int) ErrorLogic {
	return businessLogicReason[code]
}
