package transaction

import (
	"context"
	"fmt"
	"myapp/ent"
)

type TrxImpl struct {
	client *ent.Client
}

func NewTrx(client *ent.Client) *TrxImpl {
	return &TrxImpl{client: client}
}

func (r *TrxImpl) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, errRollback)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
