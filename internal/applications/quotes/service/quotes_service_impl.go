package service

import (
	"context"
	"myapp/internal/applications/quotes/dto"
	"myapp/internal/applications/quotes/repository/outbound"
)

type QuotesServiceImpl struct {
	outbound outbound.QuotesOutbound
}

func NewQuotesService(outbound outbound.QuotesOutbound) *QuotesServiceImpl {
	return &QuotesServiceImpl{
		outbound: outbound,
	}
}

func (s *QuotesServiceImpl) GetQuotes(ctx context.Context, queryParameter map[string]string) (*dto.QuoteApiResponse, error) {
	result, err := s.outbound.GetQuotes(ctx, queryParameter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *QuotesServiceImpl) PostQuotes(ctx context.Context, reqBody *dto.QuoteRequest) (*dto.QuoteApiResponse, error) {

	reqApiBody := dto.QuoteApiRequest{
		Name:   reqBody.Name,
		Author: reqBody.Author,
		Quote:  reqBody.Quote,
	}

	result, err := s.outbound.PostQuotes(ctx, reqApiBody)
	if err != nil {
		return nil, err
	}

	return result, nil
}
