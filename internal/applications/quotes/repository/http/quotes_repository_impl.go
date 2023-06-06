package http

import (
	"context"
	"net/http"
)

type QuotesRepositoryImpl struct {
	client *http.Client
}

func NewQuotesRepositoryImpl(client *http.Client) *QuotesRepositoryImpl {
	return &QuotesRepositoryImpl{client: client}
}

func (r *QuotesRepositoryImpl) GetQuotes(ctx context.Context) (*Quote, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://catfact.ninja/fact", nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

}
