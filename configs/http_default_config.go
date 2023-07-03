package configs

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"myapp/middleware"
	"net/http"
	"time"
)

func NewDefaultResty() *resty.Client {
	clientApi := resty.New()

	timeout := viper.GetInt("http.configs.default.timeout")
	retryCount := viper.GetInt("http.configs.default.retry-count")
	retryWaitTime := viper.GetInt("http.configs.default.wait-time")
	retryMaxWaitTime := viper.GetInt("http.configs.default.max-wait-time")

	clientApi.SetTimeout(time.Duration(timeout) * time.Millisecond).
		SetRetryCount(retryCount).                                               // Maximum number of retries
		SetRetryWaitTime(time.Duration(retryWaitTime) * time.Millisecond).       // Time to wait between retries
		SetRetryMaxWaitTime(time.Duration(retryMaxWaitTime) * time.Millisecond). // Maximum time to wait between retries
		AddRetryCondition(func(response *resty.Response, err error) bool {
			return response.StatusCode() != http.StatusOK
		}).
		OnRequestLog(middleware.LogRequest("")).
		OnResponseLog(middleware.LogResponse(""))

	return clientApi
}
