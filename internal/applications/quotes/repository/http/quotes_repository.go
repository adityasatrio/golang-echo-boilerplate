package http

import (
	"context"
)

type QuotesRepository interface {
	GetQuotes(ctx context.Context) error
}
