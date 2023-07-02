package outbound

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/internal/applications/quotes/dto"
	"myapp/middleware"
	"net/http"
	"time"
)

type QuoteOutboundImpl struct {
	ClientApi *resty.Client
}

func NewQuoteOutboundImpl(clientApi *resty.Client) *QuoteOutboundImpl {
	clientApi.
		SetTimeout(60 * time.Second).
		SetRetryCount(3).                            // Maximum number of retries
		SetRetryWaitTime(100 * time.Millisecond).    // Time to wait between retries
		SetRetryMaxWaitTime(500 * time.Millisecond). // Maximum time to wait between retries
		AddRetryCondition(func(response *resty.Response, err error) bool {
			return response.StatusCode() != http.StatusOK
		}).
		OnRequestLog(middleware.LogRequest("NewQuoteOutboundImpl")).
		OnResponseLog(middleware.LogResponse("NewQuoteOutboundImpl"))

	return &QuoteOutboundImpl{
		ClientApi: clientApi,
	}
}

func (o *QuoteOutboundImpl) GetQuotes(ctx context.Context) (*dto.QuoteApiResponse, error) {
	headers := map[string]string{
		"Content-Type":    "application/json",
		"x-custom-header": "custom value",
	}

	queryParameter := map[string]string{
		"sampleInt":    fmt.Sprintf("%d", 1),
		"sampleString": fmt.Sprintf("%s", "string"),
	}

	hostQuoteUrl := viper.GetString("outbound.quotes.get-url")
	response := &dto.QuoteApiResponse{}
	resp, err := o.ClientApi.R().
		SetResult(response).
		SetContext(ctx).
		SetHeaders(headers).
		SetQueryParams(queryParameter).
		Get(hostQuoteUrl)

	fmt.Printf("Response body: %s\n", resp.String())
	if err != nil {
		log.Errorf("error http libs ", err)
		return response, err
	}

	if resp.StatusCode() != 200 {
		errMsg := fmt.Sprintf("Received non-200 response code : %d", resp.RawResponse.StatusCode)
		log.Errorf(errMsg)
		return response, errors.New(errMsg)
	}

	return response, nil
}

func (o *QuoteOutboundImpl) PostQuotes(ctx context.Context, reqBody dto.QuoteApiRequest) (*dto.QuoteApiResponse, error) {
	headers := map[string]string{
		"Content-Type":    "application/json",
		"x-custom-header": "custom value",
	}

	hostQuoteUrl := viper.GetString("outbound.quotes.post-url")

	response := &dto.QuoteApiResponse{}
	resp, err := o.ClientApi.R().
		SetResult(response).
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(reqBody).
		Post(hostQuoteUrl)

	if err != nil {
		log.Errorf("error http libs ", err)
		return response, err
	}

	if resp.StatusCode() != 200 {
		errMsg := fmt.Sprintf("Received non-200 response code : %d", resp.RawResponse.StatusCode)
		log.Errorf(errMsg)
		return response, errors.New(errMsg)
	}

	return response, nil
}
