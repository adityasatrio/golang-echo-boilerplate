package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal/applications/quotes/dto"
	mock_outbound "myapp/mocks/applications/quotes/repository/outbound"
	"testing"
)

func TestGetQuotes(t *testing.T) {

	response := &dto.QuoteApiResponse{
		Author: "author",
		Quote:  "quote",
	}

	mockOutbound := new(mock_outbound.QuotesOutbound)
	service := NewQuotesService(mockOutbound)
	ctx := context.Background()

	queryParameter := map[string]string{"query1": "value1", "query2": "value2"}
	mockOutbound.On("GetQuotes", ctx, queryParameter).
		Return(response, nil)

	result, err := service.GetQuotes(ctx, queryParameter)

	assert.NoError(t, err)
	assert.Equal(t, result, response)
}

func TestGetQuotes_err(t *testing.T) {
	mockOutbound := new(mock_outbound.QuotesOutbound)
	service := NewQuotesService(mockOutbound)
	ctx := context.Background()

	queryParameter := map[string]string{"query1": "value1", "query2": "value2"}
	mockOutbound.On("GetQuotes", ctx, queryParameter).
		Return(nil, errors.New("dummy error"))

	result, err := service.GetQuotes(ctx, queryParameter)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestPostQuotes(t *testing.T) {

	request := &dto.QuoteRequest{
		Name:   "name",
		Author: "author",
		Quote:  "quote",
		Custom: "custom",
	}

	response := &dto.QuoteApiResponse{
		Author: "author",
		Quote:  "quote",
	}

	mockOutbound := new(mock_outbound.QuotesOutbound)
	service := NewQuotesService(mockOutbound)
	ctx := context.Background()

	reqApiBody := dto.QuoteApiRequest{
		Name:   "name",
		Author: "author",
		Quote:  "quote",
	}

	mockOutbound.On("PostQuotes", ctx, reqApiBody).
		Return(response, nil)

	result, err := service.PostQuotes(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, result, response)
}

func TestPostQuotes_err(t *testing.T) {

	request := &dto.QuoteRequest{
		Name:   "name",
		Author: "author",
		Quote:  "quote",
		Custom: "custom",
	}

	mockOutbound := new(mock_outbound.QuotesOutbound)
	service := NewQuotesService(mockOutbound)
	ctx := context.Background()

	reqApiBody := dto.QuoteApiRequest{
		Name:   "name",
		Author: "author",
		Quote:  "quote",
	}

	mockOutbound.On("PostQuotes", ctx, reqApiBody).
		Return(nil, errors.New("dummy error"))

	result, err := service.PostQuotes(ctx, request)

	assert.Error(t, err)
	assert.Nil(t, result)
}
