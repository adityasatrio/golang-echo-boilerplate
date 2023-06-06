package hook

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"time"
)

func VersionHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {

			if m.Op().Is(1) || m.Op().Is(2) || m.Op().Is(3) {
				start := time.Now()
				err := m.SetField("version", time.Now().UnixNano())
				if err != nil {
					// An error is returned, if the field is not defined in
					// the schema, or if the type mismatch the field type.
					log.Errorf("FAILED INIT VERSION Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
					return nil, fmt.Errorf("failed to set version: %w", err)
				}
			}

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("mutation failed: %w", err)
			}
			return v, nil
		})
	}
}
