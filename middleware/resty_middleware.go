package middleware

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
)

func LogRequest(logName string) resty.RequestLogCallback {
	return func(logRq *resty.RequestLog) error {
		logHeader := logRq.Header
		logRequestBody := logRq.Body
		log.Infof("name : %s , header: %s , reqBody : %s", logName, logHeader, logRequestBody)
		return nil
	}

}

func LogResponse(logName string) resty.ResponseLogCallback {
	return func(logRs *resty.ResponseLog) error {
		logHeader := logRs.Header
		logRequestBody := logRs.Body
		log.Infof("name : %s , header: %s , reqBody : %s", logName, logHeader, logRequestBody)
		return nil
	}
}
