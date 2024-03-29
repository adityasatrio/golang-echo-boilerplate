package outbound

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"myapp/configs/http"
	"myapp/internal/applications/quotes/dto"
	"myapp/test"
	"testing"
)

func TestNewQuoteOutbound(t *testing.T) {
	// Create a mock of the resty.Client
	mockClient := http.New()

	// Assert that the NewQuoteOutbound() function returns a QuoteOutboundImpl with the mock client
	quoteOutbound := NewQuoteOutbound()
	assert.ObjectsAreEqual(mockClient, quoteOutbound.ClientApi)
	assert.ObjectsAreEqualValues(mockClient, quoteOutbound.ClientApi)

	//if quoteOutbound.ClientApi != mockClient {
	//	t.Errorf("Expected service to be %v, but got %v", mockClient, quoteOutbound.ClientApi)
	//}
}

func TestGetQuotes(t *testing.T) {
	client := resty.New()
	ctx := context.Background()
	quoteOutboundImpl := &QuoteOutboundImpl{
		ClientApi: client,
	}

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	viper.Set("outbound.quotes.getUrl", "http://test-url")

	tests := []struct {
		name             string
		queryParameter   map[string]string
		mockStatusCode   int
		mockBodyResponse string
		expectedResp     *dto.QuoteApiResponse
	}{
		{
			name:             "Error in HTTP GET request",
			queryParameter:   map[string]string{"query1": "value1", "query2": "value2"},
			mockStatusCode:   0,
			mockBodyResponse: "",
		},
		{
			name:             "Non-200 status code",
			queryParameter:   map[string]string{"query1": "value1", "query2": "value2"},
			mockStatusCode:   404,
			mockBodyResponse: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "http://test-url",
				httpmock.NewErrorResponder(errors.New("request error")))

			resp, err := quoteOutboundImpl.GetQuotes(ctx, tt.queryParameter)

			assert.Error(t, err)
			assert.Equal(t, &dto.QuoteApiResponse{}, resp)
		})
	}
}

func TestGetQuotes_dtoResponse(t *testing.T) {
	client := resty.New()
	ctx := context.Background()
	quoteOutboundImpl := &QuoteOutboundImpl{
		ClientApi: client,
	}

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	viper.Set("outbound.quotes.getUrl", "http://test-url")

	tests := []struct {
		name             string
		queryParameter   map[string]string
		mockStatusCode   int
		mockBodyResponse string
		expectedResp     *dto.QuoteApiResponse
	}{
		{
			name:             "Successful GET request",
			queryParameter:   map[string]string{"query1": "value1", "query2": "value2"},
			mockStatusCode:   200,
			mockBodyResponse: `{"author": "test response", "quote": "test response"}`,
			expectedResp: &dto.QuoteApiResponse{
				Author: "test response",
				Quote:  "test response",
			},
		},
		{
			name:             "Failed GET request",
			queryParameter:   map[string]string{"query1": "value1", "query2": "value2"},
			mockStatusCode:   500,
			mockBodyResponse: `{"error" : true}`,
			expectedResp:     &dto.QuoteApiResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder(
				"GET",
				"http://test-url",
				test.ResponderJsonResponse(tt.mockStatusCode, tt.mockBodyResponse),
			)

			resp, err := quoteOutboundImpl.GetQuotes(ctx, tt.queryParameter)

			if tt.mockStatusCode != 200 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestPostQuotes_dtoResponse(t *testing.T) {
	client := resty.New()
	ctx := context.Background()
	quoteOutboundImpl := &QuoteOutboundImpl{
		ClientApi: client,
	}

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	viper.Set("outbound.quotes.postUrl", "http://test-url")

	tests := []struct {
		name             string
		mockRequest      dto.QuoteApiRequest
		mockStatusCode   int
		mockBodyResponse string
		expectedResp     *dto.QuoteApiResponse
	}{
		{
			name: "Successful POST request",
			mockRequest: dto.QuoteApiRequest{
				Name:   "test response",
				Author: "test response",
				Quote:  "test response",
			},
			mockStatusCode:   200,
			mockBodyResponse: `{"author": "test response", "quote": "test response"}`,
			expectedResp: &dto.QuoteApiResponse{
				Author: "test response",
				Quote:  "test response",
			},
		},
		{
			name: "Failures POST request",
			mockRequest: dto.QuoteApiRequest{
				Name:   "test response Failures",
				Author: "test response Failures",
				Quote:  "test response Failures",
			},
			mockStatusCode:   500,
			mockBodyResponse: `{"err": true}`,
			expectedResp:     &dto.QuoteApiResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder(
				"POST",
				"http://test-url",
				test.ResponderJsonResponse(tt.mockStatusCode, tt.mockBodyResponse),
			)

			resp, err := quoteOutboundImpl.PostQuotes(ctx, tt.mockRequest)

			if tt.mockStatusCode != 200 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}
