package exceptions

import "net/http"

type ErrorLogic struct {
	//ErrName  string
	ErrCode  int
	HttpCode int
	Message  string
}

const (
	DataAlreadyExist = 10001
	DataNotFound     = 10002
	DataCreateFailed = 10003
	DataUpdateFailed = 10004
	DataDeleteFailed = 10005
	DataGetFailed    = 10006
	//OtherError         = 10007
)

var businessLogicReason = map[int]ErrorLogic{
	DataAlreadyExist: {ErrCode: DataAlreadyExist, HttpCode: http.StatusUnprocessableEntity, Message: "data is already exist"},
	DataNotFound:     {ErrCode: DataNotFound, HttpCode: http.StatusNotFound, Message: "data not found"},
	DataCreateFailed: {ErrCode: DataCreateFailed, HttpCode: http.StatusUnprocessableEntity, Message: "create data failed"},
	DataUpdateFailed: {ErrCode: DataUpdateFailed, HttpCode: http.StatusUnprocessableEntity, Message: "update data failed"},
	DataDeleteFailed: {ErrCode: DataDeleteFailed, HttpCode: http.StatusUnprocessableEntity, Message: "delete data failed"},
	DataGetFailed:    {ErrCode: DataGetFailed, HttpCode: http.StatusUnprocessableEntity, Message: "get data failed"},
	//OtherError:         {ErrCode: EBL10007, HttpCode: http.StatusInternalServerError, Message: "your explanation of error EBL = error business logic"},
}

func BusinessLogicReason(code int) ErrorLogic {
	return businessLogicReason[code]
}
