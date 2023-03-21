package transaction

import (
	"context"
	"fmt"
	"myapp/ent"
)

type TrxServiceImpl struct {
	client *ent.Client
}

func NewTxServiceImpl(client *ent.Client) *TrxServiceImpl {
	return &TrxServiceImpl{client: client}
}

func (r *TrxServiceImpl) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
