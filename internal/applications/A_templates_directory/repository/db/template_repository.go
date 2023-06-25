package db

import (
	"context"
	"myapp/ent"
)

type TemplateRepository interface {
	DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error)
}
