package outbound

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/configs/http"
	"myapp/internal/applications/quotes/dto"
)

type QuoteOutboundImpl struct {
	ClientApi *resty.Client
}

func NewQuoteOutbound() *QuoteOutboundImpl {
	return &QuoteOutboundImpl{
		ClientApi: http.New(),
	}
}

func (o *QuoteOutboundImpl) GetQuotes(ctx context.Context, queryParameter map[string]string) (*dto.QuoteApiResponse, error) {
	headers := map[string]string{
		"Content-Type":    "application/json",
		"x-custom-header": "custom value",
	}

	queryParameter = map[string]string{
		"sampleInt":    fmt.Sprintf("%d", 1),
		"sampleString": fmt.Sprintf("%s", "string"),
	}

	hostQuoteUrl := viper.GetString("outbound.quotes.getUrl")
	response := &dto.QuoteApiResponse{}
	resp, err := o.ClientApi.R().
		SetResult(response).
		SetContext(ctx).
		SetHeaders(headers).
		SetQueryParams(queryParameter).
		Get(hostQuoteUrl)

	log.Infof("Response body: %s\n", resp.String())
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

	hostQuoteUrl := viper.GetString("outbound.quotes.postUrl")
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
