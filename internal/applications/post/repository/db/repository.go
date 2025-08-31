package db

import (
	"context"
	"myapp/ent"
)

type PostRepository interface {
	DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error)
}
