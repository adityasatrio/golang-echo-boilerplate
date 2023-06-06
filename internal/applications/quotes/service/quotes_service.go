package http

import (
	"context"
)

type QuotesService interface {
	GetQuotes(ctx context.Context) error
}
