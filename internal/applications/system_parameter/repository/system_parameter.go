package repository

import (
	"context"
	"myapp/ent"
)

type (
	SystemParameterRepository interface {
		Create(ctx context.Context) (*ent.System_parameter, error)
	}
)
