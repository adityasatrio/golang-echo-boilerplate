package db

import (
	"context"
	"myapp/ent"
)

type BookRepository interface {
	DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error)
}
