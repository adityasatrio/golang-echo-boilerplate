package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"time"
)

// logHeaderWithRequestIDFormat is a new logger header template with added "request_id" field.
// This template is an extended version of github.com/labstack/gommon/log `defaultHeader`.
const logHeaderWithRequestIDFormat = `{"time":"${time_rfc3339_nano}","level":"${level}","request_id":"%s","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`

// LoggerTraceRequestID is a middleware to attach a unique request id given from the incoming request / response to the Echo's global logger.
func LoggerTraceRequestID() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// try to get request id from the request or response header.
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = res.Header().Get(echo.HeaderXRequestID)
			}

			// create a new one if it still does not exist.
			if rid == "" {
				uniqueID := uuid.NewString()
				timestamp := time.Now().Unix()
				rid = fmt.Sprintf("%s:%d", uniqueID, timestamp)
			}

			// subtitute the header template with request id.
			newHeader := fmt.Sprintf(logHeaderWithRequestIDFormat, rid)
			log.SetHeader(newHeader)

			return next(c)
		}
	}

}
